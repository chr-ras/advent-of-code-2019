package main

import (
	"fmt"
	"strconv"
)

func main() {
	viablePasswords := findViablePasswords()

	fmt.Printf("Number of viable passwords: %v", len(viablePasswords))
}

func findViablePasswords() []string {
	const (
		minPasswordValue = 265275
		maxPasswordValue = 781584
	)

	var viablePasswords []string

	for i := 2; i < 8; i++ {
		for j := i; j < 10; j++ {
			for k := j; k < 10; k++ {
				for l := k; l < 10; l++ {
					for m := l; m < 10; m++ {
						for n := m; n < 10; n++ {
							passwordGuess := 100000*i + 10000*j + 1000*k + 100*l + 10*m + n

							if passwordGuess < minPasswordValue || passwordGuess > maxPasswordValue {
								continue
							}

							passwordGuessText := strconv.FormatInt(int64(passwordGuess), 10)
							if !hasAtleastTwoIdenticalAdjacentDigits(passwordGuessText) {
								continue
							}

							if !hasGroupOfExactlyTwoIdenticalDigits(passwordGuessText) {
								continue
							}

							fmt.Println(passwordGuessText)

							viablePasswords = append(viablePasswords, passwordGuessText)
						}
					}
				}
			}
		}
	}

	return viablePasswords
}

func hasAtleastTwoIdenticalAdjacentDigits(guess string) bool {
	for i := 0; i < 5; i++ {
		if guess[i] == guess[i+1] {
			return true
		}
	}

	return false
}

func hasGroupOfExactlyTwoIdenticalDigits(guess string) bool {
	groupSize := 0
	currentGroupDigit := guess[0:1]

	for i := 0; i < 6; i++ {
		currentDigit := guess[i : i+1]
		if currentDigit == currentGroupDigit {
			groupSize++
		} else {
			if groupSize == 2 {
				return true
			}

			currentGroupDigit = currentDigit
			groupSize = 1
		}
	}

	if groupSize == 2 {
		return true
	}

	return false
}
