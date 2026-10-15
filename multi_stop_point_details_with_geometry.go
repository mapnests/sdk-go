package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

type MultiStopPointDetailsWithGeometryRequest struct {
	Src         Coordinate
	StopPoints  []StopPoint
	Mode        Mode
	Traceparent *string
}

type MultiStopPointDetailsWithGeometryRoute struct {
	Distance float64 `json:"distance"`
	Duration float64 `json:"duration"`
	Geometry string  `json:"geometry"`
}

type MultiStopPointDetailsWithGeometryData struct {
	Code        string                                 `json:"code"`
	Message     string                                 `json:"message"`
	DataVersion string                                 `json:"data_version"`
	Route       MultiStopPointDetailsWithGeometryRoute `json:"route"`
}

type MultiStopPointDetailsWithGeometryResponse struct {
	Data    MultiStopPointDetailsWithGeometryData `json:"data"`
	Message string                                `json:"message"`
	Status  bool                                  `json:"status"`
}

func (s *client) MultiStopPointDetailsWithGeometry(ctx context.Context, request MultiStopPointDetailsWithGeometryRequest) (*MultiStopPointDetailsWithGeometryResponse, error) {

	if err := ValidateLatLon(request.Src.Lat, request.Src.Lon); err != nil {
		return nil, err
	}

	for _, stop := range request.StopPoints {
		if err := ValidateLatLon(stop.Lat, stop.Lon); err != nil {
			return nil, err
		}
	}

	body, err := s.request("multiStopPointDetailsWithGeometry", request)
	if err != nil {
		return nil, err
	}

	var response MultiStopPointDetailsWithGeometryResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling MultiStopPointDetailsWithGeometryResponse response: %w", err)
	}

	return &response, nil
}
