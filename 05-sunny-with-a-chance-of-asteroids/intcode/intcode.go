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
	operationPointerModifier := 0
	exit := false
	useModifierAsNewOperationPointer := false
	for i := 0; true; i += operationPointerModifier {
		if useModifierAsNewOperationPointer {
			fmt.Printf("[JMP] %v\n", operationPointerModifier)
			i = operationPointerModifier
		}

		operationPointerModifier, exit, useModifierAsNewOperationPointer = handleOperation(output, i)

		if exit {
			return output
		}
	}

	return nil
}

func handleOperation(program []int, operationPointer int) (operationPointerModifier int, exit, useModifierAsNewOperationPointer bool) {
	operation := fmt.Sprintf("%05d", program[operationPointer])
	instruction := operation[3:5]
	operationPointerModifier = 0
	exit = false
	useModifierAsNewOperationPointer = false

	switch instruction {
	case "01":
		operationPointerModifier = 4
		firstParameter, secondParameter, storeParameter := readThreeParameterValues(program, operationPointer, operation)
		fmt.Printf("[ADD] Operand #1: %v Operand #2 %v Target: %v\n", firstParameter, secondParameter, storeParameter)
		handleAddition(program, firstParameter, secondParameter, storeParameter)
	case "02":
		operationPointerModifier = 4
		firstParameter, secondParameter, storeParameter := readThreeParameterValues(program, operationPointer, operation)
		fmt.Printf("[MUL] Operand #1: %v Operand #2 %v Target: %v\n", firstParameter, secondParameter, storeParameter)
		handleMultiplication(program, firstParameter, secondParameter, storeParameter)
	case "03":
		operationPointerModifier = 2
		storeParameter := readParameterValue(program, operationPointer, operation, 1, true)
		fmt.Printf("[INP] Input: %v Target: %v\n", Input, storeParameter)
		handleInput(program, storeParameter)
	case "04":
		operationPointerModifier = 2
		output := readParameterValue(program, operationPointer, operation, 1, false)
		fmt.Printf("[OUT] %v\n", output)
	case "05":
		operationPointerModifier = 3
		conditionParameter := readParameterValue(program, operationPointer, operation, 1, false)
		jumpParameter := readParameterValue(program, operationPointer, operation, 2, false)
		fmt.Printf("[JPT] Condition: %v Jump address: %v\n", conditionParameter, jumpParameter)

		if conditionParameter != 0 {
			operationPointerModifier = jumpParameter
			useModifierAsNewOperationPointer = true
		}
	case "06":
		operationPointerModifier = 3
		conditionParameter := readParameterValue(program, operationPointer, operation, 1, false)
		jumpParameter := readParameterValue(program, operationPointer, operation, 2, false)
		fmt.Printf("[JPF] Condition: %v Jump address: %v\n", conditionParameter, jumpParameter)

		if conditionParameter == 0 {
			operationPointerModifier = jumpParameter
			useModifierAsNewOperationPointer = true
		}
	case "07":
		operationPointerModifier = 4
		firstParameter, secondParameter, storeParameter := readThreeParameterValues(program, operationPointer, operation)
		fmt.Printf("[LTH] left: %v right: %v target: %v\n", firstParameter, secondParameter, storeParameter)
		handleLessThan(program, firstParameter, secondParameter, storeParameter)
	case "08":
		operationPointerModifier = 4
		firstParameter, secondParameter, storeParameter := readThreeParameterValues(program, operationPointer, operation)
		fmt.Printf("[EQL] left: %v right: %v target: %v\n", firstParameter, secondParameter, storeParameter)
		handleEquals(program, firstParameter, secondParameter, storeParameter)
	case "99":
		operationPointerModifier = 0
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

func handleLessThan(program []int, firstParameter, secondParameter, storeAddress int) {
	if firstParameter < secondParameter {
		program[storeAddress] = 1
	} else {
		program[storeAddress] = 0
	}
}

func handleEquals(program []int, firstParameter, secondParameter, storeAddress int) {
	if firstParameter == secondParameter {
		program[storeAddress] = 1
	} else {
		program[storeAddress] = 0
	}
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
