package sdk

import (
	"fmt"
)

func (s *client) request(label string, request interface{}) ([]byte, error) {
	status, resp, errMsg, err := callSecureRequest(label, s.apiKey, s.packageName, s.timeoutMs, request)
	if err != nil {
		return nil, fmt.Errorf("[%s] native error: %s", label, errMsg)
	}
	if status != 200 {
		return nil, fmt.Errorf("[%s] unexpected status code %d", label, status)
	}
	return []byte(resp), nil
}
