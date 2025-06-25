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

📧 [support@example.com](mailto:support@Example.com)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
