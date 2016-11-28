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
	expectedDiagonalPath := []AStarNodeState{
		TracePoint(0, 0, 0, 0),
		TracePoint(1, 1, 1, 1),
		TracePoint(2, 2, 1, 1),
		TracePoint(2, 2, 0, 0)}

	AssertTraceEqual(t, actualPath, expectedDiagonalPath)
}

func Test__Works_For_Square_Map_With_Obstacles(t *testing.T) {
	// * O O X *
	// O * O X *
	// O X O X O
	// O O * X *
	// O O O * O

	// given
	_, start, end := GraphFromBitmap("resources/5x5_terrain_with_obstacles")

	// when
	distance, actualPath := AStar(start, end)

	// then
	PersistGraphToBitmap(actualPath, "resources/5x5_terrain_with_obstacles")

	// and:
	AssertDistanceEqual(t, distance, math.Sqrt(2) + math.Sqrt(5) + + 2 * math.Sqrt(2) + 3)

	// and
	expectedPath := []AStarNodeState{
		TracePoint(0, 0, 0, 0),
		TracePoint(1, 1, 1, 1),
		TracePoint(2, 3, 1, 2),
		TracePoint(3, 4, 1, 1),
		TracePoint(3, 4, 0, 0),
		TracePoint(4, 3, 1, -1),
		TracePoint(4, 2, 0, -1),
		TracePoint(4, 1, 0, -1),
		TracePoint(4, 0, 0, -1),
		TracePoint(4, 0, 0, 0)}

	AssertTraceEqual(t, actualPath, expectedPath)
}

func Test__Jumps_Over_Obstacles_When_Velocity_Is_Sufficient(t *testing.T) {
	// * * O * X * * *
	// O O O O X O O O
	// O O X O X O O O
	// O O O O X O O O
	// O O O O O O O O

	// given
	_, start, end := GraphFromBitmap("resources/8x5_terrain_with_obstacles")

	// when
	distance, actualPath := AStar(start, end)

	// then
	PersistGraphToBitmap(actualPath, "resources/8x5_terrain_with_obstacles")

	// and:
	AssertDistanceEqual(t, distance, 7)

	// and
	expectedPath := []AStarNodeState{
		TracePoint(0, 0, 0, 0),
		TracePoint(1, 0, 1, 0),
		TracePoint(2, 0, 1, 0),
		TracePoint(3, 0, 1, 0),
		TracePoint(5, 0, 2, 0),
		TracePoint(6, 0, 1, 0),
		TracePoint(7, 0, 1, 0),
		TracePoint(7, 0, 0, 0)}

	AssertTraceEqual(t, actualPath, expectedPath)
}
