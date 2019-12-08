package spaceimageformat

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	input := "123456789012"
	expected := [][]int{
		{
			1, 2, 3, 4, 5, 6,
		},
		{
			7, 8, 9, 0, 1, 2,
		},
	}

	actual := Parse(input, 3, 2)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Decode(\"%v\") == %v, expected %v", input, actual, expected)
	}
}

func TestCheckIsValid(t *testing.T) {
	input := [][]int{
		{1, 1, 2, 3},
		{0, 1, 2, 2},
		{0, 0, 1, 2},
	}

	expected := 2

	actual := CheckIsValid(input)

	if expected != actual {
		t.Errorf("CheckIsValid(%v) == %v, expected %v", input, actual, expected)
	}
}
