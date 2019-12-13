package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/chr-ras/advent-of-code-2019/13-care-package/arcade"
)

func main() {
	file, err := os.Open("./arcade_intcode.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	scanner := bufio.NewScanner(file)

	programText := ""
	for scanner.Scan() {
		programText += scanner.Text()
	}

	programTextElements := strings.Split(programText, ",")
	program := []int64{}
	for _, element := range programTextElements {
		codeElement, _ := strconv.ParseInt(element, 10, 64)
		program = append(program, codeElement)
	}

	program[0] = 2

	gameFinished := make(chan struct{})
	go arcade.RunGame(program, gameFinished)

	<-gameFinished
}
