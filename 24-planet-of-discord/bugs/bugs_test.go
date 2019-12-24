package bugs

import (
	"reflect"
	"testing"
)

func TestCalcBiodiversityForRecurringLayout(t *testing.T) {
	testCase := struct {
		input                [][]bool
		expectedBiodiversity int
		expectedLayout       [][]bool
	}{
		[][]bool{
			{false, false, false, false, true},
			{true, false, false, true, false},
			{true, false, false, true, true},
			{false, false, true, false, false},
			{true, false, false, false, false},
		},
		2129920,
		[][]bool{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{true, false, false, false, false},
			{false, true, false, false, false},
		},
	}

	actualBiodiversity, actualLayout := CalcBiodiversityForRecurringLayout(testCase.input, false)

	if actualBiodiversity != testCase.expectedBiodiversity || !reflect.DeepEqual(actualLayout, testCase.expectedLayout) {
		t.Errorf("CalcBiodiversityForRecurringLayout(%v) == (%d, %v), expected (%d, %v)", testCase.input, actualBiodiversity, actualLayout, testCase.expectedBiodiversity, testCase.expectedLayout)
	}
}
