package bugs

import (
	"fmt"
	"math"
	"time"

	"github.com/gosuri/uilive"
)

// CalcBiodiversityForRecurringLayout determines the first recurring layout and calculates the biodiversity for that layout.
func CalcBiodiversityForRecurringLayout(initialLayout [][]bool, print bool) (int, [][]bool) {
	currentLayout := append([][]bool{}, initialLayout...)
	layoutLookup := make(map[string]struct{})

	writer := uilive.New()
	writer.Start()

	for {
		simulateIteration(currentLayout)

		if print {
			prettyPrint(currentLayout, writer)
			time.Sleep(250 * time.Millisecond)
		}

		if _, recurringLayout := layoutLookup[getKey(currentLayout)]; recurringLayout {
			break
		}

		layoutLookup[getKey(currentLayout)] = struct{}{}
	}

	writer.Stop()

	return calcBiodiversity(currentLayout), currentLayout
}

func simulateIteration(currentLayout [][]bool) {
	output := make([][]bool, len(currentLayout))

	for i := 0; i < len(output); i++ {
		output[i] = make([]bool, len(currentLayout[i]))
	}

	for row := 0; row < len(currentLayout); row++ {
		for col := 0; col < len(currentLayout[row]); col++ {
			adjacentBugs := 0

			if col != 0 && currentLayout[row][col-1] {
				adjacentBugs++
			}

			if col != len(currentLayout[row])-1 && currentLayout[row][col+1] {
				adjacentBugs++
			}

			if row != 0 && currentLayout[row-1][col] {
				adjacentBugs++
			}

			if row != len(currentLayout)-1 && currentLayout[row+1][col] {
				adjacentBugs++
			}

			switch {
			case currentLayout[row][col] && adjacentBugs != 1:
				output[row][col] = false

			case !currentLayout[row][col] && (adjacentBugs == 1 || adjacentBugs == 2):
				output[row][col] = true

			default:
				output[row][col] = currentLayout[row][col]
			}
		}
	}

	for row := 0; row < len(currentLayout); row++ {
		for col := 0; col < len(currentLayout[row]); col++ {
			currentLayout[row][col] = output[row][col]
		}
	}
}

func calcBiodiversity(layout [][]bool) int {
	biodiversity := 0

	for row := 0; row < len(layout); row++ {
		for col := 0; col < len(layout[row]); col++ {
			if !layout[row][col] {
				continue
			}

			biodiversity += int(math.Pow(2, float64(row*len(layout[0])+col)))
		}
	}

	return biodiversity
}

func prettyPrint(layout [][]bool, writer *uilive.Writer) {
	output := ""

	for row := 0; row < len(layout); row++ {
		for col := 0; col < len(layout[row]); col++ {
			if layout[row][col] {
				output += "#"
			} else {
				output += "."
			}
		}

		output += "\n"
	}

	fmt.Fprintf(writer, output+"\n")
}

func getKey(layout [][]bool) string {
	return fmt.Sprintf("%v", layout)
}
