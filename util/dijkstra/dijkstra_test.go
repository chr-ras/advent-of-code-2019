package dijkstra

import "testing"

func TestShortestPath(t *testing.T) {
	cases := []struct {
		from, to string
		graph    Graph
		expected int
	}{
		{
			"A", "F",
			Graph{
				"A": {
					Edge{To: "B", Length: 10},
					Edge{To: "C", Length: 14},
				},
				"B": {
					Edge{To: "C", Length: 3},
					Edge{To: "D", Length: 5},
					Edge{To: "E", Length: 6},
				},
				"C": {
					Edge{To: "E", Length: 5},
					Edge{To: "F", Length: 9},
				},
				"D": {
					Edge{To: "F", Length: 7},
				},
				"E": {
					Edge{To: "D", Length: 4},
					Edge{To: "F", Length: 3},
				},
				"F": {},
			},
			19,
		},
		{
			"A", "E",
			Graph{
				"A": {
					Edge{To: "B", Length: 20},
					Edge{To: "C", Length: 40},
					Edge{To: "D", Length: 30},
				},
				"B": {
					Edge{To: "D", Length: 34},
					Edge{To: "F", Length: 10},
				},
				"C": {
					Edge{To: "D", Length: 10},
					Edge{To: "E", Length: 50},
				},
				"D": {
					Edge{To: "E", Length: 30},
					Edge{To: "F", Length: 24},
				},
				"E": {},
				"F": {
					Edge{To: "E", Length: 10},
				},
			},
			40,
		},
	}

	for _, c := range cases {
		actual := ShortestPath(c.from, c.to, c.graph)

		if actual != c.expected {
			t.Errorf("ShortestPath(%v, %v, %v) == %d, expected %d", c.from, c.to, c.graph, actual, c.expected)
		}
	}
}
