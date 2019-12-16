package main

import (
	"fmt"

	"github.com/chr-ras/advent-of-code-2019/15-oxygen-system/oxygen"
	"github.com/chr-ras/advent-of-code-2019/15-oxygen-system/repairdroid"
	"github.com/chr-ras/advent-of-code-2019/util/aoc"
)

func main() {
	program := aoc.ReadIntcode("repair_droid.txt")

	oxygenStation, shipMap := repairdroid.FindShortestWayToOxygenTank(program)

	fmt.Printf("Oxygen station found after %d steps at position %v.\n", oxygenStation.Distance, oxygenStation.Position)

	minutesToFillShipWithOxygen := oxygen.SimulateOxygenExpansion(shipMap, oxygenStation)

	fmt.Printf("It took %d minutes for the oxygen to fill the entire space ship.\n", minutesToFillShipWithOxygen)
}
