package intcode

import (
	"fmt"
	"time"

	queue "github.com/enriquebris/goconcurrentqueue"
)

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
//			9 relative base offset,
//			99 return
//			Parameter modes:
//				0: positional
//				1: immediate (interpret as value)
//				2: relative (relative base + offset from parameter)
// [1]: First operand
// [2]: Second operand
// [3]: Save target index
func ExecuteProgram(program []int64, finalMemory chan []int64, inputQueue, outputQueue queue.Queue, extraMemorySize int64) {
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

		operationPointerModifier, relativeBase, exit, useModifierAsNewOperationPointer = handleOperation(output, i, relativeBase, inputQueue, outputQueue)

		if exit {
			finalMemory <- output
			return
		}
	}

	finalMemory <- nil
}

func handleOperation(program []int64, operationPointer, currentRelativeBase int64, inputQueue, outputQueue queue.Queue) (operationPointerModifier, relativeBase int64, exit, useModifierAsNewOperationPointer bool) {
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

		handleInput(program, operationPointer, currentRelativeBase, operation, inputQueue)
	case "04":
		operationPointerModifier = 2

		handleOutput(program, operationPointer, currentRelativeBase, operation, outputQueue)
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

func handleInput(program []int64, operationPointer, relativeBase int64, operation string, inputQueue queue.Queue) {
	storeParameter := readParameterValue(program, operationPointer, relativeBase, operation, 1, true)
	inputElement, _ := inputQueue.DequeueOrWaitForNextElement()
	inputValue, _ := inputElement.(int64)

	program[storeParameter] = inputValue
}

func handleOutput(program []int64, operationPointer, relativeBase int64, operation string, outputQueue queue.Queue) {
	output := readParameterValue(program, operationPointer, relativeBase, operation, 1, false)

	outputQueue.Enqueue(output)

	// There seems to be a race condition in conjunction with the painting robot. Sometimes the outputs get mixed up and thus the robot takes a different path.
	// A short sleeping peroid appears to work around this problem (at least on my machine) though it is not a fix for this issue.
	time.Sleep(5 * time.Millisecond)
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
