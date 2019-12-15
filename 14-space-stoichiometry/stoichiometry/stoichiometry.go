package stoichiometry

import (
	"math"
	"strconv"
	"strings"
)

// MaxFuelForOre returns the maximum amount of fuel that can be produced wih the given ore.
func MaxFuelForOre(reactions string, ore int64) int64 {
	reactionLookup := parseReactions(reactions)

	oreForOneFuel := calcMinOreForFuel(reactionLookup, 1)

	lowerFuelGuess := ore / oreForOneFuel
	upperFuelGuess := 2 * lowerFuelGuess

	for upperFuelGuess-lowerFuelGuess > 1 {
		centerFuelGuess := lowerFuelGuess + (upperFuelGuess-lowerFuelGuess)/2

		centerOre := calcMinOreForFuel(reactionLookup, centerFuelGuess)
		if centerOre < ore {
			lowerFuelGuess = centerFuelGuess
		} else {
			upperFuelGuess = centerFuelGuess
		}
	}

	return lowerFuelGuess
}

// MinOreForFuel calculates the ORE needed to produce a target amount of FUEL.
func MinOreForFuel(reactions string, targetFuel int64) int64 {
	reactionLookup := parseReactions(reactions)

	return calcMinOreForFuel(reactionLookup, targetFuel)
}

func calcMinOreForFuel(reactionLookup map[string]reaction, targetFuel int64) int64 {
	neededElements := make(map[string]int64)
	leftovers := make(map[string]int64)

	calcNeededElements(element{name: "FUEL", amount: targetFuel}, reactionLookup, neededElements, leftovers)

	return calcOreFromElements(reactionLookup, neededElements)
}

func calcNeededElements(currentElement element, reactions map[string]reaction, neededElements map[string]int64, leftovers map[string]int64) {
	reaction := reactions[currentElement.name]

	for _, reactingElement := range reaction.from {
		if reactingElement.name == "ORE" {
			return
		}

		parentElementFactor := int64(math.Ceil(float64(currentElement.amount) / float64(reaction.to.amount)))

		reactionForReactingElement := reactions[reactingElement.name]
		neededElementsForReactingElement := reactingElement.amount*parentElementFactor - leftovers[reactingElement.name]

		reactionFactor := int64(math.Ceil(float64(neededElementsForReactingElement) / float64(reactionForReactingElement.to.amount)))

		actualElementsForReaction := reactionFactor * reactionForReactingElement.to.amount

		leftovers[reactingElement.name] = actualElementsForReaction - neededElementsForReactingElement
		neededElements[reactingElement.name] += actualElementsForReaction

		calcNeededElements(element{name: reactingElement.name, amount: neededElementsForReactingElement}, reactions, neededElements, leftovers)
	}
}

func calcOreFromElements(reactions map[string]reaction, neededElements map[string]int64) int64 {
	amountOfOre := int64(0)
	for element, neededAmount := range neededElements {
		reactionForElement := reactions[element]
		if reactionForElement.from[0].name != "ORE" {
			continue
		}

		factor := int64(math.Ceil(float64(neededAmount) / float64(reactionForElement.to.amount)))

		amountOfOre += factor * reactionForElement.from[0].amount
	}

	return amountOfOre
}

func parseReactions(reactionsString string) map[string]reaction {
	elements := make(map[string]reaction)

	lines := strings.Split(reactionsString, "\n")
	for _, line := range lines {
		equationSides := strings.Split(line, " => ")

		rightSideElement := parseElement(equationSides[1])

		leftSideElementParts := strings.Split(equationSides[0], ", ")
		leftSideElements := []element{}
		for _, leftPart := range leftSideElementParts {
			leftSideElements = append(leftSideElements, parseElement(leftPart))
		}

		elements[rightSideElement.name] = reaction{from: leftSideElements, to: rightSideElement}
	}

	return elements
}

func parseElement(elementString string) element {
	elementParts := strings.Split(elementString, " ")
	amount, _ := strconv.ParseInt(elementParts[0], 10, 64)
	return element{
		name:   elementParts[1],
		amount: int64(amount),
	}
}

type reaction struct {
	from []element
	to   element
}

type element struct {
	name   string
	amount int64
}
