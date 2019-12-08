package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/chr-ras/advent-of-code-2019/08-space-image-format/spaceimageformat"
)

func main() {
	file, err := os.Open("./password_image.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	scanner := bufio.NewScanner(file)
	image := ""

	for scanner.Scan() {
		image += scanner.Text()
	}

	width := 25
	height := 6
	parsedImage := spaceimageformat.Parse(image, width, height)
	checksum := spaceimageformat.CheckIsValid(parsedImage)

	fmt.Printf("Checksum: %v\n", checksum)

	decodedImage := spaceimageformat.Decode(parsedImage, width, height)
	spaceimageformat.PrettyPrint(decodedImage, width, height)
}
