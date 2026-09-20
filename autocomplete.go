package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

type AutoCompleteRequest struct {
	Query            string
	ActiveZone       *bool
	Lat              *float64
	Lon              *float64
	Limit            *int64
	Radius           *int64
	AcceptedLanguage *string
	XRequestID       *string
}

type AutoCompleteResponse struct {
	Status  bool           `json:"status"`
	Message string         `json:"message"`
	Data    []ResponseData `json:"data"`
}

type ResponseData struct {
	PlaceID string   `json:"placeId"`
	Address string   `json:"address"`
	Name    string   `json:"name"`
	Types   []string `json:"types"`
}

func (s *client) Autocomplete(ctx context.Context, request AutoCompleteRequest) (*AutoCompleteResponse, error) {
	err := ValidateLatLonPtr(request.Lat, request.Lon)
	if err != nil {
		return nil, err
	}

	if err := ValidateXRequestID(request.XRequestID); err != nil {
		return nil, err
	}

	if err := ValidateAcceptedLanguage(request.AcceptedLanguage); err != nil {
		return nil, err
	}

	if err := ValidateLimit(request.Limit, AutocompleteMaxLimit); err != nil {
		return nil, err
	}

	if err := ValidateRadiusPtr(request.Radius, AutocompleteMaxRadius); err != nil {
		return nil, err
	}

	if err := ValidateRadiusRequiresLatLon(request.Radius, request.Lat, request.Lon); err != nil {
		return nil, err
	}

	normalizedQuery, err := ValidateAndNormalizeQuery(request.Query)
	if err != nil {
		return nil, err
	}
	request.Query = normalizedQuery

	body, err := s.request("autocomplete", request)
	if err != nil {
		return nil, err
	}

	var response AutoCompleteResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling AutoComplete response: %w", err)
	}

	return &response, nil
}
