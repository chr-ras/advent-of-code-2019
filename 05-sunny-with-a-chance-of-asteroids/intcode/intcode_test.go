package intcode

import "testing"

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
		actual := ExecuteProgram(c.input)
		if !sliceEquals(actual, c.expected) {
			t.Errorf("ExecuteProgram(%v) == %v, expected %v", c.input, actual, c.expected)
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
