package sdk
import "C"
import (
	"encoding/json"
	"unsafe"
)

// SecureResult holds the outcome of a secure HTTP request.
type SecureResult struct {
	Success      bool
	StatusCode   int
	Response     string
	ErrorMessage string
}

func XRequestIDFromJSON(jsonRequest string) string {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonRequest), &data); err != nil {
		return ""
	}

	v, ok := data["XRequestID"]
	if !ok || v == nil {
		return ""
	}

	id, ok := v.(string)
	if !ok {
		return ""
	}

	return id
}

func performSecureRequest(label, apiKey, origin string, timeoutMs int32, jsonRequest string) SecureResult {
	cLabel := C.CString(label)
	cKey := C.CString(apiKey)
	cOrigin := C.CString(origin)
	cJSON := C.CString(jsonRequest)
	cReqID := C.CString(XRequestIDFromJSON(jsonRequest))
	defer C.free(unsafe.Pointer(cLabel))
	defer C.free(unsafe.Pointer(cKey))
	defer C.free(unsafe.Pointer(cOrigin))
	defer C.free(unsafe.Pointer(cJSON))
	defer C.free(unsafe.Pointer(cReqID))

	cr := C.perform_secure_request(cLabel, cKey, cOrigin, C.int(timeoutMs), cJSON, cReqID)
	defer C.free_secure_result(&cr)

	return SecureResult{
		Success:      cr.success != 0,
		StatusCode:   int(cr.status_code),
		Response:     C.GoString(cr.response),
		ErrorMessage: C.GoString(cr.error_message),
	}
}
