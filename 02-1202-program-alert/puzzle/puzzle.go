package main

import (
	"fmt"

	"github.com/chr-ras/advent-of-code-2019/02-1202-program-alert/intcode"
)

func main() {
	initialMemoryState := []int{
		1, 0, 0, 3,
		1, 1, 2, 3,
		1, 3, 4, 3,
		1, 5, 0, 3,
		2, 13, 1, 19,
		1, 19, 9, 23,
		1, 5, 23, 27,
		1, 27, 9, 31,
		1, 6, 31, 35,
		2, 35, 9, 39,
		1, 39, 6, 43,
		2, 9, 43, 47,
		1, 47, 6, 51,
		2, 51, 9, 55,
		1, 5, 55, 59,
		2, 59, 6, 63,
		1, 9, 63, 67,
		1, 67, 10, 71,
		1, 71, 13, 75,
		2, 13, 75, 79,
		1, 6, 79, 83,
		2, 9, 83, 87,
		1, 87, 6, 91,
		2, 10, 91, 95,
		2, 13, 95, 99,
		1, 9, 99, 103,
		1, 5, 103, 107,
		2, 9, 107, 111,
		1, 111, 5, 115,
		1, 115, 5, 119,
		1, 10, 119, 123,
		1, 13, 123, 127,
		1, 2, 127, 131,
		1, 131, 13, 0,
		99, 2, 14, 0,
		0,
	}

	// Restore gravity assist program to the 1202 program alarm state
	program := append([]int(nil), initialMemoryState...)
	program[1] = 12
	program[2] = 2

	output := intcode.ExecuteProgram(program)

	fmt.Printf("Intcode program output: \n%v\n", output)

	verb, noun := intcode.DetermineGravityAssistParameters(program, 19690720)
	result := 100*noun + verb
	fmt.Printf("Gravity assist parameters for output 19690720: Verb %v, Noun %v\nResult 100 * noun + verb = %v", verb, noun, result)
}
