package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

type DistanceMatrixRequest struct {
	OriginLat float64
	OriginLon float64
	DestLat   float64
	DestLon   float64
	Mode      Mode
}

type DistanceMatrixData struct {
	DistanceInMetres float64 `json:"distanceInMetres"`
	EtaInSeconds     float64 `json:"etaInSeconds"`
}

type DistanceMatrixResponse struct {
	Data DistanceMatrixData `json:"data"`
}

func (s *client) DistanceMatrix(ctx context.Context, request DistanceMatrixRequest) (*DistanceMatrixResponse, error) {
	fmt.Println("📍 DistanceMatrix request:", request)

	body, err := s.request("distanceMatrix", request)
	if err != nil {
		return nil, err
	}

	var response DistanceMatrixResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling DistanceMatrixResponse: %w", err)
	}

	fmt.Println("📍 DistanceMatrix response:", response)
	return &response, nil
}
