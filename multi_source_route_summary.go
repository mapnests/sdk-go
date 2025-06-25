package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

type Source struct {
	ID   int     `json:"id"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
	Mode string  `json:"mode"`
}

type Destination struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type MultiSourceRouteSummaryRequest struct {
	Sources     []Source    `json:"sources"`
	Destination Destination `json:"destination"`
}

type RouteSummary struct {
	ID             int     `json:"id"`
	DistanceMeters float64 `json:"distanceInMeters"`
	EtaSeconds     float64 `json:"etaInSeconds"`
	Geometry       string  `json:"geometry"`
}

type DistanceMatrixResponseData struct {
	RouteSummaries []RouteSummary `json:"routeSummaries"`
}

type MultiSourceRouteSummaryResponse struct {
	Data    DistanceMatrixResponseData `json:"data"`
	Message string                     `json:"message"`
	Status  bool                       `json:"status"`
}

func (s *client) MultiSourceRouteSummary(ctx context.Context, request MultiSourceRouteSummaryRequest) (*MultiSourceRouteSummaryResponse, error) {
	fmt.Println("📍 MultiSourceRouteSummaryRequest request:", request)

	body, err := s.request("multiSourceRouteSummary", request)
	if err != nil {
		return nil, err
	}

	var response MultiSourceRouteSummaryResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling MultiSourceRouteSummaryRequest response: %w", err)
	}

	return &response, nil
}
