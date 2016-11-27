package astar

import (
	"testing"
	"math"
)

func Test__Works_For_Square_Map_With_Diagonal_Path_Without_Obstacles(t *testing.T) {
	// * O O
	// O * O
	// O O *

	// given
	_, start, end := GraphFromBitmap("resources/3x3_empty_terrain")

	// when
	distance, actualPath := AStar(start, end)

	// then
	PersistGraphToBitmap(actualPath, "resources/3x3_empty_terrain")

	// and
	AssertDistanceEqual(t, distance, 2 * math.Sqrt(2))

	// and
	pointInBetween := TracePoint(1, 1, 1, 1)
	expectedDiagonalPath := []AStarNodeState{AsTracePoint(start), pointInBetween, AsTracePoint(end)}
	AssertTraceEqual(t, actualPath, expectedDiagonalPath)
}

func Test__Works_For_Square_Map_With_Obstacles(t *testing.T) {
	// * O O X *
	// O * O X *
	// O X * X *
	// O O * X *
	// O O O * O

	// given
	_, start, end := GraphFromBitmap("resources/5x5_terrain_with_obstacles")

	// when
	distance, diagonalPath := AStar(start, end)

	// then
	PersistGraphToBitmap(diagonalPath, "resources/5x5_terrain_with_obstacles")
	AssertDistanceEqual(t, distance, 4 * math.Sqrt(2) + 4)

	// and
	//expectedPath := []Node{
	//	mapWithObstacles.PointOf(coordinates(0, 0)),
	//	mapWithObstacles.PointOf(coordinates(1, 1)),
	//	mapWithObstacles.PointOf(coordinates(2, 2)),
	//	mapWithObstacles.PointOf(coordinates(2, 3)),
	//	mapWithObstacles.PointOf(coordinates(3, 4)),
	//	mapWithObstacles.PointOf(coordinates(4, 3)),
	//	mapWithObstacles.PointOf(coordinates(4, 2)),
	//	mapWithObstacles.PointOf(coordinates(4, 1)),
	//	mapWithObstacles.PointOf(coordinates(4, 0))}
	//
	//assertPathEqual(t, actualPath, expectedPath)
}
