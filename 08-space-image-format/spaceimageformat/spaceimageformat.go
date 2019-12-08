package spaceimageformat

import (
	"math"
	"strconv"
)

// Parse calculates the space image format layers of the encoded image using the given image width and height.
func Parse(encodedImage string, width, height int) [][]int {
	layers := [][]int{}

	for layer := 0; layer < len(encodedImage)/width/height; layer++ {
		imageLayer := []int{}
		pixelOffset := layer * width * height
		for pixelIndex := 0; pixelIndex < width*height; pixelIndex++ {
			pixelValue := encodedImage[pixelIndex+pixelOffset : pixelIndex+pixelOffset+1]
			pixel, _ := strconv.ParseInt(pixelValue, 10, 64)
			imageLayer = append(imageLayer, int(pixel))
		}

		layers = append(layers, imageLayer)
	}

	return layers
}

// CheckIsValid determines a value showing whether the image has been sent correctly or has become corrupted.
func CheckIsValid(layers [][]int) int {
	numberOfZeros := math.MaxInt32
	numberOfOnes := 0
	numberOfTwos := 0

	for _, layer := range layers {
		currentLayerZeros := 0
		currentLayerOnes := 0
		currentLayerTwos := 0
		for _, pixel := range layer {
			switch pixel {
			case 0:
				currentLayerZeros++
			case 1:
				currentLayerOnes++
			case 2:
				currentLayerTwos++
			}
		}

		if currentLayerZeros < numberOfZeros {
			numberOfZeros = currentLayerZeros
			numberOfOnes = currentLayerOnes
			numberOfTwos = currentLayerTwos
		}
	}

	return numberOfOnes * numberOfTwos
}
