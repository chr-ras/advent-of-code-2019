package intcode

import (
	"reflect"
	"testing"

	queue "github.com/enriquebris/goconcurrentqueue"
)

func TestExecuteProgram(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{
			[]int{
				1, 9, 10, 3,
				2, 3, 11, 0,
				99, 30, 40, 50,
			},
			[]int{
				3500, 9, 10, 70,
				2, 3, 11, 0,
				99, 30, 40, 50,
			},
		},
		{
			[]int{
				1, 0, 0, 0,
				99,
			},
			[]int{
				2, 0, 0, 0,
				99,
			},
		},
		{
			[]int{
				2, 3, 0, 3,
				99,
			},
			[]int{
				2, 3, 0, 6,
				99,
			},
		},
		{
			[]int{
				2, 4, 4, 5,
				99, 0,
			},
			[]int{
				2, 4, 4, 5,
				99, 9801,
			},
		},
		{
			[]int{
				1, 1, 1, 4,
				99, 5, 6, 0,
				99,
			},
			[]int{
				30, 1, 1, 4,
				2, 5, 6, 0,
				99,
			},
		},
	}

	for _, c := range cases {
		memoryChannel := make(chan []int)
		inputQueue := queue.NewFIFO()
		outputQueue := queue.NewFIFO()

		go ExecuteProgram(c.input, memoryChannel, inputQueue, outputQueue)

		actual := <-memoryChannel

		if !sliceEquals(actual, c.expected) {
			t.Errorf("ExecuteProgram(%v) == %v, expected %v", c.input, actual, c.expected)
		}
	}
}

func TestExecuteProgramInput(t *testing.T) {
	cases := []struct {
		inputValue   []int
		inputProgram []int
		expected     []int
	}{
		{
			[]int{1},
			[]int{
				3, 0, 99,
			},
			[]int{
				1, 0, 99,
			},
		},
	}

	for _, c := range cases {
		memoryChannel := make(chan []int)
		inputQueue := queue.NewFIFO()
		outputQueue := queue.NewFIFO()

		for _, input := range c.inputValue {
			inputQueue.Enqueue(input)
		}

		go ExecuteProgram(c.inputProgram, memoryChannel, inputQueue, outputQueue)

		actual := <-memoryChannel

		if !sliceEquals(actual, c.expected) {
			t.Errorf("ExecuteProgram(%v) == %v with input %v, expected %v", c.inputProgram, actual, c.inputValue, c.expected)
		}
	}
}

func TestExecuteProgramOutput(t *testing.T) {
	input := []int{
		4, 0, 4, 1, 99,
	}

	expectedProgramState := []int{
		4, 0, 4, 1, 99,
	}

	expectedOutput := []int{4, 0}

	memoryChannel := make(chan []int)
	inputQueue := queue.NewFIFO()
	outputQueue := queue.NewFIFO()

	go ExecuteProgram(input, memoryChannel, inputQueue, outputQueue)

	actualState := <-memoryChannel
	actualOutput := fetchAllElementsFromQueue(outputQueue)

	if !reflect.DeepEqual(actualState, expectedProgramState) || !reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Errorf("ExecuteProgram(%v) == %v, %v, expected %v, %v", input, actualState, actualOutput, expectedProgramState, expectedOutput)
	}
}

func TestExecuteProgramJumpIfTrue(t *testing.T) {
	cases := []struct {
		inputProgram []int
		expected     []int
	}{
		{
			[]int{
				1105, 2, 3, 99,
			},
			[]int{
				1105, 2, 3, 99,
			},
		},
		{
			[]int{
				1105, 0, 4, 99, 0,
			},
			[]int{
				1105, 0, 4, 99, 0,
			},
		},
	}

	for _, c := range cases {
		memoryChannel := make(chan []int)
		inputQueue := queue.NewFIFO()
		outputQueue := queue.NewFIFO()

		go ExecuteProgram(c.inputProgram, memoryChannel, inputQueue, outputQueue)

		actual := <-memoryChannel

		if !sliceEquals(actual, c.expected) {
			t.Errorf("ExecuteProgram(%v) == %v, expected %v", c.inputProgram, actual, c.expected)
		}
	}
}

func TestExecuteProgramJumpIfFalse(t *testing.T) {
	cases := []struct {
		inputProgram []int
		expected     []int
	}{
		{
			[]int{
				1106, 0, 3, 99,
			},
			[]int{
				1106, 0, 3, 99,
			},
		},
		{
			[]int{
				1106, 2, 4, 99, 0,
			},
			[]int{
				1106, 2, 4, 99, 0,
			},
		},
	}

	for _, c := range cases {
		memoryChannel := make(chan []int)
		inputQueue := queue.NewFIFO()
		outputQueue := queue.NewFIFO()

		go ExecuteProgram(c.inputProgram, memoryChannel, inputQueue, outputQueue)

		actual := <-memoryChannel

		if !sliceEquals(actual, c.expected) {
			t.Errorf("ExecuteProgram(%v) == %v, expected %v", c.inputProgram, actual, c.expected)
		}
	}
}

func TestExecuteProgramLessThan(t *testing.T) {
	cases := []struct {
		inputProgram []int
		expected     []int
	}{
		{
			[]int{
				1107, 0, 3, 0, 99,
			},
			[]int{
				1, 0, 3, 0, 99,
			},
		},
		{
			[]int{
				1107, 3, 0, 0, 99,
			},
			[]int{
				0, 3, 0, 0, 99,
			},
		},
	}

	for _, c := range cases {
		memoryChannel := make(chan []int)
		inputQueue := queue.NewFIFO()
		outputQueue := queue.NewFIFO()

		go ExecuteProgram(c.inputProgram, memoryChannel, inputQueue, outputQueue)

		actual := <-memoryChannel

		if !sliceEquals(actual, c.expected) {
			t.Errorf("ExecuteProgram(%v) == %v, expected %v", c.inputProgram, actual, c.expected)
		}
	}
}

func TestExecuteProgramEquals(t *testing.T) {
	cases := []struct {
		inputProgram []int
		expected     []int
	}{
		{
			[]int{
				1108, 3, 3, 0, 99,
			},
			[]int{
				1, 3, 3, 0, 99,
			},
		},
		{
			[]int{
				1108, 3, 0, 0, 99,
			},
			[]int{
				0, 3, 0, 0, 99,
			},
		},
	}

	for _, c := range cases {
		memoryChannel := make(chan []int)
		inputQueue := queue.NewFIFO()
		outputQueue := queue.NewFIFO()

		go ExecuteProgram(c.inputProgram, memoryChannel, inputQueue, outputQueue)

		actual := <-memoryChannel

		if !sliceEquals(actual, c.expected) {
			t.Errorf("ExecuteProgram(%v) == %v, expected %v", c.inputProgram, actual, c.expected)
		}
	}
}

func TestReadParameterValue(t *testing.T) {
	cases := []struct {
		inputState             []int
		inputOperationPointer  int
		inputOperation         string
		inputParameterPosition int
		expectedValue          int
	}{
		// positional mode
		{
			[]int{2, 4, 3, 2, 33},
			0,
			"00002",
			1,
			33,
		},
		{
			[]int{2, 4, 3, 2, 33},
			0,
			"00002",
			2,
			2,
		},
		{
			[]int{2, 4, 3, 2, 33},
			0,
			"00002",
			3,
			3,
		},
		// immediate mode
		{
			[]int{102, 4, 3, 2, 33},
			0,
			"00102",
			1,
			4,
		},
		{
			[]int{1002, 4, 3, 2, 33},
			0,
			"01002",
			2,
			3,
		},
		{
			[]int{10002, 4, 3, 2, 33},
			0,
			"10002",
			3,
			2,
		},
	}

	for _, c := range cases {
		actualValue := readParameterValue(c.inputState, c.inputOperationPointer, c.inputOperation, c.inputParameterPosition, false)

		if actualValue != c.expectedValue {
			t.Errorf("readParameterValue({%v}, %v, %v, %v) == %v, expected %v", c.inputState, c.inputOperationPointer, c.inputOperation, c.inputParameterPosition, actualValue, c.expectedValue)
		}
	}
}

func TestCompleteInputOutputPrograms(t *testing.T) {
	cases := []struct {
		program  []int
		input    []int
		expected []int
	}{
		{
			[]int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1001, 191, 50, 224, 101, -64, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 5, 224, 224, 1, 224, 223, 223, 2, 150, 218, 224, 1001, 224, -1537, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 2, 224, 1, 223, 224, 223, 1002, 154, 5, 224, 101, -35, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 5, 224, 1, 224, 223, 223, 1102, 76, 17, 225, 1102, 21, 44, 224, 1001, 224, -924, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 4, 224, 1, 224, 223, 223, 101, 37, 161, 224, 101, -70, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 6, 224, 224, 1, 223, 224, 223, 102, 46, 157, 224, 1001, 224, -1978, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 5, 224, 1, 224, 223, 223, 1102, 5, 29, 225, 1101, 10, 7, 225, 1101, 43, 38, 225, 1102, 33, 46, 225, 1, 80, 188, 224, 1001, 224, -73, 224, 4, 224, 102, 8, 223, 223, 101, 4, 224, 224, 1, 224, 223, 223, 1101, 52, 56, 225, 1101, 14, 22, 225, 1101, 66, 49, 224, 1001, 224, -115, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 7, 224, 1, 224, 223, 223, 1101, 25, 53, 225, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 108, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 329, 101, 1, 223, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 344, 1001, 223, 1, 223, 8, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 359, 101, 1, 223, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 374, 101, 1, 223, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 389, 101, 1, 223, 223, 7, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 404, 1001, 223, 1, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 419, 1001, 223, 1, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 434, 101, 1, 223, 223, 1008, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 449, 1001, 223, 1, 223, 1007, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 464, 1001, 223, 1, 223, 1008, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 479, 101, 1, 223, 223, 1007, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 494, 1001, 223, 1, 223, 108, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 509, 101, 1, 223, 223, 8, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 524, 1001, 223, 1, 223, 107, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 539, 101, 1, 223, 223, 107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 554, 101, 1, 223, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 569, 1001, 223, 1, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 584, 1001, 223, 1, 223, 1008, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 599, 1001, 223, 1, 223, 1107, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 614, 101, 1, 223, 223, 7, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 629, 1001, 223, 1, 223, 1108, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 644, 1001, 223, 1, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 659, 101, 1, 223, 223, 1108, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226},
			[]int{1},
			[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 11193703},
		},
		{
			[]int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1001, 191, 50, 224, 101, -64, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 5, 224, 224, 1, 224, 223, 223, 2, 150, 218, 224, 1001, 224, -1537, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 2, 224, 1, 223, 224, 223, 1002, 154, 5, 224, 101, -35, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 5, 224, 1, 224, 223, 223, 1102, 76, 17, 225, 1102, 21, 44, 224, 1001, 224, -924, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 4, 224, 1, 224, 223, 223, 101, 37, 161, 224, 101, -70, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 6, 224, 224, 1, 223, 224, 223, 102, 46, 157, 224, 1001, 224, -1978, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 5, 224, 1, 224, 223, 223, 1102, 5, 29, 225, 1101, 10, 7, 225, 1101, 43, 38, 225, 1102, 33, 46, 225, 1, 80, 188, 224, 1001, 224, -73, 224, 4, 224, 102, 8, 223, 223, 101, 4, 224, 224, 1, 224, 223, 223, 1101, 52, 56, 225, 1101, 14, 22, 225, 1101, 66, 49, 224, 1001, 224, -115, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 7, 224, 1, 224, 223, 223, 1101, 25, 53, 225, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 108, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 329, 101, 1, 223, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 344, 1001, 223, 1, 223, 8, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 359, 101, 1, 223, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 374, 101, 1, 223, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 389, 101, 1, 223, 223, 7, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 404, 1001, 223, 1, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 419, 1001, 223, 1, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 434, 101, 1, 223, 223, 1008, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 449, 1001, 223, 1, 223, 1007, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 464, 1001, 223, 1, 223, 1008, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 479, 101, 1, 223, 223, 1007, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 494, 1001, 223, 1, 223, 108, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 509, 101, 1, 223, 223, 8, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 524, 1001, 223, 1, 223, 107, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 539, 101, 1, 223, 223, 107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 554, 101, 1, 223, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 569, 1001, 223, 1, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 584, 1001, 223, 1, 223, 1008, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 599, 1001, 223, 1, 223, 1107, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 614, 101, 1, 223, 223, 7, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 629, 1001, 223, 1, 223, 1108, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 644, 1001, 223, 1, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 659, 101, 1, 223, 223, 1108, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226},
			[]int{5},
			[]int{12410607},
		},
	}

	for _, c := range cases {
		memoryChannel := make(chan []int)
		inputQueue := queue.NewFIFO()
		outputQueue := queue.NewFIFO()

		for _, input := range c.input {
			inputQueue.Enqueue(input)
		}

		go ExecuteProgram(c.program, memoryChannel, inputQueue, outputQueue)

		state := <-memoryChannel
		actual := fetchAllElementsFromQueue(outputQueue)

		if state != nil && !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("ExecuteProgram(%v) output == %v with input %v, expected %v", c.program, actual, c.input, c.expected)
		}
	}
}

func sliceEquals(first, second []int) bool {
	if len(first) != len(second) {
		return false
	}

	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}

	return true
}

func fetchAllElementsFromQueue(queue queue.Queue) []int {
	values := []int{}
	for queue.GetLen() > 0 {
		element, _ := queue.Dequeue()
		value, _ := element.(int)

		values = append(values, value)
	}

	return values
}
