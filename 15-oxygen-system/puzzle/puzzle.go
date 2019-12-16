package main

import (
	"github.com/chr-ras/advent-of-code-2019/15-oxygen-system/repairdroid"
	"github.com/chr-ras/advent-of-code-2019/util/aoc"
)

func main() {
	program := aoc.ReadIntcode("repair_droid.txt")

	repairdroid.FindShortestWayToOxygenTank(program)
}
