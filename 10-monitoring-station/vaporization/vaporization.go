package vaporization

import (
	"fmt"
	"sort"

	g "github.com/chr-ras/advent-of-code-2019/util/geometry"
)

// VaporizeAsteroids simulates a laser going clockwise vaporizing the asteroids one by one.
func VaporizeAsteroids(asteroidMap []string, monitoringStation g.Point) {
	asteroidsByAngleMap, angles, totalAsteroids := getAsteroidsByAngle(asteroidMap, monitoringStation)

	sort.Float64s(angles)

	numberOfAsteroidsVaporized := 0

	for true {
		for _, angle := range angles {
			asteroidsForAngle := asteroidsByAngleMap[angle]
			if len(asteroidsForAngle) == 0 {
				continue
			}

			numberOfAsteroidsVaporized++
			fmt.Printf("[%03d] Vaporizing %v\n", numberOfAsteroidsVaporized, asteroidsForAngle[0])

			if len(asteroidsForAngle) == 1 {
				asteroidsByAngleMap[angle] = []g.Point{}
				fmt.Printf("[%03d] All asteroids for angle %v vaporized\n", numberOfAsteroidsVaporized, angle)

				continue
			}

			asteroidsByAngleMap[angle] = asteroidsForAngle[1:len(asteroidsForAngle)]
		}

		if numberOfAsteroidsVaporized == totalAsteroids {
			fmt.Printf("Vaporized all %v asteroids!", totalAsteroids)

			return
		}
	}
}

func getAsteroidsByAngle(asteroidMap []string, monitoringStation g.Point) (map[float64][]g.Point, []float64, int) {
	asteroidsByAngleMap := make(map[float64][]g.Point)
	angles := []float64{}
	totalAsteroids := 0
	mapWidth := len(asteroidMap[0])

	for y := 0; y < len(asteroidMap); y++ {
		for x := 0; x < mapWidth; x++ {
			if x == monitoringStation.X && y == monitoringStation.Y {
				continue
			}

			if asteroidMap[y][x:x+1] == "#" {
				asteroidPoint := g.Point{X: x, Y: y}
				angleToStation := monitoringStation.GetAngle(asteroidPoint)
				asteroidsByAngleMap[angleToStation] = append(asteroidsByAngleMap[angleToStation], asteroidPoint)
				totalAsteroids++
			}
		}
	}

	for angle, asteroidsForAngle := range asteroidsByAngleMap {

		sort.Slice(asteroidsForAngle, func(i, j int) bool {
			firstPoint := asteroidsForAngle[i]
			secondPoint := asteroidsForAngle[j]

			firstVectorLength := g.GetDirectionVector(monitoringStation, firstPoint).Length()
			secondVectorLength := g.GetDirectionVector(monitoringStation, secondPoint).Length()

			return firstVectorLength < secondVectorLength
		})

		angles = append(angles, angle)
	}

	return asteroidsByAngleMap, angles, totalAsteroids
}
