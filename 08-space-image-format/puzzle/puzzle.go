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

	decodedImage := spaceimageformat.Parse(image, 25, 6)
	checksum := spaceimageformat.CheckIsValid(decodedImage)

	fmt.Printf("Checksum: %v", checksum)
}
