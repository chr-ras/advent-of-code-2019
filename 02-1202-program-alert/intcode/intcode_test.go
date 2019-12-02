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
