package sdk

import (
	"errors"
	"fmt"
	"math"
)

const (
	minLat  = 20.5
	maxLat  = 26.7
	minLon  = 87.9
	maxLon  = 92.8
	epsilon = 1e-9 // Tolerance for floating point comparison
)

var (
	ErrInvalidLatitude  = errors.New("Invalid Query Params, latitude must be between 20.5 and 26.7 (Only Inside Bangladesh)")
	ErrInvalidLongitude = errors.New("Invalid Query Params, longitude must be between 87.9 and 92.8 (Only Inside Bangladesh)")
	ErrNaNLatitude      = errors.New("latitude cannot be NaN")
	ErrNaNLongitude     = errors.New("longitude cannot be NaN")
	ErrSameLatLon       = errors.New("latitude and longitude have the same value")
)

type ValidationResult struct {
	Valid       bool
	IsSameValue bool
	Error       error
}

func ValidateLatLon(lat, lon float64) error {
	if math.IsNaN(lat) {
		return ErrNaNLatitude
	}
	if math.IsNaN(lon) {
		return ErrNaNLongitude
	}
	if lat < minLat || lat > maxLat {
		return fmt.Errorf("%w: got %.6f", ErrInvalidLatitude, lat)
	}
	if lon < minLon || lon > maxLon {
		return fmt.Errorf("%w: got %.6f", ErrInvalidLongitude, lon)
	}
	if areSameValues(lat, lon) {
		return ErrSameLatLon
	}

	return nil
}

func areSameValues(lat, lon float64) bool {
	return math.Abs(lat-lon) < epsilon
}

func ValidateLatLonPtr(lat, lon *float64) error {
	if lat == nil && lon == nil {
		return nil
	}
	if lat == nil {
		return ErrNaNLatitude
	}
	if lon == nil {
		return ErrNaNLongitude
	}
	return ValidateLatLon(*lat, *lon)
}
