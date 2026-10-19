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

	tp := TraceparentFromJSON(string(jsonBytes))
	if tp != "" {
		if err := ValidateTraceparent(&tp); err != nil {
			fmt.Printf("Invalid traceparent ignored: %v The request will be sent without it. \n", err)
			tp = ""
		}
	}
	// Traceparent travels as a header, never in the body (strict servers reject unknown fields).
	jsonBytes = stripTraceparent(jsonBytes)

	result := performSecureRequest(label, apiKey, origin, timeoutMs, string(jsonBytes), tp)

	if !result.Success {
		errMsg := result.ErrorMessage
		if errMsg == "" {
			errMsg = result.Response
		}
		return result.StatusCode, "", errMsg, fmt.Errorf("secure request failed")
	}

	return result.StatusCode, result.Response, "", nil
}

// stripTraceparent removes the Traceparent field so the request is sent without it.
func stripTraceparent(jsonBytes []byte) []byte {
	var data map[string]any
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		return jsonBytes
	}
	delete(data, "Traceparent")
	out, err := json.Marshal(data)
	if err != nil {
		return jsonBytes
	}
	return out
}
