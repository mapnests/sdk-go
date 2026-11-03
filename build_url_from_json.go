package sdk

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

const (
	baseURL = "https://engine.mapnests.com"
)

func buildURLFromJSON(label, jsonStr string) (string, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return "", err
	}

	switch label {
	//Geo Map APIs

	case "reverse":
		lat := data["Lat"].(float64)
		lon := data["Lon"].(float64)
		return fmt.Sprintf("%s/geomap/api/v2/reverse?lat=%f&lon=%f%s", baseURL, lat, lon, acceptLanguageParam(data)), nil

	case "search":
		var activeLocations string
		limit := positiveNumParam(data, "Limit", "limit", 0)
		lat := positiveNumParam(data, "Lat", "lat", 2)
		lon := positiveNumParam(data, "Lon", "lon", 2)
		radius := positiveNumParam(data, "Radius", "radius", 0)
		page := positiveNumParam(data, "Page", "page", 0)
		if v, _ := data["ActiveLocations"].(bool); v {
			activeLocations = "&activeLocations=true"
		}
		query := url.QueryEscape(data["Query"].(string))
		return fmt.Sprintf("%s/geomap/api/v2/search?q=%s%s%s%s%s%s%s%s", baseURL, query, lat, lon, radius, page, limit, activeLocations, acceptLanguageParam(data)), nil

	case "searchByRadius":
		var activeLocations string
		limit := positiveNumParam(data, "Limit", "limit", 0)
		lat := positiveNumParam(data, "Lat", "lat", 2)
		lon := positiveNumParam(data, "Lon", "lon", 2)
		radius := positiveNumParam(data, "Radius", "radius", 0)
		page := positiveNumParam(data, "Page", "page", 0)
		if v, _ := data["ActiveLocations"].(bool); v {
			activeLocations = "&activeLocations=true"
		}
		query := url.QueryEscape(data["Query"].(string))
		return fmt.Sprintf("%s/geomap/api/v2/search/radius?q=%s%s%s%s%s%s%s%s", baseURL, query, lat, lon, radius, page, limit, activeLocations, acceptLanguageParam(data)), nil

	case "autocomplete":
		query := url.QueryEscape(data["Query"].(string))
		var zoneActiveOnly string
		limit := positiveNumParam(data, "Limit", "limit", 0)
		lat := positiveNumParam(data, "Lat", "lat", 2)
		lon := positiveNumParam(data, "Lon", "lon", 2)
		if data["ActiveZone"] != nil {
			zoneActiveOnly = "&zoneActiveOnly=" + strconv.FormatBool(data["ActiveZone"].(bool))
		}
		return fmt.Sprintf("%s/geomap/api/v2/autocomplete?q=%s%s%s%s%s%s", baseURL, query, zoneActiveOnly, lat, lon, limit, acceptLanguageParam(data)), nil

	case "detailsByPlaceId":
		placeID := url.QueryEscape(data["PlaceID"].(string))
		return fmt.Sprintf("%s/geomap/api/v1/details/%s", baseURL, placeID), nil

	case "geocode":
		query := url.QueryEscape(data["Query"].(string))
		limit := positiveNumParam(data, "Limit", "limit", 0)
		return fmt.Sprintf("%s/geomap/api/v2/geocode?q=%s%s%s", baseURL, query, limit, acceptLanguageParam(data)), nil

		// Route Map APIs

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

	case "pairWiseRouteSummary":
		return fmt.Sprintf("%s/routemap/api/v2/routes/pairwise-summary", baseURL), nil

	case "snapToRoad":
		return fmt.Sprintf("%s/routemap/api/v1/nearest/road", baseURL), nil

	case "multiStopPoints":
		return fmt.Sprintf("%s/routemap/api/v1/routes/multi-stoppoints", baseURL), nil

	case "multiSourceRouteSummaryWithoutGeometry":
		return fmt.Sprintf("%s/routemap/api/v1/routes/multi-source-summary-without-geometry", baseURL), nil

	case "pairwiseDistanceMatrix":
		return fmt.Sprintf("%s/routemap/api/v1/routes/pairwise-distancematrix", baseURL), nil

	case "multiStopPointDetailsWithGeometry":
		return fmt.Sprintf("%s/routemap/api/v1/routes/multi-stoppoint-details-with-geometry", baseURL), nil

	// ETA APIs

	case "etaWithoutGeometryByMode":
		mode := url.QueryEscape(data["Mode"].(string))
		routeCount := positiveNumParam(data, "NoOfRoutes", "route_count", 0)
		return fmt.Sprintf("%s/eta/api/v1/get-eta?from_lat=%s&from_lon=%s&to_lat=%s&to_lon=%s&mode=%s%s",
			baseURL, num(data["FromLat"]), num(data["FromLon"]), num(data["ToLat"]), num(data["ToLon"]), mode, routeCount), nil

	case "etaWithGeometryByMode":
		mode := url.QueryEscape(data["Mode"].(string))
		var tripID string
		if v, ok := data["TripId"].(string); ok && v != "" {
			tripID = "&trip_id=" + url.QueryEscape(v)
		}
		return fmt.Sprintf("%s/eta/api/v1/get-eta/details?from_lat=%s&from_lon=%s&to_lat=%s&to_lon=%s&mode=%s%s",
			baseURL, num(data["FromLat"]), num(data["FromLon"]), num(data["ToLat"]), num(data["ToLon"]), mode, tripID), nil

	case "etaWithoutStoppage":
		return fmt.Sprintf("%s/eta/api/v1/get-eta/details/multiple-modes?from_lat=%s&from_lon=%s&to_lat=%s&to_lon=%s",
			baseURL, num(data["FromLat"]), num(data["FromLon"]), num(data["ToLat"]), num(data["ToLon"])), nil

	case "etaMultipleStoppage":
		return fmt.Sprintf("%s/eta/api/v1/get-eta/details/multi-stoppage", baseURL), nil

	case "etaWithMultiSourceMultiStopPoint":
		return fmt.Sprintf("%s/routemap/api/v1/routes/eta-with-multi-source-multi-stoppoint", baseURL), nil

	// Static Map APIs

	case "staticMap":
		return fmt.Sprintf("%s/staticmap/api/v1/get-static-map", baseURL), nil

	default:
		return "", fmt.Errorf("unsupported label: %s", label)
	}
}

// num formats a JSON number without trailing zeros or exponent.
func num(v interface{}) string {
	f, _ := v.(float64)
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// acceptLanguageParam returns "&acceptLanguage=<lang>" when the optional
// AcceptedLanguage field is set, otherwise an empty string.
func acceptLanguageParam(data map[string]interface{}) string {
	if v, ok := data["AcceptedLanguage"].(string); ok && v != "" {
		return "&acceptLanguage=" + url.QueryEscape(v)
	}
	return ""
}

// positiveNumParam returns "&<param>=<value>" when data[key] is a number
// greater than 0, formatted with the given decimals, otherwise "".
func positiveNumParam(data map[string]interface{}, key, param string, decimals int) string {
	if v, ok := data[key].(float64); ok && v > 0 {
		return "&" + param + "=" + strconv.FormatFloat(v, 'f', decimals, 64)
	}
	return ""
}
