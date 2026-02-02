package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

func (s *client) AutocompleteWithoutZone(ctx context.Context, request AutoCompleteRequest) (*AutoCompleteResponse, error) {

	err := ValidateLatLonPtr(request.Lat,request.Lon)
    if err != nil {
        return nil,err
    }
	
	normalizedQuery, err := ValidateAndNormalizeQuery(request.Query)
	if err != nil {
		return nil, err
	}
	request.Query = normalizedQuery
	
	fmt.Println("📍 Autocomplete Without Zone request:", request)
	body, err := s.request("autocompleteWithoutZone", request)
	if err != nil {
		return nil, err
	}

	var response AutoCompleteResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling Autocomplete Without Zone response: %w", err)
	}

	return &response, nil
}