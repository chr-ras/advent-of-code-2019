package repairdroid

import (
	"fmt"
	"math"
	"time"

	"github.com/chr-ras/advent-of-code-2019/util/geometry"
	"github.com/chr-ras/advent-of-code-2019/util/intcode"
	q "github.com/enriquebris/goconcurrentqueue"
	"github.com/gosuri/uilive"
)

// FindShortestWayToOxygenTank controls the repair droid to explore the ship and find the shortest route to the oxygen tank.
func FindShortestWayToOxygenTank(remoteControlProgram []int64) (OxygenStation, map[geometry.Vector]Position) {
	movementQueue := q.NewFIFO()
	resultQueue := q.NewFIFO()
	finalMemory := make(chan []int64)

	go intcode.ExecuteProgram(remoteControlProgram, finalMemory, movementQueue, resultQueue, 1024)

	oxygenStation := make(chan OxygenStation)
	explorationFinished := make(chan map[geometry.Vector]Position)

	go exploreShip(movementQueue, resultQueue, oxygenStation, explorationFinished)

	station := <-oxygenStation
	shipMap := <-explorationFinished

	return station, shipMap
}

func exploreShip(movementQueue, resultQueue q.Queue, oxygenStation chan OxygenStation, explorationFinished chan map[geometry.Vector]Position) {
	startPositionVector := geometry.Vector{X: 0, Y: 0}
	shipMap := make(map[geometry.Vector]Position)
	shipMap[startPositionVector] = Position{Status: StartPosition, AdjacentPositions: []geometry.Vector{}}

	writer := uilive.New()
	writer.Start()

	moveDroid(0, geometry.Vector{}, startPositionVector, shipMap, movementQueue, resultQueue, oxygenStation, writer)

	writer.Stop()

	explorationFinished <- shipMap
}

func moveDroid(currentStepsTaken int64, previousDirection, currentPosition geometry.Vector, shipMap map[geometry.Vector]Position, movementQueue, resultQueue q.Queue, oxygenStation chan OxygenStation, writer *uilive.Writer) {
	northDirectionVector := geometry.Vector{X: 0, Y: -1}
	eastDirectionVector := geometry.Vector{X: 1, Y: 0}
	southDirectionVector := geometry.Vector{X: 0, Y: 1}
	westDirectionVector := geometry.Vector{X: -1, Y: 0}

	movedNorth := goIntoDirection(currentStepsTaken, previousDirection, northDirectionVector, currentPosition, northCommand, shipMap, movementQueue, resultQueue, oxygenStation, writer)
	if movedNorth {
		reverseMove(currentPosition, shipMap, southCommand, movementQueue, resultQueue, writer)
	}

	movedEast := goIntoDirection(currentStepsTaken, previousDirection, eastDirectionVector, currentPosition, eastCommand, shipMap, movementQueue, resultQueue, oxygenStation, writer)
	if movedEast {
		reverseMove(currentPosition, shipMap, westCommand, movementQueue, resultQueue, writer)
	}

	movedSouth := goIntoDirection(currentStepsTaken, previousDirection, southDirectionVector, currentPosition, southCommand, shipMap, movementQueue, resultQueue, oxygenStation, writer)
	if movedSouth {
		reverseMove(currentPosition, shipMap, northCommand, movementQueue, resultQueue, writer)
	}

	movedWest := goIntoDirection(currentStepsTaken, previousDirection, westDirectionVector, currentPosition, westCommand, shipMap, movementQueue, resultQueue, oxygenStation, writer)
	if movedWest {
		reverseMove(currentPosition, shipMap, eastCommand, movementQueue, resultQueue, writer)
	}
}

func goIntoDirection(currentStepsTaken int64, previousDirection, newDirection, currentPosition geometry.Vector, droidCommand int64, shipMap map[geometry.Vector]Position, movementQueue, resultQueue q.Queue, oxygenStation chan OxygenStation, writer *uilive.Writer) bool {
	if previousDirection.ScalarMult(-1) != newDirection {
		newPosition := currentPosition.Add(newDirection)

		if _, alreadyVisited := shipMap[newPosition]; !alreadyVisited {
			movementQueue.Enqueue(droidCommand)
			resultElement, _ := resultQueue.DequeueOrWaitForNextElement()
			result := resultElement.(int64)

			var newPositionAdjacentPositions []geometry.Vector
			if result != HitWall {
				newPositionAdjacentPositions = []geometry.Vector{currentPosition}
			}

			shipMap[newPosition] = Position{Status: result, AdjacentPositions: newPositionAdjacentPositions}

			currentPositionInfo := shipMap[currentPosition]
			currentPositionInfo.AdjacentPositions = append(currentPositionInfo.AdjacentPositions, newPosition)

			shipMap[currentPosition] = currentPositionInfo

			if result == HitWall {
				PrettyPrint(currentPosition, shipMap, writer, 10)
				return false
			}

			currentStepsTaken++

			PrettyPrint(newPosition, shipMap, writer, 10)
			if result == OxygenSystem {
				oxygenStation <- OxygenStation{Distance: currentStepsTaken, Position: newPosition}
			}

			moveDroid(currentStepsTaken, newDirection, newPosition, shipMap, movementQueue, resultQueue, oxygenStation, writer)

			return true
		}
	}

	return false
}

func reverseMove(currentPosition geometry.Vector, shipMap map[geometry.Vector]Position, command int64, movementQueue, resultQueue q.Queue, writer *uilive.Writer) {
	movementQueue.Enqueue(command)
	resultQueue.DequeueOrWaitForNextElement() // ignore result because the result is already in the ship map

	PrettyPrint(currentPosition, shipMap, writer, 10)
}

// PrettyPrint prints the current ship map.
func PrettyPrint(currentPosition geometry.Vector, shipMap map[geometry.Vector]Position, writer *uilive.Writer, sleepTime int) {
	minX, minY, maxX, maxY := math.MaxInt32, math.MaxInt32, math.MinInt32, math.MinInt32

	for position := range shipMap {
		if position.X < minX {
			minX = position.X
		}

		if position.X > maxX {
			maxX = position.X
		}

		if position.Y < minY {
			minY = position.Y
		}

		if position.Y > maxY {
			maxY = position.Y
		}
	}

	output := make([][]string, maxY-minY+1)
	for i := range output {
		output[i] = make([]string, maxX-minX+1)
		for j := range output[i] {
			output[i][j] = " "
		}
	}

	xOffset := 0 - minX
	yOffset := 0 - minY

	for position, positionInfo := range shipMap {
		x := position.X + xOffset
		y := position.Y + yOffset
		if positionInfo.Status == Oxygen {
			output[y][x] = "O"
			continue
		}

		if position.X == currentPosition.X && position.Y == currentPosition.Y {
			output[y][x] = "o"
		} else {
			switch positionInfo.Status {
			case HitWall:
				output[y][x] = "░"
			case Moved:
				output[y][x] = "."
			case OxygenSystem:
				output[y][x] = "X"
			case StartPosition:
				output[y][x] = "S"
			}
		}
	}

	renderedOutput := ""
	for _, row := range output {
		for _, cell := range row {
			renderedOutput += cell
		}

		renderedOutput += "\n"
	}

	fmt.Fprintf(writer, renderedOutput)
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
}

const (
	northCommand = int64(1)
	southCommand = int64(2)
	westCommand  = int64(3)
	eastCommand  = int64(4)

	// HitWall indicates that a position is occupied by a wall
	HitWall = int64(0)
	// Moved indicates that a position is accessible
	Moved = int64(1)
	// OxygenSystem indicates that a position is accessible and the oxygen system is there
	OxygenSystem = int64(2)
	// StartPosition indicates that a position is accessible and that it is the starting position
	StartPosition = int64(3)
	// Oxygen indicates that a position is accessible and that oxygen is available
	Oxygen = int64(4)
)

// OxygenStation defines the position of the oxygen station and its distance from the starting position.
type OxygenStation struct {
	Distance int64
	Position geometry.Vector
}

// Position defines a position on the space ship and includes a statu (wall, ...) as well as the adjacent positions.
type Position struct {
	Status            int64
	AdjacentPositions []geometry.Vector
}
