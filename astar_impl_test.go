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
	squareMap, start, end := GraphFromBitmap("resources/3x3_empty_terrain.bmp")

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
	mapWithObstacles, start, end := GraphFromBitmap("resources/5x5_terrain_with_obstacles.bmp")

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

func coordinates(x, y int) CartesianCoordinates {
	return CartesianCoordinates{x: x, y: y}
}