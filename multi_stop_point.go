package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

type MultiStopPointsRequest struct {
	Src        Coordinate
	StopPoints []StopPoint
	Mode       Mode
	XRequestID *string
}

type StopPoint struct {
	ID  int
	Lat float64
	Lon float64
}

type MultiStopPointsRouteSummary struct {
	RouteSummary
	Source    Coordinate
	StopPoint Coordinate
}

type MultiStopPointsRouteData struct {
	DistanceInMeters float64
	EtaInSeconds     float64
	RouteSummaries   []MultiStopPointsRouteSummary
}

type MultiStopPointsResponse struct {
	Data    MultiStopPointsRouteData `json:"data"`
	Message string                   `json:"message"`
	Status  bool                     `json:"status"`
}

func (s *client) MultiStopPoints(ctx context.Context, request MultiStopPointsRequest) (*MultiStopPointsResponse, error) {

	if err := ValidateLatLon(request.Src.Lat, request.Src.Lon); err != nil {
		return nil, err
	}

	for _, src := range request.StopPoints {
		if err := ValidateLatLon(src.Lat, src.Lon); err != nil {
			return nil, err
		}
	}

	fmt.Println("📍 MultiStopPointsRequest request:", request)

	body, err := s.request("multiStopPoints", request)
	if err != nil {
		return nil, err
	}

	var response MultiStopPointsResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshalling MultiStopPointsResponse response: %v", err)
	}

	return &response, nil
}
