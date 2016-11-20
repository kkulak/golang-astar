package astar

import "math"

type MapShape struct {
	xSize, ySize int
	obstacles    *[]AStarNodeState
}

func (aMap MapShape) containsPoint(coords AStarNodeState) bool {
	return !aMap.containsObstacle(coords) && aMap.withinBorders(coords)
}

func (aMap MapShape) containsObstacle(point AStarNodeState) bool {
	for _, o := range *aMap.obstacles {
		if o.x == point.x && o.y == point.y {
			return true
		}
	}
	return false
}

func (aMap MapShape) withinBorders(point AStarNodeState) bool {
	return point.x >= 0 && point.x < aMap.xSize && point.y >= 0 && point.y < aMap.ySize
}

type AStarNodeState struct {
	x, y   int
	vx, vy int
}

func (state AStarNodeState) availableStates() []AStarNodeState {
	return []AStarNodeState{
		{state.x + state.vx, state.y + state.vy, state.vx, state.vy},
		{state.x + state.vx, state.y + state.vy + 1, state.vx, state.vy + 1},
		{state.x + state.vx, state.y + state.vy - 1, state.vx, state.vy - 1},
		{state.x + state.vx + 1, state.y + state.vy, state.vx + 1, state.vy},
		{state.x + state.vx + 1, state.y + state.vy + 1, state.vx + 1, state.vy + 1},
		{state.x + state.vx + 1, state.y + state.vy - 1, state.vx + 1, state.vy - 1},
		{state.x + state.vx - 1, state.y + state.vy, state.vx - 1, state.vy},
		{state.x + state.vx - 1, state.y + state.vy + 1, state.vx - 1, state.vy + 1},
		{state.x + state.vx - 1, state.y + state.vy - 1, state.vx - 1, state.vy - 1}}
}

type Graph struct {
	aMap MapShape
}

func (graph Graph) PointOf(coordinates AStarNodeState) AStarNode {
	if graph.aMap.containsPoint(coordinates) {
		return AStarNode{state: coordinates, aMap: graph.aMap}
	}
	// todo should be an error
	return AStarNode{}
}

func MapOfSize(x int, y int, obstacles []AStarNodeState) Graph {
	return Graph{aMap: MapShape{xSize: x, ySize: y, obstacles: &obstacles}}
}

type AStarNode struct {
	state AStarNodeState
	aMap  MapShape
}

func (point AStarNode) AdjacentNodes() []Node {
	availableStatesIgnoringMapBoundaries := point.state.availableStates()
	takeMapBoundariesIntoAccount := func(c AStarNodeState) bool {
		return point.aMap.containsPoint(c)
	}
	coordinatesOfNeighbours := filter(availableStatesIgnoringMapBoundaries, takeMapBoundariesIntoAccount)
	return coordinatesToNodes(coordinatesOfNeighbours, point.aMap)
}

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

func (source AStarNode) Cost(destination Node) float64 {
	destinationCasted := destination.(AStarNode)
	if contains(source.AdjacentNodes(), destinationCasted) {
		return cartesianDistance(source, destinationCasted)
	}

	// todo refactor
	return math.MaxFloat64
}

func (source AStarNode) EstimatedCost(destination Node) float64 {
	destinationCasted := destination.(AStarNode)
	return manhattanDistance(source, destinationCasted)
}

func cartesianDistance(from, to AStarNode) float64 {
	return math.Sqrt(math.Pow(float64(to.state.x - from.state.x), 2) + math.Pow(float64(to.state.y - from.state.y), 2))
}

func manhattanDistance(from, to AStarNode) float64 {
	return math.Abs((float64(to.state.x - from.state.x)) + math.Abs(float64(to.state.y - from.state.y)))
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