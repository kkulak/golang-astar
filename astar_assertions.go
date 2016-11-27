package astar

import (
	"reflect"
	"testing"
	"math"
)

func AssertDistanceEqual(t *testing.T, actualDistance, expectedDistance float64) {
	epsilon := 10e-3
	if math.Abs(actualDistance - expectedDistance) >= epsilon {
		t.Errorf("Wrong distance. Got %f, expected %f.", actualDistance, expectedDistance)
	}

}
func AssertTraceEqual(t *testing.T, actualTrace []Node, expectedTrace []AStarNodeState) {
	assertTraceAsNodeStateEqual(t, toAStarNodeState(actualTrace), expectedTrace)
}

func assertTraceAsNodeStateEqual(t *testing.T, actualTrace, expectedTrace []AStarNodeState) {
	if !reflect.DeepEqual(actualTrace, expectedTrace) {
		t.Errorf("Wrong trace. Got %v, expected %v.", actualTrace, expectedTrace)
	}
}

func toAStarNodeState(trace []Node) []AStarNodeState {
	traceWithVelocity := make([]AStarNodeState, 0)
	for _, node := range trace {
		traceWithVelocity = append(traceWithVelocity, asTracePoint(node))
	}
	return traceWithVelocity
}

func asTracePoint(node Node) AStarNodeState {
	return node.(AStarNode).state
}

func TracePoint(x, y, vx, vy int) AStarNodeState {
	return AStarNodeState{coordinates: Coordinates{x: x, y: y}, velocity: Velocity{x: vx, y: vy}}
}