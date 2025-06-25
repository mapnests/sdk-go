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
}

type client struct {
	apiKey      string
	packageName string
}

func NewClient(apiKey, packageName string) Client {
	return &client{
		apiKey:      apiKey,
		packageName: packageName,
	}
}
