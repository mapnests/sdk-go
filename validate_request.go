package sdk

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	ErrQueryLength = errors.New("query must be more than 3 characters long")
	multiSpaceRegex = regexp.MustCompile(`\s+`)
)

func ValidateAndNormalizeQuery(query string) (string, error) {
	query = strings.TrimSpace(query)

	if len(query) < 3 {
		return "", fmt.Errorf("Error: %s", ErrQueryLength)
	}
	query = multiSpaceRegex.ReplaceAllString(query, " ")
	return query, nil
}