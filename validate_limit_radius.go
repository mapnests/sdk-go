package sdk

import (
	"errors"
	"fmt"
)

const (
	SearchMaxLimit  = 20
	SearchMaxRadius = 150000

	SearchByRadiusMaxLimit  = 20
	SearchByRadiusMaxRadius = 150000

	GeocodeMaxLimit = 10

	AutocompleteMaxLimit  = 20
	AutocompleteMaxRadius = 5000

	ETAMaxRoutes = 5
)

var (
	ErrInvalidLimit  = errors.New("invalid query params, limit must be greater than 0")
	ErrLimitExceeded = errors.New("invalid query params, limit exceeds maximum allowed value")

	ErrInvalidRadius  = errors.New("invalid query params, radius must be greater than 0")
	ErrRadiusExceeded = errors.New("invalid query params, radius exceeds maximum allowed value")

	ErrRadiusRequiresLatLon = errors.New("invalid query params, lat and lon are required when radius is provided")
)

func ValidateLimit(limit *int64, maxLimit int64) error {
	if limit == nil {
		return nil
	}
	if *limit <= 0 {
		return ErrInvalidLimit
	}
	if *limit > maxLimit {
		return fmt.Errorf("%w: got %d, max allowed is %d", ErrLimitExceeded, *limit, maxLimit)
	}
	return nil
}

func ValidateRadiusPtr(radius *int64, maxRadius int64) error {
	if radius == nil {
		return nil
	}
	return ValidateRadius(*radius, maxRadius)
}

func ValidateRadius(radius int64, maxRadius int64) error {
	if radius <= 0 {
		return ErrInvalidRadius
	}
	if radius > maxRadius {
		return fmt.Errorf("%w: got %d, max allowed is %d", ErrRadiusExceeded, radius, maxRadius)
	}
	return nil
}

// ValidateRadiusRequiresLatLon checks that, when radius is provided, both
// lat and lon are also provided, since a radius search has no center point
// without them.
func ValidateRadiusRequiresLatLon(radius *int64, lat, lon *float64) error {
	if radius == nil {
		return nil
	}
	if lat == nil || lon == nil {
		return ErrRadiusRequiresLatLon
	}
	return nil
}
