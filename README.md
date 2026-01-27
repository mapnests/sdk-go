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
  * [Autocomplete Without Zone](#autocomplete-without-zone)
  * [Search By Radius](#search-by-radius)
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
    "lat": 23.7952951,
    "lon": 90.41680379038863,
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
	
//With optional value (Latitude, Longitude, Limit)
autocompleteRes, err := mapClient.Autocomplete(ctx, mapnests.AutoCompleteRequest{
		Query: "Gulshan Road",
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

---


## License

<a name="license"></a>

This project is licensed under the [MIT License](LICENSE).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## Contact

<a name="contact"></a>

📧 [dev@mapnests.com](mailto:dev@mapnests.com)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
