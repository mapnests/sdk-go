package sdk

var underMaintenanceMap = map[string]bool{
    "Search":					true,
    "Reverse": 					true,
	"DistanceMatrix": 			false,
	"DistanceMatrixDetails": 	false,
	"PairWiseRouteSummary": 	false,
	"MultiSourceRouteSummary": 	false,
	"Autocomplete": 			false,
	"AutocompleteWithoutZone" :	false,
}

func isUnderMaintenance(fn string) bool {
    return underMaintenanceMap[fn]
}