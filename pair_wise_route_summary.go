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

type PairWiseRouteSummaryResponse struct {
	Status  bool                     `json:"status"`
	Message string                   `json:"message"`
	Data    RouteSummaryResponseData `json:"data"`
}

type RouteSummaryResponseData struct {
	RouteSummaries []RouteSummary `json:"routeSummaries"`
}

func (s *client) PairWiseRouteSummary(ctx context.Context, request PairWiseRouteSummaryRequest) (*PairWiseRouteSummaryResponse, error) {
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
