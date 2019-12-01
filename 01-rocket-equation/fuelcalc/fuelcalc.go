package fuelcalc

import "math"

// CalculateFuelForModule determines fuel for a module by its mass.
func CalculateFuelForModule(mass int64) float64 {
	fuelForMass := math.Floor(float64(mass)/float64(3)) - 2.0
	if fuelForMass < 0.0 {
		return 0.0
	}

	return fuelForMass + CalculateFuelForModule(int64(fuelForMass))
}

// CalculateFuelForRocket determnines the fuel for the whole rocket based on the module masses.
func CalculateFuelForRocket(moduleMasses []int64) float64 {
	sum := 0.0
	for _, mass := range moduleMasses {
		sum += CalculateFuelForModule(mass)
	}

	return sum
}
