package astar

import (
	"testing"
	"math"
)

func AssertDistanceEqual(t *testing.T, actualDistance, expectedDistance float64) {
	epsilon := 10e-3
	if math.Abs(actualDistance - expectedDistance) >= epsilon {
		t.Errorf("Wrong distance. Got %f, expected %f.", actualDistance, expectedDistance)
	}

}
