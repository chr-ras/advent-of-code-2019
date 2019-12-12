package robot

import (
	"fmt"

	"github.com/chr-ras/advent-of-code-2019/11-space-police/intcode"
	g "github.com/chr-ras/advent-of-code-2019/util/geometry"
	queue "github.com/enriquebris/goconcurrentqueue"
)

func RunPaintRobot(program []int64) int {
	memoryChannel := make(chan []int64)
	inputQueue := queue.NewFIFO()
	outputQueue := queue.NewFIFO()

	navigationChannel := make(chan map[g.Point]color)

	go intcode.ExecuteProgram(program, memoryChannel, inputQueue, outputQueue, 1024)
	go navigateRobot(inputQueue, outputQueue, navigationChannel)

	<-memoryChannel

	outputQueue.Enqueue(int64(99))

	panels := <-navigationChannel

	for panelPoint := range panels {
		fmt.Println(panelPoint)
	}

	return len(panels)
}

func navigateRobot(inputQueue, outputQueue queue.Queue, navigationChannel chan map[g.Point]color) {
	paintedPlates := make(map[g.Point]color)

	directionVector := g.Vector{X: 0, Y: -1}
	currentPoint := g.Point{X: 0, Y: 0}

	for {
		currentColor := paintedPlates[currentPoint]
		inputQueue.Enqueue(int64(currentColor))

		colorToPaintElement, _ := outputQueue.DequeueOrWaitForNextElement()
		colorToPaint := colorToPaintElement.(int64)
		if colorToPaint == 99 {
			navigationChannel <- paintedPlates
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

type color int64

const (
	black color = iota
	white
)
