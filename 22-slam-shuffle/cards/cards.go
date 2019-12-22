package cards

import (
	"regexp"
	"strconv"
)

var numberRegex = regexp.MustCompile(`[-]?\d+`)
var cutNRegex = regexp.MustCompile(`cut [-]?\d+`)
var dealWithIncrementRegex = regexp.MustCompile(`deal with increment \d+`)

// ShuffleDeck applies the shuffle techniques specified in the shuffle process to a card deck of the given decksize.
func ShuffleDeck(shuffleProcess []string, deckSize int) []int {
	currentDeck := make([]int, deckSize)

	for i := 0; i < deckSize; i++ {
		currentDeck[i] = i
	}

	for _, technique := range shuffleProcess {
		switch {
		case technique == "deal into new stack":
			currentDeck = applyDealIntoNewStack(currentDeck)

		case cutNRegex.MatchString(technique):
			currentDeck = applyCutN(technique, currentDeck)

		case dealWithIncrementRegex.MatchString(technique):
			currentDeck = applyDealWithIncrement(technique, currentDeck, deckSize)
		}
	}

	return currentDeck
}

func applyDealIntoNewStack(currentDeck []int) []int {
	output := []int{}

	for i := len(currentDeck) - 1; i >= 0; i-- {
		output = append(output, currentDeck[i])
	}

	return output
}

func applyCutN(technique string, currentDeck []int) []int {
	numberToCut := extractNumber(technique)

	if numberToCut < 0 {
		return append(currentDeck[len(currentDeck)+numberToCut:len(currentDeck)], currentDeck[0:len(currentDeck)+numberToCut]...)
	}

	return append(currentDeck[numberToCut:len(currentDeck)], currentDeck[0:numberToCut]...)
}

func applyDealWithIncrement(technique string, currentDeck []int, deckSize int) []int {
	output := make([]int, deckSize)

	increment := extractNumber(technique)

	for i := 0; i < deckSize; i++ {
		currentIndex := (i * increment) % deckSize

		output[currentIndex] = currentDeck[i]
	}

	return output
}

func extractNumber(technique string) int {
	number, _ := strconv.ParseInt(numberRegex.FindString(technique), 10, 64)
	return int(number)
}
