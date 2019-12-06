package orbitmap

import (
	"fmt"
	"strings"
)

// CalcOrbits determines the total number of direct and indirect orbits in the given orbit map.
func CalcOrbits(orbitMap string) int {
	fmt.Println()
	parsedMap := parseOrbitMap(orbitMap)

	objectOrbitsCenterMap := make(map[string]string)
	for _, mapElement := range parsedMap {
		objectOrbitsCenterMap[mapElement.object] = mapElement.center
	}

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

func parseOrbitMap(orbitMap string) []orbit {
	var parsedMap []orbit
	mapParts := strings.Split(orbitMap, "\n")

	for _, mapPart := range mapParts {
		orbitParts := strings.Split(mapPart, ")")

		parsedMap = append(parsedMap, orbit{orbitParts[0], orbitParts[1]})
	}

	return parsedMap
}

type orbit struct {
	center, object string
}
