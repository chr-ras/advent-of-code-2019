package main

import (
	"github.com/chr-ras/advent-of-code-2019/23-category-six/cat6"
	"github.com/chr-ras/advent-of-code-2019/util/aoc"
)

func main() {
	program := aoc.ReadIntcode("./nic_software.txt")

	cat6.RunNetwork(program)
}
