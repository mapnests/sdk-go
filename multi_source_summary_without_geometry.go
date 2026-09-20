package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

// type MultiSourceRouteSummaryWithoutGeometryRequest struct {
// 	Sources     []Source
// 	Destination Destination
// 	XRequestID  *string
// }

type RouteSummaryWithoutGeometry struct {
	ID             int     `json:"id"`
	DistanceMeters float64 `json:"distanceInMeters"`
	EtaSeconds     float64 `json:"etaInSeconds"`
	Reachable      bool    `json:"reachable"`
}

type DistanceMatrixResponseDataWithoutGeometry struct {
	RouteSummaries []RouteSummaryWithoutGeometry `json:"routeSummaries"`
}

type MultiSourceRouteSummaryWithoutGeometryResponse struct {
	Data    DistanceMatrixResponseDataWithoutGeometry `json:"data"`
	Message string                                    `json:"message"`
	Status  bool                                      `json:"status"`
}

func (s *client) MultiSourceRouteSummaryWithoutGeometry(ctx context.Context, request MultiSourceRouteSummaryRequest) (*MultiSourceRouteSummaryWithoutGeometryResponse, error) {

	if err := ValidateXRequestID(request.XRequestID); err != nil {
		return nil, err
	}

	for _, src := range request.Sources {
		if err := ValidateLatLon(src.Lat, src.Lon); err != nil {
			return nil, err
		}
	}

	if err := ValidateLatLon(request.Destination.Lat, request.Destination.Lon); err != nil {
		return nil, err
	}

	body, err := s.request("multiSourceRouteSummaryWithoutGeometry", request)
	if err != nil {
		return nil, err
	}

	var response MultiSourceRouteSummaryWithoutGeometryResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling MultiSourceRouteSummaryRequest response: %w", err)
	}

	return &response, nil
}
