package sdk

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
