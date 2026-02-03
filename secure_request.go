package sdk

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

type SecureResult struct {
	Success      bool
	StatusCode   int
	Response     string
	ErrorMessage string
}

var HTTPMethodMap = map[string]string{
	"reverse":                 http.MethodGet,
	"search":                  http.MethodGet,
	"distanceMatrix":          http.MethodGet,
	"distanceMatrixDetails":   http.MethodGet,
	"pairWiseRouteSummary":    http.MethodPost,
	"multiSourceRouteSummary": http.MethodPost,
	"autocomplete":            http.MethodGet,
	"autocompleteWithoutZone": http.MethodGet,
	"searchByRadius":          http.MethodGet,
	"detailsByPlaceId":        http.MethodGet,
}

var token string

func performSecureRequest(label string, apiKey string, origin string, timeoutMs int32, jsonRequest string) SecureResult {

	urlStr, err := buildURLFromJSON(label, jsonRequest)
	if err != nil {
		return SecureResult{false, 0, "", err.Error()}
	}

	method := HTTPMethodMap[label]
	if method == "" {
		method = http.MethodGet
	}

	tokenHeader, err := generateToken(apiKey)
	if err != nil {
		return SecureResult{false, 0, "", err.Error()}
	}

	headers := map[string]string{
		"X-API-KEY":    apiKey,
		"Origin":       origin,
		"fxsrf":        tokenHeader,
		"Content-Type": "application/json",
	}

	client := &http.Client{Timeout: time.Duration(timeoutMs) * time.Millisecond}
	var resp *http.Response

	makeRequest := func() (*http.Response, error) {
		var req *http.Request
		if method == http.MethodPost {
			req, err = http.NewRequest(http.MethodPost, urlStr, bytes.NewBuffer([]byte(jsonRequest)))
		} else {
			req, err = http.NewRequest(http.MethodGet, urlStr, nil)
		}
		if err != nil {
			return nil, err
		}
		for k, v := range headers {
			req.Header.Set(k, v)
		}
		return client.Do(req)
	}

	resp, err = makeRequest()
	if err != nil {
		return SecureResult{false, 0, "", err.Error()}
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return SecureResult{false, resp.StatusCode, "", err.Error()}
	}

	if resp.StatusCode == http.StatusUnauthorized {
		h := resp.Header.Get("cf-ray-status-id-tn")
		if h != "" {
			token = doubleBase64Encode(h)
			tokenHeader, _ = generateToken(apiKey)
			headers["fxsrf"] = tokenHeader

			resp, err = makeRequest()
			if err != nil {
				return SecureResult{false, 0, "", err.Error()}
			}
			defer resp.Body.Close()
			bodyBytes, _ = io.ReadAll(resp.Body)
		}
	}

	return SecureResult{
		Success:      resp.StatusCode >= 200 && resp.StatusCode < 300,
		StatusCode:   resp.StatusCode,
		Response:     string(bodyBytes),
		ErrorMessage: "",
	}
}
