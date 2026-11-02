package sdk

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"math"
	"strconv"
)

const (
	defaultPolylinePrecision = 6
	earthRadiusMeters        = 6371008.8
	RDPToleranceMeters       = 1.0
)

var (
	ErrEmptyPolyline    = errors.New("polyline is required")
	ErrInvalidPolyline  = errors.New("invalid polyline encoding")
	ErrInvalidPrecision = errors.New("only precision 5 or 6 is allowed")
)

type StaticMapRequest struct {
	Polyline    string       `json:"polyline"`
	Precision   string       `json:"precision"`
	Markers     []Coordinate `json:"markers"`
	Traceparent *string
}

type StaticMapResponse struct {
	ImageBase64    string
	DistanceMeters float64
}

// StaticMap renders the map image and also returns the total polyline distance.
// request.Polyline is the raw (unencoded) polyline; Precision defaults to 6.
func (s *client) StaticMap(ctx context.Context, request StaticMapRequest) (*StaticMapResponse, error) {
	if request.Polyline == "" {
		return nil, ErrEmptyPolyline
	}
	for _, m := range request.Markers {
		if err := ValidateLatLon(m.Lat, m.Lon); err != nil {
			return nil, err
		}
	}

	precision := defaultPolylinePrecision
	if request.Precision != "" {
		p, err := strconv.Atoi(request.Precision)
		if err != nil || (p != 5 && p != 6) {
			return nil, fmt.Errorf("%w: got %q", ErrInvalidPrecision, request.Precision)
		}
		precision = p
	}

	points, err := DecodePolyline(request.Polyline, precision)
	if err != nil {
		return nil, err
	}
	distance := pointsDistance(points) // distance is measured on the original (unrounded) points
	// polyline6 -> polyline5; any other precision is kept as is.
	if precision == defaultPolylinePrecision {
		precision = 5
		points = roundPoints(points, precision)
	}
	request.Precision = strconv.Itoa(precision)

	// The API expects the simplified polyline base64-encoded.
	request.Polyline = base64.StdEncoding.EncodeToString(encodePolyline(simplifyPolyline(points, RDPToleranceMeters), precision))

	image, err := s.request("staticMap", request)
	if err != nil {
		return nil, err
	}

	return &StaticMapResponse{
		ImageBase64:    base64.StdEncoding.EncodeToString(image),
		DistanceMeters: distance,
	}, nil
}

// PolylineDistance returns the total length in meters of an encoded polyline
// (Google polyline algorithm) with the given coordinate precision.
func PolylineDistance(polyline string, precision int) (float64, error) {
	points, err := DecodePolyline(polyline, precision)
	if err != nil {
		return 0, err
	}

	return pointsDistance(points), nil
}

// pointsDistance returns the total length in meters of the path through points.
func pointsDistance(points []Coordinate) float64 {
	var total float64
	for i := 1; i < len(points); i++ {
		total += haversineMeters(points[i-1], points[i])
	}
	return total
}

// DecodePolyline decodes an encoded polyline into coordinates.
func DecodePolyline(polyline string, precision int) ([]Coordinate, error) {
	factor := math.Pow10(precision)
	points := make([]Coordinate, 0, len(polyline)/4) // decoded output; pre-sized (each point takes at least 2 bytes)
	var lat, lon int64

	for i := 0; i < len(polyline); {
		dLat, next, ok := decodePolylineValue(polyline, i)
		if !ok {
			return nil, fmt.Errorf("%w at byte %d", ErrInvalidPolyline, i)
		}
		dLon, end, ok := decodePolylineValue(polyline, next)
		if !ok {
			return nil, fmt.Errorf("%w at byte %d", ErrInvalidPolyline, next)
		}
		i = end

		lat += dLat
		lon += dLon
		points = append(points, Coordinate{Lat: float64(lat) / factor, Lon: float64(lon) / factor})
	}
	return points, nil
}

// decodePolylineValue reads one varint-encoded signed value starting at index i.
func decodePolylineValue(s string, i int) (value int64, next int, ok bool) {
	var result int64
	var shift uint
	for ; i < len(s); i++ {
		c := s[i]
		if c < 63 || c > 126 || shift >= 60 { // valid chars are '?'..'~'; 12 chunks (60 bits) is the max accepted
			return 0, 0, false
		}
		b := int64(c) - 63
		result |= (b & 0x1f) << shift
		shift += 5
		if b < 0x20 {
			if result&1 != 0 { // lowest bit marks a negative number
				return ^(result >> 1), i + 1, true
			}
			return result >> 1, i + 1, true
		}
	}
	return 0, 0, false
}

// encodePolyline encodes coordinates into polyline bytes.
func encodePolyline(points []Coordinate, precision int) []byte {
	factor := math.Pow10(precision)
	buf := make([]byte, 0, len(points)*8)
	var prevLat, prevLon int64
	for _, p := range points {
		lat, lon := int64(math.Round(p.Lat*factor)), int64(math.Round(p.Lon*factor))
		buf = appendPolylineValue(buf, lat-prevLat)
		buf = appendPolylineValue(buf, lon-prevLon)
		prevLat, prevLon = lat, lon
	}
	return buf
}

// appendPolylineValue appends one signed value in polyline varint format.
func appendPolylineValue(buf []byte, v int64) []byte {
	v <<= 1 // zigzag: sign goes in the lowest bit
	if v < 0 {
		v = ^v
	}
	for ; v >= 0x20; v >>= 5 {
		buf = append(buf, byte((0x20|(v&0x1f))+63))
	}
	return append(buf, byte(v+63))
}

// roundPoints returns a new slice with every coordinate rounded to the given
// precision (decimal places). The input slice is not modified.
func roundPoints(points []Coordinate, precision int) []Coordinate {
	factor := math.Pow10(precision)
	rounded := make([]Coordinate, len(points))
	for i, p := range points {
		rounded[i] = Coordinate{
			Lat: math.Round(p.Lat*factor) / factor,
			Lon: math.Round(p.Lon*factor) / factor,
		}
	}
	return rounded
}

// simplifyPolyline applies Ramer-Douglas-Peucker with a tolerance in meters.
func simplifyPolyline(points []Coordinate, toleranceMeters float64) []Coordinate {
	if len(points) < 3 {
		return points
	}
	first, last := points[0], points[len(points)-1]
	index, maxDist := 0, 0.0
	for i := 1; i < len(points)-1; i++ {
		if d := segmentDistanceMeters(points[i], first, last); d > maxDist {
			index, maxDist = i, d
		}
	}
	if maxDist <= toleranceMeters {
		return []Coordinate{first, last}
	}
	left := simplifyPolyline(points[:index+1], toleranceMeters)
	right := simplifyPolyline(points[index:], toleranceMeters)
	return append(left[:len(left):len(left)], right[1:]...) // join, dropping duplicated split point; capped cap forces a copy
}

// segmentDistanceMeters returns the distance from p to segment a-b, using a
// local equirectangular projection around a.
func segmentDistanceMeters(p, a, b Coordinate) float64 {
	const rad = math.Pi / 180
	k := math.Cos(a.Lat*rad) * earthRadiusMeters * rad
	bx, by := (b.Lon-a.Lon)*k, (b.Lat-a.Lat)*earthRadiusMeters*rad
	px, py := (p.Lon-a.Lon)*k, (p.Lat-a.Lat)*earthRadiusMeters*rad
	l2 := bx*bx + by*by
	if l2 == 0 {
		return haversineMeters(p, a)
	}
	t := math.Max(0, math.Min(1, (px*bx+py*by)/l2))
	return math.Hypot(px-t*bx, py-t*by)
}

// haversineMeters returns the great-circle distance between two coordinates.
func haversineMeters(a, b Coordinate) float64 {
	const rad = math.Pi / 180
	dLat := (b.Lat - a.Lat) * rad
	dLon := (b.Lon - a.Lon) * rad
	h := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(a.Lat*rad)*math.Cos(b.Lat*rad)*math.Sin(dLon/2)*math.Sin(dLon/2)
	return 2 * earthRadiusMeters * math.Asin(math.Sqrt(h))
}
