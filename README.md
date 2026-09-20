<a name="readme-top"></a>

##### Readme Top


<br />
<div align="center">
  <a>
	<img src="https://cdn.pixabay.com/photo/2015/06/23/11/13/globe-818583_1280.png" alt="Logo" width="100">
  </a>

<h1 align="center">Mapnests Go SDK</h1>

  <p align="center">
	
Map Nests
  </p>
</div>




A secure and efficient Go SDK for the **Mapnests Platform**, enabling powerful geospatial capabilities such as **Search (Geocoding)**, **Reverse (Reverse Geocoding)**, **Distance Matrix**, **Autocomplete** and **Autocomplet Without Zone**.

---

## 📚 Table of Contents

* [Installation](#installation)
* [Quick Start](#quick-start)
* [Core Features](#core-features)

  * [Distance Matrix](#distance-matrix)
  * [Distance Matrix Details](#distance-matrix-details)
  * [Pairwise Route Summary](#pairwise-route-summary)
  * [Multi Source Route Summary](#multi-source-route-summary)
  * [Search](#search)
  * [Reverse](#reverse)
  * [Autocomplete](#autocomplete)
  * [Search By Radius](#search-by-radius)
  * [Detailed Search By PlaceId ](#detailed-search-by-placeId)
  * [Snap to Road](#snap-to-road)
  * [Multi Stop Point](#multi-stop-point)
  * [Multi Source Route Summary Without Geometry](#multi-source-route-summary-without-geometry)
  * [Geocode](#geocode)
  * [ETA With Geometry](#eta-with-geometry)
  * [ETA With Geometry By Mode](#eta-with-geometry-by-mode)
  * [ETA Without Geometry By Mode](#eta-without-geometry-by-mode)
  * [ETA Multiple Stoppage](#eta-multiple-stoppage)
* [License](#license)
* [Contact](#contact)

---

## Installation

```bash
go get github.com/mapnests/sdk-go
```-   [Search By Radius](#search-by-radius)

Import into your project:

```go
import mapnests "github.com/mapnests/sdk-go"
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"
	mapnests "github.com/mapnests/sdk-go"
)

func main() {
  //Without custom timeout
	client := mapnests.NewClient("YOUR_API_KEY", "com.your.app")

  //With Custom timeout in millisecond 
  clientWithTimeoutMs := mapnests.NewClient("YOUR_API_KEY", "com.your.app", "timeout in millisecond")

	res, err := client.Search(context.Background(), mapnests.SearchRequest{
		Query: "Dhaka",
	})
	if err != nil {
		log.Fatal("Search failed:", err)
	}

	fmt.Println("Search result:", res)
}
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## Core Features

### Distance Matrix

<a name="distance-matrix"></a>

> Calculates the distance and estimated time of arrival (ETA) between origin and destination points.

**Example Input:**

```go
client.DistanceMatrix(ctx, mapnests.DistanceMatrixRequest{
	OriginLat: 23.8103,  // Dhaka
	OriginLon: 90.4125,
	DestLat:   23.7500,  // Jatrabari
	DestLon:   90.4200,
	Mode:      mapnests.TravelModeCar,
})
```

**Example Output:**

```json
{
  "data": {
	"distanceInMetres": 8900,
	"etaInSeconds": 1300
  }
}
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Distance Matrix Details

<a name="route-details"></a>

> Returns step-by-step routing metadata, including geometry, waypoints, and navigation instructions.

**Example Input:**

```go
client.DistanceMatrixDetails(ctx, mapnests.DistanceMatrixDetailsRequest{
	OriginLat: 23.7806,  // Farmgate
	OriginLon: 90.3984,
	DestLat:   23.7740,  // Dhanmondi 32
	DestLon:   90.3681,
	Mode:      mapnests.TravelModeCar,
})
```

**Example Output (simplified):**

```json
{
  "status": true,
  "message": "success",
  "data": {
	"routeResponse": {
	  "routes": [
		{
		  "distance": 4800,
		  "duration": 700,
		  "geometry": "encoded_polyline",
		  "legs": [ ... ]
		}
	  ]
	}
  }
}
```

📘 **For detailed documentation on all response fields (e.g., `routes`, `legs`, `steps`, `maneuver`, etc.), check the [Distance Matrix Response Reference](docs/distance_matrix_details.md).**

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Pairwise Route Summary

> Computes distances, ETAs and geometries  for multiple source-destination pairs in a single request. This is ideal for optimizing batch operations and comparing route statistics efficiently.

**Example Input:**

```golang
pairwiseRes, err := mapClient.PairWiseRouteSummary(ctx, mapnests.PairWiseRouteSummaryRequest{
		Pairs: []mapnests.PairWiseRoute{
			{ID: 1, Src: mapnests.Coordinate{Lat: 23.8113, Lon: 90.4135}, Dest: mapnests.Coordinate{Lat: 23.7815, Lon: 90.4123}, Mode: mapnests.TravelModeBicycling},
			{ID: 2, Src: mapnests.Coordinate{Lat: 23.8123, Lon: 90.4145}, Dest: mapnests.Coordinate{Lat: 23.7825, Lon: 90.4133}, Mode: mapnests.TravelModeBicycling},
			{ID: 3, Src: mapnests.Coordinate{Lat: 23.8133, Lon: 90.4155}, Dest: mapnests.Coordinate{Lat: 23.7835, Lon: 90.4143}, Mode: mapnests.TravelModeBicycling},
		},
	})
	if err != nil {
		log.Fatal("PairwiseRouteSummary error:", err)
	}
	fmt.Println("PairwiseRouteSummary result:", *pairwiseRes)

```

**Example Output:**

```json
{
  "status": true,
  "message": "success",
  "data": [
    {
      "id": 1,
      "distanceInMeters": 8900,
      "etaInSeconds": 1300,
      "geometry": "encoded_polyline_string"
    },
    {
      "id": 2,
      "distanceInMeters": 4800,
      "etaInSeconds": 700,
      "geometry": "another_encoded_polyline"
    }
  ]
}
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Multi Source Route Summary

> Computes distances, ETAs and geometries for multiple source-destination pairs in a single request. This is ideal for optimizing batch operations and comparing route statistics efficiently.

**Example Input:**

```golang
multiRes, err := mapClient.MultiSourceRouteSummary(ctx, mapnests.MultiSourceRouteSummaryRequest{
		Sources: []mapnests.Source{
			{ID: 1, Lat: 23.7805733, Lon: 90.2792399, Mode: string(mapnests.TravelModeCar)},
			{ID: 2, Lat: 23.75, Lon: 90.36, Mode: string(mapnests.TravelModeCar)},
			{ID: 3, Lat: 23.7, Lon: 90.42, Mode: string(mapnests.TravelModeCar)},
			{ID: 4, Lat: 23.7654321, Lon: 90.3456789, Mode: string(mapnests.TravelModeCar)},
			{ID: 5, Lat: 23.7123456, Lon: 90.3765432, Mode: string(mapnests.TravelModeCar)},
		},
		Destination: mapnests.Destination{Lat: 23.810332, Lon: 90.412518},
	})
	if err != nil {
		log.Fatal("MultiSourceRouteSummary error:", err)
	}
	fmt.Println("MultiSourceRouteSummary result:", *multiRes)
```

**Example Output:**

```json
{
    "data": {
        "routeSummaries": [
            {
                "id": 1,
                "distanceInMeters": 23782.9,
                "etaInSeconds": 1720,
                "geometry": "encoded_polyline_string"
            },
            {
                "id": 2,
                "distanceInMeters": 13421.9,
                "etaInSeconds": 1084.9,
                "geometry": "encoded_polyline_string"
            },
            {
                "id": 3,
                "distanceInMeters": 15212.3,
                "etaInSeconds": 1285.3,
                "geometry": "encoded_polyline_string"
            },
            {
                "id": 4,
                "distanceInMeters": 14120.2,
                "etaInSeconds": 1129.3,
                "geometry": "encoded_polyline_string"
            },
            {
                "id": 5,
                "distanceInMeters": 16555.4,
                "etaInSeconds": 1388,
                "geometry": "encoded_polyline_string"
            }
        ]
    },
    "message": "Success",
    "status": true
}
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Search

<a name="search"></a>

> Finds places, streets, and landmarks using text-based queries. The search can also be refined by providing latitude and longitude with radius support.

**Example Input:**

```go
client.Search(ctx, mapnests.SearchRequest{
	Query: "Uttara, Dhaka",
})
```
*To get only Zone Data provide ActiveLocations*

```go
client.Search(ctx, mapnests.SearchRequest{
	Query: "Uttara, Dhaka",
  ActiveLocations: true
})
```

**Example Output:**

```json
{
  "data": {
    "items": [
      {
        "placeId": "ebc9a1b56224de67dee16d967424915375dccaa69e9bd120f4f9c905445808c9",
        "lat": 23.858248,
        "lon": 90.4015501,
        "types": [
          "amenity",
          "school",
          "amenity"
        ],
        "address": "Scholastica (School) Senior Uttara Campus, Uttara, Dhaka",
        "name": "Scholastica (School) Senior Uttara Campus, Uttara, Dhaka",
        "houseNumber": "",
        "houseName": "",
        "street": "",
        "phone": "",
        "website": "",
        "country": "Bangladesh",
        "city": "Dhaka",
        "thana": "",
        "division": "",
        "district": "",
        "postalCode": "1230",
        "plusCode": "",
        "sublocality": "",
        "localArea": ""
      }
    ],
    "itemsPerPage": 1,
    "pageNumber": 1,
    "totalItems": 2996,
    "totalPages": 2996
  },
  "message": "Success",
  "status": true
}

```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Reverse

<a name="reverse"></a>

> Reverse geocodes coordinates to return detailed location information for a given latitude and longitude.

**Example Input:**

```go
client.Reverse(ctx, mapnests.ReverseRequest{
	Lat: 23.8103,
	Lon: 90.4125,
})
```

**Example Output:**

```json
{
  "data": {
    "placeId": "a4d8c105e24fbc3d91bb0486e1701fa20b9b56329df0fc5c47f7df003e3cc579",
    "category": "",
    "type": "yes",
    "class": "building",
    "name": "Concord Ik Tower",
    "address": "Concord Ik Tower, House#2, Road 94, Gulshan North Avenue, Gulshan 2, Gulshan, Dhaka-1212",
    "country": "Bangladesh",
    "city": "Gulshan 2, Dhaka",
    "thana": "Gulshan",
    "district": "Dhaka",
    "division": "",
    "postalCode": "1212",
    "website": "",
    "houseNumber": "2",
    "houseName": "",
    "subLocality": "",
    "localArea": "",
    "types": [
      "building",
      "yes",
      "building"
    ]
  },
  "message": "Success",
  "status": true
}
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Autocomplete

<a name="autocomplete"></a>

> Auto Complete suggests relevant places, streets, and landmarks as you type a partial search query.

**Example Input:**

```go
//Without optional value
autocompleteRes, err := mapClient.Autocomplete(ctx, mapnests.AutoCompleteRequest{
		Query: "Gulshan Road"
	})
	
// Optional parameter: ActiveZone.
// If ActiveZone is set to true, the search results will be returned within the zone. (By Default ActiveZone is true)
// If ActiveZone is set to false, the search results will be not consider any zone data.
autocompleteRes, err := mapClient.Autocomplete(ctx, mapnests.AutoCompleteRequest{
		Query: "Gulshan Road",
    ActiveZone: &activeZone,	
	}) 

// Optional parameters: Latitude, Longitude, and Radius.
// If provided, the search results will be returned within a specified radius,
// using the given latitude and longitude as the center point.
autocompleteRes, err := mapClient.Autocomplete(ctx, mapnests.AutoCompleteRequest{
		Query: "Gulshan Road",
		Lat: &lat,
		Lon: &lon,
    Radius: &radius,	
	})

// Optional parameter: Limit.
// If provided, the search results will be returned within a specified limit.
autocompleteRes, err := mapClient.Autocomplete(ctx, mapnests.AutoCompleteRequest{
		Query: "Gulshan Road",
    Limit: &limit,	
	})  
```

**Example Output:**

```json
{
  "data": [
    {
      "placeId": "4e7820118661ce107f308dff7648bf0a9d2847b78b720b08c9d39fe3662c4a8c",
      "name": "Gulshan",
      "address": "Gulshan, House#76, Palolika, Road-24, Gulshan-1, Gulshan, Dhaka-1212",
      "types": [
        "landuse",
        "residential",
        "landuse"
      ]
    }
  ],
  "message": "Success",
  "status": true
}
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>


### Autocomplete without Zone

<a name="autocomplete-without-zone"></a>

> Auto Complete suggests relevant places, streets, and landmarks as you type a partial search query.

**Example Input:**

```go
//Without optional value
autocompleteWithoutZoneRes, err := mapClient.AutocompleteWithoutZone(ctx, mapnests.AutoCompleteRequest{
		Query: "Uttara"
	})
	
//With optional value (Latitude, Longitude, Limit)
autocompleteWithoutZoneRes, err := mapClient.AutocompleteWithoutZone(ctx, mapnests.AutoCompleteRequest{
		Query: "Uttara",
		Lat: &lat,
		Lon: &lon,
		Limit: &limit,	
	})
```

**Example Output:**

```json
{
  "data": [
    {
      "placeId": "7d7e8fd275bfd9be9853ada14417d104e824d1c11600599bd326fb858429d83c",
      "name": "Uttara",
      "address": "Uttara, House#21, Road 17, Sector 11, Uttara, Dhaka-1230",
      "types": [
        "amenity",
        "restaurant",
        "amenity"
      ]
    }
  ],
  "message": "Success",
  "status": true
}

```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Search By Radius

<a name="search-by-radius"></a>

> Finds places, streets, and landmarks using text-based queries. All searches are restricted to results within the specified radius, which requires latitude, longitude, and radius parameters.

**Example Input:**

```go
client.SearchByRadius(ctx, mapnests.SearchByRadiusRequest{
	Query: "Bashundhara Residential Area, Dhaka",
	Lat:   23.8156,
	Lon:   90.4287,
	Radius: 1000,
})
```

**Example Output:**

```json
{
  "data": {
    "items": [
      {
        "placeId": "f6e5bf556e2163d89f65d634b6456d28736a6a72dc1c1df933a8b13d0597956a",
        "lat": 23.714221,
        "lon": 90.4059638,
        "types": [ "amenity", "bank", "amenity" ],
        "address": "Uttara Bank",
        "name": "Uttara Bank",
        "houseNumber": "",
        "houseName": "",
        "street": "Nawab Yousuf Sarak",
        "phone": "",
        "website": "",
        "country": "Bangladesh",
        "city": "",
        "thana": "",
        "division": "",
        "district": "",
        "postalCode": "1100",
        "plusCode": "",
        "sublocality": "",
        "localArea": ""
      }
    ],
    "itemsPerPage": 1,
    "pageNumber": 1,
    "totalItems": 12,
    "totalPages": 12
  },
  "message": "Success",
  "status": true
}

```
<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Detailed Search By PlaceId

<a name="detailed-search-by-placeId"></a>

> Provide detailed geolocation information for a place using its unique place_id, including address, coordinates, administrative regions, contact details, and metadata.

**Example Input:**

```go
client.DetailsByPlaceID(ctx, mapnests.DetailsByPlaceIDRequest{
		PlaceID: "4355aad6b8eb0b4f0ee3fa972ff9ac3fdc2d7f86f634d81f79dcf396f21826a0",
	})
```

**Example Output:**

```json
{
  "data": {
    "placeId": "4355aad6b8eb0b4f0ee3fa972ff9ac3fdc2d7f86f634d81f79dcf396f21826a0",
    "lat": 23.8060476,
    "lon": 90.3744551,
    "types": [
      "office",
      "office"
    ],
    "address": "Notari Public, House# 42, Mirpur Road, Senpara Parbata, Mirpur, Dhaka-1216",
    "name": "Mirpur",
    "houseNumber": "42",
    "houseName": "",
    "street": "Mirpur Road",
    "phone": "",
    "website": "",
    "country": "Bangladesh",
    "city": "Mirpur, Dhaka",
    "thana": "Kafrul",
    "division": "",
    "district": "Dhaka",
    "postalCode": "1216",
    "plusCode": "",
    "sublocality": "",
    "localArea": ""
  },
  "message": "Success",
  "status": true
}
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Snap to Road

<a name="snap-to-road"></a>

> Returns the nearest road’s latitude and longitude for the given input coordinates.

**Example Input:**

```ts
client.SnapToRoad(ctx, mapnests.SnapToRoadRequest{
		Mode: "walking",
		Latitude: 23.8103,
		Longitude: 90.4125,
		},
	)
```

**Example Output:**

```json
{
  "data": {
    "waypoints": [
      {
        "address": "Lane 11 East, বারিধারা ডিওএইচএস, জোয়ার সাহারা, ঢাকা, ঢাকা মহানগর, ঢাকা জেলা, ঢাকা বিভাগ, 1229, বাংলাদেশ",
        "location": [
          90.412496,
          23.810403
        ],
        "distance_meters": 11.41511481
      }
    ]
  },
  "message": "Success",
  "status": true
}
```
<p align="right">(<a href="#readme-top">back to top</a>)</p>

---


### Multi Stop Point

<a name="snap-to-road"></a>

> Calculates distance, ETA, and route geometry for a journey starting from a single source and passing through multiple stop points in sequence. It returns both total and per-segment results, processes routes concurrently for efficiency, and supports multiple stops.

**Example Input:**

```ts

  client.MultiStopPoints(ctx, mapnests.MultiStopPointsRequest{
		Src: mapnests.Coordinate{Lat: 23.7805733, Lon: 90.2792399},
		StopPoints: []mapnests.StopPoint{
			{ID: 1, Lat: 23.7805733, Lon: 90.2792399},
			{ID: 2, Lat: 23.75, Lon: 90.36},
		},
		Mode: mapnests.TravelModeBicycling,
	})
```

**Example Output:**

```json
{
  "data": {
    "DistanceInMeters": 10783.300170898438,
    "EtaInSeconds": 2903.2998962402344,
    "RouteSummaries": [
      {
        "id": 1,
        "distanceInMeters": 780.1,
        "etaInSeconds": 243.9,
        "geometry": "qjipCiznfPQ_A~Dy@zCy@{CmJiBiHaAL]wDmAN",
        "Source": {
          "lat": 23.809973415982903,
          "lon": 90.35697149649764
        },
        "StopPoint": {
          "lat": 23.81038738311683,
          "lon": 90.36203008256733
        }
      },
      {
        "id": 2,
        "distanceInMeters": 10003.2,
        "etaInSeconds": 2659.4,
        "geometry": "{mipCoyofP^oXoAcHbTuE?}IpEwYr_@}h@fNgd@zZmC`NZ~Co@a@kJbB_Kw@sSl@}^l@SuEc@zC{a@mReBT_VqXqpAsNpAhN}ByDmv@SwNbBq@",
        "Source": {
          "lat": 23.81038738311683,
          "lon": 90.36203008256733
        },
        "StopPoint": {
          "lat": 23.798308134287165,
          "lon": 90.43522641639149
        }
      }
    ]
  },
  "message": "Success",
  "status": true
}
```
<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Multi Source Route Summary Without Geometry

<a name="multi-source-route-summary-without-geometry"></a>

> Computes multiple source-destination pairs in a single request, without returning route geometry. This is ideal when only summary metrics are needed and the extra geometry payload should be avoided.

**Example Input:**

```golang
multiRes, err := mapClient.MultiSourceRouteSummaryWithoutGeometry(ctx, mapnests.MultiSourceRouteSummaryRequest{
		Sources: []mapnests.Source{
			{ID: 1, Lat: 23.7805733, Lon: 90.2792399, Mode: "car"},
			{ID: 2, Lat: 23.75, Lon: 90.36, Mode: "car"},
			{ID: 3, Lat: 23.7, Lon: 90.42, Mode: "car"},
			{ID: 4, Lat: 23.7654321, Lon: 90.3456789, Mode: "car"},
			{ID: 5, Lat: 23.7123456, Lon: 90.3765432, Mode: "car"},
		},
		Destination: mapnests.Destination{Lat: 23.810332, Lon: 90.412518},
	})
	if err != nil {
		log.Fatal("MultiSourceRouteSummaryWithoutGeometry error:", err)
	}
	fmt.Println("MultiSourceRouteSummaryWithoutGeometry result:", *multiRes)
```

**Example Output:**

```json
{
  "data": {
    "routeSummaries": [
      {
        "id": 1,
        "distanceInMeters": 23766.3,
        "etaInSeconds": 1717.7,
        "reachable": true
      },
      {
        "id": 2,
        "distanceInMeters": 13392.8,
        "etaInSeconds": 1082,
        "reachable": true
      },
      {
        "id": 3,
        "distanceInMeters": 15163.7,
        "etaInSeconds": 1283.1,
        "reachable": true
      },
      {
        "id": 4,
        "distanceInMeters": 14197.6,
        "etaInSeconds": 1130.9,
        "reachable": true
      },
      {
        "id": 5,
        "distanceInMeters": 16513.4,
        "etaInSeconds": 1384.6,
        "reachable": true
      }
    ]
  },
  "message": "Success",
  "status": true
}
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Geocode

<a name="geocode"></a>

> Finds places, streets, and landmarks using text-based queries, returning geocoded coordinates and address metadata.

**Example Input:**

```go
client.Geocode(ctx, mapnests.GeocodeRequest{
	Query: "Uttara, Dhaka",
})

// Optional parameters: Limit, AcceptedLanguage, XRequestID
limit := int64(5)
lang := "en"
client.Geocode(ctx, mapnests.GeocodeRequest{
	Query:            "Uttara, Dhaka",
	Limit:            &limit,
	AcceptedLanguage: &lang,
})
```

**Example Output:**

```json
{
  "data": [
    {
      "place_id": 528793,
      "lat": "23.8693275",
      "lon": "90.3926893",
      "category": "place",
      "type": "suburb",
      "place_rank": 19,
      "importance": 0.14667666666666662,
      "addresstype": "suburb",
      "name": "Uttara",
      "display_name": "Uttara, Dhaka, Dhaka Metropolitan, Dhaka District, Dhaka Division, 1230, Bangladesh",
      "address": ""
    },
    {
      "place_id": 426690,
      "lat": "23.73929105",
      "lon": "90.40722897981641",
      "category": "building",
      "type": "residential",
      "place_rank": 30,
      "importance": 0.00000999999999995449,
      "addresstype": "building",
      "name": "Uttara (Oficers Quarter)",
      "display_name": "Uttara (Oficers Quarter), Petrol Pump Road, Kakrail, Segunbagicha, Dhaka, Dhaka Metropolitan, Dhaka District, Dhaka Division, 1000, Bangladesh",
      "address": ""
    }
  ],
  "message": "Success",
  "status": true
}
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### ETA With Geometry

<a name="eta-with-geometry"></a>

> Calculates distance, ETA, and route geometry for a single source-destination pair across all travel modes (car, cng, motorcycle) in one request.

**Example Input:**

```go
res, err := client.EtaWithGeometry(ctx, mapnests.ETAsRequest{
	FromLat: 23.759248321873432,
	FromLon: 90.41691646513276,
	ToLat:   23.770148309864418,
	ToLon:   90.40842674496818,
})
if err != nil {
	log.Fatal("EtaWithGeometry error:", err)
}
fmt.Println("EtaWithGeometry result:", *res)
```

**Example Output:**

```json
{
  "status": true,
  "message": "success",
  "data": {
    "car": {
      "distance": 2971.4,
      "duration": 277.00000000000006,
      "geometry": "gxcil@_qrmkDbO{BmJq\\{C{OeD}Sc@qBh@yByh@wW}SmKs@YoVoLaReJkfAkg@qEuByy@ca@et@_\\w_@sQgNuGot@y\\cXeMwDxBqBvA{CtAiBb@sARmAFyB@wAUqAc@sCmA{HmEwSyKcHqDiDkBsCaBgB_AeA_@mAOeA@cA`@}@z@i@~@{FpK}CnF}AxH_AxF[lEKpAChB?|COhd@UnoAk@zM{AxNgDfTiB~OuG~t@gGjj@KdEA~Rd@tH|AjNtApJjCpInFdNdD`HpIjMrUd^d[t^hAzBTt@JbA@pBO`EEvBKrBQbBW`BwBtB}pAlnAiGnF{I`GmDuCuDwBmJeDiHeC{E{AiGgB}F{AaKwBbAzIjA|IxBdPn@|AFfAGfAkAlMoBzWaAfM}@dFcBzDb@~BLnCAxBe@tGoAvPaDzc@eGzy@w@dK",
      "waypoints": [
        {
          "name": "",
          "location": [90.416928, 23.759252],
          "distance": 0.84497434,
          "hint": "ttITgNjSE4BpAAAAewAAAAAAAAAAAAAAOx_qQffDB0IAAAAAAAAAAGkAAAB7AAAAAAAAAAAAAABsAAAAIKdjBZSJagEYp2MFkolqAQAAvxUAAAAA"
        },
        {
          "name": "Bir Uttam Mir Shawkat Sarak",
          "location": [90.408445, 23.770238],
          "distance": 9.86563,
          "hint": "6rgBgP___38WAAAARQAAAAAAAADsAAAAau6gQYBKJEIAAAAAvJhRQxYAAABFAAAAAAAAAOwAAABsAAAA_YVjBX60agHuhWMFJrRqAQAAbxEAAAAA"
        }
      ]
    },
    "cng": {
      "distance": 2971.4,
      "duration": 277.00000000000006,
      "geometry": "gxcil@_qrmkDbO{BmJq\\{C{OeD}Sc@qBh@yByh@wW}SmKs@YoVoLaReJkfAkg@qEuByy@ca@et@_\\w_@sQgNuGot@y\\cXeMwDxBqBvA{CtAiBb@sARmAFyB@wAUqAc@sCmA{HmEwSyKcHqDiDkBsCaBgB_AeA_@mAOeA@cA`@}@z@i@~@{FpK}CnF}AxH_AxF[lEKpAChB?|COhd@UnoAk@zM{AxNgDfTiB~OuG~t@gGjj@KdEA~Rd@tH|AjNtApJjCpInFdNdD`HpIjMrUd^d[t^hAzBTt@JbA@pBO`EEvBKrBQbBW`BwBtB}pAlnAiGnF{I`GmDuCuDwBmJeDiHeC{E{AiGgB}F{AaKwBbAzIjA|IxBdPn@|AFfAGfAkAlMoBzWaAfM}@dFcBzDb@~BLnCAxBe@tGoAvPaDzc@eGzy@w@dK",
      "waypoints": [
        {
          "name": "",
          "location": [90.416928, 23.759252],
          "distance": 0.84497434,
          "hint": "ttITgNjSE4BpAAAAewAAAAAAAAAAAAAAOx_qQffDB0IAAAAAAAAAAGkAAAB7AAAAAAAAAAAAAABsAAAAIKdjBZSJagEYp2MFkolqAQAAvxUAAAAA"
        },
        {
          "name": "Bir Uttam Mir Shawkat Sarak",
          "location": [90.408445, 23.770238],
          "distance": 9.86563,
          "hint": "6rgBgP___38WAAAARQAAAAAAAADsAAAAau6gQYBKJEIAAAAAvJhRQxYAAABFAAAAAAAAAOwAAABsAAAA_YVjBX60agHuhWMFJrRqAQAAbxEAAAAA"
        }
      ]
    },
    "motorcycle": {
      "distance": 2971.4,
      "duration": 207.75,
      "geometry": "gxcil@_qrmkDbO{BmJq\\{C{OeD}Sc@qBh@yByh@wW}SmKs@YoVoLaReJkfAkg@qEuByy@ca@et@_\\w_@sQgNuGot@y\\cXeMwDxBqBvA{CtAiBb@sARmAFyB@wAUqAc@sCmA{HmEwSyKcHqDiDkBsCaBgB_AeA_@mAOeA@cA`@}@z@i@~@{FpK}CnF}AxH_AxF[lEKpAChB?|COhd@UnoAk@zM{AxNgDfTiB~OuG~t@gGjj@KdEA~Rd@tH|AjNtApJjCpInFdNdD`HpIjMrUd^d[t^hAzBTt@JbA@pBO`EEvBKrBQbBW`BwBtB}pAlnAiGnF{I`GmDuCuDwBmJeDiHeC{E{AiGgB}F{AaKwBbAzIjA|IxBdPn@|AFfAGfAkAlMoBzWaAfM}@dFcBzDb@~BLnCAxBe@tGoAvPaDzc@eGzy@w@dK",
      "waypoints": [
        {
          "name": "",
          "location": [90.416928, 23.759252],
          "distance": 0.84497434,
          "hint": "ttITgNjSE4BpAAAAewAAAAAAAAAAAAAAOx_qQffDB0IAAAAAAAAAAGkAAAB7AAAAAAAAAAAAAABsAAAAIKdjBZSJagEYp2MFkolqAQAAvxUAAAAA"
        },
        {
          "name": "Bir Uttam Mir Shawkat Sarak",
          "location": [90.408445, 23.770238],
          "distance": 9.86563,
          "hint": "6rgBgP___38WAAAARQAAAAAAAADsAAAAau6gQYBKJEIAAAAAvJhRQxYAAABFAAAAAAAAAOwAAABsAAAA_YVjBX60agHuhWMFJrRqAQAAbxEAAAAA"
        }
      ]
    }
  }
}
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### ETA With Geometry By Mode

<a name="eta-with-geometry-by-mode"></a>

> Calculates distance, ETA, and route geometry for a single source-destination pair for one specific travel mode.

**Example Input:**

```go
res, err := client.EtaWithGeometryByMode(ctx, mapnests.ETAWithGeometryByModeRequest{
	FromLat: 23.7805733,
	FromLon: 90.2792399,
	ToLat:   23.810332,
	ToLon:   90.412518,
	TripId:  "550e8400-e29b-41d4-a716-446655440000",
	Mode:    mapnests.TravelModeCar,
})
if err != nil {
	log.Fatal("EtaWithGeometryByMode error:", err)
}
fmt.Println("EtaWithGeometryByMode result:", *res)
```

**Example Output:**

```json
{
  "status": true,
  "message": "Success",
  "data": {
    "distance": 23766.3,
    "duration": 3537.316298415542,
    "geometry": "gtpjl@wfgekD_R`Iqj@`Leo@jGw`@`F_VtCkTbIoUxHeg@nZoLlS{Kj`@mMxVaQfM_M~F{C?_GiD}@eD}Ee@kHuBeJ_A{WXyHt@y@jD_EzJ_Ii@qc@_Bua@xAiWnCkWpDeo@xHkjA~KuFKeBmByAgFqAwFwCwDqHLiIz@yIaCeJ{AiNjDxTdlBtDd\\z@rHjGlg@zF|o@rCzUbCrc@Lrc@D~JgA`f@mE|g@cIri@eOlo@}EpRcG|SwNz^kHjOgP~XkQl\\{Sj_@eJjNuXjb@uExEyC^gBEcAUcA_A]iA?qAd@iBnAgCdDyF|GmGnTmTjMqPtSeYdHyLnFmIpHcOhGmLbFoJ~G_OtLs]zMug@rKcf@vFqa@bCyY~Ak]~@w\\A{W{@wYqBm\\cEqh@}Foh@wGqc@{AuNaVkwBu\\acCiCeOqBsQcAcLcBuc@cCsZkDo^c@yFsBoZ}AqSw@qPSgc@iAeTw@af@kA_c@qCmrAc@c\\{Aqx@p@ciBLaE`Ckx@nE_l@jFer@fAuNdQmyAlT}mArc@{cCd^msBxJ_g@`ZeiB`]k_ChLw`A|Iaw@xYi|BnHcm@`AgH`CkS^eD`CiRbFk`@lAuJhByPlE{\\hAiIpMgeAhEa]hEqb@XyBjBwMhCgWhEik@xAmW`AeSnBu_@hBir@lA{eA?oMG{|@YqTmAcVeCmWgHmj@}BaToCga@aA_\\_C{m@}@oV@aPPuNr@kLrAsNdB}J`B_InCuKdCwIhD_MjFcPpF_QjIyTbDwHh@gA~DsIbDyFjDeG~DaGvD{FrL_OvKyN`S}WfQoV`[qa@bEoFtP_UvIqMpIiNtBoExJiQzCmGdCoGzDaK`BiFjBoG|@mDhAmElCyKhH}YtJ_c@xSq`ApDkRhFcVhMgh@dEmNfEcNj{@i{CvOqo@f@cBzD{P`BoHdE_RfB}MxDyk@j@mNP_NIaMyB{h@mByo@kCir@oC}o@k@oSa@mOKmS?oE~@uS|@wRvAcVz@_NbFcj@bKsmAef@gEioAuKsw@{Hqb@gD{fAsIuo@eFat@yF_Ki@}TeB_QsAoE[uRuA}\\mBqLw@_n@eEaOiAmJs@uRmA{HYob@mCwa@kCmeAqHqLk@eOmAuW}BcPiAuFaAgBk@{A}@o@aAc@cA}@}BqN{a@wK{ZcOe^iJ}RyTcb@yRg_@}a@gw@sHoMmCqDeWgg@mJyQuWmd@mA}AoAkAqBeAi@OgCuAsAkAu@{@u@oAy@wBg@sBUuAKyBDqBaAeDcAyBkA_CcFgJyMiW}NeXm[{l@{GcMcJiPwf@o}@qDwGot@srAkM_VoOkXcFuIg[ok@aWme@{B_EmR}]iEaI_BoCwS{b@k]qr@aEuIu@aBoIuQ_C{F_F}MkCwImCaJw@wD{ByMwSw|AwCySs@iDcAiDyAi@cAcAm@wAQcBHeBf@{A~@gAiEoXgAwWRmSfCs]nGoc@xAuPfHgi@fDsUxLay@n@cEv@kHtHig@zAyFfIwZlKoVhPkZbYa`@xQeXzGcJdNoPbEaFpPqSlAsAdAmAr^md@x[ia@vlAo_Bzk@yu@|NqWxSeb@zIaTpQ{g@jTat@vUsw@xh@olBvJif@dC}K^q@~@m@zAUv@KlBElPYbPa@vCQtRoAhMeB`TkE`^{D|[mCbAKfPwA|^iDt^sDx]mDvRq@|MTdMbAtIXtIj@rk@`CxJb@bWHnJu@xYaCxCm@jC{Ah@}A\\eAi@}RIsC]aHWqEiA{SS}D{B{a@AeFAmCDg@v@}H|[}cBsViHoQcE}KaCcfAgJeMcAu]qDyi@iGij@yIoOwC_MkC{TsEsr@yJsf@{IqrA}Uoa@qHu]{Fw@O_[eFwI}Bw@Ssq@iT_c@uOsPeH{HqFcL}J}MkNkZg]ae@gg@m[e^oYwYoFgFmIoEmNoH}ByAoj@oYuQ_JgNwEsVyF}P}Bix@oE}PgAmaAeG}g@iCkt@qF{LqAOkEcDk~@aK{zBwC}u@MiH{FNyF`@gJ`BXwAT}AI}Am@_C_BeDmBkD_@w@QiAQqCGgE?kAFmCvFeBrCg@rBWtCUrD`@vLhBeKgPg^m^ij@_w@aSw^w`@_t@oL}TuIuQmr@mfBucC_qEluAq@r\\{@x\\Gl]q@~[m@t^k@h\\oAd\\Wb]g@n\\y@x\\w@`@r_@",
    "waypoints": [
      {
        "name": "",
        "location": [90.28006, 23.782228],
        "distance": 201.75824,
        "hint": "P8ERgF7BEYA2AAAAAAAAAMABAAAAAAAAvtIVQgAAAACsIJtDAAAAADYAAAAAAAAAwAEAAAAAAABsAAAAfJBhBVTjagFIjWEF2txqAQoArwkAAAAA"
      },
      {
        "name": "Lane 11 East",
        "location": [90.412517, 23.810403],
        "distance": 8.090942,
        "hint": "lJoTgHObE4B8AAAAeAAAAAAAAAAAAAAAdzVcQpbgVEIAAAAAAAAAAHwAAAB4AAAAAAAAAAAAAABsAAAA5ZVjBWNRawHolWMFGlFrAQAAjwUAAAAA"
      }
    ]
  }
}
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### ETA Without Geometry By Mode

<a name="eta-without-geometry-by-mode"></a>

> Calculates distance and ETA for a single source-destination pair for one specific travel mode, without route geometry. Optionally returns multiple candidate routes via `NoOfRoutes`.

**Example Input:**

```go
numberOfRoutes := int64(2)

res, err := client.EtaWithoutGeometryByMode(ctx, mapnests.ETAByModeRequest{
	FromLat:    23.82440246492232,
	FromLon:    90.35823219240551,
	ToLat:      23.77683932043941,
	ToLon:      90.40510870368077,
	Mode:       mapnests.TravelModeCar,
	NoOfRoutes: &numberOfRoutes,
})
if err != nil {
	log.Fatal("EtaWithoutGeometryByMode error:", err)
}
fmt.Println("EtaWithoutGeometryByMode result:", *res)
```

**Example Output:**

```json
{
  "data": [
    {
      "from": {
        "lat": 23.824402,
        "lon": 90.358232
      },
      "to": {
        "lat": 23.776839,
        "lon": 90.405109
      },
      "distance": 10739.9,
      "eta": 905.3920000000002
    },
    {
      "from": {
        "lat": 23.824402,
        "lon": 90.358232
      },
      "to": {
        "lat": 23.776839,
        "lon": 90.405109
      },
      "distance": 11800.2,
      "eta": 1127.792
    }
  ],
  "message": "Success",
  "status": true
}
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### ETA Multiple Stoppage

<a name="eta-multiple-stoppage"></a>

> Calculates distance, ETA, and route geometry for a journey from a source to a dropoff through one or more ordered stoppages, per requested travel mode. Each mode's response is broken down into per-hop segments plus totals.

**Example Input:**

```go
res, err := client.EtaMultiStoppage(ctx, mapnests.ETAMultipleStoppageRequest{
	Source: mapnests.LatLng{
		Latitude:  23.8115,
		Longitude: 90.414,
	},
	DropOff: mapnests.LatLng{
		Latitude:  23.8115,
		Longitude: 90.414,
	},
	Modes: []mapnests.Mode{
		mapnests.TravelModeCar,
		mapnests.TravelModeMotorcycle,
	},
	Stoppages: []mapnests.ETAStoppage{
		{
			Latitude:      23.795,
			Longitude:     90.41,
			StoppageOrder: 1,
		},
	},
	TripId: "550e8400-e29b-41d4-a716-446655440000",
})
if err != nil {
	log.Fatal("EtaMultiStoppage error:", err)
}
fmt.Println("EtaMultiStoppage result:", *res)
```

**Example Output:**

```json
{
  "status": true,
  "message": "success",
  "trip_id": "781473024756609000",
  "data": {
    "car": {
      "hops": [
        {
          "order": 1,
          "from": "source",
          "to": "stoppage_1",
          "distance": 2821.9,
          "duration": 382.5,
          "geometry": "wsill@gzlmkDdAz~@n\\y@x\\w@`]yAh_@oB|NyFbZkYbMgLxQiRnQkU|N_SfKkLdLwIjMoEfYyIxC_AhDnBxFx@fE`@hs@eFz@Bj@v@JnWl@leABlBx@rBjALlA?fG?~{@MzD??hCAhxAAjq@~Jvo@tbByb@j_Cmo@fTqF~SkGpKiC~f@mM~@|ErF|WDNlClV`@db@Rt{@`FEhn@u@dkAkH`z@iCzd@wCP|CNhb@"
        },
        {
          "order": 2,
          "from": "stoppage_1",
          "to": "dropoff",
          "distance": 3335,
          "duration": 454.9,
          "geometry": "u`ikl@i`emkDOib@Q}CjTg@dg@cA}Ime@oKqi@cDsOyKcl@aHia@qCeMcD_IeGiLeMfImHnE}KhDyMvDuG~AqThFa\\tHgVfHs]bJwSnFkk@~N_Dx@sr@zRiLdDw\\rI}TfGgl@nOyUvGoVfHaD|@aFpA_Bb@aa@tKya@fKu]fIQDwBf@aAiFiLuo@Ayr@EsxAgkAi@aAu@]qAFkCIofAQyTYu@q@OyABsAFgEXuf@fFqBQiAQwEi@aEiAaDiBmIeNqHwOmCkH_S|D_O`@_LqGmBo@_CrKeAbTyu@tFsS~CkMlDmGd@oHt@iLHeDZc@RLhGr@xEClTVhPe@zm@Bbh@a]xAy\\v@o\\x@eA{~@"
        }
      ],
      "total_distance": 6156.9,
      "total_duration": 837.4
    }
  }
}
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## License

<a name="license"></a>

This project is licensed under the [MIT License](LICENSE).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## Contact

<a name="contact"></a>

📧 [dev@mapnests.com](mailto:dev@mapnests.com)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
