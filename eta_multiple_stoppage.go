package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

type LatLng struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ETAStoppage struct {
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	StoppageOrder int     `json:"stoppage_order"`
}

type ETAMultipleStoppageRequest struct {
	Source     LatLng        `json:"source"`
	DropOff    LatLng        `json:"dropoff"`
	Stoppages  []ETAStoppage `json:"stoppages"`
	Modes      []Mode        `json:"modes"`
	TripId     string        `json:"trip_id"`
	XRequestID *string       `json:"XRequestID"`
}

type HopsData struct {
	Order    int     `json:"order"`
	From     string  `json:"from"`
	To       string  `json:"to"`
	Distance float64 `json:"distance"`
	Duration float64 `json:"duration"`
	Geometry string  `json:"geometry"`
}

type ModeETAData struct {
	Hops          []*HopsData `json:"hops"`
	TotalDistance float64     `json:"total_distance"`
	TotalDuration float64     `json:"total_duration"`
}

type ETAMultipleStoppageData map[Mode]*ModeETAData

type ETAMultipleStoppageResponse struct {
	Status  bool                    `json:"status"`
	Message string                  `json:"message"`
	TripId  string                  `json:"trip_id"`
	Data    ETAMultipleStoppageData `json:"data"`
}

func (s *client) EtaMultiStoppage(ctx context.Context, request ETAMultipleStoppageRequest) (*ETAMultipleStoppageResponse, error) {
	if err := ValidateLatLon(request.Source.Latitude, request.Source.Longitude); err != nil {
		return nil, err
	}

	if err := ValidateLatLon(request.DropOff.Latitude, request.DropOff.Longitude); err != nil {
		return nil, err
	}

	for _, stoppage := range request.Stoppages {
		if err := ValidateLatLon(stoppage.Latitude, stoppage.Longitude); err != nil {
			return nil, err
		}
	}

	for _, mode := range request.Modes {
		if err := ValidateEtaMode(mode); err != nil {
			return nil, err
		}
	}

	if err := ValidateTripID(request.TripId); err != nil {
		return nil, err
	}

	if err := ValidateXRequestID(request.XRequestID); err != nil {
		return nil, err
	}

	body, err := s.request("etaMultipleStoppage", request)
	if err != nil {
		return nil, err
	}

	var response ETAMultipleStoppageResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling ETAMultipleStoppageResponse response: %w", err)
	}

	return &response, nil
}
