package sdk

import "net/http"

// HTTPMethodMap maps API labels to their HTTP method.
var HTTPMethodMap = map[string]string{
	"reverse":                                http.MethodGet,
	"search":                                 http.MethodGet,
	"distanceMatrix":                         http.MethodGet,
	"distanceMatrixDetails":                  http.MethodGet,
	"pairWiseRouteSummary":                   http.MethodPost,
	"multiSourceRouteSummary":                http.MethodPost,
	"autocomplete":                           http.MethodGet,
	"searchByRadius":                         http.MethodGet,
	"detailsByPlaceId":                       http.MethodGet,
	"snapToRoad":                             http.MethodPost,
	"multiStopPoints":                        http.MethodPost,
	"multiSourceRouteSummaryWithoutGeometry": http.MethodPost,
	"geocode":                                http.MethodGet,
	"etaWithoutGeometryByMode":               http.MethodGet,
	"etaWithGeometryByMode":                  http.MethodGet,
	"etaWithGeometry":                        http.MethodGet,
	"etaMultipleStoppage":                    http.MethodPost,
}
