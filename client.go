package sdk

import (
	"context"
	_ "embed"
)

type Client interface {
	// Geo Map APIs
	Reverse(ctx context.Context, request ReverseRequest) (*ReverseResponse, error)
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)
	Autocomplete(ctx context.Context, request AutoCompleteRequest) (*AutoCompleteResponse, error)
	SearchByRadius(ctx context.Context, request SearchByRadiusRequest) (*SearchResponse, error)
	DetailsByPlaceID(ctx context.Context, request DetailsByPlaceIDRequest) (*DetailsByPlaceIDResponse, error)
	Geocode(ctx context.Context, request GeocodeRequest) (*GeocodeResponse, error)
	
	// Route APIs
	DistanceMatrix(ctx context.Context, request DistanceMatrixRequest) (*DistanceMatrixResponse, error)
	DistanceMatrixDetails(ctx context.Context, request DistanceMatrixDetailsRequest) (*DistanceMatrixDetailsResponse, error)
	PairWiseRouteSummary(ctx context.Context, request PairWiseRouteSummaryRequest) (*PairWiseRouteSummaryResponse, error)
	MultiSourceRouteSummary(ctx context.Context, request MultiSourceRouteSummaryRequest) (*MultiSourceRouteSummaryResponse, error)
	MultiStopPoints(ctx context.Context, request MultiStopPointsRequest) (*MultiStopPointsResponse, error)
	SnapToRoad(ctx context.Context, request SnapToRoadRequest) (*SnapToRoadResponse, error)
	MultiSourceRouteSummaryWithoutGeometry(ctx context.Context, request MultiSourceRouteSummaryRequest) (*MultiSourceRouteSummaryWithoutGeometryResponse, error)
	PairwiseDistanceMatrix(ctx context.Context, request PairwiseDistanceMatrixRequest) (*PairwiseDistanceMatrixResponse, error)
	MultiStopPointDetailsWithGeometry(ctx context.Context, request MultiStopPointDetailsWithGeometryRequest) (*MultiStopPointDetailsWithGeometryResponse, error)
	
	// ETA APIs
	EtaWithoutGeometryByMode(ctx context.Context, request ETAByModeRequest) (*EtaWithoutGeometryByModeResponse, error)
	EtaWithGeometryByMode(ctx context.Context, request ETAWithGeometryByModeRequest) (*EtaWithGeometryByModeResponse, error)
	EtaWithoutStoppage(ctx context.Context, request ETAsRequest) (*ETAsResponse, error)
	EtaMultiStoppage(ctx context.Context, request ETAMultipleStoppageRequest) (*ETAMultipleStoppageResponse, error)
	EtaWithMultiSourceMultiStopPoint(ctx context.Context, request EtaWithMultiSourceMultiStopPointRequest) (*EtaWithMultiSourceMultiStopPointResponse, error)
	
	// Static Map APIs
	StaticMap(ctx context.Context, request StaticMapRequest) (*StaticMapResponse, error)
}

type client struct {
	apiKey      string
	packageName string
	timeoutMs   int32
}

func NewClient(apiKey, packageName string, timeoutMs ...int32) Client {
	var timeout int32 = 30000
	if len(timeoutMs) > 0 {
		timeout = timeoutMs[0]
	}
	return &client{
		apiKey:      apiKey,
		packageName: packageName,
		timeoutMs:   timeout,
	}
}
