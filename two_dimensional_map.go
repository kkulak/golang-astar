package astar

import "math"

type MapShape struct {
	xSize, ySize int
	obstacles    *[]CartesianCoordinates
}

func (aMap MapShape) containsPoint(coords CartesianCoordinates) bool {
	return !aMap.containsObstacle(coords) && aMap.withinBorders(coords)
}

func (aMap MapShape) containsObstacle(point CartesianCoordinates) bool {
	for _, o := range *aMap.obstacles {
		if o == point {
			return true
		}
	}
	return false
}

func (aMap MapShape) withinBorders(point CartesianCoordinates) bool {
	return point.x >= 0 && point.x < aMap.xSize && point.y >= 0 && point.y < aMap.ySize
}

type CartesianCoordinates struct {
	x, y int
}

func (coords CartesianCoordinates) surroundingCoordinates() []CartesianCoordinates {
	return []CartesianCoordinates{
		{coords.x - 1, coords.y - 1},
		{coords.x - 1, coords.y},
		{coords.x - 1, coords.y + 1},
		{coords.x, coords.y - 1 },
		{coords.x, coords.y + 1},
		{coords.x + 1, coords.y - 1},
		{coords.x + 1, coords.y},
		{coords.x + 1, coords.y + 1}}
}

type Graph struct {
	aMap MapShape
}

func (graph Graph) PointOf(coordinates CartesianCoordinates) TwoDimensionalPoint {
	if graph.aMap.containsPoint(coordinates) {
		return TwoDimensionalPoint{coordinates: coordinates, aMap: graph.aMap}
	}
	// todo should be an error
	return TwoDimensionalPoint{}
}

func MapOfSize(x int, y int, obstacles []CartesianCoordinates) Graph {
	return Graph{aMap: MapShape{xSize: x, ySize: y, obstacles: &obstacles}}
}

func SquareMapOfSize(x int, obstacles []CartesianCoordinates) Graph {
	return MapOfSize(x, x, obstacles)
}

type TwoDimensionalPoint struct {
	coordinates CartesianCoordinates
	aMap        MapShape
}

func (point TwoDimensionalPoint) AdjacentNodes() []Node {
	surroundingCoordinatesIgnoringMapBoundaries := point.coordinates.surroundingCoordinates()
	takeMapBoundariesIntoAccount := func(c CartesianCoordinates) bool {
		return point.aMap.containsPoint(c)
	}
	coordinatesOfNeighbours := filter(surroundingCoordinatesIgnoringMapBoundaries, takeMapBoundariesIntoAccount)
	return coordinatesToNodes(coordinatesOfNeighbours, point.aMap)
}

func filter(coordinates []CartesianCoordinates, f func(CartesianCoordinates) bool) []CartesianCoordinates {
	filteredCoordinates := make([]CartesianCoordinates, 0)
	for _, point := range coordinates {
		if f(point) {
			filteredCoordinates = append(filteredCoordinates, point)
		}
	}
	return filteredCoordinates
}

func coordinatesToNodes(coordinates []CartesianCoordinates, boundaries MapShape) []Node {
	points := make([]Node, 0)
	for _, singleCoordinate := range coordinates {
		points = append(points, TwoDimensionalPoint{coordinates: singleCoordinate, aMap: boundaries})
	}
	return points
}

func (source TwoDimensionalPoint) Cost(destination Node) float64 {
	destinationCasted := destination.(TwoDimensionalPoint)
	if contains(source.AdjacentNodes(), destinationCasted) {
		return cartesianDistance(source, destinationCasted)
	}

	// todo refactor
	return math.MaxFloat64
}

func (source TwoDimensionalPoint) EstimatedCost(destination Node) float64 {
	destinationCasted := destination.(TwoDimensionalPoint)
	return manhattanDistance(source, destinationCasted)
}

func cartesianDistance(from, to TwoDimensionalPoint) float64 {
	return math.Sqrt(math.Pow(float64(to.coordinates.x - from.coordinates.x), 2) + math.Pow(float64(to.coordinates.y - from.coordinates.y), 2))
}

func manhattanDistance(from, to TwoDimensionalPoint) float64 {
	return math.Abs((float64(to.coordinates.x - from.coordinates.x)) + math.Abs(float64(to.coordinates.y - from.coordinates.y)))
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