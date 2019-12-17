package fft

import (
	"math"
	"strconv"
)

// ApplyFft applies the FFT algorithm to the input signal and returns the reconstructed signal.
func ApplyFft(signalText string, phases int) string {
	signal := parseSignal(signalText)

	for phase := 0; phase < phases; phase++ {

		newSignal := []int{}

		for element := 0; element < len(signal); element++ {

			currentPattern := buildCurrentPattern(element, len(signal))

			newElementRawValue := 0

			for i := 0; i < len(signal); i++ {
				newElementRawValue += signal[i] * currentPattern[i%len(currentPattern)]
			}

			newElementValue := int(math.Abs(float64(newElementRawValue % 10)))

			newSignal = append(newSignal, newElementValue)
		}

		signal = newSignal
	}

	output := ""

	for _, element := range signal {
		output += strconv.FormatInt(int64(element), 10)
	}

	return output
}

func parseSignal(signalText string) []int {
	signal := []int{}

	for i := 0; i < len(signalText); i++ {
		element, _ := strconv.ParseInt(signalText[i:i+1], 10, 64)
		signal = append(signal, int(element))
	}

	return signal
}

func buildCurrentPattern(elementIndex, signalLength int) []int {
	initialPattern := []int{0, 1, 0, -1}

	currentElementPattern := []int{}

	for patternRound := 0; patternRound < signalLength/len(initialPattern)+1; patternRound++ {
		for patternElement := 0; patternElement < len(initialPattern); patternElement++ {
			for element := 0; element < elementIndex+1; element++ {
				currentElementPattern = append(currentElementPattern, initialPattern[patternElement])
			}
		}
	}

	return currentElementPattern[1:len(currentElementPattern)]
}
