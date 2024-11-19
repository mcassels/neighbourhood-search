package backend

import (
	"context"
	"fmt"
	"log"

	"googlemaps.github.io/maps"
)

func Geocode(apiKey string, address string) string {
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	// Restrict the geocoding to the area around Victoria, BC
	// Northeast: 48.712274, -123.276955
	// Southwest: 48.315447, -123.811767
	r := &maps.GeocodingRequest{
		Address: address,
		Bounds: &maps.LatLngBounds{
			NorthEast: maps.LatLng{
				Lat: 48.712274,
				Lng: -123.276955,
			},
			SouthWest: maps.LatLng{
				Lat: 48.315447,
				Lng: -123.811767,
			},
		},
	}
	fmt.Println("making request")
	results, err := c.Geocode(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	if len(results) == 0 {
		return "No results found"
	}
	return results[0].Geometry.Location.String()
}
