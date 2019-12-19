package main

import (
	"fmt"

	"github.com/chr-ras/advent-of-code-2019/19-tractor-beam/tractorbeam"
	"github.com/chr-ras/advent-of-code-2019/util/aoc"
)

func main() {
	program := aoc.ReadIntcode("./drone_system_program.txt")

	affectedCells := tractorbeam.ScanArea(program, 1000)

	fmt.Printf("Number of cells affected by tractor beam: %d\n", affectedCells)
}
