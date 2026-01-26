package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

type SearchRequest struct {
	Query           string
	Lat             *float64
	Lon             *float64
	Radius          *int64
	Page            *int64
	Limit           *int64
	ActiveLocations bool
}

type SearchResponse struct {
	Data    SearchResponseData `json:"data"`
	Message string             `json:"message"`
	Status  bool               `json:"status"`
}

type SearchResponseData struct {
	Items        []SearchData `json:"items"`
	ItemsPerPage int          `json:"itemsPerPage"`
	PageNumber   int          `json:"pageNumber"`
	TotalItems   int          `json:"totalItems"`
	TotalPages   int          `json:"totalPages"`
}

type SearchData struct {
	PlaceID     string   `json:"placeId"`
	Lat         float64  `json:"lat"`
	Lon         float64  `json:"lon"`
	Types       []string `json:"types"`
	Address     string   `json:"address"`
	Name        string   `json:"name"`
	HouseNumber string   `json:"houseNumber"`
	HouseName   string   `json:"houseName"`
	Street      string   `json:"street"`
	Phone       string   `json:"phone"`
	Website     string   `json:"website"`
	Country     string   `json:"country"`
	City        string   `json:"city"`
	Thana       string   `json:"thana"`
	Division    string   `json:"division"`
	District    string   `json:"district"`
	PostalCode  string   `json:"postalCode"`
	PlusCode    string   `json:"plusCode"`
	Sublocality string   `json:"sublocality"`
	LocalArea   string   `json:"localArea"`
}

func (s *client) Search(ctx context.Context, request SearchRequest) (*SearchResponse, error) {

	if isUnderMaintenance("Search") {
		return &SearchResponse{
			Message: "Search service is under maintenance",
			Status:  false,
		}, nil
	}

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
