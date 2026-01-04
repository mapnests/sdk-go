package sdk

import (
	"encoding/json"
	"fmt"
)

func callSecureRequest(label, apiKey, origin string, timeoutMs int32, request any) (int, string, string, error) {
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		return 0, "", "", fmt.Errorf("failed to marshal request: %v", err)
	}

	result := performSecureRequest(label, apiKey, origin, timeoutMs, string(jsonBytes))

	if !result.Success {
		return result.StatusCode, "", result.ErrorMessage, fmt.Errorf("secure request failed")
	}

	return result.StatusCode, result.Response, "", nil
}
