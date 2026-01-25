package sdk

import (
	"context"
	_ "embed"
)

type Client interface {
	DistanceMatrix(ctx context.Context, request DistanceMatrixRequest) (*DistanceMatrixResponse, error)
	DistanceMatrixDetails(ctx context.Context, request DistanceMatrixDetailsRequest) (*DistanceMatrixDetailsResponse, error)
	Reverse(ctx context.Context, request ReverseRequest) (*ReverseResponse, error)
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)
	PairWiseRouteSummary(ctx context.Context, request PairWiseRouteSummaryRequest) (*PairWiseRouteSummaryResponse, error)
	MultiSourceRouteSummary(ctx context.Context, request MultiSourceRouteSummaryRequest) (*MultiSourceRouteSummaryResponse, error)
	Autocomplete(ctx context.Context, request AutoCompleteRequest) (*AutoCompleteResponse, error)
	AutocompleteWithoutZone(ctx context.Context, request AutoCompleteRequest) (*AutoCompleteResponse, error)
	SearchByRadius(ctx context.Context, request SearchByRadiusRequest) (*SearchResponse, error)
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
