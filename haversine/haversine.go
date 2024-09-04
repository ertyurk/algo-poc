package main

import (
	"fmt"
	"math"
)

const earthRadiusKm = 6371.0

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	lat1 = degreesToRadians(lat1)
	lon1 = degreesToRadians(lon1)
	lat2 = degreesToRadians(lat2)
	lon2 = degreesToRadians(lon2)

	dlat := lat2 - lat1
	dlon := lon2 - lon1

	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadiusKm * c
}

func main() {
	// Example: Distance between New York City and Los Angeles
	nycLat, nycLon := 40.7128, -74.0060
	laLat, laLon := 34.0522, -118.2437

	distance := haversine(nycLat, nycLon, laLat, laLon)
	fmt.Printf("Distance between NYC and LA: %.2f km\n", distance)
}
