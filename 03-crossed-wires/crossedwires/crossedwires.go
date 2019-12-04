package crossedwires

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Direction enum
type Direction int

const (
	Up Direction = iota
	Left
	Down
	Right
)

// WirePathElement struct represents a path element of the wires.
type WirePathElement struct {
	Direction Direction
	Length    int
}

// Coordinate struct represents a 2D coordinate
type Coordinate struct {
	X, Y int
}

// Edge struct is represented by two points in a 2D coordinate system
type Edge struct {
	firstPoint, secondPoint Coordinate
}

// ParsePathString parses an input string in the form of U7,R6,D4,L4
// and returns the parsed path elements.
func ParsePathString(path string) []WirePathElement {
	parts := strings.Split(path, ",")
	var parsedPath []WirePathElement

	for _, part := range parts {
		var direction Direction
		switch {
		case strings.HasPrefix(part, "U"):
			direction = Up
		case strings.HasPrefix(part, "L"):
			direction = Left
		case strings.HasPrefix(part, "D"):
			direction = Down
		case strings.HasPrefix(part, "R"):
			direction = Right
		}

		lengthString := part[1:len(part)]
		length, _ := strconv.ParseInt(lengthString, 10, 64)

		parsedPath = append(parsedPath, WirePathElement{direction, int(length)})
	}

	return parsedPath
}

// FindClosestCrossingDistance determines the closest crossing of the two specified path strings
// in relation to the start point {0, 0}. The path length is determined using the Manhattan distance.
// The start point does not count as a crossing, neither does if a path crosses itself.
func FindClosestCrossingDistance(firstPathString, secondPathString string) int {
	firstPath := ParsePathString(firstPathString)
	secondPath := ParsePathString(secondPathString)

	firstPathEdgePoints := travelPath(firstPath)
	secondPathEdgePoints := travelPath(secondPath)

	currentMinDistance := math.MaxInt32

	for i := 0; i < len(firstPathEdgePoints)-1; i++ {
		for j := 0; j < len(secondPathEdgePoints)-1; j++ {
			firstEdge := Edge{firstPathEdgePoints[i], firstPathEdgePoints[i+1]}
			secondEdge := Edge{secondPathEdgePoints[j], secondPathEdgePoints[j+1]}

			fmt.Printf("Check edges %v and %v for intersection.\n", firstEdge, secondEdge)

			intersects, intersection := intersection(firstEdge, secondEdge)

			if !intersects {
				continue
			}

			if intersection.X == 0 && intersection.Y == 0 {
				continue
			}

			distanceFromStart := CalcManhattanDistance(Coordinate{0, 0}, intersection)
			fmt.Printf("Intersection found at %v with distance from start of %v.\n", intersection, distanceFromStart)

			if distanceFromStart < currentMinDistance {
				currentMinDistance = distanceFromStart
			}
		}
	}

	return currentMinDistance
}

func travelPath(path []WirePathElement) []Coordinate {
	coordinates := []Coordinate{Coordinate{0, 0}}

	currentCoordinate := Coordinate{0, 0}
	for _, pathElement := range path {
		switch pathElement.Direction {
		case Up:
			currentCoordinate.Y += pathElement.Length
		case Left:
			currentCoordinate.X -= pathElement.Length
		case Down:
			currentCoordinate.Y -= pathElement.Length
		case Right:
			currentCoordinate.X += pathElement.Length
		}

		coordinates = append(coordinates, Coordinate{currentCoordinate.X, currentCoordinate.Y})
	}

	return coordinates
}

func intersection(firstEdge, secondEdge Edge) (bool, Coordinate) {
	var intersection Coordinate

	firstEdgeMinX := int(math.Min(float64(firstEdge.firstPoint.X), float64(firstEdge.secondPoint.X)))
	firstEdgeMaxX := int(math.Max(float64(firstEdge.firstPoint.X), float64(firstEdge.secondPoint.X)))
	firstEdgeMinY := int(math.Min(float64(firstEdge.firstPoint.Y), float64(firstEdge.secondPoint.Y)))
	firstEdgeMaxY := int(math.Max(float64(firstEdge.firstPoint.Y), float64(firstEdge.secondPoint.Y)))

	secondEdgeMinX := int(math.Min(float64(secondEdge.firstPoint.X), float64(secondEdge.secondPoint.X)))
	secondEdgeMaxX := int(math.Max(float64(secondEdge.firstPoint.X), float64(secondEdge.secondPoint.X)))
	secondEdgeMinY := int(math.Min(float64(secondEdge.firstPoint.Y), float64(secondEdge.secondPoint.Y)))
	secondEdgeMaxY := int(math.Max(float64(secondEdge.firstPoint.Y), float64(secondEdge.secondPoint.Y)))

	if (firstEdgeMinX <= secondEdgeMaxX) &&
		(firstEdgeMaxX >= secondEdgeMinX) {
		if (secondEdgeMinY <= firstEdgeMaxY) &&
			(secondEdgeMaxY >= firstEdgeMinY) {
			if firstEdgeMinX == firstEdgeMaxX {
				intersectionX := firstEdgeMinX
				intersectionY := 0
				if firstEdgeMinY >= secondEdgeMaxY {
					intersectionY = firstEdgeMinY
				} else {
					intersectionY = secondEdgeMinY
				}

				return true, Coordinate{intersectionX, intersectionY}
			}

			intersectionX := 0
			if firstEdgeMinX >= secondEdgeMaxX {
				intersectionX = firstEdgeMinX
			} else {
				intersectionX = secondEdgeMinX
			}

			intersectionY := firstEdgeMinY

			return true, Coordinate{intersectionX, intersectionY}
		}
	}

	return false, intersection
}

// CalcManhattanDistance determines the length of the path between two points.
// See: https://en.wikipedia.org/wiki/Taxicab_geometry
func CalcManhattanDistance(first, second Coordinate) int {
	return int(math.Abs(float64(first.X-second.X)) + math.Abs(float64(first.Y-second.Y)))
}
