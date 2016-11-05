package astar

import (
	"testing"
	"reflect"
	"math"
)

func Test__Works_For_Square_Map_With_Diagonal_Path_Without_Obstacles(t *testing.T) {
	// * O O
	// O * O
	// O O *

	// given
	squareMap := SquareMapOfSize(3, noObstacles())
	start := squareMap.PointOf(coordinates(0, 0))
	end := squareMap.PointOf(coordinates(2, 2))

	// when
	distance, diagonalPath := AStar(start, end)

	// then
	assertDistanceEqual(t, distance, 2 * math.Sqrt(2))

	// and
	pointInBetween := squareMap.PointOf(coordinates(1, 1))
	assertPathEqual(t, diagonalPath, []Node{start, pointInBetween, end})
}

func Test__Works_For_Square_Map_With_Obstacles(t *testing.T) {
	// * O O X *
	// O * O X *
	// O X * X *
	// O O * X *
	// O O O * O

	// given
	obstacles := obstaclePoints(obstaclePoint(coordinates(1, 2)), verticalObstacleLine(coordinates(3, 0), coordinates(3, 3)))
	mapWithObstacles := SquareMapOfSize(5, obstacles)
	start := mapWithObstacles.PointOf(coordinates(0, 0))
	end := mapWithObstacles.PointOf(coordinates(4, 0))

	// when
	distance, actualPath := AStar(start, end)

	// then
	assertDistanceEqual(t, distance, 4 * math.Sqrt(2) + 4)

	// and
	expectedPath := []Node{
		mapWithObstacles.PointOf(coordinates(0, 0)),
		mapWithObstacles.PointOf(coordinates(1, 1)),
		mapWithObstacles.PointOf(coordinates(2, 2)),
		mapWithObstacles.PointOf(coordinates(2, 3)),
		mapWithObstacles.PointOf(coordinates(3, 4)),
		mapWithObstacles.PointOf(coordinates(4, 3)),
		mapWithObstacles.PointOf(coordinates(4, 2)),
		mapWithObstacles.PointOf(coordinates(4, 1)),
		mapWithObstacles.PointOf(coordinates(4, 0))}

	assertPathEqual(t, actualPath, expectedPath)
}

func assertDistanceEqual(t *testing.T, actualDistance, expectedDistance float64) {
	if actualDistance != expectedDistance {
		t.Errorf("Wrong distance. Got %f, expected %f.", actualDistance, expectedDistance)
	}

}
func assertPathEqual(t *testing.T, actualPath, expectedPath []Node) {
	if !reflect.DeepEqual(actualPath, expectedPath) {
		t.Errorf("Wrong path. Got %s, expected %s.", actualPath, expectedPath)
	}
}

func noObstacles() []CartesianCoordinates {
	return []CartesianCoordinates{}
}

func obstaclePoints(obstacles ...[]CartesianCoordinates) []CartesianCoordinates {
	allObstacles := make([]CartesianCoordinates, 0)
	for _, obstacle := range obstacles {
		allObstacles = append(allObstacles, obstacle...)
	}
	return allObstacles
}

func obstaclePoint(coordinate CartesianCoordinates) []CartesianCoordinates {
	return []CartesianCoordinates{coordinate}
}

func verticalObstacleLine(from, to CartesianCoordinates) []CartesianCoordinates {
	obstacles := make([]CartesianCoordinates, 0)
	xPosition := from.x
	yPosition := from.y
	for yPosition <= to.y {
		obstacles = append(obstacles, coordinates(xPosition, yPosition))
		yPosition = yPosition + 1
	}
	return obstacles
}

func coordinates(x, y int) CartesianCoordinates {
	return CartesianCoordinates{x: x, y: y}
}