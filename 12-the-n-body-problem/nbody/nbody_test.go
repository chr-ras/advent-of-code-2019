package nbody

import (
	"testing"

	v "github.com/chr-ras/advent-of-code-2019/util/geometry/vector3"
)

func TestSimulateJupiterMoons(t *testing.T) {
	cases := []struct {
		positions []v.Vector3
		steps     int
		expected  float64
	}{
		{
			[]v.Vector3{
				v.Vector3{X: -1, Y: 0, Z: 2},
				v.Vector3{X: 2, Y: -10, Z: -7},
				v.Vector3{X: 4, Y: -8, Z: 8},
				v.Vector3{X: 3, Y: 5, Z: -1},
			},
			10,
			179.0,
		},
		{
			[]v.Vector3{
				v.Vector3{X: -8, Y: -10, Z: 0},
				v.Vector3{X: 5, Y: 5, Z: 10},
				v.Vector3{X: 2, Y: -7, Z: 3},
				v.Vector3{X: 9, Y: -8, Z: -3},
			},
			100,
			1940.0,
		},
	}

	for _, c := range cases {
		actual := SimulateJupiterMoons(append([]v.Vector3(nil), c.positions...), c.steps)

		if actual != c.expected {
			t.Errorf("SimulateJupiterMoons(%v, %v) == %v, expected %v", c.positions, c.steps, actual, c.expected)
		}
	}
}
