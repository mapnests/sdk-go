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
	case "reverse":
		lat := data["Lat"].(float64)
		lon := data["Lon"].(float64)
		return fmt.Sprintf("%s/geomap/api/v1/reverse?lat=%f&lon=%f", baseURL, lat, lon), nil
	case "distanceMatrix", "distanceMatrixDetails":
		fromLat := data["OriginLat"].(float64)
		fromLon := data["OriginLon"].(float64)
		toLat 	:= data["DestLat"].(float64)
		toLon 	:= data["DestLon"].(float64)
		mode 	:= url.QueryEscape(data["Mode"].(string))
		path 	:= "distancematrix"
		version:= "v3"
		if label == "distanceMatrixDetails" {
			path = "distancematrixdetails"
			version = "v1"
		}
		return fmt.Sprintf("%s/routemap/api/%s/routes/%s?fromLat=%f&fromLong=%f&toLat=%f&toLong=%f&mode=%s", baseURL,version, path, fromLat, fromLon, toLat, toLon, mode), nil
	case "search":
		query := url.QueryEscape(data["Query"].(string))
		return fmt.Sprintf("%s/geomap/api/v1/search?q=%s", baseURL, query), nil
	case "multiSourceRouteSummary":
		return fmt.Sprintf("%s/routemap/api/v1/routes/multi-source-summary", baseURL), nil
	case "pairWiseRouteSummary":
		return fmt.Sprintf("%s/routemap/api/v1/routes/pairwise-summary", baseURL), nil
	case "autocomplete","autocompleteWithoutZone":
		query 	:= url.QueryEscape(data["Query"].(string))
		path 	:= "/all"
		var lat string
		var lon string
		var limit string
		if label == "autocomplete" {
			path = ""
		}
		if(data["Limit"]!= nil && data["Limit"].(float64) > 0){
			limit = "&limit=" + strconv.FormatFloat(data["Limit"].(float64),'f', 0, 64)
		}
		if(data["Lat"] != nil && data["Lat"].(float64) > 0 ){
			lat = "&lan="+ strconv.FormatFloat(data["Lat"].(float64),'f', 2, 64)
		}
		if(data["Lon"] != nil && data["Lon"].(float64) > 0){
			lon = "&lon="+ strconv.FormatFloat(data["Lon"].(float64),'f', 2, 64)
		}
		return fmt.Sprintf("%s/geomap/api/v1/autocomplete%s?q=%s%s%s%s", baseURL,path,query,lat,lon,limit), nil

	default:
		return "", fmt.Errorf("unsupported label: %s", label)
	}
}
