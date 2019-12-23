package intcode2

import (
	"fmt"
)

// ExecuteProgram executes an Intcode program.
// Difference to implementation in "intcode" package: This implementation uses buffered channels as queues.
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
//			9 relative base offset,
//			99 return
//			Parameter modes:
//				0: positional
//				1: immediate (interpret as value)
//				2: relative (relative base + offset from parameter)
// [1]: First operand
// [2]: Second operand
// [3]: Save target index
func ExecuteProgram(name string, program []int64, finalMemory chan []int64, inputQueue, outputQueue chan int64, extraMemorySize int64) {
	// https://github.com/go101/go101/wiki/How-to-perfectly-clone-a-slice%3F
	output := append([]int64(nil), program...)
	output = append(output, make([]int64, extraMemorySize)...)

	operationPointerModifier := int64(0)
	relativeBase := int64(0)
	exit := false
	useModifierAsNewOperationPointer := false

	for i := int64(0); true; i += operationPointerModifier {
		if useModifierAsNewOperationPointer {
			i = operationPointerModifier
		}

		operationPointerModifier, relativeBase, exit, useModifierAsNewOperationPointer = handleOperation(name, output, i, relativeBase, inputQueue, outputQueue)

		if exit {
			finalMemory <- output
			close(outputQueue)
			return
		}
	}

	finalMemory <- nil
}

func handleOperation(name string, program []int64, operationPointer, currentRelativeBase int64, inputQueue, outputQueue chan int64) (operationPointerModifier, relativeBase int64, exit, useModifierAsNewOperationPointer bool) {
	operation := fmt.Sprintf("%05d", program[operationPointer])
	instruction := operation[3:5]
	operationPointerModifier = 0
	relativeBase = currentRelativeBase
	exit = false
	useModifierAsNewOperationPointer = false

	switch instruction {
	case "01":
		operationPointerModifier = 4
		firstParameter, secondParameter, storeParameter := readThreeParameterValues(program, operationPointer, currentRelativeBase, operation)
		handleAddition(program, firstParameter, secondParameter, storeParameter)
	case "02":
		operationPointerModifier = 4
		firstParameter, secondParameter, storeParameter := readThreeParameterValues(program, operationPointer, currentRelativeBase, operation)
		handleMultiplication(program, firstParameter, secondParameter, storeParameter)
	case "03":
		operationPointerModifier = 2

		handleInput(name, program, operationPointer, currentRelativeBase, operation, inputQueue)
	case "04":
		operationPointerModifier = 2

		handleOutput(name, program, operationPointer, currentRelativeBase, operation, outputQueue)
	case "05":
		operationPointerModifier = 3
		conditionParameter := readParameterValue(program, operationPointer, currentRelativeBase, operation, 1, false)
		jumpParameter := readParameterValue(program, operationPointer, currentRelativeBase, operation, 2, false)

		if conditionParameter != 0 {
			operationPointerModifier = jumpParameter
			useModifierAsNewOperationPointer = true
		}
	case "06":
		operationPointerModifier = 3
		conditionParameter := readParameterValue(program, operationPointer, currentRelativeBase, operation, 1, false)
		jumpParameter := readParameterValue(program, operationPointer, currentRelativeBase, operation, 2, false)

		if conditionParameter == 0 {
			operationPointerModifier = jumpParameter
			useModifierAsNewOperationPointer = true
		}
	case "07":
		operationPointerModifier = 4
		firstParameter, secondParameter, storeParameter := readThreeParameterValues(program, operationPointer, currentRelativeBase, operation)
		handleLessThan(program, firstParameter, secondParameter, storeParameter)
	case "08":
		operationPointerModifier = 4
		firstParameter, secondParameter, storeParameter := readThreeParameterValues(program, operationPointer, currentRelativeBase, operation)
		handleEquals(program, firstParameter, secondParameter, storeParameter)
	case "09":
		operationPointerModifier = 2
		offset := readParameterValue(program, operationPointer, currentRelativeBase, operation, 1, false)
		relativeBase += offset
	case "99":
		operationPointerModifier = 0
		exit = true
		fmt.Println("[RET ]")
	default:
		panic(fmt.Errorf("Unexpected opcode %v", operation))
	}

	return
}

func handleAddition(program []int64, firstAddend, secondAddend, storeAddress int64) {
	program[storeAddress] = firstAddend + secondAddend
}

func handleMultiplication(program []int64, firstFactor, secondFactor, storeAddress int64) {
	program[storeAddress] = firstFactor * secondFactor
}

func handleInput(name string, program []int64, operationPointer, relativeBase int64, operation string, inputQueue chan int64) {
	storeParameter := readParameterValue(program, operationPointer, relativeBase, operation, 1, true)
	inputValue := <-inputQueue
	program[storeParameter] = inputValue

	fmt.Printf("[%v] [INP] %d target = %d\n", name, inputValue, storeParameter)
}

func handleOutput(name string, program []int64, operationPointer, relativeBase int64, operation string, outputQueue chan int64) {
	output := readParameterValue(program, operationPointer, relativeBase, operation, 1, false)
	outputQueue <- output

	fmt.Printf("[%v] [OUT] %d\n", name, output)
}

func handleLessThan(program []int64, firstParameter, secondParameter, storeAddress int64) {
	if firstParameter < secondParameter {
		program[storeAddress] = 1
	} else {
		program[storeAddress] = 0
	}
}

func handleEquals(program []int64, firstParameter, secondParameter, storeAddress int64) {
	if firstParameter == secondParameter {
		program[storeAddress] = 1
	} else {
		program[storeAddress] = 0
	}
}

func readThreeParameterValues(program []int64, operationPointer, relativeBase int64, operation string) (firstParameter, secondParameter, storeParameter int64) {
	firstParameter = readParameterValue(program, operationPointer, relativeBase, operation, 1, false)
	secondParameter = readParameterValue(program, operationPointer, relativeBase, operation, 2, false)
	storeParameter = readParameterValue(program, operationPointer, relativeBase, operation, 3, true)

	return
}

func readParameterValue(program []int64, operationPointer, relativeBase int64, operation string, parameterPosition int64, isStoreParameter bool) int64 {
	parameterModeEnd := 4 - parameterPosition
	parameterModeStart := parameterModeEnd - 1

	parameterValue := program[operationPointer+parameterPosition]

	parameterMode := operation[parameterModeStart:parameterModeEnd]

	switch parameterMode {
	case "0":
		if isStoreParameter {
			return parameterValue
		}

		return program[parameterValue]
	case "1":
		return parameterValue
	case "2":
		if isStoreParameter {
			return parameterValue + relativeBase
		}

		return program[relativeBase+parameterValue]

	default:
		panic(fmt.Errorf("Unexpected parameter mode %v", parameterMode))
	}
}
