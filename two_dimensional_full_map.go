package astar

import "math"

type MapBoundaries struct {
	xSize, ySize int
}

func (boundaries MapBoundaries) hasCoordinates(coords CartesianCoordinates) bool {
	return coords.x >= 0 && coords.x < boundaries.xSize && coords.y >= 0 && coords.y < boundaries.ySize
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
	boundaries MapBoundaries
}

func (graph Graph) PointOf(coordinates CartesianCoordinates) TwoDimensionalPoint {
	if graph.boundaries.hasCoordinates(coordinates) {
		return TwoDimensionalPoint{coordinates: coordinates, mapBoundaries: graph.boundaries}
	}
	// todo should be an error
	return TwoDimensionalPoint{}
}

func MapOfSize(x, y int) Graph {
	return Graph{boundaries: MapBoundaries{xSize: x, ySize: y}}
}

func SquareMapOfSize(x int) Graph {
	return MapOfSize(x, x)
}

type TwoDimensionalPoint struct {
	coordinates   CartesianCoordinates
	mapBoundaries MapBoundaries
}

func (point TwoDimensionalPoint) AdjacentNodes() []Node {
	surroundingCoordinatesIgnoringMapBoundaries := point.coordinates.surroundingCoordinates()
	takeMapBoundariesIntoAccount := func(c CartesianCoordinates) bool {
		return point.mapBoundaries.hasCoordinates(c)
	}
	coordinatesOfNeighbours := filter(surroundingCoordinatesIgnoringMapBoundaries, takeMapBoundariesIntoAccount)
	return coordinatesToNodes(coordinatesOfNeighbours, point.mapBoundaries)
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

func coordinatesToNodes(coordinates []CartesianCoordinates, boundaries MapBoundaries) []Node {
	points := make([]Node, 0)
	for _, singleCoordinate := range coordinates {
		points = append(points, TwoDimensionalPoint{coordinates: singleCoordinate, mapBoundaries: boundaries})
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

func contains(nodes []Node, node Node) bool {
	for _, t := range nodes {
		if t == node {
			return true
		}
	}
	return false
}