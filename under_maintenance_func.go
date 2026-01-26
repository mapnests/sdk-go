package sdk

var underMaintenanceMap = map[string]bool{
	"Search":                  false,
	"Reverse":                 false,
	"DistanceMatrix":          false,
	"DistanceMatrixDetails":   false,
	"PairWiseRouteSummary":    false,
	"MultiSourceRouteSummary": false,
	"Autocomplete":            false,
	"AutocompleteWithoutZone": false,
	"SearchByRadius":          false,
}

func isUnderMaintenance(fn string) bool {
	return underMaintenanceMap[fn]
}
