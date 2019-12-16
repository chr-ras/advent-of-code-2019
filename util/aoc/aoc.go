package aoc

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ReadIntcode reads the given file and splits it into int64 intcode instructions.
func ReadIntcode(filePath string) []int64 {
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		panic(err)
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

	return program
}
