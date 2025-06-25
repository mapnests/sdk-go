


##### Readme Top
<div>
<a align="left" name="readme-top"></a>
<p align="right"><a href="../README.md#distance-matrix-details">Back to main readme file</a></p>
</div>

<br />
<div align="center">
  <a>
	<img src="https://cdn.pixabay.com/photo/2015/06/23/11/13/globe-818583_1280.png" alt="Logo" width="100">
  </a>

<h1 align="center">Distance Matrix Response Details</h1>

  <p align="center">
	
Map Nests Go SDK
  </p>
</div>


## 📑 Table of Contents

* [SDK Example Usage](#sdk-example-usage)
* [Sample Output](#sample-output)
* [Response Structure Overview](#response-structure-overview)

  * [Routes](#routes)

    * [Legs](#legs)

      * [Steps](#steps)

        * [Maneuver](#maneuver)
        * [Intersections](#intersections)
  * [Waypoints](#waypoints)

---

## SDK Example Usage

```go
client.DistanceMatrixDetails(ctx, mapnests.DistanceMatrixDetailsRequest{
  OriginLat: 23.7806,  // Farmgate
  OriginLon: 90.3984,
  DestLat:   23.7740,  // Dhanmondi 32
  DestLon:   90.3681,
  Mode:      mapnests.TravelModeCar,
})
```
<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## Sample Output

<details>
<summary>Click to expand sample JSON</summary>

```json
{
  "status": true,
  "message": "Success",
  "data": {
    "routeResponse": {
      "code": "Ok",
      "message": "",
      "data_version": "",
      "routes": [
        {
          "distance": 3939.2,
          "duration": 591.4,
          "weight_name": "routability",
          "weight": 591.4,
          "geometry": "uvwxFzvpkNfD|FfHrJzJzNzAdDlEjDnJ~GhB|@",
          "legs": [
            {
              "distance": 3939.2,
              "duration": 591.4,
              "summary": "মাদানী এভিনিউ, প্রগতি সরণী",
              "weight": 591.4,
              "steps": [
                {
                  "distance": 238,
                  "duration": 28.3,
                  "geometry": "ev|kl@ml|mkDcAqKnvB_Z",
                  "name": "প্রগতি সরণী",
                  "mode": "driving",
                  "driving_side": "right",
                  "weight": 28.3,
                  "maneuver": {
                    "location": [90.421975, 23.804787],
                    "bearing_before": 348,
                    "bearing_after": 78,
                    "type": "continue",
                    "modifier": "uturn"
                  },
                  "intersections": [
                    {
                      "location": [90.421975, 23.804787],
                      "bearings": [75, 165, 345],
                      "entry": [true, false, true],
                      "in": 1,
                      "out": 0
                    }
                  ]
                }
              ]
            }
          ]
        }
      ],
      "waypoints": [
        {
          "name": "Swapnil Road",
          "location": [90.41644, 23.810034],
          "distance": 0.443003,
          "hint": "CusTgA3r..."
        }
      ]
    }
  }
}
```

</details>

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## Response Structure Overview

### Routes

| Name         | Definition                                                   | Purpose                                                                 | Example           |
| ------------ | ------------------------------------------------------------ | ----------------------------------------------------------------------- | ----------------- |
| distance     | Total distance of the full route in meters                   | Helps calculate trip length, useful for cost, fuel, and time estimation | 3939.2            |
| duration     | Estimated time to complete the route (in seconds)            | Used to show ETA and plan schedules                                     | 591.4             |
| weight_name  | Metric used during routing (e.g., "duration", "routability") | Indicates what factor the route optimization was based on               | "routability"     |
| weight       | Route cost using the optimization metric                     | Helps prioritize routes based on efficiency, not just time              | 591.4             |
| geometry     | Encoded polyline of the entire route                         | Can be decoded to draw the path on a map                                | "uvwxFzvpkNfDFfHrJz..." |
| legs[]       | Segments between waypoints                                   | Breaks the full route into parts for step-wise navigation               | See [Legs](#legs) |

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Legs

| Name     | Definition                               | Purpose                                               | Example                      |
| -------- | ---------------------------------------- | ----------------------------------------------------- | ---------------------------- |
| distance | Distance in meters of this segment       | Defines how far this portion of the route is          | 3939.2                       |
| duration | Time estimate for this leg in seconds    | Allows estimating per-leg travel time                 | 591.4                        |
| summary  | Major road names covered in this segment | Useful for displaying a quick path summary in UIs     | "মাদানী এভিনিউ, প্রগতি সরণী" |
| weight   | Weighted cost of the leg                 | Used internally to rank or evaluate route performance | 591.4                        |
| steps\[] | Turn-by-turn navigation instructions     | Provides step-level breakdown for navigation systems  | See [Steps](#steps)          |

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Steps

| Name          | Definition                           | Purpose                                                               | Example                             |
| ------------- | ------------------------------------ | --------------------------------------------------------------------- | ----------------------------------- |
| distance      | Length of this instruction step      | Shows how far to travel before the next action                        | 238                                 |
| duration      | Time for this step (in seconds)      | Used for showing countdowns and progress in navigation UIs            | 28.3                                |
| geometry      | Encoded polyline for this step       | Represents only the portion of the route covered in this step         | "evkl\@mlmkDcAqKnvB\_Z"             |
| name          | Street or road name                  | Describes the current road in turn instructions                       | "প্রগতি সরণী"                       |
| mode          | Travel mode                          | Informs if user is walking, driving, biking, etc.                     | "driving"                           |
| driving_side  | Side of road to drive on             | Relevant in countries with left/right driving regulations             | "right"                             |
| weight        | Step weight/cost                     | Helps evaluate internal routing logic per instruction                 | 28.3                                |
| maneuver      | The action to perform                | Tells the user what movement to make at the beginning of the step     | See [Maneuver](#maneuver)           |
| intersections | Relevant intersections for this step | Provides network junction context where decisions or checks may occur | See [Intersections](#intersections) |

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Maneuver

| Name            | Definition                              | Purpose                                                     | Example                 |
| --------------- | --------------------------------------- | ----------------------------------------------------------- | ----------------------- |
| location        | Coordinates of the maneuver             | Exact point where user must act (turn, arrive, depart)      | \[90.421975, 23.804787] |
| bearing\_before | Direction vehicle is facing before turn | Helps determine approach angle into maneuver                | 348                     |
| bearing\_after  | Direction vehicle is facing after turn  | Indicates resulting direction after the action              | 78                      |
| type            | Type of maneuver                        | Defines action such as "continue", "turn", "arrive", etc.   | "continue"              |
| modifier        | Optional directional modifier           | Adds nuance like "uturn", "slight left", "sharp right" etc. | "uturn"                 |

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Intersections

| Name     | Definition                               | Purpose                                            | Example                 |
| -------- | ---------------------------------------- | -------------------------------------------------- | ----------------------- |
| location | Coordinates of the intersection          | Specifies where a junction or choice-point occurs  | \[90.421975, 23.804787] |
| bearings | Possible travel directions at this point | Lists angles (0-359) available at the node         | \[75, 165, 345]         |
| entry    | Boolean flags for each direction         | Shows which directions are allowed from this point | \[true, false, true]    |
| in       | Index of direction coming into the node  | Helps determine how the vehicle arrived            | 1                       |
| out      | Index of chosen outgoing direction       | Specifies which direction will be taken next       | 0                       |

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Waypoints

| Name     | Definition                                     | Purpose                                                    | Example                |
| -------- | ---------------------------------------------- | ---------------------------------------------------------- | ---------------------- |
| name     | Nearest known road name                        | Gives readable label for the snapped location              | "Swapnil Road"         |
| location | Snapped coordinate on routable road            | Accurate point used internally for route generation        | \[90.41644, 23.810034] |
| distance | Distance in meters from input to snapped point | Shows how far original input was adjusted                  | 0.443003               |
| hint     | Encoded routing hint                           | Used for optimizing follow-up route calls; opaque to users | "CusTgA3r..."          |

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---
