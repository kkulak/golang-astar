package astar

import "math"

type MapShape struct {
	xSize, ySize int
	obstacles    *[]AStarNodeState
}

func (aMap MapShape) containsPoint(coords Coordinates) bool {
	return !aMap.containsObstacle(coords) && aMap.withinBorders(coords)
}

func (aMap MapShape) containsObstacle(point Coordinates) bool {
	for _, o := range *aMap.obstacles {
		if o.coordinates == point {
			return true
		}
	}
	return false
}

func (aMap MapShape) withinBorders(point Coordinates) bool {
	return point.x >= 0 && point.x < aMap.xSize && point.y >= 0 && point.y < aMap.ySize
}

type Coordinates struct {
	x, y int
}

type Velocity struct {
	x, y int
}

type AStarNodeState struct {
	coordinates   Coordinates
	velocity      Velocity
}

func (state AStarNodeState) availableStates() []AStarNodeState {
	return []AStarNodeState{
		{Coordinates{x: state.coordinates.x + state.velocity.x, y: state.coordinates.y + state.velocity.y},
		 Velocity   {x: state.velocity.x, y: state.velocity.y}},
		{Coordinates{x: state.coordinates.x + state.velocity.x, y: state.coordinates.y + state.velocity.y + 1},
		 Velocity   {x: state.velocity.x, y: state.velocity.y + 1}},
		{Coordinates{x: state.coordinates.x + state.velocity.x, y: state.coordinates.y + state.velocity.y - 1},
		 Velocity{x: state.velocity.x, y: state.velocity.y - 1}},
		{Coordinates{x: state.coordinates.x + state.velocity.x + 1, y: state.coordinates.y + state.velocity.y},
		 Velocity{x: state.velocity.x + 1, y: state.velocity.y}},
		{Coordinates{x: state.coordinates.x + state.velocity.x + 1, y: state.coordinates.y + state.velocity.y + 1},
		 Velocity{x: state.velocity.x + 1, y: state.velocity.y + 1}},
		{Coordinates{x: state.coordinates.x + state.velocity.x + 1, y: state.coordinates.y + state.velocity.y - 1},
		 Velocity{x: state.velocity.x + 1, y: state.velocity.y - 1}},
		{Coordinates{x: state.coordinates.x + state.velocity.x - 1, y: state.coordinates.y + state.velocity.y},
		 Velocity{x: state.velocity.x - 1, y: state.velocity.y}},
		{Coordinates{x: state.coordinates.x + state.velocity.x - 1, y: state.coordinates.y + state.velocity.y + 1},
		 Velocity{x: state.velocity.x - 1, y: state.velocity.y + 1}},
		{Coordinates{x: state.coordinates.x + state.velocity.x - 1, y: state.coordinates.y + state.velocity.y - 1},
		 Velocity{x: state.velocity.x - 1, y: state.velocity.y - 1}}}
}

type Graph struct {
	aMap MapShape
}

func (graph Graph) PointOf(coordinates AStarNodeState) AStarNode {
	if graph.aMap.containsPoint(coordinates.coordinates) {
		return AStarNode{state: coordinates, aMap: graph.aMap}
	}
	// todo should be an error
	return AStarNode{}
}

type AStarNode struct {
	state AStarNodeState
	aMap  MapShape
}

func (point AStarNode) AdjacentNodes() []Node {
	availableStatesIgnoringMapBoundaries := point.state.availableStates()
	takeMapBoundariesIntoAccount := func(c AStarNodeState) bool {
		return point.aMap.containsPoint(c.coordinates)
	}
	coordinatesOfNeighbours := filter(availableStatesIgnoringMapBoundaries, takeMapBoundariesIntoAccount)
	return coordinatesToNodes(coordinatesOfNeighbours, point.aMap)
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
	return math.Sqrt(math.Pow(float64(to.state.coordinates.x - from.state.coordinates.x), 2) + math.Pow(float64(to.state.coordinates.y - from.state.coordinates.y), 2))
}

func manhattanDistance(from, to AStarNode) float64 {
	return math.Abs((float64(to.state.coordinates.x - from.state.coordinates.x)) + math.Abs(float64(to.state.coordinates.y - from.state.coordinates.y)))
}