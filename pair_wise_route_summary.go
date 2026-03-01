package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

type Coordinate struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type PairWiseRoute struct {
	ID   int        `json:"id"`
	Src  Coordinate `json:"src"`
	Dest Coordinate `json:"dest"`
	Mode Mode       `json:"mode"`
}

type PairWiseRouteSummaryRequest struct {
	Pairs []PairWiseRoute `json:"pairs"`
}

type RouteStep struct {
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

type RouteLeg struct {
	Annotation map[string]interface{} `json:"annotation"`
	Distance   float64                `json:"distance"`
	Duration   float64                `json:"duration"`
	Summary    string                 `json:"summary"`
	Weight     float64                `json:"weight"`
	Steps      []RouteStep            `json:"steps"`
}

type PairWiseRouteSummary struct {
	ID             int     		`json:"id"`
	DistanceMeters float64 		`json:"distanceInMeters"`
	EtaSeconds     float64 		`json:"etaInSeconds"`
	Geometry       string  		`json:"geometry"`
	Legs           []RouteLeg 	`json:"legs"`
}


type RouteSummaryResponseData struct {
	RouteSummaries []PairWiseRouteSummary `json:"routeSummaries"`
	
}

type PairWiseRouteSummaryResponse struct {
	Status  bool                     `json:"status"`
	Message string                   `json:"message"`
	Data    RouteSummaryResponseData `json:"data"`
}

func (s *client) PairWiseRouteSummary(ctx context.Context, request PairWiseRouteSummaryRequest) (*PairWiseRouteSummaryResponse, error) {
	for _, pair := range request.Pairs {
		if err := ValidateLatLon(pair.Src.Lat, pair.Src.Lon); err != nil {
			return nil, err
		}
		if err := ValidateLatLon(pair.Dest.Lat, pair.Dest.Lon); err != nil {
			return nil, err
		}
	}

	fmt.Println("📍 PairWiseRouteSummary request:", request)

	body, err := s.request("pairWiseRouteSummary", request)
	if err != nil {
		return nil, err
	}

	var response PairWiseRouteSummaryResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling PairWiseRouteSummary response: %w", err)
	}

	return &response, nil
}
