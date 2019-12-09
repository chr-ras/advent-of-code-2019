package intcode

import (
	"fmt"

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

	fmt.Printf("Start execute program with intcode %v\n", output)
	operationPointerModifier := int64(0)
	relativeBase := int64(0)
	exit := false
	useModifierAsNewOperationPointer := false

	for i := int64(0); true; i += operationPointerModifier {
		if useModifierAsNewOperationPointer {
			fmt.Printf("[JMP ] %v\n", operationPointerModifier)
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
		fmt.Printf("[ADD ] Operand #1: %v Operand #2 %v Target: %v\n", firstParameter, secondParameter, storeParameter)
		handleAddition(program, firstParameter, secondParameter, storeParameter)
	case "02":
		operationPointerModifier = 4
		firstParameter, secondParameter, storeParameter := readThreeParameterValues(program, operationPointer, currentRelativeBase, operation)
		fmt.Printf("[MULT] Operand #1: %v Operand #2 %v Target: %v\n", firstParameter, secondParameter, storeParameter)
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
		fmt.Printf("[JMPT] Condition: %v Jump address: %v\n", conditionParameter, jumpParameter)

		if conditionParameter != 0 {
			operationPointerModifier = jumpParameter
			useModifierAsNewOperationPointer = true
		}
	case "06":
		operationPointerModifier = 3
		conditionParameter := readParameterValue(program, operationPointer, currentRelativeBase, operation, 1, false)
		jumpParameter := readParameterValue(program, operationPointer, currentRelativeBase, operation, 2, false)
		fmt.Printf("[JMPF] Condition: %v Jump address: %v\n", conditionParameter, jumpParameter)

		if conditionParameter == 0 {
			operationPointerModifier = jumpParameter
			useModifierAsNewOperationPointer = true
		}
	case "07":
		operationPointerModifier = 4
		firstParameter, secondParameter, storeParameter := readThreeParameterValues(program, operationPointer, currentRelativeBase, operation)
		fmt.Printf("[LESS] left: %v right: %v target: %v\n", firstParameter, secondParameter, storeParameter)
		handleLessThan(program, firstParameter, secondParameter, storeParameter)
	case "08":
		operationPointerModifier = 4
		firstParameter, secondParameter, storeParameter := readThreeParameterValues(program, operationPointer, currentRelativeBase, operation)
		fmt.Printf("[EQUL] left: %v right: %v target: %v\n", firstParameter, secondParameter, storeParameter)
		handleEquals(program, firstParameter, secondParameter, storeParameter)
	case "09":
		operationPointerModifier = 2
		offset := readParameterValue(program, operationPointer, currentRelativeBase, operation, 1, false)
		fmt.Printf("[BASE] previous %v, parameter %v\n", relativeBase, offset)
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
	fmt.Println("[INP ] Dequeueing or waiting for input")
	inputElement, _ := inputQueue.DequeueOrWaitForNextElement()
	inputValue, _ := inputElement.(int64)

	fmt.Printf("[INP ] Input: %v Target: %v\n", inputValue, storeParameter)

	program[storeParameter] = inputValue
}

func handleOutput(program []int64, operationPointer, relativeBase int64, operation string, outputQueue queue.Queue) {
	output := readParameterValue(program, operationPointer, relativeBase, operation, 1, false)
	fmt.Printf("[OUT ] %v\n", output)

	outputQueue.Enqueue(output)
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

	// if isStoreParameter {
	// 	return parameterValue
	// }

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
