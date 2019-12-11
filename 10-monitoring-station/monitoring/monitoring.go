package monitoring

import (
	"fmt"

	"github.com/chr-ras/advent-of-code-2019/util/calc"
	g "github.com/chr-ras/advent-of-code-2019/util/geometry"
)

// BestAsteroidForMonitoringStation determines the best (== from where one can see the most asteroids) asteroid on the asteroid map to build a monitoring station on.
func BestAsteroidForMonitoringStation(asteroidMap []string) (asteroid g.Point, visibleAsteroids int) {
	visibleAsteroids = 0

	for row := 0; row < len(asteroidMap); row++ {
		for column := 0; column < len(asteroidMap[row]); column++ {
			if asteroidMap[row][column:column+1] == "#" {
				currentAsteroid := g.Point{X: column, Y: row}
				visibleAsteroidsForCurrentAsteroid, _ := CheckLineOfSight(asteroidMap, currentAsteroid)
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
func CheckLineOfSight(asteroidMap []string, asteroidToCheck g.Point) (countOfVisibleAsteroids int, visibilityStatus [][]CellStatus) {
	countOfVisibleAsteroids = 0
	visibilityStatus, mapHeight, mapWidth := prepareOutputSlice(asteroidMap)

	for x := asteroidToCheck.X; x >= 0; x-- {
		for y := asteroidToCheck.Y; y >= 0; y-- {
			updateVisibilityStatus(asteroidMap, g.Point{X: x, Y: y}, asteroidToCheck, visibilityStatus, &countOfVisibleAsteroids, false)
		}

		for y := asteroidToCheck.Y + 1; y < mapHeight; y++ {
			updateVisibilityStatus(asteroidMap, g.Point{X: x, Y: y}, asteroidToCheck, visibilityStatus, &countOfVisibleAsteroids, false)
		}
	}

	for x := asteroidToCheck.X + 1; x < mapWidth; x++ {
		for y := asteroidToCheck.Y; y >= 0; y-- {
			updateVisibilityStatus(asteroidMap, g.Point{X: x, Y: y}, asteroidToCheck, visibilityStatus, &countOfVisibleAsteroids, false)
		}

		for y := asteroidToCheck.Y + 1; y < mapHeight; y++ {
			updateVisibilityStatus(asteroidMap, g.Point{X: x, Y: y}, asteroidToCheck, visibilityStatus, &countOfVisibleAsteroids, false)
		}
	}

	prettyPrint(visibilityStatus)

	return
}

func updateVisibilityStatus(asteroidMap []string, spaceToCheck, startAsteroid g.Point, visibilityStatus [][]CellStatus, countOfVisibleAsteroids *int, isBlockingCheck bool) {
	if spaceToCheck.X == startAsteroid.X && spaceToCheck.Y == startAsteroid.Y {
		return
	}

	if visibilityStatus[spaceToCheck.Y][spaceToCheck.X] == BlockedAsteroid || visibilityStatus[spaceToCheck.Y][spaceToCheck.X] == BlockedSpace {
		return
	}

	mapPosition := asteroidMap[spaceToCheck.Y][spaceToCheck.X : spaceToCheck.X+1]
	positionIsAsteroid := false
	switch mapPosition {
	case ".":
		if isBlockingCheck {
			visibilityStatus[spaceToCheck.Y][spaceToCheck.X] = BlockedSpace
		} else {
			visibilityStatus[spaceToCheck.Y][spaceToCheck.X] = VisibleSpace
			return
		}
	case "#":
		positionIsAsteroid = true
		if isBlockingCheck {
			visibilityStatus[spaceToCheck.Y][spaceToCheck.X] = BlockedAsteroid
		} else {
			visibilityStatus[spaceToCheck.Y][spaceToCheck.X] = VisibleAsteroid
			*countOfVisibleAsteroids++
		}
	default:
		panic(fmt.Errorf("Unexpected mapPosition %v", mapPosition))
	}

	vectorX := spaceToCheck.X - startAsteroid.X
	vectorY := spaceToCheck.Y - startAsteroid.Y
	xYGreatestCommonDivisior := calc.GreatestCommonDivisor(vectorX, vectorY)
	newSpaceToCheck := g.Point{X: spaceToCheck.X + vectorX/xYGreatestCommonDivisior, Y: spaceToCheck.Y + vectorY/xYGreatestCommonDivisior}
	if newSpaceToCheck.Y < 0 || newSpaceToCheck.Y >= len(asteroidMap) || newSpaceToCheck.X < 0 || newSpaceToCheck.X >= len(asteroidMap[0]) {
		// The new position is out of the map (expected to happen)
		return
	}

	updateVisibilityStatus(asteroidMap, newSpaceToCheck, spaceToCheck, visibilityStatus, countOfVisibleAsteroids, isBlockingCheck || positionIsAsteroid)
}

func prettyPrint(visibilityStatus [][]CellStatus) {
	for row := 0; row < len(visibilityStatus); row++ {
		for column := 0; column < len(visibilityStatus[row]); column++ {
			switch visibilityStatus[row][column] {
			case Unvisited:
				fmt.Print("?")
			case VisibleSpace:
				fmt.Print(".")
			case VisibleAsteroid:
				fmt.Print("X")
			case BlockedSpace:
				fmt.Print("░")
			case BlockedAsteroid:
				fmt.Print("x")
			default:
				fmt.Print("_")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func prepareOutputSlice(asteroidMap []string) (outputSlice [][]CellStatus, asteroidMapHeight, asteroidMapWidth int) {
	asteroidMapHeight = len(asteroidMap)
	asteroidMapWidth = len(asteroidMap[0])

	outputSlice = make([][]CellStatus, asteroidMapHeight)
	for i := 0; i < asteroidMapHeight; i++ {
		outputSlice[i] = make([]CellStatus, asteroidMapWidth)
	}

	return
}

// CellStatus enum
type CellStatus int

const (
	// Unvisited means: The cell has not been visited
	Unvisited CellStatus = iota
	// VisibleSpace means: The cell contains no asteroid and is visible from the specified asteroid
	VisibleSpace
	// VisibleAsteroid means: The cell contains an asteroid and is visible from the specified asteroid
	VisibleAsteroid
	// BlockedSpace means: The cell contains no asteroid and is not visible from the specified asteroid
	BlockedSpace
	// BlockedAsteroid means: The cell contains an asteroid and is not visible from the specified asteroid
	BlockedAsteroid
)
