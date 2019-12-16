package main

import (
	"fmt"

	"github.com/chr-ras/advent-of-code-2019/15-oxygen-system/repairdroid"
	"github.com/chr-ras/advent-of-code-2019/util/aoc"
)

func main() {
	program := aoc.ReadIntcode("repair_droid.txt")

	oxygenStation := repairdroid.FindShortestWayToOxygenTank(program)

	fmt.Printf("Oxygen station found after %d steps at position %v.\n", oxygenStation.Distance, oxygenStation.Position)
}
