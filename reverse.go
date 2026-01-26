package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

type ReverseRequest struct {
	Lat float64
	Lon float64
}

type ReverseResponse struct {
	Data    ReverseData `json:"data"`
	Message string      `json:"message"`
	Status  bool        `json:"status"`
}

type ReverseData struct {
	PlaceID        string   `json:"placeId"`
	Lat            float64  `json:"lat"`
	Lon            float64  `json:"lon"`
	Category       string   `json:"category"`
	Type           string   `json:"type"`
	Class          string   `json:"class"`
	AddressType    string   `json:"addressType"`
	Name           string   `json:"name"`
	DisplayName    string   `json:"displayName"`
	DisplayAddress string   `json:"displayAddress"`
	Address        string   `json:"address"`
	Country        string   `json:"country"`
	City           string   `json:"city"`
	Thana          string   `json:"thana"`
	District       string   `json:"district"`
	Division       string   `json:"division"`
	PostalCode     string   `json:"postalCode"`
	Website        string   `json:"website"`
	HouseNumber    string   `json:"houseNumber"`
	HouseName      string   `json:"houseName"`
	SubLocality    string   `json:"subLocality"`
	LocalArea      string   `json:"localArea"`
	Types          []string `json:"types"`
}

func (s *client) Reverse(ctx context.Context, request ReverseRequest) (*ReverseResponse, error) {

	if isUnderMaintenance("Reverse") {
		return &ReverseResponse{
			Message: "Reverse service is under maintenance",
			Status:  false,
		}, nil
	}

	err := ValidateLatLon(request.Lat, request.Lon)
	if err != nil {
		return nil, err
	}

	fmt.Println("📍 Reverse request:", request)

	body, err := s.request("reverse", request)
	if err != nil {
		return nil, err
	}

	var response ReverseResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling Reverse response: %w", err)
	}

	return &response, nil

}
