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
	PlaceRank   int     `json:"placeRank"`
	Importance  float64 `json:"importance"`
	Name        string  `json:"name"`
	AddressType string  `json:"addresstype"`
	Lon         string  `json:"lon"`
	Lat         string  `json:"lat"`
	Category    string  `json:"category"`
	Type        string  `json:"type"`
	DisplayName string  `json:"displayName"`
	PlaceID     int     `json:"placeid"`
	Address     Address `json:"address"`
}

type Address struct {
	Country       string `json:"country"`
	CountryCode   string `json:"countryCode"`
	City          string `json:"city"`
	Road          string `json:"road"`
	StateDistrict string `json:"stateDistrict"`
	ISO31662Lvl4  string `json:"iso31662Lvl4"`
	ISO31662Lvl5  string `json:"iso31662Lvl5"`
	Municipality  string `json:"municipality"`
	Postcode      string `json:"postcode"`
	Suburb        string `json:"suburb"`
	Borough       string `json:"borough"`
	State         string `json:"state"`
}

func (s *client) Reverse(ctx context.Context, request ReverseRequest) (*ReverseResponse, error) {
	fmt.Println("📍 Reverse request:", request)

	body, err := s.request("reverseGeocode", request)
	if err != nil {
		return nil, err
	}

	var response ReverseResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling Reverse response: %w", err)
	}

	return &response, nil
}
