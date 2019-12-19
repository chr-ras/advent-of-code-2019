package tractorbeam

import (
	"fmt"

	"github.com/chr-ras/advent-of-code-2019/util/intcode"
	q "github.com/enriquebris/goconcurrentqueue"
)

// ScanArea scans the specified area using the drone system and returns the total number of cells affected by the tractor beam.
func ScanArea(droneSystemProgram []int64, gridSize int) int {
	totalAffectedCells := 0

	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			inputQueue := q.NewFIFO()
			outputQueue := q.NewFIFO()
			finalMemory := make(chan []int64)

			go intcode.ExecuteProgram(droneSystemProgram, finalMemory, inputQueue, outputQueue, 1024)

			inputQueue.Enqueue(int64(x))
			inputQueue.Enqueue(int64(y))

			resultElement, _ := outputQueue.DequeueOrWaitForNextElement()
			result := resultElement.(int64)

			if result == 1 {
				totalAffectedCells++
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

			<-finalMemory
		}

		fmt.Println()
	}

	return totalAffectedCells
}
