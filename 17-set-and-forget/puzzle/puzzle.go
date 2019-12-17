package main

import (
	"fmt"

	"github.com/chr-ras/advent-of-code-2019/17-set-and-forget/ascii"
	"github.com/chr-ras/advent-of-code-2019/util/aoc"
)

func main() {
	asciiProgram := aoc.ReadIntcode("./ascii_program.txt")

	calibrationResult := ascii.Calibrate(asciiProgram)

	fmt.Printf("Calibration finished with result %d.\n", calibrationResult)
}
