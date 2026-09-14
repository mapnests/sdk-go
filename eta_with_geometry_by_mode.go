package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

type EtaWithGeometryByModeResponse struct {
	Status  bool      `json:"status"`
	Message string    `json:"message"`
	Data    RouteData `json:"data"`
}

func (s *client) EtaWithGeometryByMode(ctx context.Context, request ETAByModeRequest) (*EtaWithGeometryByModeResponse, error) {
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
