package sdk

import "fmt"

type APIError struct {
	Label      string
	StatusCode int
	Body       string
}

func (e *APIError) Error() string {
	if e.StatusCode == 0 {
		return fmt.Sprintf("[%s] request failed: %s", e.Label, e.Body)
	}
	return fmt.Sprintf("[%s] status %d: %s", e.Label, e.StatusCode, e.Body)
}
