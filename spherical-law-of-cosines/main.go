package main

import (
	"fmt"
	"math"
)

const earthRadiusKm = 6371.0

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func sphericalLawOfCosines(lat1, lon1, lat2, lon2 float64) float64 {
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

	distance := sphericalLawOfCosines(nycLat, nycLon, laLat, laLon)
	fmt.Printf("Distance between NYC and LA (Spherical Law of Cosines): %.2f km\n", distance)

	// Example: Small distance calculation
	lat1, lon1 := 40.7128, -74.0060 // NYC
	lat2, lon2 := 40.7130, -74.0062 // Very close to NYC
	smallDistance := sphericalLawOfCosines(lat1, lon1, lat2, lon2)
	fmt.Printf("Small distance calculation: %.5f km\n", smallDistance)
}
