package intcode

import "fmt"

// ExecuteProgram executes an Intcode program.
// A operation consists of 4 integers:
// [0]: Operation. 1 add, 2 multiply, 99 return
// [1]: First operand index
// [2]: Second operand index
// [3]: Save target index
func ExecuteProgram(program []int) []int {
	// https://github.com/go101/go101/wiki/How-to-perfectly-clone-a-slice%3F
	output := append([]int(nil), program...)
	fmt.Printf("Start execute program with intcode %v\n", output)
	for i := 0; true; i += 4 {
		switch output[i] {
		case 1:
			fmt.Printf("Operation: %v Operand #1: %v Operand #2 %v Target: %v\n", output[i], output[i+1], output[i+2], output[i+3])
			output[output[i+3]] = output[output[i+1]] + output[output[i+2]]
		case 2:
			fmt.Printf("Operation: %v Operand #1: %v Operand #2 %v Target: %v\n", output[i], output[i+1], output[i+2], output[i+3])
			output[output[i+3]] = output[output[i+1]] * output[output[i+2]]
		case 99:
			fmt.Printf("RETURN\n")
			return output
		default:
			return nil
		}
	}

	return nil
}
