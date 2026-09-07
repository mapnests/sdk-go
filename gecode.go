package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

type GeocodeRequest struct {
	Query            string
	Limit            *int64
	AcceptedLanguage *string
	XRequestID       *string
}

type GeocodeResponse struct {
	Status  bool          `json:"status"`
	Message string        `json:"message"`
	Data    []GeocodeData `json:"data"`
}

type GeocodeData struct {
	PlaceID     int64   `json:"place_id"`
	Lat         string  `json:"lat"`
	Lon         string  `json:"lon"`
	Category    string  `json:"category"`
	Type        string  `json:"type"`
	PlaceRank   int     `json:"place_rank"`
	Importance  float64 `json:"importance"`
	AddressType string  `json:"addresstype"`
	Name        string  `json:"name"`
	DisplayName string  `json:"display_name"`
	Address     string  `json:"address"`
}

func (s *client) Geocode(ctx context.Context, request GeocodeRequest) (*GeocodeResponse, error) {

	if isUnderMaintenance("Geocode") {
		return &GeocodeResponse{
			Message: "Geocode service is under maintenance",
			Status:  false,
		}, nil
	}

	if err := ValidateXRequestID(request.XRequestID); err != nil {
		return nil, err
	}

	if err := ValidateAcceptedLanguage(request.AcceptedLanguage); err != nil {
		return nil, err
	}

	if err := ValidateLimit(request.Limit, GeocodeMaxLimit); err != nil {
		return nil, err
	}

	normalizedQuery, err := ValidateAndNormalizeQuery(request.Query)
	if err != nil {
		return nil, err
	}
	request.Query = normalizedQuery

	body, err := s.request("geocode", request)
	if err != nil {
		return nil, err
	}

	var response GeocodeResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling Geocode response: %w", err)
	}

	return &response, nil
}
