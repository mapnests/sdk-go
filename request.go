package sdk

import "encoding/json"

// traceparentFromJSON extracts the optional "Traceparent" field from a JSON
// request body. It returns nil when the field is missing or null.
func traceparentFromJSON(jsonStr string) *string {
	var data struct {
		Traceparent *string `json:"Traceparent"`
	}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return nil
	}
	return data.Traceparent
}

func (s *client) request(label string, req interface{}) ([]byte, error) {

	status, resp, errMsg, err := callSecureRequest(label, s.apiKey, s.packageName, s.timeoutMs, req)
	if err != nil {
		return nil, &APIError{Label: label, StatusCode: status, Body: errMsg}
	}
	if status != 200 {
		return nil, &APIError{Label: label, StatusCode: status, Body: resp}
	}
	return []byte(resp), nil
}
