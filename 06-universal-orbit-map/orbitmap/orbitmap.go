package orbitmap

import (
	"fmt"
	"strings"
)

// CalcOrbits determines the total number of direct and indirect orbits in the given orbit map.
func CalcOrbits(orbitMap string) int {
	parsedMap := parseOrbitMap(orbitMap)

	objectOrbitsCenterMap, _ := createOrbitLookups(parsedMap)

	totalOrbits := 0
	for object, center := range objectOrbitsCenterMap {
		for {
			totalOrbits++

			object = center

			if object == "COM" {
				break
			}

			center = objectOrbitsCenterMap[object]
		}
	}

	return totalOrbits
}

// CalcOrbitTransitions determines the number of orbit transitions required to get from YOU to SAN.
func CalcOrbitTransitions(orbitMap string) int {
	parsedMap := parseOrbitMap(orbitMap)
	objectToCenterMap, centerToObjectsMap := createOrbitLookups(parsedMap)

	startOrbitCenter := objectToCenterMap["YOU"]
	startObject := "YOU"

	foundSanta, transitions := calcOrbitTransitions(objectToCenterMap, centerToObjectsMap, startOrbitCenter, startObject, 0)

	if !foundSanta {
		fmt.Println("Something went seriously wrong here. We did't find Santa!")
		return -1
	}

	return transitions
}

func calcOrbitTransitions(objectToCenterMap map[string]string, centerToObjectsMap map[string][]string, currentCenter, cameFromObject string, currentTransitionCount int) (foundSanta bool, transitions int) {
	if orbitHasSanta(centerToObjectsMap, currentCenter) {
		fmt.Println("Santa was found")
		return true, currentTransitionCount
	}

	// First we go "up the tree" - if it is not the root [COM]
	if currentCenter != "COM" {
		orbitOfCenter := objectToCenterMap[currentCenter]
		if orbitOfCenter != cameFromObject {
			fmt.Printf("Going up the tree from [%v] to [%v]\n", currentCenter, orbitOfCenter)
			currentTransitionCount++

			// The current center becomes the new object and its orbit becomes the new center
			foundSanta, transitions = calcOrbitTransitions(objectToCenterMap, centerToObjectsMap, orbitOfCenter, currentCenter, currentTransitionCount)
			if foundSanta {
				return true, transitions
			}

			// This wasn't the way to go so remove this transition from the count
			currentTransitionCount--
		} else {
			fmt.Printf("NOT going up the tree because we came from [%v]\n", cameFromObject)
		}
	}

	// Second, we check other objects orbiting the current center
	objectsOrbitingCurrentCenter := centerToObjectsMap[currentCenter]
	for _, object := range objectsOrbitingCurrentCenter {
		// Ignore the object we came from because that has already been checked.
		if object == cameFromObject {
			fmt.Printf("NOT checking children of [%v] because we came from there\n", object)
			continue
		}

		fmt.Printf("Checking children of [%v] : [%v]\n", currentCenter, object)
		currentTransitionCount++
		foundSanta, transitions = calcOrbitTransitions(objectToCenterMap, centerToObjectsMap, object, currentCenter, currentTransitionCount)
		if foundSanta {
			return true, transitions
		}

		currentTransitionCount--
	}

	return false, -1
}

func orbitHasSanta(centerToObjectsMap map[string][]string, currentCenter string) bool {
	currentCenterObjects := centerToObjectsMap[currentCenter]
	for _, object := range currentCenterObjects {
		if object == "SAN" {
			return true
		}
	}

	return false
}

func parseOrbitMap(orbitMap string) []orbit {
	var parsedMap []orbit
	mapParts := strings.Split(orbitMap, "\n")

	for _, mapPart := range mapParts {
		orbitParts := strings.Split(mapPart, ")")

		parsedMap = append(parsedMap, orbit{orbitParts[0], orbitParts[1]})
	}

	return parsedMap
}

func createOrbitLookups(orbitMap []orbit) (objectToCenterMap map[string]string, centerToObjectsMap map[string][]string) {
	objectToCenterMap = make(map[string]string)
	centerToObjectsMap = make(map[string][]string)

	for _, mapElement := range orbitMap {
		objectToCenterMap[mapElement.object] = mapElement.center
		objectsForCenter := centerToObjectsMap[mapElement.center]
		if objectsForCenter == nil {
			centerToObjectsMap[mapElement.center] = []string{mapElement.object}
		} else {
			centerToObjectsMap[mapElement.center] = append(objectsForCenter, mapElement.object)
		}
	}

	return
}

type orbit struct {
	center, object string
}
