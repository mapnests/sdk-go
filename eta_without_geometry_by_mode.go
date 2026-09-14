package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

type ETAByModeRequest struct {
	FromLat    float64
	FromLon    float64
	ToLat      float64
	ToLon      float64
	Mode       Mode
	NoOfRoutes *int64
	XRequestID *string
}

type EtaWithoutGeometryByModeData struct {
	From     Coordinate `json:"from"`
	To       Coordinate `json:"to"`
	Distance float64    `json:"distance"`
	Eta      float64    `json:"eta"`
}

type EtaWithoutGeometryByModeResponse struct {
	Data    []EtaWithoutGeometryByModeData `json:"data"`
	Message string                         `json:"message"`
	Status  bool                           `json:"status"`
}

func (s *client) EtaWithoutGeometryByMode(ctx context.Context, request ETAByModeRequest) (*EtaWithoutGeometryByModeResponse, error) {

	if err := ValidateLatLon(request.FromLat, request.FromLon); err != nil {
		return nil, err
	}

	if err := ValidateLatLon(request.ToLat, request.ToLon); err != nil {
		return nil, err
	}

	if err := ValidateEtaMode(request.Mode); err != nil {
		return nil, err
	}

	if err := ValidateLimit(request.NoOfRoutes, ETAMaxRoutes); err != nil {
		return nil, err
	}

	if err := ValidateXRequestID(request.XRequestID); err != nil {
		return nil, err
	}

	body, err := s.request("etaWithoutGeometryByMode", request)
	if err != nil {
		return nil, err
	}

	var response EtaWithoutGeometryByModeResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling EtaWithoutGeometryByModeResponse response: %w", err)
	}

	return &response, nil
}
