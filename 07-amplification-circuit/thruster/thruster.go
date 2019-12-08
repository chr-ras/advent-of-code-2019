package thruster

import (
	"fmt"

	"github.com/chr-ras/advent-of-code-2019/07-amplification-circuit/intcode"

	queue "github.com/enriquebris/goconcurrentqueue"
	"github.com/gitchander/permutation"
)

// CalcHighestSignal determines the best signal for the thruster by trying all phase setting permutations.
// Phase settings: {0, 1, 2, 3, 4}
func CalcHighestSignal(amplifierControllerSoftware []int) int {
	phaseSetting := []int{0, 1, 2, 3, 4}

	return calcHighestSignal(amplifierControllerSoftware, phaseSetting)
}

// CalcHighestSignalInFeedbackLoopMode determines the best signal for the thruster by trying all phase setting permutations.
// Phase settings: {5, 6, 7, 8, 9}
func CalcHighestSignalInFeedbackLoopMode(amplifierControllerSoftware []int) int {
	phaseSetting := []int{5, 6, 7, 8, 9}

	return calcHighestSignal(amplifierControllerSoftware, phaseSetting)
}

func calcHighestSignal(amplifierControllerSoftware []int, phaseSetting []int) int {
	bestSignal := 0

	phasePermutations := permutation.New(permutation.IntSlice(phaseSetting))

	for phasePermutations.Next() {
		signal := executeAmplifierControllerSoftware(amplifierControllerSoftware, phaseSetting)

		if signal > bestSignal {
			bestSignal = signal
		}
	}

	return bestSignal
}

func executeAmplifierControllerSoftware(amplifierControllerSoftware []int, phaseSettings []int) int {
	lastAmplifierMemoryChannel := make(chan []int)

	fifthOutputQueue := queue.NewFIFO()
	firstOutputQueue := queue.NewFIFO()
	secondOutputQueue := queue.NewFIFO()
	thirdOutputQueue := queue.NewFIFO()
	fourthOutputQueue := queue.NewFIFO()

	// Initialize output queues with phase settings, starting with the last one as the input for the first amplifier.
	fifthOutputQueue.Enqueue(phaseSettings[0])
	// First input queue also needs the initial signal of 0
	fifthOutputQueue.Enqueue(0)
	firstOutputQueue.Enqueue(phaseSettings[1])
	secondOutputQueue.Enqueue(phaseSettings[2])
	thirdOutputQueue.Enqueue(phaseSettings[3])
	fourthOutputQueue.Enqueue(phaseSettings[4])

	go intcode.ExecuteProgram(append([]int(nil), amplifierControllerSoftware...), make(chan []int), fifthOutputQueue, firstOutputQueue)
	go intcode.ExecuteProgram(append([]int(nil), amplifierControllerSoftware...), make(chan []int), firstOutputQueue, secondOutputQueue)
	go intcode.ExecuteProgram(append([]int(nil), amplifierControllerSoftware...), make(chan []int), secondOutputQueue, thirdOutputQueue)
	go intcode.ExecuteProgram(append([]int(nil), amplifierControllerSoftware...), make(chan []int), thirdOutputQueue, fourthOutputQueue)
	go intcode.ExecuteProgram(append([]int(nil), amplifierControllerSoftware...), lastAmplifierMemoryChannel, fourthOutputQueue, fifthOutputQueue)

	lastAmplifierFinalState := <-lastAmplifierMemoryChannel

	if lastAmplifierFinalState == nil {
		fmt.Printf("Last amplifier state is nil")
		return -1
	}

	outputElement, err := fifthOutputQueue.Dequeue()
	if err != nil {
		fmt.Printf("Could not fetch final output from last amplifier: %v\n", err.Error())
		return -1
	}

	outputValue, _ := outputElement.(int)

	return outputValue
}
