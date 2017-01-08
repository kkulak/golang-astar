package astar

import (
	"testing"
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
	PersistGraphToFile(distance, actualPath, "resources/3x3_empty_terrain")

	// and
	AssertDistanceEqual(t, distance, 3)
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
	PersistGraphToFile(distance, actualPath, "resources/5x5_terrain_with_obstacles")

	// and:
	AssertDistanceEqual(t, distance, 8)
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
	PersistGraphToFile(distance, actualPath, "resources/8x5_terrain_with_obstacles")

	// and:
	AssertDistanceEqual(t, distance, 6)
}
