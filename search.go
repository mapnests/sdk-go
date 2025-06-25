package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

type SearchRequest struct {
	Query string
}

type SearchResponse struct {
	Data    []SearchData `json:"data"`
	Message string       `json:"message"`
	Status  bool         `json:"status"`
}

type SearchData struct {
	PlaceID        string   `json:"place_id"`
	Lat            string   `json:"lat"`
	Lon            string   `json:"lon"`
	AddressTypes   []string `json:"addresstypes"`
	DisplayName    string   `json:"display_name"`
	DisplayAddress string   `json:"display_address"`
}

func (s *client) Search(ctx context.Context, request SearchRequest) (*SearchResponse, error) {
	body, err := s.request("search", request)
	if err != nil {
		return nil, err
	}

	var response SearchResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling Search response: %w", err)
	}

	return &response, nil
}
