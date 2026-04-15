package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

type SnapToRoadRequest struct{
    Mode Mode
    Latitude float64
    Longitude float64
}

type SnapToRoadWaypoint struct {
    Address         string    `json:"address"`
    Location        []float64 `json:"location"`
    DistanceMeters  float64   `json:"distance_meters"`
}

type SnapToRoadData struct {
    Waypoints []SnapToRoadWaypoint `json:"waypoints"`
}

type SnapToRoadResponse struct {
    Status  bool           `json:"status"`
    Message string         `json:"message"`
    Data    SnapToRoadData `json:"data"`
}

func (s *client) SnapToRoad(ctx context.Context, request SnapToRoadRequest) (*SnapToRoadResponse, error) {
    
    if isUnderMaintenance("SnapToRoad") {
		return &SnapToRoadResponse{
			Message: "SnapToRoad service is under maintenance",
			Status:  false,
		}, nil
	}
    
    if err := ValidateLatLon(request.Latitude, request.Longitude); err != nil {
        return nil, err
    }

    fmt.Println("📍 SnapToRoad request:", request)

	body, err := s.request("snapToRoad", request)
	if err != nil {
		return nil, err
	}

    var response SnapToRoadResponse
    if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling SnapToRoad response: %w", err)
	}

    return &response, nil
}  