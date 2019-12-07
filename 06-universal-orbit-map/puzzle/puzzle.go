package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/chr-ras/advent-of-code-2019/06-universal-orbit-map/orbitmap"
)

func main() {
	file, err := os.Open("./orbitmap.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	scanner := bufio.NewScanner(file)

	orbitMap := ""
	for scanner.Scan() {
		orbitMap += (scanner.Text() + "\n")
	}

	orbitMap = strings.TrimRight(orbitMap, "\n")

	totalOrbits := orbitmap.CalcOrbits(orbitMap)

	transitions := orbitmap.CalcOrbitTransitions(orbitMap)

	fmt.Printf("Total orbits: %v required transitions to santa: %v", totalOrbits, transitions)
}
