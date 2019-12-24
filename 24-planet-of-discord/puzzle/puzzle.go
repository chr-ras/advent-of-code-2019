package main

import (
	"fmt"
	"strings"

	"github.com/chr-ras/advent-of-code-2019/24-planet-of-discord/bugs"
)

func main() {
	rawInitialLayout := "..##.\n..#..\n##...\n#....\n...##"
	initialLayoutRows := strings.Split(rawInitialLayout, "\n")

	initialLayout := [][]bool{}

	for row := 0; row < len(initialLayoutRows); row++ {
		initialLayoutRow := []bool{}
		for col := 0; col < len(initialLayoutRows[row]); col++ {
			if initialLayoutRows[row][col:col+1] == "#" {
				initialLayoutRow = append(initialLayoutRow, true)
			} else {
				initialLayoutRow = append(initialLayoutRow, false)
			}
		}

		initialLayout = append(initialLayout, initialLayoutRow)
	}

	biodiversity, _ := bugs.CalcBiodiversityForRecurringLayout(initialLayout, true)

	fmt.Printf("The biodiversity of the first recurring layout is %d.\n", biodiversity)
}
