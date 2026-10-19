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
)

var (
	ErrEmptyPolyline    = errors.New("polyline is required")
	ErrInvalidPolyline  = errors.New("invalid polyline encoding")
	ErrInvalidPrecision = errors.New("precision must be an integer between 1 and 10")
)

type StaticMapRequest struct {
	Polyline    string       `json:"polyline"`
	Precision   string       `json:"precision"`
	Markers     []Coordinate `json:"markers"`
	Traceparent *string
}

type StaticMapResponse struct {
	Image          []byte  // rendered map image (PNG)
	DistanceMeters float64 // total length of the polyline
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
		if err != nil || p < 1 || p > 10 {
			return nil, fmt.Errorf("%w: got %q", ErrInvalidPrecision, request.Precision)
		}
		precision = p
	}
	request.Precision = strconv.Itoa(precision)

	distance, err := PolylineDistance(request.Polyline, precision)
	if err != nil {
		return nil, err
	}

	// The API expects the polyline base64-encoded; callers pass it raw.
	request.Polyline = base64.StdEncoding.EncodeToString([]byte(request.Polyline))

	image, err := s.request("staticMap", request)
	if err != nil {
		return nil, err
	}

	return &StaticMapResponse{Image: image, DistanceMeters: distance}, nil
}

// PolylineDistance returns the total length in meters of an encoded polyline
// (Google polyline algorithm) with the given coordinate precision.
func PolylineDistance(polyline string, precision int) (float64, error) {
	points, err := DecodePolyline(polyline, precision)
	if err != nil {
		return 0, err
	}

	var total float64
	for i := 1; i < len(points); i++ {
		total += haversineMeters(points[i-1], points[i])
	}
	return total, nil
}

// DecodePolyline decodes an encoded polyline into coordinates.
func DecodePolyline(polyline string, precision int) ([]Coordinate, error) {
	factor := math.Pow10(precision)
	var points []Coordinate
	var lat, lon int64

	for i := 0; i < len(polyline); {
		dLat, next, ok := decodePolylineValue(polyline, i)
		if !ok {
			return nil, ErrInvalidPolyline
		}
		dLon, next, ok := decodePolylineValue(polyline, next)
		if !ok {
			return nil, ErrInvalidPolyline
		}
		i = next

		lat += dLat
		lon += dLon
		points = append(points, Coordinate{Lat: float64(lat) / factor, Lon: float64(lon) / factor})
	}
	return points, nil
}

func decodePolylineValue(s string, i int) (value int64, next int, ok bool) {
	var result int64
	var shift uint
	for ; i < len(s); i++ {
		b := int64(s[i]) - 63
		if b < 0 || shift > 60 {
			return 0, 0, false
		}
		result |= (b & 0x1f) << shift
		shift += 5
		if b < 0x20 {
			if result&1 != 0 {
				return ^(result >> 1), i + 1, true
			}
			return result >> 1, i + 1, true
		}
	}
	return 0, 0, false
}

func haversineMeters(a, b Coordinate) float64 {
	const rad = math.Pi / 180
	dLat := (b.Lat - a.Lat) * rad
	dLon := (b.Lon - a.Lon) * rad
	h := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(a.Lat*rad)*math.Cos(b.Lat*rad)*math.Sin(dLon/2)*math.Sin(dLon/2)
	return 2 * earthRadiusMeters * math.Asin(math.Sqrt(h))
}
