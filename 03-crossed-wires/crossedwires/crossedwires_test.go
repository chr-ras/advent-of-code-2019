package crossedwires

import (
	"reflect"
	"testing"
)

func TestFindClosestCrossingDistance(t *testing.T) {
	cases := []struct {
		firstPath, secondPath string
		expected              int
	}{
		{
			"R8,U5,L5,D3",
			"U7,R6,D4,L4",
			6,
		},
		{
			"R75,D30,R83,U83,L12,D49,R71,U7,L72",
			"U62,R66,U55,R34,D71,R55,D58,R83",
			159,
		},
		{
			"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			135,
		},
	}

	for _, c := range cases {
		actual := FindClosestCrossingDistance(c.firstPath, c.secondPath)

		if actual != c.expected {
			t.Errorf("FindClosestCrossingDistance(%v, %v) == %v, expected %v", c.firstPath, c.secondPath, actual, c.expected)
		}
	}
}

func TestIntersection(t *testing.T) {
	cases := []struct {
		firstEdge, secondEdge Edge
		expectedIntersects    bool
		expectedIntersection  Coordinate
	}{
		{
			Edge{Coordinate{0, 2}, Coordinate{2, 2}},
			Edge{Coordinate{0, 1}, Coordinate{2, 1}},
			false,
			Coordinate{},
		},
		{
			Edge{Coordinate{0, 2}, Coordinate{2, 2}},
			Edge{Coordinate{-1, 1}, Coordinate{-1, -1}},
			false,
			Coordinate{},
		},
		{
			Edge{Coordinate{0, 2}, Coordinate{2, 2}},
			Edge{Coordinate{-1, 3}, Coordinate{-1, 5}},
			false,
			Coordinate{},
		},
		{
			Edge{Coordinate{0, 2}, Coordinate{2, 2}},
			Edge{Coordinate{0, 0}, Coordinate{0, 2}},
			true,
			Coordinate{0, 2},
		},
		{
			Edge{Coordinate{0, 2}, Coordinate{2, 2}},
			Edge{Coordinate{1, 1}, Coordinate{1, 3}},
			true,
			Coordinate{1, 2},
		},
		{
			Edge{Coordinate{0, 2}, Coordinate{2, 2}},
			Edge{Coordinate{2, 2}, Coordinate{2, 4}},
			true,
			Coordinate{2, 2},
		},
		{
			Edge{Coordinate{0, 2}, Coordinate{2, 2}},
			Edge{Coordinate{0, 2}, Coordinate{-2, 2}},
			true,
			Coordinate{0, 2},
		},
		{
			Edge{Coordinate{2, 4}, Coordinate{2, 2}},
			Edge{Coordinate{2, 2}, Coordinate{2, 0}},
			true,
			Coordinate{2, 2},
		},
	}

	for _, c := range cases {
		acutalIntersects, actualIntersection := intersection(c.firstEdge, c.secondEdge)

		if acutalIntersects != c.expectedIntersects || actualIntersection != c.expectedIntersection {
			t.Errorf("intersection(%v, %v) == (%v, %v), expected (%v, %v)",
				c.firstEdge, c.secondEdge, acutalIntersects, actualIntersection, c.expectedIntersects, c.expectedIntersection)
		}

		acutalIntersects, actualIntersection = intersection(c.secondEdge, c.firstEdge)

		if acutalIntersects != c.expectedIntersects || actualIntersection != c.expectedIntersection {
			t.Errorf("intersection(%v, %v) == (%v, %v), expected (%v, %v), swapping the edges must not change the intersection",
				c.secondEdge, c.firstEdge, acutalIntersects, actualIntersection, c.expectedIntersects, c.expectedIntersection)
		}
	}
}

func TestCalcManhattanDistanceFromZeroPoint(t *testing.T) {
	cases := []struct {
		input    Coordinate
		expected int
	}{
		{
			Coordinate{0, 0},
			0,
		},
		{
			Coordinate{1, 0},
			1,
		},
		{
			Coordinate{0, -1},
			1,
		},
		{
			Coordinate{3, -3},
			6,
		},
		{
			Coordinate{-3, 3},
			6,
		},
		{
			Coordinate{2, 7},
			9,
		},
		{
			Coordinate{-2, -7},
			9,
		},
	}

	for _, c := range cases {
		actual := CalcManhattanDistance(Coordinate{0, 0}, c.input)

		if actual != c.expected {
			t.Errorf("CalcManhattanDistance(Coordinate{0, 0}, %v) == %v, expected %v)", c.input, actual, c.expected)
		}
	}
}

func TestCalcManhattanDistanceFromNonZeroPoint(t *testing.T) {
	firstInput := Coordinate{-2, -7}
	secondInput := Coordinate{3, 3}

	expected := 15

	actual := CalcManhattanDistance(firstInput, secondInput)

	if actual != expected {
		t.Errorf("CalcManhattanDistance(%v, %v) == %v, expected %v", firstInput, secondInput, actual, expected)
	}

	actual = CalcManhattanDistance(secondInput, firstInput)

	if actual != expected {
		t.Errorf("CalcManhattanDistance(%v, %v) == %v, expected %v. Swapping parameters must not change the result.", secondInput, firstInput, actual, expected)
	}
}

func TestParsePathString(t *testing.T) {
	cases := []struct {
		input    string
		expected []WirePathElement
	}{
		{
			"U7,R6,D4,L4",
			[]WirePathElement{WirePathElement{Up, 7}, WirePathElement{Right, 6}, WirePathElement{Down, 4}, WirePathElement{Left, 4}},
		},
		{
			"R42,L2,D4711",
			[]WirePathElement{WirePathElement{Right, 42}, WirePathElement{Left, 2}, WirePathElement{Down, 4711}},
		},
	}

	for _, c := range cases {
		actual := ParsePathString(c.input)

		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("ParsePathString(%v) == %v, expected %v", c.input, actual, c.expected)
		}
	}
}

func TestTravelPath(t *testing.T) {
	input := ParsePathString("U7,R6,D4,L4")
	expected := []Coordinate{
		Coordinate{0, 0},
		Coordinate{0, 7},
		Coordinate{6, 7},
		Coordinate{6, 3},
		Coordinate{2, 3},
	}

	actual := travelPath(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("travelPath(%v) == %v, expected %v", input, actual, expected)
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
