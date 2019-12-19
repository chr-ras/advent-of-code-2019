package tractorbeam

import (
	"fmt"

	"github.com/chr-ras/advent-of-code-2019/util/geometry"

	"github.com/chr-ras/advent-of-code-2019/util/intcode"
	q "github.com/enriquebris/goconcurrentqueue"
)

// ScanArea scans the specified area using the drone system and returns the total number of cells affected by the tractor beam.
func ScanArea(droneSystemProgram []int64, gridSize int) int {
	totalAffectedCells := 0

	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			if isInBeam(x, y, droneSystemProgram) {
				totalAffectedCells++
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}

	return totalAffectedCells
}

// FindClosestSquare searches for the closest square of the given size which is contained in the tractor beam.
// Returns the top left coordinate.
func FindClosestSquare(droneSystemProgram []int64, squareSize int) geometry.Vector {
	x := findInitialXOfEdge(squareSize*4, droneSystemProgram)

	for y := squareSize * 4; ; y++ {
		if !isInBeam(x, y, droneSystemProgram) {
			x++

			if !isInBeam(x, y, droneSystemProgram) {
				panic(fmt.Errorf("%d, %d should be in the beam", x, y))
			}
		}

		squareOffset := squareSize - 1

		if isInBeam(x, y-squareOffset, droneSystemProgram) && isInBeam(x+squareOffset, y-squareOffset, droneSystemProgram) {
			return geometry.Vector{X: x, Y: y - squareOffset}
		}
	}
}

func findInitialXOfEdge(y int, droneSystemProgram []int64) int {
	for x := 0; ; x++ {
		if isInBeam(x, y, droneSystemProgram) {
			return x
		}
	}
}

func isInBeam(x, y int, droneSystemProgram []int64) bool {
	inputQueue := q.NewFIFO()
	outputQueue := q.NewFIFO()
	finalMemory := make(chan []int64)

	go intcode.ExecuteProgram(droneSystemProgram, finalMemory, inputQueue, outputQueue, 1024)

	inputQueue.Enqueue(int64(x))
	inputQueue.Enqueue(int64(y))

	resultElement, _ := outputQueue.DequeueOrWaitForNextElement()
	result := resultElement.(int64)

	<-finalMemory

	return result == 1
}
