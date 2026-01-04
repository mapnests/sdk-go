package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

type AutoCompleteRequest struct{
	Query string   
	Lat   *float64  
	Lon   *float64   
	Limit *int64
}

type AutoCompleteResponse struct{
	Status 	bool 			`json:"status"`
	Message string			`json:"message"`
	Data 	[]ResponseData 	`json:"data"`
}

type ResponseData struct{
	PlaceID string 		`json:"place_id"`
	Address string 		`json:"address"`
	Types 	[]string 	`json:"types"`
}

func (s *client) Autocomplete(ctx context.Context, request AutoCompleteRequest) (*AutoCompleteResponse, error) {
	err := ValidateLatLonPtr(request.Lat,request.Lon)
    if err != nil {
        return nil,err
    }
	
	fmt.Println("📍 Autocomplete request:", request)

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