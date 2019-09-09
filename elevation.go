package main

import (
	"context"
	"fmt"
	"log"

	"googlemaps.github.io/maps"
)

func main() {
	client, err := maps.NewClient(maps.WithAPIKey("YOUR_KEY_HERE"))

	if err != nil {
		log.Fatalf("Fatal error: %s", err)
	}

	request := &maps.ElevationRequest{
		Locations: []maps.LatLng{
			{
				Lat: 39,
				Lng: -104,
			},
		},
	}

	response, err := client.Elevation(context.Background(), request)

	if err != nil {
		log.Fatalf("Request failed with: %s", err)
	}

	fmt.Println(response[0].Elevation)
}
