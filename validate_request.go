package sdk

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	ErrQueryLength = errors.New("Please provide a query parameter with more than 2 characters.")
	multiSpaceRegex = regexp.MustCompile(`\s+`)
)

func ValidateAndNormalizeQuery(query string) (string, error) {
	query = multiSpaceRegex.ReplaceAllString(strings.TrimSpace(query), " ")
	if len(query) < 3 {
		return "", fmt.Errorf("Error: %s", ErrQueryLength)
	}
	return query, nil
}