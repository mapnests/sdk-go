package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

type DistanceMatrixDetailsRequest struct {
	OriginLat float64
	OriginLon float64
	DestLat   float64
	DestLon   float64
	Mode      Mode
}

type DistanceMatrixDetailsResponse struct {
	Status  bool      `json:"status"`
	Message string    `json:"message"`
	Data    RouteData `json:"data"`
}

type RouteData struct {
	RouteResponse RouteResponse `json:"routeResponse"`
}

type RouteResponse struct {
	Code        string     `json:"code"`
	Message     string     `json:"message"`
	DataVersion string     `json:"data_version"`
	Routes      []Route    `json:"routes"`
	Waypoints   []Waypoint `json:"waypoints"`
}

type Route struct {
	Distance   float64 `json:"distance"`
	Duration   float64 `json:"duration"`
	WeightName string  `json:"weight_name"`
	Weight     float64 `json:"weight"`
	Geometry   string  `json:"geometry"`
	Legs       []Leg   `json:"legs"`
}

type Leg struct {
	Distance float64 `json:"distance"`
	Duration float64 `json:"duration"`
	Summary  string  `json:"summary"`
	Weight   float64 `json:"weight"`
	Steps    []Step  `json:"steps"`
}

type Step struct {
	Distance      float64        `json:"distance"`
	Duration      float64        `json:"duration"`
	Geometry      string         `json:"geometry"`
	Name          string         `json:"name"`
	Mode          string         `json:"mode"`
	DrivingSide   string         `json:"driving_side"`
	Weight        float64        `json:"weight"`
	Maneuver      Maneuver       `json:"maneuver"`
	Intersections []Intersection `json:"intersections"`
}

type Maneuver struct {
	Location      []float64 `json:"location"`
	BearingBefore int       `json:"bearing_before"`
	BearingAfter  int       `json:"bearing_after"`
	Type          string    `json:"type"`
	Modifier      string    `json:"modifier"`
}

type Intersection struct {
	Location []float64 `json:"location"`
	Bearings []int     `json:"bearings"`
	Entry    []bool    `json:"entry"`
	In       int       `json:"in"`
	Out      int       `json:"out"`
}

type Waypoint struct {
	Name     string    `json:"name"`
	Location []float64 `json:"location"`
	Distance float64   `json:"distance"`
	Hint     string    `json:"hint"`
}

func (s *client) DistanceMatrixDetails(ctx context.Context, request DistanceMatrixDetailsRequest) (*DistanceMatrixDetailsResponse, error) {
	
	err := ValidateLatLon(request.OriginLat,request.OriginLon)
    if err != nil {
        return nil,err
    }
	err = ValidateLatLon(request.DestLat,request.DestLon)
    if err != nil {
        return nil,err
    }
	
	fmt.Println("📍 DistanceMatrixDetails request:", request)

	body, err := s.request("distanceMatrixDetails", request)
	if err != nil {
		return nil, err
	}

	var response DistanceMatrixDetailsResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling DistanceMatrixDetails response: %w", err)
	}

	return &response, nil
}
