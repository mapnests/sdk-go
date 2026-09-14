package sdk

import "errors"

var ErrInvalidAcceptedLanguage = errors.New("accepted language must be either \"en\" or \"bn\"")

func ValidateAcceptedLanguage(acceptedLanguage *string) error {
	if acceptedLanguage == nil {
		return nil
	}

	switch *acceptedLanguage {
	case "en", "bn":
		return nil
	default:
		return ErrInvalidAcceptedLanguage
	}
}
