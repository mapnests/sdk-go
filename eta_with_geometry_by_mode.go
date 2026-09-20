package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

type ETAWithGeometryByModeRequest struct {
	FromLat    float64
	FromLon    float64
	ToLat      float64
	ToLon      float64
	Mode       Mode
	TripId     string
	XRequestID *string
}

type ETARouteData struct {
	Distance  float64    `json:"distance"`
	Duration  float64    `json:"duration"`
	Geometry  string     `json:"geometry"`
	Waypoints []Waypoint `json:"waypoints"`
}
type EtaWithGeometryByModeResponse struct {
	Status  bool         `json:"status"`
	Message string       `json:"message"`
	Data    ETARouteData `json:"data"`
}

func (s *client) EtaWithGeometryByMode(ctx context.Context, request ETAWithGeometryByModeRequest) (*EtaWithGeometryByModeResponse, error) {
	if err := ValidateLatLon(request.FromLat, request.FromLon); err != nil {
		return nil, err
	}

	if err := ValidateLatLon(request.ToLat, request.ToLon); err != nil {
		return nil, err
	}

	if err := ValidateEtaMode(request.Mode); err != nil {
		return nil, err
	}

	if err := ValidateTripID(request.TripId); err != nil {
		return nil, err
	}

	if err := ValidateXRequestID(request.XRequestID); err != nil {
		return nil, err
	}

	body, err := s.request("etaWithGeometryByMode", request)
	if err != nil {
		return nil, err
	}

	var response EtaWithGeometryByModeResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling EtaWithGeometryByModeResponse response: %w", err)
	}

	return &response, nil
}
