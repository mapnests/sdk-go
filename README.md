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




A secure and efficient Go SDK for the **Mapnests Platform**, enabling powerful geospatial capabilities such as **Search (Geocoding)**, **Reverse (Reverse Geocoding)**, and **Distance Matrix**.

---

## 📚 Table of Contents

* [Installation](#installation)
* [Quick Start](#quick-start)
* [Core Features](#core-features)

  * [Distance Matrix](#distance-matrix)
  * [Distance Matrix Details](#distance-matrix-details)
  * [Pairwise Route Summary](#pairwise-route-summary)
  * [Multi Source Route Summary](#multi-source-route-summary)
  * [Search (Geocoding)](#search-geocoding)
  * [Reverse Geocoding](#reverse-geocoding)
* [License](#license)
* [Contact](#contact)

---

## Installation

```bash
go get github.com/mapnests/sdk-go
```

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
	client := mapnests.NewClient("YOUR_API_KEY", "com.your.app")

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

### Search (Geocoding)

<a name="search-geocoding"></a>

> Finds places, streets, and landmarks by text query.

**Example Input:**

```go
client.Search(ctx, mapnests.SearchRequest{
	Query: "Bashundhara Residential Area, Dhaka",
})
```

**Example Output:**

```json
{
  "data": [
	{
	  "place_id": "123456",
	  "lat": "23.8156",
	  "lon": "90.4287",
	  "display_name": "Bashundhara Residential Area, Dhaka, Bangladesh"
	}
  ],
  "status": true,
  "message": "success"
}
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Reverse Geocoding

<a name="reverse-geocoding"></a>

> Converts GPS coordinates into a readable address.

**Example Input:**

```go
client.Reverse(ctx, mapnests.ReverseRequest{
	Lat: 23.7806,
	Lon: 90.3984,
})
```

**Example Output:**

```json
{
  "data": {
	"displayName": "Farmgate, Tejgaon, Dhaka, Bangladesh",
	"address": {
	  "country": "Bangladesh",
	  "state": "Dhaka Division",
	  "city": "Dhaka"
	}
  },
  "status": true,
  "message": "success"
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
