package main

import (
	"fmt"
	"math"
)

const earthRadiusKm = 6371.0

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func greatCircleDistance(lat1, lon1, lat2, lon2 float64) float64 {
	φ1 := degreesToRadians(lat1)
	φ2 := degreesToRadians(lat2)
	Δλ := degreesToRadians(lon2 - lon1)

	return earthRadiusKm * math.Acos(
		math.Sin(φ1)*math.Sin(φ2)+
			math.Cos(φ1)*math.Cos(φ2)*math.Cos(Δλ),
	)
}

func main() {
	// Example: Distance between New York City and Los Angeles
	nycLat, nycLon := 40.7128, -74.0060
	laLat, laLon := 34.0522, -118.2437

	distance := greatCircleDistance(nycLat, nycLon, laLat, laLon)
	fmt.Printf("Great Circle Distance between NYC and LA: %.2f km\n", distance)
}
