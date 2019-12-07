package orbitmap

import (
	"reflect"
	"testing"
)

func TestCalcOrbits(t *testing.T) {
	input := "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L"

	expected := 42

	actual := CalcOrbits(input)

	if actual != expected {
		t.Errorf("CalcOrbits(\"%v\") == %v, expected %v", input, actual, expected)
	}
}

func TestCalcOrbitTransitions(t *testing.T) {
	input := "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L\nK)YOU\nI)SAN"
	expected := 4

	actual := CalcOrbitTransitions(input)

	if actual != expected {
		t.Errorf("CalcOrbitTransitions(\"%v\") == %v, expected %v", input, actual, expected)
	}
}

func TestParseOrbitMap(t *testing.T) {
	cases := []struct {
		input    string
		expected []orbit
	}{
		{
			"COM)B\nB)C\nC)D\nD)E",
			[]orbit{orbit{"COM", "B"}, orbit{"B", "C"}, orbit{"C", "D"}, orbit{"D", "E"}},
		},
	}

	for _, c := range cases {
		actual := parseOrbitMap(c.input)

		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("parseOrbitMap(\"%v\") == %v, expected %v", c.input, actual, c.expected)
		}
	}
}
