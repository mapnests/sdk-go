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

type ETAsModesData struct {
	Car        EtaWithGeometryByModeResponse `json:"car"`
	CNG        EtaWithGeometryByModeResponse `json:"cng"`
	Motorcycle EtaWithGeometryByModeResponse `json:"motorcycle"`
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
