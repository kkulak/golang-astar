package astar

import "github.com/deckarep/golang-set"

func filter(coordinates []AStarNodeState, f func(AStarNodeState) bool) []AStarNodeState {
	filteredCoordinates := make([]AStarNodeState, 0)
	for _, point := range coordinates {
		if f(point) {
			filteredCoordinates = append(filteredCoordinates, point)
		}
	}
	return filteredCoordinates
}

func coordinatesToNodes(coordinates []AStarNodeState, boundaries MapShape) []Node {
	points := make([]Node, 0)
	for _, aStarNodeState := range coordinates {
		points = append(points, AStarNode{state: aStarNodeState, aMap: boundaries})
	}
	return points
}

// todo duplication
func contains(nodes []Node, node Node) bool {
	for _, t := range nodes {
		if t == node {
			return true
		}
	}
	return false
}

func MapOfSize(x int, y int, obstacles mapset.Set) Graph {
	return Graph{aMap: MapShape{xSize: x, ySize: y, obstacles: obstacles}}
}
