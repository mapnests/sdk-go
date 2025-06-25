# 🗺️ Mapnests Go SDK

A lightweight, efficient Go SDK for interacting with the **Mapnests API**, offering search (geocoding), reverse (reverse geocoding) and distance matrix services.

---

## 📦 Installation

To add the SDK to your Go project:

```bash
go get github.com/mapnests/sdk-go
```

---

## ⚙️ Usage Example

```go
package main

import (
	"context"
	"fmt"
	"log"
	mapnests "github.com/mapnests/sdk-go"
)

func main() {
	mapClient := mapnests.NewClient("YOUR_API_KEY", "com.your.app")

	res, err := mapClient.Search(context.Background(), mapnests.SearchRequest{
		Query:    "Dhaka"
	})
	if err != nil {
		log.Fatal("Search failed:", err)
	}

	fmt.Println("Search Result:", res)
}
```

---

## 🧾 Request & Response Models

### 🔹 Distance Matrix

```go
type DistanceMatrixRequest struct {
	OriginLat float64
	OriginLon float64
	DestLat   float64
	DestLon   float64
	Mode      Mode
}

type DistanceMatrixData struct {
	DistanceInMetres float64 `json:"distanceInMetres"`
	EtaInSeconds     float64 `json:"etaInSeconds"`
}

type DistanceMatrixResponse struct {
	Data DistanceMatrixData `json:"data"`
}
```

---

### 🔹 Search (Geocoding)

```go
type SearchRequest struct {
	Query string
}

type SearchData struct {
	PlaceID        string   `json:"place_id"`
	Lat            string   `json:"lat"`
	Lon            string   `json:"lon"`
	AddressTypes   []string `json:"addresstypes"`
	DisplayName    string   `json:"display_name"`
	DisplayAddress string   `json:"display_address"`
}

type SearchResponse struct {
	Data    []SearchData `json:"data"`
	Message string       `json:"message"`
	Status  bool         `json:"status"`
}
```

---

### 🔹 Reverse (Reverse Geocoding)

```go
type ReverseRequest struct {
	Lat float64
	Lon float64
}

type ReverseData struct {
	PlaceRank   int     `json:"placeRank"`
	Importance  float64 `json:"importance"`
	Name        string  `json:"name"`
	AddressType string  `json:"addresstype"`
	Lon         string  `json:"lon"`
	Lat         string  `json:"lat"`
	Category    string  `json:"category"`
	Type        string  `json:"type"`
	DisplayName string  `json:"displayName"`
	PlaceID     int     `json:"placeid"`
	Address     Address `json:"address"`
}

type ReverseResponse struct {
	Data    ReverseData `json:"data"`
	Message string      `json:"message"`
	Status  bool        `json:"status"`
}
```

---

### 🔹 Distance Matrix with Route Details

```go
type DistanceMatrixDetailsRequest struct {
	OriginLat float64
	OriginLon float64
	DestLat   float64
	DestLon   float64
	Mode      Mode
}

type DistanceMatrixDetailsResponse struct {
	Status  bool      `json:"status"`
	Message string    `json:"message"`
	Data    RouteData `json:"data"`
}

type RouteData struct {
	RouteResponse RouteResponse `json:"routeResponse"`
}

type RouteResponse struct {
	Code        string     `json:"code"`
	Message     string     `json:"message"`
	DataVersion string     `json:"data_version"`
	Routes      []Route    `json:"routes"`
	Waypoints   []Waypoint `json:"waypoints"`
}

type Route struct {
	Distance   float64 `json:"distance"`
	Duration   float64 `json:"duration"`
	WeightName string  `json:"weight_name"`
	Weight     float64 `json:"weight"`
	Geometry   string  `json:"geometry"`
	Legs       []Leg   `json:"legs"`
}

type Leg struct {
	Distance float64 `json:"distance"`
	Duration float64 `json:"duration"`
	Summary  string  `json:"summary"`
	Weight   float64 `json:"weight"`
	Steps    []Step  `json:"steps"`
}

type Step struct {
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

type Maneuver struct {
	Location      []float64 `json:"location"`
	BearingBefore int       `json:"bearing_before"`
	BearingAfter  int       `json:"bearing_after"`
	Type          string    `json:"type"`
	Modifier      string    `json:"modifier"`
}

type Intersection struct {
	Location []float64 `json:"location"`
	Bearings []int     `json:"bearings"`
	Entry    []bool    `json:"entry"`
	In       int       `json:"in"`
	Out      int       `json:"out"`
}

type Waypoint struct {
	Name     string    `json:"name"`
	Location []float64 `json:"location"`
	Distance float64   `json:"distance"`
	Hint     string    `json:"hint"`
}
```