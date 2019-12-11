package monitoring

import (
	"reflect"
	"testing"

	g "github.com/chr-ras/advent-of-code-2019/util/geometry"
)

var lineOfSightCases = []struct {
	asteroidMap          []string
	asteroidX, asteroidY int
	expectedVisibleCount int
}{
	{
		[]string{
			".#..#",
			".....",
			"#####",
			"....#",
			"...##",
		},
		3, 4, 8,
	},
	{
		[]string{
			"......#.#.",
			"#..#.#....",
			"..#######.",
			".#.#.###..",
			".#..#.....",
			"..#....#.#",
			"#..#....#.",
			".##.#..###",
			"##...#..#.",
			".#....####",
		},
		5, 8, 33,
	},
	{
		[]string{
			"#.#...#.#.",
			".###....#.",
			".#....#...",
			"##.#.#.#.#",
			"....#.#.#.",
			".##..###.#",
			"..#...##..",
			"..##....##",
			"......#...",
			".####.###.",
		},
		1, 2, 35,
	},
	{
		[]string{
			".#..#..###",
			"####.###.#",
			"....###.#.",
			"..###.##.#",
			"##.##.#.#.",
			"....###..#",
			"..#.#..#.#",
			"#..#.#.###",
			".##...##.#",
			".....#.#..",
		},
		6, 3, 41,
	},
	{
		[]string{
			".#..##.###...#######",
			"##.############..##.",
			".#.######.########.#",
			".###.#######.####.#.",
			"#####.##.#.##.###.##",
			"..#####..#.#########",
			"####################",
			"#.####....###.#.#.##",
			"##.#################",
			"#####.##.###..####..",
			"..######..##.#######",
			"####.##.####...##..#",
			".#####..#.######.###",
			"##...#.##########...",
			"#.##########.#######",
			".####.#.###.###.#.##",
			"....##.##.###..#####",
			".#.#.###########.###",
			"#.#.#.#####.####.###",
			"###.##.####.##.#..##",
		},
		11, 13, 210,
	},
}

func TestCheckLineOfSight(t *testing.T) {
	// Adding the demonstration case from the puzzle description showing which asteroids will be visible.
	cases := append(lineOfSightCases, struct {
		asteroidMap          []string
		asteroidX, asteroidY int
		expectedVisibleCount int
	}{
		[]string{
			"#.........",
			"...#......",
			"...#..#...",
			".####....#",
			"..#.#.#...",
			".....#....",
			"..###.#.##",
			".......#..",
			"....#...#.",
			"...#..#..#",
		},
		0, 0, 7,
	})

	for _, c := range cases {
		asteroidToCheck := g.Point{X: c.asteroidX, Y: c.asteroidY}
		actual, _ := CheckLineOfSight(c.asteroidMap, asteroidToCheck)

		if actual != c.expectedVisibleCount {
			t.Errorf("CheckLineOfSight(%v, %v) == %v, expected %v", c.asteroidMap, asteroidToCheck, actual, c.expectedVisibleCount)
		}
	}
}

func TestBestAsteroidForMonitoringStation(t *testing.T) {
	for _, c := range lineOfSightCases {
		expectedAsteroid := g.Point{X: c.asteroidX, Y: c.asteroidY}
		actualAsteroid, actualVisibleAsteroids := BestAsteroidForMonitoringStation(c.asteroidMap)
		if !reflect.DeepEqual(actualAsteroid, expectedAsteroid) || actualVisibleAsteroids != c.expectedVisibleCount {
			t.Errorf("BestAsteroidForMonitoringStation(%v) == %v, %v, expected %v, %v", c.asteroidMap, actualAsteroid, actualVisibleAsteroids, expectedAsteroid, c.expectedVisibleCount)
		}
	}
}
