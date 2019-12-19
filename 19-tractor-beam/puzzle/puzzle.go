package main

import (
	"fmt"

	"github.com/chr-ras/advent-of-code-2019/19-tractor-beam/tractorbeam"
	"github.com/chr-ras/advent-of-code-2019/util/aoc"
)

func main() {
	program := aoc.ReadIntcode("./drone_system_program.txt")

	affectedCells := tractorbeam.ScanArea(program, 50)

	fmt.Printf("Number of cells affected by tractor beam: %d\n", affectedCells)

	squareSize := 100
	closestSquareVector := tractorbeam.FindClosestSquare(program, squareSize)

	fmt.Printf("Closest square in tractor beam of size %d at %v, puzzle result: %d\n", squareSize, closestSquareVector, closestSquareVector.X*10000+closestSquareVector.Y)
}
