package robot

import (
	"fmt"
	"math"

	"github.com/chr-ras/advent-of-code-2019/11-space-police/intcode"
	g "github.com/chr-ras/advent-of-code-2019/util/geometry"
	queue "github.com/enriquebris/goconcurrentqueue"
)

// RunPaintRobot simulates the painting robot.
func RunPaintRobot(program []int64) int {
	memoryChannel := make(chan []int64)
	inputQueue := queue.NewFIFO()
	outputQueue := queue.NewFIFO()

	panelChannel := make(chan map[g.Point]color)

	go intcode.ExecuteProgram(program, memoryChannel, inputQueue, outputQueue, 1024)
	go navigateRobot(inputQueue, outputQueue, panelChannel)

	<-memoryChannel

	outputQueue.Enqueue(int64(99))

	panels := <-panelChannel

	prettyPrint(panels)

	return len(panels)
}

func navigateRobot(inputQueue, outputQueue queue.Queue, panelChannel chan map[g.Point]color) {
	paintedPlates := make(map[g.Point]color)

	directionVector := g.Vector{X: 0, Y: -1}
	currentPoint := g.Point{X: 0, Y: 0}
	paintedPlates[currentPoint] = white

	for {
		currentColor := paintedPlates[currentPoint]
		inputQueue.Enqueue(int64(currentColor))

		colorToPaintElement, _ := outputQueue.DequeueOrWaitForNextElement()
		colorToPaint := colorToPaintElement.(int64)
		if colorToPaint == 99 {
			panelChannel <- paintedPlates
			return
		} else if colorToPaint == 0 {
			paintedPlates[currentPoint] = black
		} else {
			paintedPlates[currentPoint] = white
		}

		rotationElement, _ := outputQueue.DequeueOrWaitForNextElement()
		rotation := rotationElement.(int64)
		if rotation == 0 {
			directionVector = directionVector.RotateLeft()
		} else {
			directionVector = directionVector.RotateRight()
		}

		currentPoint = currentPoint.AsVector().Add(directionVector).AsPoint()
	}
}

func prettyPrint(plates map[g.Point]color) {
	minX, minY, maxX, maxY := math.MaxInt32, math.MaxInt32, math.MinInt32, math.MinInt32

	for plate := range plates {
		if plate.X < minX {
			minX = plate.X
		}

		if plate.X > maxX {
			maxX = plate.X
		}

		if plate.Y < minY {
			minY = plate.Y
		}

		if plate.Y > maxY {
			maxY = plate.Y
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

	for plate, color := range plates {
		if color == white {
			output[plate.Y+yOffset][plate.X+xOffset] = "░"
		} else {
			output[plate.Y+yOffset][plate.X+xOffset] = " "
		}
	}

	for row := 0; row < len(output); row++ {
		for col := 0; col < len(output[0]); col++ {
			fmt.Print(output[row][col])
		}

		fmt.Println()
	}
}

type color int64

const (
	black color = iota
	white
)
