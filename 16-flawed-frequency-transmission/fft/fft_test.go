package fft

import "testing"

func TestApplyFft(t *testing.T) {
	cases := []struct {
		inputSignal              string
		inputPhases              int
		expectedFirstEigthDigits string
	}{
		{
			"12345678",
			4,
			"01029498",
		},
		{
			"80871224585914546619083218645595",
			100,
			"24176176",
		},
		{
			"19617804207202209144916044189917",
			100,
			"73745418",
		},
		{
			"69317163492948606335995924319873",
			100,
			"52432133",
		},
	}

	for _, c := range cases {
		actual := ApplyFft(c.inputSignal, c.inputPhases)

		if actual[0:8] != c.expectedFirstEigthDigits {
			t.Errorf("ApplyFft(%v, %d) == %v, expected %v", c.inputSignal, c.inputPhases, actual[0:8], c.expectedFirstEigthDigits)
		}
	}
}
