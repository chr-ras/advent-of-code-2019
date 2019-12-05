package intcode

import (
	"fmt"
)

// Input used for the input operation
var Input int

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
	operationLength := 0
	exit := false
	for i := 0; true; i += operationLength {
		operationLength, exit = handleOperation(output, i)

		if exit {
			return output
		}
	}

	return nil
}

func handleOperation(program []int, operationPointer int) (operationLength int, exit bool) {
	operation := fmt.Sprintf("%05d", program[operationPointer])
	instruction := operation[3:5]
	operationLength = 0
	exit = false

	switch instruction {
	case "01":
		operationLength = 4
		firstParameter, secondParameter, storeParameter := readThreeParameterValues(program, operationPointer, operation)
		fmt.Printf("[ADD] Operand #1: %v Operand #2 %v Target: %v\n", firstParameter, secondParameter, storeParameter)
		handleAddition(program, firstParameter, secondParameter, storeParameter)
	case "02":
		operationLength = 4
		firstParameter, secondParameter, storeParameter := readThreeParameterValues(program, operationPointer, operation)
		fmt.Printf("[MUL] Operand #1: %v Operand #2 %v Target: %v\n", firstParameter, secondParameter, storeParameter)
		handleMultiplication(program, firstParameter, secondParameter, storeParameter)
	case "03":
		operationLength = 2
		storeParameter := readParameterValue(program, operationPointer, operation, 1, true)
		fmt.Printf("[INP] Target: %v\n", storeParameter)
		handleInput(program, storeParameter)
	case "04":
		operationLength = 2
		output := readParameterValue(program, operationPointer, operation, 1, false)
		fmt.Printf("[OUT] %v\n", output)
	case "99":
		operationLength = 0
		exit = true
	}

	return
}

func handleAddition(program []int, firstAddend, secondAddend, storeAddress int) {
	program[storeAddress] = firstAddend + secondAddend
}

func handleMultiplication(program []int, firstFactor, secondFactor, storeAddress int) {
	program[storeAddress] = firstFactor * secondFactor
}

func handleInput(program []int, storeAddress int) {
	program[storeAddress] = Input
}

func handleOutput(output int) {
	fmt.Printf("[OUT] %v\n", output)
}

func readThreeParameterValues(program []int, operationPointer int, operation string) (firstParameter, secondParameter, storeParameter int) {
	firstParameter = readParameterValue(program, operationPointer, operation, 1, false)
	secondParameter = readParameterValue(program, operationPointer, operation, 2, false)
	storeParameter = readParameterValue(program, operationPointer, operation, 3, true)

	return
}

func readParameterValue(program []int, operationPointer int, operation string, parameterPosition int, isStoreParameter bool) int {
	parameterModeEnd := 4 - parameterPosition
	parameterModeStart := parameterModeEnd - 1

	isParameterPositional := operation[parameterModeStart:parameterModeEnd] == "0"

	parameterValue := program[operationPointer+parameterPosition]
	if !isParameterPositional || isStoreParameter {
		return parameterValue
	}

	return program[parameterValue]
}
