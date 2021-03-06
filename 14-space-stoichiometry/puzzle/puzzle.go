package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/chr-ras/advent-of-code-2019/14-space-stoichiometry/stoichiometry"
)

func main() {
	file, err := os.Open("./reactions.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	scanner := bufio.NewScanner(file)

	input := ""
	for scanner.Scan() {
		input += (scanner.Text() + "\n")
	}

	input = strings.TrimSuffix(input, "\n")

	requiredOre := stoichiometry.MinOreForFuel(input, 1)

	fmt.Printf("Required ore for 1 unit of fuel: %d\n", requiredOre)

	collectedOre := int64(1000000000000)

	maxFuel := stoichiometry.MaxFuelForOre(input, collectedOre)

	fmt.Printf("Max fuel: %d\n", maxFuel)
}
