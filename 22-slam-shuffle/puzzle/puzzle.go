package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/chr-ras/advent-of-code-2019/22-slam-shuffle/cards"
)

func main() {
	file, err := os.Open("./shuffle_process.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	scanner := bufio.NewScanner(file)

	shuffleProcess := []string{}
	for scanner.Scan() {
		shuffleProcess = append(shuffleProcess, scanner.Text())
	}

	shuffledDeck := cards.ShuffleDeck(shuffleProcess, 10007)

	index := -1
	for i, card := range shuffledDeck {
		if card == 2019 {
			index = i
			break
		}
	}

	fmt.Printf("Card index of card 2019: %d\n", index)
}
