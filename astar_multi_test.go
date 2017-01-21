package astar

import (
	"testing"
)

func Test__08(t *testing.T) {
	// given
	start, end := ReadMapForMultiplePoints("resources/test08")

	// when
	distance, actualPath := AStar(start, end)

	// then
	PersistGraphWithMultiplePointsToBitmap(actualPath, "resources/test08")

	// and
	AssertDistanceEqual(t, distance, 5)
}
