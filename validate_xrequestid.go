package sdk

import (
	"errors"
	"regexp"
	"strings"
)

var (
	traceparentRegex = regexp.MustCompile(`^([0-9a-f]{2})-([0-9a-f]{32})-([0-9a-f]{16})-([0-9a-f]{2})$`)

	errTraceparentFormat   = errors.New("must be a valid W3C traceparent header (version-traceid-parentid-traceflags).")
	errTraceparentVersion  = errors.New("version must not be ff.")
	errTraceparentTraceID  = errors.New("traceid must not be all zeros.")
	errTraceparentParentID = errors.New("parentid must not be all zeros.")

	zeroTraceID  = strings.Repeat("0", 32)
	zeroParentID = strings.Repeat("0", 16)
)

// ValidateTraceparent returns a descriptive error if the value is not a valid
// W3C traceparent, or nil if it is valid (or nil).
func ValidateTraceparent(traceparent *string) error {
	if traceparent == nil {
		return nil
	}

	match := traceparentRegex.FindStringSubmatch(*traceparent)
	if match == nil {
		return errTraceparentFormat
	}

	version, traceID, parentID := match[1], match[2], match[3]
	switch {
	case version == "ff":
		return errTraceparentVersion
	case traceID == zeroTraceID:
		return errTraceparentTraceID
	case parentID == zeroParentID:
		return errTraceparentParentID
	}

	return nil
}
