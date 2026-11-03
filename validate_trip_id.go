package sdk

import (
	"errors"
	"strings"
)

var ErrInvalidTripID = errors.New("tripId must not be empty")

func ValidateTripID(tripID string) error {
	if strings.TrimSpace(tripID) == "" {
		return ErrInvalidTripID
	}
	return nil
}
