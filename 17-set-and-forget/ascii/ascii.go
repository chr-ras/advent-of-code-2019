package ascii

import (
	"fmt"

	"github.com/chr-ras/advent-of-code-2019/util/geometry"

	"github.com/chr-ras/advent-of-code-2019/util/intcode"
	q "github.com/enriquebris/goconcurrentqueue"
)

// Calibrate runs the ASCII program and returns the sum of the alignment parameters.
func Calibrate(program []int64) int {
	finalMemory := make(chan []int64)
	inputQueue := q.NewFIFO()
	cameraQueue := q.NewFIFO()

	imageChannel := make(chan []int64)

	go intcode.ExecuteProgram(program, finalMemory, inputQueue, cameraQueue, 4096)
	go readCameraImage(cameraQueue, imageChannel)

	<-finalMemory

	cameraQueue.Enqueue(int64(-1))

	rawImage := <-imageChannel

	pathElements := extractPathElements(rawImage)

	prettyPrint(rawImage)

	return calcAlignmentParameters(pathElements)
}

func readCameraImage(cameraQueue q.Queue, imageChannel chan []int64) {
	image := []int64{}

	for {
		pixelElement, _ := cameraQueue.DequeueOrWaitForNextElement()
		pixel := pixelElement.(int64)

		if pixel == -1 {
			imageChannel <- image

			return
		}

		image = append(image, pixel)
	}
}

func extractPathElements(rawImage []int64) map[geometry.Vector][]geometry.Vector {
	image := [][]int64{}

	currentLine := []int64{}
	for _, pixel := range rawImage {
		if pixel == 10 {
			if len(currentLine) == 0 {
				continue
			}

			image = append(image, currentLine)
			currentLine = []int64{}

			continue
		}

		currentLine = append(currentLine, pixel)
	}

	pathElements := make(map[geometry.Vector][]geometry.Vector)

	northDirectionVector := geometry.Vector{X: 0, Y: -1}
	southDirectionVector := geometry.Vector{X: 0, Y: 1}
	westDirectionVector := geometry.Vector{X: -1, Y: 0}
	eastDirectionVector := geometry.Vector{X: 1, Y: 0}

	for y := 0; y < len(image); y++ {
		for x := 0; x < len(image[y]); x++ {
			elementVector := geometry.Vector{X: x, Y: y}

			findNeighbor(image, northDirectionVector, elementVector, pathElements)
			findNeighbor(image, southDirectionVector, elementVector, pathElements)
			findNeighbor(image, westDirectionVector, elementVector, pathElements)
			findNeighbor(image, eastDirectionVector, elementVector, pathElements)
		}
	}

	return pathElements
}

func findNeighbor(image [][]int64, direction geometry.Vector, currentElement geometry.Vector, pathElements map[geometry.Vector][]geometry.Vector) {
	neighborVector := currentElement.Add(direction)
	if neighborVector.Y >= 0 && neighborVector.Y < len(image) && neighborVector.X >= 0 && neighborVector.X < len(image[0]) {
		if neighbor := image[currentElement.Y][currentElement.X]; isScaffold(neighbor) {
			adjacentElements := pathElements[neighborVector]
			adjacentElements = append(adjacentElements, currentElement)
			pathElements[neighborVector] = adjacentElements
		}
	}
}

func isScaffold(value int64) bool {
	// Treat the robot as a scaffold for calibrating purposes.
	return value == 35 || value == 94
}

func calcAlignmentParameters(pathElements map[geometry.Vector][]geometry.Vector) int {
	result := 0

	for point, adjacentPoints := range pathElements {
		if len(adjacentPoints) != 4 {
			continue
		}

		result += (point.X * point.Y)
	}

	return result
}

func prettyPrint(image []int64) {
	output := ""

	for _, pixel := range image {
		switch pixel {
		case 10:
			output += "\n"
		case 35:
			output += "#"
		case 46:
			output += "."
		case 94:
			output += "^"
		}
	}

	fmt.Println(output)
}
