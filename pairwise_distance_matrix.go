package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

type PairwiseDistanceMatrixRoute struct {
	ID          int        `json:"id"`
	Origin      Coordinate `json:"origin"`
	Destination Coordinate `json:"destination"`
}

type SortOrder string

const (
	SortDistanceInMetresAsc  = SortOrder("distanceInMetres_asc")
	SortDistanceInMetresDesc = SortOrder("distanceInMetres_desc")
	SortEtaInSecondsAsc      = SortOrder("etaInSeconds_asc")
	SortEtaInSecondsDesc     = SortOrder("etaInSeconds_desc")
)

type PairwiseDistanceMatrixRequest struct {
	Mode       Mode                          `json:"mode"`
	Routes     []PairwiseDistanceMatrixRoute `json:"routes"`
	Sort       SortOrder                     `json:"sort"`
	XRequestID *string
}

type PairwiseDistanceMatrixResult struct {
	ID               int     `json:"id"`
	DistanceInMetres float64 `json:"distanceInMetres"`
	EtaInSeconds     float64 `json:"etaInSeconds"`
}

type PairwiseDistanceMatrixResponse struct {
	Status  bool                           `json:"status"`
	Message string                         `json:"message"`
	Data    []PairwiseDistanceMatrixResult `json:"data"`
}

func (s *client) PairwiseDistanceMatrix(ctx context.Context, request PairwiseDistanceMatrixRequest) (*PairwiseDistanceMatrixResponse, error) {
	if err := ValidateXRequestID(request.XRequestID); err != nil {
		return nil, err
	}

	for _, route := range request.Routes {
		if err := ValidateLatLon(route.Origin.Lat, route.Origin.Lon); err != nil {
			return nil, err
		}

		if err := ValidateLatLon(route.Destination.Lat, route.Destination.Lon); err != nil {
			return nil, err
		}
	}

	body, err := s.request("pairwiseDistanceMatrix", request)
	if err != nil {
		return nil, err
	}

	var response PairwiseDistanceMatrixResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling PairwiseDistanceMatrixResponse response: %w", err)
	}

	return &response, nil
}
