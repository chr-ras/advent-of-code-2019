package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/chr-ras/advent-of-code-2019/01-rocket-equation/fuelcalc"
)

func main() {
	file, err := os.Open("./module_masses.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	moduleMasses := []int64{}
	for scanner.Scan() {
		scannedInput := scanner.Text()
		parsedMass, _ := strconv.ParseInt(scannedInput, 10, 64)
		moduleMasses = append(moduleMasses, parsedMass)
	}

	requiredFuel := fuelcalc.CalculateFuelForRocket(moduleMasses)

	fmt.Printf("Required fuel: %v", requiredFuel)
}
