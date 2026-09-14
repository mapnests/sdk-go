package sdk

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

const (
	baseURL = "http://192.168.169.58:9080"
)

func buildURLFromJSON(label, jsonStr string) (string, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return "", err
	}

	switch label {
	//Geomap APIs

	case "reverse":
		lat := data["Lat"].(float64)
		lon := data["Lon"].(float64)
		var acceptedLanguage string
		if data["AcceptedLanguage"] != nil {
			if lang, ok := data["AcceptedLanguage"].(string); ok && lang != "" {
				acceptedLanguage = "&acceptedLanguage=" + url.QueryEscape(lang)
			}
		}
		return fmt.Sprintf("%s/geomap/api/v2/reverse?lat=%f&lon=%f%s", baseURL, lat, lon, acceptedLanguage), nil

	case "geocode":
		query := url.QueryEscape(data["Query"].(string))
		var limit string
		if data["Limit"] != nil {
			if l, ok := data["Limit"].(float64); ok && l > 0 {
				limit = "&limit=" + strconv.FormatFloat(l, 'f', 0, 64)
			}
		}
		var acceptedLanguage string
		if data["AcceptedLanguage"] != nil {
			if lang, ok := data["AcceptedLanguage"].(string); ok && lang != "" {
				acceptedLanguage = "&acceptedLanguage=" + url.QueryEscape(lang)
			}
		}
		return fmt.Sprintf("%s/geomap/api/v2/geocode?q=%s%s%s", baseURL, query, limit, acceptedLanguage), nil

	case "search":
		var lat string
		var lon string
		var limit string
		var page string
		var activeLocations string
		var radius string
		var acceptedLanguage string
		if data["Limit"] != nil && data["Limit"].(float64) > 0 {
			limit = "&limit=" + strconv.FormatFloat(data["Limit"].(float64), 'f', 0, 64)
		}
		if data["Lat"] != nil && data["Lat"].(float64) > 0 {
			lat = "&lat=" + strconv.FormatFloat(data["Lat"].(float64), 'f', 2, 64)
		}
		if data["Lon"] != nil && data["Lon"].(float64) > 0 {
			lon = "&lon=" + strconv.FormatFloat(data["Lon"].(float64), 'f', 2, 64)
		}
		if data["Radius"] != nil && data["Radius"].(float64) > 0 {
			radius = "&radius=" + strconv.FormatFloat(data["Radius"].(float64), 'f', 0, 64)
		}
		if data["Page"] != nil && data["Page"].(float64) > 0 {
			page = "&page=" + strconv.FormatFloat(data["Page"].(float64), 'f', 0, 64)
		}
		if data["ActiveLocations"] != nil && data["ActiveLocations"].(bool) {
			activeLocations = "&activeLocations=true"
		}
		if data["AcceptedLanguage"] != nil {
			if lang, ok := data["AcceptedLanguage"].(string); ok && lang != "" {
				acceptedLanguage = "&acceptedLanguage=" + url.QueryEscape(lang)
			}
		}
		query := url.QueryEscape(data["Query"].(string))
		return fmt.Sprintf("%s/geomap/api/v2/search?q=%s%s%s%s%s%s%s%s", baseURL, query, lat, lon, radius, page, limit, activeLocations, acceptedLanguage), nil

	case "searchByRadius":
		var lat string
		var lon string
		var limit string
		var page string
		var activeLocations string
		var radius string
		var acceptedLanguage string
		if data["Limit"] != nil && data["Limit"].(float64) > 0 {
			limit = "&limit=" + strconv.FormatFloat(data["Limit"].(float64), 'f', 0, 64)
		}
		if data["Lat"] != nil && data["Lat"].(float64) > 0 {
			lat = "&lat=" + strconv.FormatFloat(data["Lat"].(float64), 'f', 2, 64)
		}
		if data["Lon"] != nil && data["Lon"].(float64) > 0 {
			lon = "&lon=" + strconv.FormatFloat(data["Lon"].(float64), 'f', 2, 64)
		}
		if data["Radius"] != nil && data["Radius"].(float64) > 0 {
			radius = "&radius=" + strconv.FormatFloat(data["Radius"].(float64), 'f', 0, 64)
		}
		if data["Page"] != nil && data["Page"].(float64) > 0 {
			page = "&page=" + strconv.FormatFloat(data["Page"].(float64), 'f', 0, 64)
		}
		if data["ActiveLocations"] != nil && data["ActiveLocations"].(bool) {
			activeLocations = "&activeLocations=true"
		}
		if data["AcceptedLanguage"] != nil {
			if lang, ok := data["AcceptedLanguage"].(string); ok && lang != "" {
				acceptedLanguage = "&acceptedLanguage=" + url.QueryEscape(lang)
			}
		}
		query := url.QueryEscape(data["Query"].(string))
		return fmt.Sprintf("%s/geomap/api/v2/search/radius?q=%s%s%s%s%s%s%s%s", baseURL, query, lat, lon, radius, page, limit, activeLocations, acceptedLanguage), nil

	case "autocomplete":
		query := url.QueryEscape(data["Query"].(string))
		var lat string
		var lon string
		var limit string
		var zoneActiveOnly string
		var acceptedLanguage string
		if data["Limit"] != nil && data["Limit"].(float64) > 0 {
			limit = "&limit=" + strconv.FormatFloat(data["Limit"].(float64), 'f', 0, 64)
		}
		if data["Lat"] != nil && data["Lat"].(float64) > 0 {
			lat = "&lat=" + strconv.FormatFloat(data["Lat"].(float64), 'f', 2, 64)
		}
		if data["Lon"] != nil && data["Lon"].(float64) > 0 {
			lon = "&lon=" + strconv.FormatFloat(data["Lon"].(float64), 'f', 2, 64)
		}
		if data["ActiveZone"] != nil {
			zoneActiveOnly = "&zoneActiveOnly=" + strconv.FormatBool(data["ActiveZone"].(bool))
		}
		if data["AcceptedLanguage"] != nil {
			if lang, ok := data["AcceptedLanguage"].(string); ok && lang != "" {
				acceptedLanguage = "&acceptedLanguage=" + url.QueryEscape(lang)
			}
		}
		return fmt.Sprintf("%s/geomap/api/v2/autocomplete?q=%s%s%s%s%s%s", baseURL, query, zoneActiveOnly, lat, lon, limit, acceptedLanguage), nil

	case "detailsByPlaceId":
		placeID := url.QueryEscape(data["PlaceID"].(string))
		return fmt.Sprintf("%s/geomap/api/v1/details/%s", baseURL, placeID), nil

	//Routemap APIs

	case "distanceMatrix":
		fromLat := data["OriginLat"].(float64)
		fromLon := data["OriginLon"].(float64)
		toLat := data["DestLat"].(float64)
		toLon := data["DestLon"].(float64)
		mode := url.QueryEscape(data["Mode"].(string))
		return fmt.Sprintf("%s/routemap/api/v3/routes/distancematrix?fromLat=%f&fromLong=%f&toLat=%f&toLong=%f&mode=%s", baseURL, fromLat, fromLon, toLat, toLon, mode), nil

	case "distanceMatrixDetails":
		fromLat := data["OriginLat"].(float64)
		fromLon := data["OriginLon"].(float64)
		toLat := data["DestLat"].(float64)
		toLon := data["DestLon"].(float64)
		mode := url.QueryEscape(data["Mode"].(string))
		return fmt.Sprintf("%s/routemap/api/v2/routes/distancematrixdetails?fromLat=%f&fromLong=%f&toLat=%f&toLong=%f&mode=%s", baseURL, fromLat, fromLon, toLat, toLon, mode), nil

	case "multiSourceRouteSummary":
		return fmt.Sprintf("%s/routemap/api/v1/routes/multi-source-summary", baseURL), nil
	case "multiSourceRouteSummaryWithoutGeometry":
		return fmt.Sprintf("%s/routemap/api/v1/routes/multi-source-summary-without-geometry", baseURL), nil
	case "pairWiseRouteSummary":
		return fmt.Sprintf("%s/routemap/api/v2/routes/pairwise-summary", baseURL), nil

	case "snapToRoad":
		return fmt.Sprintf("%s/routemap/api/v1/nearest/road", baseURL), nil

	case "multiStopPoints":
		return fmt.Sprintf("%s/routemap/api/v1/routes/multi-stoppoints", baseURL), nil

	// ETA APIs
	case "etaWithoutGeometryByMode":
		fromLat := data["FromLat"].(float64)
		fromLon := data["FromLon"].(float64)
		toLat := data["ToLat"].(float64)
		toLon := data["ToLon"].(float64)
		mode := url.QueryEscape(data["Mode"].(string))
		var route_count string
		if data["NoOfRoutes"] != nil && data["NoOfRoutes"].(float64) > 0 {
			route_count = "&route_count=" + strconv.FormatFloat(data["NoOfRoutes"].(float64), 'f', 0, 64)
		}
		return fmt.Sprintf("%s/eta/api/v1/get-eta?from_lat=%f&from_lon=%f&to_lat=%f&to_lon=%f&mode=%s%s", baseURL, fromLat, fromLon, toLat, toLon, mode, route_count), nil

	case "etaWithGeometryByMode":
		fromLat := data["FromLat"].(float64)
		fromLon := data["FromLon"].(float64)
		toLat := data["ToLat"].(float64)
		toLon := data["ToLon"].(float64)
		mode := url.QueryEscape(data["Mode"].(string))
		return fmt.Sprintf("%s/eta/api/v1/get-eta/details?from_lat=%f&from_lon=%f&to_lat=%f&to_lon=%f&mode=%s", baseURL, fromLat, fromLon, toLat, toLon, mode), nil

	case "etaWithGeometry":
		fromLat := data["FromLat"].(float64)
		fromLon := data["FromLon"].(float64)
		toLat := data["ToLat"].(float64)
		toLon := data["ToLon"].(float64)
		return fmt.Sprintf("%s/eta/api/v1/get-eta/details/multiple-modes?from_lat=%f&from_lon=%f&to_lat=%f&to_lon=%f", baseURL, fromLat, fromLon, toLat, toLon), nil

	default:
		return "", fmt.Errorf("unsupported label: %s", label)
	}
}
