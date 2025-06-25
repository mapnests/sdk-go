package sdk

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"time"
)

type SecureResult struct {
	Success      bool
	StatusCode   int
	Response     string
	ErrorMessage string
}

const (
	baseUrl              = "https://engine.mapnests.com"
	TOKEN_EXPIRY_SECONDS = 30
)

var token string

func buildURLFromJSON(label, baseUrl, jsonStr string) (string, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return "", err
	}

	switch label {
	case "geocode":
		query := url.QueryEscape(data["Query"].(string))
		lang := url.QueryEscape(data["Language"].(string))
		limit := int(data["Limit"].(float64))
		return fmt.Sprintf("%s/geomap/api/v1/geocode?q=%s&language=%s&limit=%d", baseUrl, query, lang, limit), nil
	case "reverseGeocode":
		lat := data["Lat"].(float64)
		lon := data["Lon"].(float64)
		return fmt.Sprintf("%s/geomap/api/v1/reverse?lat=%f&lon=%f", baseUrl, lat, lon), nil
	case "distanceMatrix", "distanceMatrixDetails":
		fromLat := data["OriginLat"].(float64)
		fromLon := data["OriginLon"].(float64)
		toLat := data["DestLat"].(float64)
		toLon := data["DestLon"].(float64)
		mode := url.QueryEscape(data["Mode"].(string))
		path := "distancematrix"
		if label == "distanceMatrixDetails" {
			path = "distancematrixdetails"
		}
		return fmt.Sprintf("%s/routemap/api/v1/routes/%s?fromLat=%f&fromLong=%f&toLat=%f&toLong=%f&mode=%s", baseUrl, path, fromLat, fromLon, toLat, toLon, mode), nil
	case "search":
		query := url.QueryEscape(data["Query"].(string))
		return fmt.Sprintf("%s/geomap/api/v1/search?q=%s", baseUrl, query), nil
	default:
		return "", fmt.Errorf("unsupported label: %s", label)
	}
}

func generateToken(apiKey string) (string, error) {
	randInt, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	expires := time.Now().Unix() + TOKEN_EXPIRY_SECONDS

	input := fmt.Sprintf("%d:%d:%s", expires, randInt.Int64(), apiKey)
	hash := sha256.Sum256([]byte(input))
	sign := hex.EncodeToString(hash[:])

	jsonPayload := fmt.Sprintf(`{"random":%d,"expires":%d,"sign":"%s"}`,
		randInt.Int64(), expires, sign)

	return base64.StdEncoding.EncodeToString([]byte(jsonPayload)), nil
}

func doubleBase64Encode(s string) string {
	first := base64.StdEncoding.EncodeToString([]byte(s))
	return base64.StdEncoding.EncodeToString([]byte(first))
}

func performSecureRequest(label, apiKey, origin, jsonRequest string) SecureResult {
	urlStr, err := buildURLFromJSON(label, baseUrl, jsonRequest)
	if err != nil {
		return SecureResult{false, 0, "", err.Error()}
	}

	tokenHeader, _ := generateToken(apiKey)
	client := &http.Client{Timeout: 30 * time.Second}

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return SecureResult{false, 0, "", err.Error()}
	}

	req.Header.Set("X-API-KEY", apiKey)
	req.Header.Set("Origin", origin)
	req.Header.Set("fxsrf", tokenHeader)

	resp, err := client.Do(req)
	if err != nil {
		return SecureResult{false, 0, "", err.Error()}
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return SecureResult{false, 0, "", err.Error()}
	}

	if resp.StatusCode == 401 {
		h := resp.Header.Get("cf-ray-status-id-tn")
		if h != "" {
			encoded := doubleBase64Encode(h)
			token = encoded

			tokenHeader, _ = generateToken(apiKey)
			req.Header.Set("fxsrf", tokenHeader)
			resp, err = client.Do(req)
			if err != nil {
				return SecureResult{false, 0, "", err.Error()}
			}
			defer resp.Body.Close()
			bodyBytes, _ = io.ReadAll(resp.Body)
		}
	}

	return SecureResult{
		Success:    resp.StatusCode >= 200 && resp.StatusCode < 300,
		StatusCode: resp.StatusCode,
		Response:   string(bodyBytes),
	}
}
