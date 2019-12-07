package thruster

import (
	"github.com/chr-ras/advent-of-code-2019/07-amplification-circuit/intcode"

	"github.com/gitchander/permutation"
)

// CalcHighestSignal determines the best signal for the thruster by trying all phase setting permutations.
func CalcHighestSignal(amplifierControllerSoftware []int) int {
	phaseSetting := []int{0, 1, 2, 3, 4}
	phasePermutations := permutation.New(permutation.IntSlice(phaseSetting))

	bestSignal := 0

	for phasePermutations.Next() {
		signal := executeAmplifierControllerSoftware(amplifierControllerSoftware, phaseSetting)

		if signal > bestSignal {
			bestSignal = signal
		}
	}

	intcode.ExecuteProgram(amplifierControllerSoftware)

	return bestSignal
}

func executeAmplifierControllerSoftware(amplifierControllerSoftware []int, phaseSettings []int) int {
	programOutput := 0

	for i := 0; i < 5; i++ {
		intcode.Input = []int{phaseSettings[i], programOutput}

		program := append([]int(nil), amplifierControllerSoftware...)
		_, outputs := intcode.ExecuteProgram(program)

		programOutput = outputs[len(outputs)-1]
	}

	return programOutput
}
