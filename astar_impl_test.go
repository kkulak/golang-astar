package astar

import (
	"testing"
	"reflect"
	"math"
)

func Test__Works_For_Square_Map_With_Diagonal_Path(t *testing.T) {
	// given
	squareMap := SquareMapOfSize(3)
	start := squareMap.PointOf(CartesianCoordinates{x: 0, y: 0})
	end := squareMap.PointOf(CartesianCoordinates{x: 2, y: 2})

	// when
	distance, diagonalPath := AStar(start, end)

	// then
	if distance != 2 * math.Sqrt(2) {
		t.Error("wrong path length")
	}

	// and
	pointInBetween := squareMap.PointOf(CartesianCoordinates{x: 1, y: 1})
	if !reflect.DeepEqual(diagonalPath, []Node{start, pointInBetween, end}) {
		t.Error("invalid path")
	}
}

