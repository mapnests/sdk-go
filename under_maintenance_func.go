package sdk

var underMaintenanceMap = map[string]bool{
	"Search":                                 false,
	"Reverse":                                false,
	"DistanceMatrix":                         false,
	"DistanceMatrixDetails":                  false,
	"PairWiseRouteSummary":                   false,
	"MultiSourceRouteSummary":                false,
	"MultiSourceRouteSummaryWithoutGeometry": false,
	"Autocomplete":                           false,
	"SearchByRadius":                         false,
	"DetailsByPlaceID":                       false,
	"SnapToRoad":                             false,
	"MultiStopPoints":                        false,
	"Geocode":                                false,
}

func isUnderMaintenance(fn string) bool {
	return underMaintenanceMap[fn]
}
