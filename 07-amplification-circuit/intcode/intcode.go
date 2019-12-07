package intcode

import (
	"fmt"
)

// Input used for the input operation
var Input []int
var inputIndex = 0

// ExecuteProgram executes an Intcode program.
// An operation consists of up to 4 integers, the last three being parameters with varying length:
// [0]: Operation + Parameter modes:
//			1 add,
//			2 multiply,
//			3 input,
//			4 output
//			5 jump if true,
//			6 jump if false,
// 			7 less than,
//			8 equals,
//			99 return
//			Parameter modes:
//				0: positional
//				1: immediate (interpret as value)
// [1]: First operand
// [2]: Second operand
// [3]: Save target index
func ExecuteProgram(program []int) (finalMemory, outputValues []int) {
	// https://github.com/go101/go101/wiki/How-to-perfectly-clone-a-slice%3F
	output := append([]int(nil), program...)
	fmt.Printf("Start execute program with intcode %v\n", output)
	operationPointerModifier := 0
	exit := false
	useModifierAsNewOperationPointer := false
	outputValues = []int{}
	inputIndex = 0

	for i := 0; true; i += operationPointerModifier {
		if useModifierAsNewOperationPointer {
			fmt.Printf("[JMP] %v\n", operationPointerModifier)
			i = operationPointerModifier
		}

		operationPointerModifier, exit, useModifierAsNewOperationPointer, outputValues = handleOperation(output, i, outputValues)

		if exit {
			return output, outputValues
		}
	}

	return nil, nil
}

func handleOperation(program []int, operationPointer int, outputValues []int) (operationPointerModifier int, exit, useModifierAsNewOperationPointer bool, outputValuesResult []int) {
	operation := fmt.Sprintf("%05d", program[operationPointer])
	instruction := operation[3:5]
	operationPointerModifier = 0
	exit = false
	useModifierAsNewOperationPointer = false
	outputValuesResult = outputValues

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
		outputValuesResult = append(outputValues, output)
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
	program[storeAddress] = Input[inputIndex]
	inputIndex++
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
