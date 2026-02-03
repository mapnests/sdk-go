package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

type DetailsByPlaceIDRequest struct {
	PlaceID string
}

type PlaceDetail struct {
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

type DetailsByPlaceIDResponse struct {
	Data    PlaceDetail			 `json:"data"`
	Message string               `json:"message"`
	Status  bool                 `json:"status"`
}

func (s *client) DetailsByPlaceID(ctx context.Context, request DetailsByPlaceIDRequest) (*DetailsByPlaceIDResponse, error) {
	if isUnderMaintenance("DetailsByPlaceID") {
		return &DetailsByPlaceIDResponse{
			Message: "DetailsByPlaceID service is under maintenance",
			Status:  false,
		}, nil
	}

	fmt.Println("📍 DetailsByPlaceID request:", request)
	
	body, err := s.request("detailsByPlaceId", request)

	if err != nil {
		return nil, err
	}

	var response DetailsByPlaceIDResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling DetailsByPlaceID response: %w", err)
	}

	return &response, nil
}