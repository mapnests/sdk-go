package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

type SearchByRadiusRequest struct {
	Query           string
	Lat             float64
	Lon             float64
	Radius          int64
	ActiveLocations bool
	Page            *int64
	Limit           *int64
}

func (s *client) SearchByRadius(ctx context.Context, request SearchByRadiusRequest) (*SearchResponse, error) {

	if isUnderMaintenance("SearchByRadius") {
		return &SearchResponse{
			Message: "SearchByRadius service is under maintenance",
			Status:  false,
		}, nil
	}
	
	normalizedQuery, err := ValidateAndNormalizeQuery(request.Query)
	if err != nil {
		
		return nil, fmt.Errorf("Error normalizing query: %v", err)
	}
	request.Query = normalizedQuery

	body, err := s.request("searchByRadius", request)
	if err != nil {
		return nil, err
	}

	var response SearchResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling Search response: %w", err)
	}

	return &response, nil
}
