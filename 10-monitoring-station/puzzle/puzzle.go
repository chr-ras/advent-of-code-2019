package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/chr-ras/advent-of-code-2019/10-monitoring-station/monitoring"
	"github.com/chr-ras/advent-of-code-2019/10-monitoring-station/vaporization"
	g "github.com/chr-ras/advent-of-code-2019/util/geometry"
)

func main() {
	file, err := os.Open("./asteroid_map.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	scanner := bufio.NewScanner(file)
	asteroidMap := []string{}

	for scanner.Scan() {
		asteroidMap = append(asteroidMap, scanner.Text())
	}

	part1(asteroidMap)
	part2(asteroidMap)
}

func part1(asteroidMap []string) {
	asteroid, visibleAsteroids := monitoring.BestAsteroidForMonitoringStation(asteroidMap)

	fmt.Printf("Best asteroid %v with %v visible asteroids.\n", asteroid, visibleAsteroids)
}

func part2(asteroidMap []string) {
	vaporization.VaporizeAsteroids(asteroidMap, g.Point{X: 29, Y: 28})
}
