package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

type ETAsRequest struct {
	FromLat    float64
	FromLon    float64
	ToLat      float64
	ToLon      float64
	XRequestID *string
}

type ETAModeData struct {
	Distance  float64    `json:"distance"`
	Duration  float64    `json:"duration"`
	Geometry  string     `json:"geometry"`
	Waypoints []Waypoint `json:"waypoints"`
}

type ETAsModesData struct {
	Car        ETAModeData `json:"car"`
	CNG        ETAModeData `json:"cng"`
	Motorcycle ETAModeData `json:"motorcycle"`
}

type ETAsResponse struct {
	Status  bool          `json:"status"`
	Message string        `json:"message"`
	Data    ETAsModesData `json:"data"`
}

func (s *client) EtaWithGeometry(ctx context.Context, request ETAsRequest) (*ETAsResponse, error) {
	if err := ValidateLatLon(request.FromLat, request.FromLon); err != nil {
		return nil, err
	}

	if err := ValidateLatLon(request.ToLat, request.ToLon); err != nil {
		return nil, err
	}

	if err := ValidateXRequestID(request.XRequestID); err != nil {
		return nil, err
	}

	body, err := s.request("etaWithGeometry", request)
	if err != nil {
		return nil, err
	}

	var response ETAsResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling ETAsResponse response: %w", err)
	}

	return &response, nil
}
