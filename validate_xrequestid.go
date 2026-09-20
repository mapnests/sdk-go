package sdk

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidXRequestID = errors.New("x-request-id must be a valid W3C traceparent value: <2 hex version>-<32 hex trace-id>-<16 hex parent-id>-<2 hex flags>, e.g. 00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01")

	traceparentRegex = regexp.MustCompile(`^[0-9a-f]{2}-[0-9a-f]{32}-[0-9a-f]{16}-[0-9a-f]{2}$`)

	invalidTraceparentVersion = "ff"
	zeroTraceID               = strings.Repeat("0", 32)
	zeroParentID              = strings.Repeat("0", 16)
)

func ValidateXRequestID(xRequestID *string) error {
	if xRequestID == nil {
		return nil
	}

	value := *xRequestID
	if !traceparentRegex.MatchString(value) {
		return ErrInvalidXRequestID
	}

	parts := strings.Split(value, "-")
	version, traceID, parentID := parts[0], parts[1], parts[2]

	if version == invalidTraceparentVersion || traceID == zeroTraceID || parentID == zeroParentID {
		return ErrInvalidXRequestID
	}

	return nil
}
