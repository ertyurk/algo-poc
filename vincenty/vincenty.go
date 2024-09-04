package main

import (
	"fmt"
	"math"
)

const (
	a                    = 6378137.0         // WGS-84 semi-major axis
	f                    = 1 / 298.257223563 // WGS-84 flattening
	b                    = a * (1 - f)       // semi-minor axis
	maxIterations        = 200
	convergenceThreshold = 1e-12
)

func vincenty(lat1, lon1, lat2, lon2 float64) (float64, error) {
	φ1 := lat1 * math.Pi / 180.0
	λ1 := lon1 * math.Pi / 180.0
	φ2 := lat2 * math.Pi / 180.0
	λ2 := lon2 * math.Pi / 180.0

	L := λ2 - λ1
	tanU1 := (1 - f) * math.Tan(φ1)
	cosU1 := 1 / math.Sqrt(1+tanU1*tanU1)
	sinU1 := tanU1 * cosU1
	tanU2 := (1 - f) * math.Tan(φ2)
	cosU2 := 1 / math.Sqrt(1+tanU2*tanU2)
	sinU2 := tanU2 * cosU2

	λ := L
	var sinλ, cosλ, sinσ, cosσ, cos2σM, σ, C float64

	for i := 0; i < maxIterations; i++ {
		sinλ = math.Sin(λ)
		cosλ = math.Cos(λ)
		sinσ = math.Sqrt(
			(cosU2*sinλ)*(cosU2*sinλ) + (cosU1*sinU2-sinU1*cosU2*cosλ)*(cosU1*sinU2-sinU1*cosU2*cosλ),
		)
		if sinσ == 0 {
			return 0, nil // coincident points
		}
		cosσ = sinU1*sinU2 + cosU1*cosU2*cosλ
		σ = math.Atan2(sinσ, cosσ)
		sinα := cosU1 * cosU2 * sinλ / sinσ
		cos2α := 1 - sinα*sinα
		cos2σM = cosσ - 2*sinU1*sinU2/cos2α
		C = f / 16 * cos2α * (4 + f*(4-3*cos2α))
		λPrev := λ
		λ = L + (1-C)*f*sinα*(σ+C*sinσ*(cos2σM+C*cosσ*(-1+2*cos2σM*cos2σM)))
		if math.Abs(λ-λPrev) < convergenceThreshold {
			u2 := cos2α * (a*a - b*b) / (b * b)
			A := 1 + u2/16384*(4096+u2*(-768+u2*(320-175*u2)))
			B := u2 / 1024 * (256 + u2*(-128+u2*(74-47*u2)))
			Δσ := B * sinσ * (cos2σM + B/4*(cosσ*(-1+2*cos2σM*cos2σM)-B/6*cos2σM*(-3+4*sinσ*sinσ)*(-3+4*cos2σM*cos2σM)))
			s := b * A * (σ - Δσ)
			return s, nil
		}
	}
	return 0, fmt.Errorf("vincenty formula failed to converge")
}

func main() {
	// Example: Distance between New York City and Los Angeles
	nycLat, nycLon := 40.7128, -74.0060
	laLat, laLon := 34.0522, -118.2437

	distance, err := vincenty(nycLat, nycLon, laLat, laLon)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Distance between NYC and LA: %.3f km\n", distance/1000)
	}
}
