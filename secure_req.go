// secure_req.go
package sdk

/*
#cgo LDFLAGS: -ldl
#include <stdlib.h>

typedef struct {
	int success;
	int statusCode;
	char *response;
	char *errorMessage;
} SecureResult;

SecureResult call_secure_func(const char *label, const char *key, const char *origin, const char *json);
*/
import "C"

import (
	"encoding/json"
	"fmt"
)

func callSecureRequest(label, apiKey, origin string, request interface{}) (int, string, string, error) {
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		return 0, "", "", fmt.Errorf("failed to marshal request: %v", err)
	}

	result := performSecureRequest(label, apiKey, origin, string(jsonBytes))
	if result.Success == false {
		return result.StatusCode, "", result.ErrorMessage, fmt.Errorf("secure request failed")
	}

	return result.StatusCode, result.Response, "", nil
}
