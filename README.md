# Local Search API

This API allows searching for locations within a given radius of a longitude and latitude.

## Endpoints
### GET 
Allows searching for locations within a circular or square area.

## Query Parameters

- long: The longitude of the search origin.
- lat: The latitude of the search origin.
- distance: The radius in meters of the search area.
- circle: Set to true for a circular search area, (default) false for square bounding box around coordinates.

## Example Requests

- Search within 100m of (-88.33, 36.33):
`/search?long=-88.33&lat=36.33&distance=100`
- Circular search within 200m:
`/search?long=-88.33&lat=36.33&distance=200&circle=true`

Response

Returns a JSON array of location objects, ordered by distance from the search origin (closest first), then by rating (highest first) for locations within 50m of the query coordinates.

A location object contains:

- id
- name
- website
- coordinates
- description
- rating
- distance (in meters)
  
## Error Responses

400 Bad Request if invalid query parameters are passed.
500 Internal Server Error for any server-side issues.
## Local Setup
To run locally:

- Install Go.
- Clone this repo.
- Install dependencies: go get ./...
- The main data management logic is handled in `data.go`. The database path can be changed in `get.go`. The `main.go` file initializes the API, starts the server and provides error handling.
- Start the server: go run main.go.
- The API will be running on http://localhost:8000.


