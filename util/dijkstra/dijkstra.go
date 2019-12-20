package dijkstra

import (
	"math"
	"sort"
)

// ShortestPath applies the Dijkstra algorithm to the graph and finds the shortest path from the given start to the destination.
func ShortestPath(from, to interface{}, graph Graph) int {
	nodeWeights := make(map[interface{}]int)
	for node := range graph {
		nodeWeights[node] = math.MaxInt32
	}

	nodeWeights[from] = 0

	shortestPathImpl(to, from, graph, make(map[interface{}]struct{}), nodeWeights)

	return nodeWeights[to]
}

func shortestPathImpl(to, current interface{}, graph Graph, completedNodes map[interface{}]struct{}, nodeWeights map[interface{}]int) {
	if current == to {
		return
	}

	if _, visited := completedNodes[current]; visited {
		return
	}

	currentNodeEdges := graph[current]
	if len(currentNodeEdges) == 0 {
		return
	}

	sort.Slice(currentNodeEdges, func(i, j int) bool {
		firstEdge := currentNodeEdges[i]
		secondEdge := currentNodeEdges[j]

		return firstEdge.Length < secondEdge.Length
	})

	currentNodeWeight := nodeWeights[current]
	nodesVisited := []node{}

	for _, edge := range currentNodeEdges {
		targetNodeWeight := nodeWeights[edge.To]
		if targetNodeWeight > currentNodeWeight+edge.Length {
			targetNodeWeight = currentNodeWeight + edge.Length
			nodeWeights[edge.To] = targetNodeWeight
		}

		nodesVisited = append(nodesVisited, node{name: edge.To, weight: targetNodeWeight})
	}

	completedNodes[current] = struct{}{}

	sort.Slice(nodesVisited, func(i, j int) bool {
		firstNode := nodesVisited[i]
		secondNode := nodesVisited[j]

		return firstNode.weight < secondNode.weight
	})

	for _, node := range nodesVisited {
		shortestPathImpl(to, node.name, graph, completedNodes, nodeWeights)
	}
}

// Graph describes a weighted graph.
type Graph map[interface{}][]Edge

// Edge describes an edge in a weighted graph.
type Edge struct {
	To     interface{}
	Length int
}

type node struct {
	name   interface{}
	weight int
}
