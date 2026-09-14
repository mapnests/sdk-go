package sdk

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidEtaMode = errors.New("invalid mode, eta service only supports car, cng and motorcycle")

	allowedEtaModes = map[Mode]bool{
		TravelModeCar:        true,
		TravelModeCNG:        true,
		TravelModeMotorcycle: true,
	}
)

func ValidateEtaMode(mode Mode) error {
	if !allowedEtaModes[mode] {
		return fmt.Errorf("%w: got %q", ErrInvalidEtaMode, mode)
	}
	return nil
}
