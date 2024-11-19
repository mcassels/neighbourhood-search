package types

import "googlemaps.github.io/maps"

type AddressResult struct {
	Address       string
	Neighbourhood string
	LatLng        maps.LatLng
}
