package nbody

import (
	"fmt"
	"math"

	v "github.com/chr-ras/advent-of-code-2019/util/geometry/vector3"
)

// SimulateJupiterMoons runs the moon movement simulation for n steps and returns the total energy in the system.
func SimulateJupiterMoons(positions []v.Vector3, steps int) float64 {
	moons := prepareMoons(positions)

	for step := 0; step < steps; step++ {
		moons = applyGravitation(moons)
		moons = applyVelocity(moons)
		prettyPrint(moons)
	}

	return calcTotalEnergy(moons)
}

func prepareMoons(positions []v.Vector3) []moon {
	moons := []moon{}
	for _, position := range positions {
		moons = append(moons, moon{position: position, velocity: v.Vector3{X: 0, Y: 0, Z: 0}})
	}

	return moons
}

func applyGravitation(moons []moon) []moon {
	for i := 0; i < len(moons)-1; i++ {
		for j := i + 1; j < len(moons); j++ {
			if moons[i].position.X < moons[j].position.X {
				moons[i].velocity.X++
				moons[j].velocity.X--
			} else if moons[j].position.X < moons[i].position.X {
				moons[j].velocity.X++
				moons[i].velocity.X--
			}

			if moons[i].position.Y < moons[j].position.Y {
				moons[i].velocity.Y++
				moons[j].velocity.Y--
			} else if moons[j].position.Y < moons[i].position.Y {
				moons[j].velocity.Y++
				moons[i].velocity.Y--
			}

			if moons[i].position.Z < moons[j].position.Z {
				moons[i].velocity.Z++
				moons[j].velocity.Z--
			} else if moons[j].position.Z < moons[i].position.Z {
				moons[j].velocity.Z++
				moons[i].velocity.Z--
			}
		}
	}

	return moons
}

func applyVelocity(moons []moon) []moon {
	for i := range moons {
		moons[i].position = moons[i].position.Add(moons[i].velocity)
	}

	return moons
}

func prettyPrint(moons []moon) {
	for _, moon := range moons {
		fmt.Printf("pos=<x=%3.0f, y=%3.0f, z=%3.0f>, vel=<x=%3.0f, y=%3.0f, z=%3.0f>\n", moon.position.X, moon.position.Y, moon.position.Z, moon.velocity.X, moon.velocity.Y, moon.velocity.Z)
	}

	fmt.Println()
}

func calcTotalEnergy(moons []moon) float64 {
	totalEnergy := 0.0
	for _, moon := range moons {
		moonPotentialEnergy := math.Abs(moon.position.X) + math.Abs(moon.position.Y) + math.Abs(moon.position.Z)
		moonKineticEnergy := math.Abs(moon.velocity.X) + math.Abs(moon.velocity.Y) + math.Abs(moon.velocity.Z)
		moonTotalEnergy := moonPotentialEnergy * moonKineticEnergy

		totalEnergy += moonTotalEnergy
	}

	return totalEnergy
}

type moon struct {
	position v.Vector3
	velocity v.Vector3
}
