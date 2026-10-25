package sdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNoRiders           = errors.New("at least one rider is required")
	ErrEmptyRiderID       = errors.New("riderId is required")
	ErrDuplicateRiderID   = errors.New("riderId must be unique")
	ErrSameRiderLocations = errors.New("rider currentLocation and dropOffPoint must differ")
)

type EtaWithMultiSourceMultiStopPointRider struct {
	RiderID         string
	CurrentLocation Coordinate
	DropOffPoint    Coordinate
}

type EtaWithMultiSourceMultiStopPointRequest struct {
	PickupPoint Coordinate
	Riders      []EtaWithMultiSourceMultiStopPointRider
	Mode        Mode
	Traceparent *string
}

type EtaWithMultiSourceMultiStopPointRiderRoute struct {
	RiderID         string  `json:"id"`
	DistanceMeters  float64 `json:"distance"`
	DurationSeconds float64 `json:"duration"`
	Geometry        string  `json:"geometry"`
}

type EtaWithMultiSourceMultiStopPointResponse struct {
	Status  bool                                         `json:"status"`
	Message string                                       `json:"message"`
	Data    []EtaWithMultiSourceMultiStopPointRiderRoute `json:"data"`
}

// Wire format accepted by the service.
type etaWithMultiSourceMultiStopPointPair struct {
	ID   string     `json:"id"`
	From Coordinate `json:"from"`
	To   Coordinate `json:"to"`
}

type etaWithMultiSourceMultiStopPointBody struct {
	Dest        Coordinate                             `json:"dest"`
	Mode        Mode                                   `json:"mode"`
	Pairs       []etaWithMultiSourceMultiStopPointPair `json:"pairs"`
	Traceparent *string
}

func (s *client) EtaWithMultiSourceMultiStopPoint(ctx context.Context, request EtaWithMultiSourceMultiStopPointRequest) (*EtaWithMultiSourceMultiStopPointResponse, error) {
	if err := ValidateLatLon(request.PickupPoint.Lat, request.PickupPoint.Lon); err != nil {
		return nil, fmt.Errorf("pickupPoint: %w", err)
	}
	if len(request.Riders) == 0 {
		return nil, ErrNoRiders
	}
	if err := ValidateEtaMode(request.Mode); err != nil {
		return nil, err
	}
	if err := ValidateTraceparent(request.Traceparent); err != nil {
		return nil, err
	}

	pairs := make([]etaWithMultiSourceMultiStopPointPair, 0, len(request.Riders))
	seen := make(map[string]struct{}, len(request.Riders))
	for i, rider := range request.Riders {
		id := strings.TrimSpace(rider.RiderID)
		if id == "" {
			return nil, fmt.Errorf("riders[%d]: %w", i, ErrEmptyRiderID)
		}
		if _, dup := seen[id]; dup {
			return nil, fmt.Errorf("riders[%d]: %w: %q", i, ErrDuplicateRiderID, id)
		}
		seen[id] = struct{}{}

		if err := ValidateLatLon(rider.CurrentLocation.Lat, rider.CurrentLocation.Lon); err != nil {
			return nil, fmt.Errorf("riders[%d] (%s) currentLocation: %w", i, id, err)
		}
		if err := ValidateLatLon(rider.DropOffPoint.Lat, rider.DropOffPoint.Lon); err != nil {
			return nil, fmt.Errorf("riders[%d] (%s) dropOffPoint: %w", i, id, err)
		}
		if rider.CurrentLocation == rider.DropOffPoint {
			return nil, fmt.Errorf("riders[%d] (%s): %w", i, id, ErrSameRiderLocations)
		}

		pairs = append(pairs, etaWithMultiSourceMultiStopPointPair{
			ID:   id,
			From: rider.CurrentLocation,
			To:   rider.DropOffPoint,
		})
	}

	body, err := s.request("etaWithMultiSourceMultiStopPoint", etaWithMultiSourceMultiStopPointBody{
		Dest:        request.PickupPoint,
		Mode:        request.Mode,
		Pairs:       pairs,
		Traceparent: request.Traceparent,
	})
	if err != nil {
		return nil, err
	}

	var response EtaWithMultiSourceMultiStopPointResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling EtaWithMultiSourceMultiStopPointResponse response: %w", err)
	}

	return &response, nil
}
