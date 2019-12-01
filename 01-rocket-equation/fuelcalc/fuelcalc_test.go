package fuelcalc

import "testing"

func TestCalculateFuelForMass(t *testing.T) {
	cases := []struct {
		input    int64
		expected float64
	}{
		{12, 2},
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, c := range cases {
		actual := CalculateFuelForModule(c.input)

		if actual != c.expected {
			t.Errorf("CalculateFuelForModule(%v) == %v, expected %v", c.input, actual, c.expected)
		}
	}
}

func TestCalculateFuelForRocket(t *testing.T) {
	input := []int64{
		12, 14, 1969, 100756,
	}

	expected := 34241.0

	actual := CalculateFuelForRocket(input)

	if actual != expected {
		t.Errorf("TestCalculateFuelForRocket({12, 14, 1969, 100756}) == %v, expected %v", expected, actual)
	}
}
