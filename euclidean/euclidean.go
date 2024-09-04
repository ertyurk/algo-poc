package main

import (
	"fmt"
	"math"
)

// EuclideanDistance2D calculates the Euclidean distance between two points in 2D space
func EuclideanDistance2D(x1, y1, x2, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(dx*dx + dy*dy)
}

// EuclideanDistance3D calculates the Euclidean distance between two points in 3D space
func EuclideanDistance3D(x1, y1, z1, x2, y2, z2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	dz := z2 - z1
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// approximateEarthDistance calculates an approximation of the distance between two points on Earth
// Note: This is a rough approximation and doesn't account for Earth's curvature
func approximateEarthDistance(lat1, lon1, lat2, lon2 float64) float64 {
	// Convert latitude and longitude to x, y coordinates
	// We're using a simple approximation: 1 degree of latitude = 111 km
	// Longitude scale varies with latitude, so we adjust for the average latitude
	const kmPerDegreeLat = 111.0
	avgLat := (lat1 + lat2) / 2
	kmPerDegreeLon := math.Cos(avgLat*math.Pi/180) * kmPerDegreeLat

	x1 := lon1 * kmPerDegreeLon
	y1 := lat1 * kmPerDegreeLat
	x2 := lon2 * kmPerDegreeLon
	y2 := lat2 * kmPerDegreeLat

	return EuclideanDistance2D(x1, y1, x2, y2)
}

func main() {
	// Example: Distance between two points in 2D space
	x1, y1 := 0.0, 0.0
	x2, y2 := 3.0, 4.0
	distance2D := EuclideanDistance2D(x1, y1, x2, y2)
	fmt.Printf("2D Euclidean distance: %.2f units\n", distance2D)

	// Example: Distance between two points in 3D space
	x3, y3, z3 := 1.0, 2.0, 3.0
	x4, y4, z4 := 4.0, 5.0, 6.0
	distance3D := EuclideanDistance3D(x3, y3, z3, x4, y4, z4)
	fmt.Printf("3D Euclidean distance: %.2f units\n", distance3D)

	// Example: Approximate distance between New York City and Los Angeles
	nycLat, nycLon := 40.7128, -74.0060
	laLat, laLon := 34.0522, -118.2437
	approxDistance := approximateEarthDistance(nycLat, nycLon, laLat, laLon)
	fmt.Printf("Approximate distance between NYC and LA: %.2f km\n", approxDistance)
}
