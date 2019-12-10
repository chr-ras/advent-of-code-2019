package monitoring

import (
	"fmt"
	"math"

	c "github.com/chr-ras/advent-of-code-2019/03-crossed-wires/crossedwires"
)

// BestAsteroidForMonitoringStation determines the best (== from where one can see the most asteroids) asteroid on the asteroid map to build a monitoring station on.
func BestAsteroidForMonitoringStation(asteroidMap []string) (asteroid c.Coordinate, visibleAsteroids int) {
	visibleAsteroids = 0

	for row := 0; row < len(asteroidMap); row++ {
		for column := 0; column < len(asteroidMap[row]); column++ {
			if asteroidMap[row][column:column+1] == "#" {
				currentAsteroid := c.Coordinate{X: column, Y: row}
				visibleAsteroidsForCurrentAsteroid := CheckLineOfSight(asteroidMap, currentAsteroid)
				if visibleAsteroidsForCurrentAsteroid > visibleAsteroids {
					visibleAsteroids = visibleAsteroidsForCurrentAsteroid
					asteroid = currentAsteroid
				}
			}
		}
	}

	return
}

// CheckLineOfSight determines the count of visible asteroids from the specified asteroid.
func CheckLineOfSight(asteroidMap []string, asteroidToCheck c.Coordinate) int {
	countOfVisibleAsteroids := 0
	visibilityStatus, mapHeight, mapWidth := prepareOutputSlice(asteroidMap)

	for x := asteroidToCheck.X; x >= 0; x-- {
		for y := asteroidToCheck.Y; y >= 0; y-- {
			updateVisibilityStatus(asteroidMap, c.Coordinate{X: x, Y: y}, asteroidToCheck, visibilityStatus, &countOfVisibleAsteroids, false)
		}

		for y := asteroidToCheck.Y + 1; y < mapHeight; y++ {
			updateVisibilityStatus(asteroidMap, c.Coordinate{X: x, Y: y}, asteroidToCheck, visibilityStatus, &countOfVisibleAsteroids, false)
		}
	}

	for x := asteroidToCheck.X + 1; x < mapWidth; x++ {
		for y := asteroidToCheck.Y; y >= 0; y-- {
			updateVisibilityStatus(asteroidMap, c.Coordinate{X: x, Y: y}, asteroidToCheck, visibilityStatus, &countOfVisibleAsteroids, false)
		}

		for y := asteroidToCheck.Y + 1; y < mapHeight; y++ {
			updateVisibilityStatus(asteroidMap, c.Coordinate{X: x, Y: y}, asteroidToCheck, visibilityStatus, &countOfVisibleAsteroids, false)
		}
	}

	prettyPrint(visibilityStatus)

	return countOfVisibleAsteroids
}

func updateVisibilityStatus(asteroidMap []string, spaceToCheck c.Coordinate, startAsteroid c.Coordinate, visibilityStatus [][]cellStatus, countOfVisibleAsteroids *int, isBlockingCheck bool) {
	if spaceToCheck.X == startAsteroid.X && spaceToCheck.Y == startAsteroid.Y {
		return
	}

	if visibilityStatus[spaceToCheck.Y][spaceToCheck.X] == blockedAsteroid || visibilityStatus[spaceToCheck.Y][spaceToCheck.X] == blockedSpace {
		return
	}

	mapPosition := asteroidMap[spaceToCheck.Y][spaceToCheck.X : spaceToCheck.X+1]
	positionIsAsteroid := false
	switch mapPosition {
	case ".":
		if isBlockingCheck {
			visibilityStatus[spaceToCheck.Y][spaceToCheck.X] = blockedSpace
		} else {
			visibilityStatus[spaceToCheck.Y][spaceToCheck.X] = visibleSpace
			return
		}
	case "#":
		positionIsAsteroid = true
		if isBlockingCheck {
			visibilityStatus[spaceToCheck.Y][spaceToCheck.X] = blockedAsteroid
		} else {
			visibilityStatus[spaceToCheck.Y][spaceToCheck.X] = visibleAsteroid
			*countOfVisibleAsteroids++
		}
	default:
		panic(fmt.Errorf("Unexpected mapPosition %v", mapPosition))
	}

	vectorX := spaceToCheck.X - startAsteroid.X
	vectorY := spaceToCheck.Y - startAsteroid.Y
	xYGreatestCommonDivisior := greatestCommonDivisor(vectorX, vectorY)
	newSpaceToCheck := c.Coordinate{X: spaceToCheck.X + vectorX/xYGreatestCommonDivisior, Y: spaceToCheck.Y + vectorY/xYGreatestCommonDivisior}
	if newSpaceToCheck.Y < 0 || newSpaceToCheck.Y >= len(asteroidMap) || newSpaceToCheck.X < 0 || newSpaceToCheck.X >= len(asteroidMap[0]) {
		// The new position is out of the map (expected to happen)
		return
	}

	updateVisibilityStatus(asteroidMap, newSpaceToCheck, spaceToCheck, visibilityStatus, countOfVisibleAsteroids, isBlockingCheck || positionIsAsteroid)
}

func prettyPrint(visibilityStatus [][]cellStatus) {
	for row := 0; row < len(visibilityStatus); row++ {
		for column := 0; column < len(visibilityStatus[row]); column++ {
			switch visibilityStatus[row][column] {
			case unvisited:
				fmt.Print("?")
			case visibleSpace:
				fmt.Print(".")
			case visibleAsteroid:
				fmt.Print("X")
			case blockedSpace:
				fmt.Print("░")
			case blockedAsteroid:
				fmt.Print("x")
			default:
				fmt.Print("_")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func prepareOutputSlice(asteroidMap []string) (outputSlice [][]cellStatus, asteroidMapHeight, asteroidMapWidth int) {
	asteroidMapHeight = len(asteroidMap)
	asteroidMapWidth = len(asteroidMap[0])

	outputSlice = make([][]cellStatus, asteroidMapHeight)
	for i := 0; i < asteroidMapHeight; i++ {
		outputSlice[i] = make([]cellStatus, asteroidMapWidth)
	}

	return
}

func greatestCommonDivisor(a, b int) int {
	a = int(math.Abs(float64(a)))
	b = int(math.Abs(float64(b)))
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

type cellStatus int

const (
	unvisited cellStatus = iota
	visibleSpace
	visibleAsteroid
	blockedSpace
	blockedAsteroid
)
