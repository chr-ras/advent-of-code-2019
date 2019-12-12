package main

import (
	"fmt"

	"github.com/chr-ras/advent-of-code-2019/12-the-n-body-problem/nbody"
	v "github.com/chr-ras/advent-of-code-2019/util/geometry/vector3"
)

func main() {
	moonPositions := []v.Vector3{
		v.Vector3{X: -10, Y: -10, Z: -13},
		v.Vector3{X: 5, Y: 5, Z: -9},
		v.Vector3{X: 3, Y: 8, Z: -16},
		v.Vector3{X: 1, Y: 3, Z: -3},
	}

	totalEnergy := nbody.SimulateJupiterMoons(moonPositions, 1000)

	fmt.Printf("Total energy: %10.0f\n", totalEnergy)
}
