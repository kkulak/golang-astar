package astar

import (
	"reflect"
	"testing"
)

func AssertDistanceEqual(t *testing.T, actualDistance, expectedDistance float64) {
	if actualDistance != expectedDistance {
		t.Errorf("Wrong distance. Got %f, expected %f.", actualDistance, expectedDistance)
	}

}
func AssertTraceEqual(t *testing.T, actualTrace []Node, expectedTrace []AStarNodeState) {
	assertTraceAsNodeStateEqual(t, normalized(toAStarNodeState(actualTrace)), expectedTrace)
}

func normalized(trace []AStarNodeState) []AStarNodeState {
	if len(trace) < 2 {
		return trace
	}

	tracePrefix := trace[:len(trace) - 2]
	lastTwoElements := trace[len(trace) - 2:]
	if sameTracePointsConsideringDeceleration(lastTwoElements[0], lastTwoElements[1]) {
		return append(tracePrefix, lastTwoElements[1])
	} else {
		return append(tracePrefix, lastTwoElements...)
	}
}

func sameTracePointsConsideringDeceleration(first, second AStarNodeState) bool {
	sameCoordinates := first.x == second.x && first.y == second.y

	velocityDifference := (first.vx - second.vx) + (first.vy - second.vy)
	ableToStopInNextStep := velocityDifference >= 1 && velocityDifference <= 2

	return sameCoordinates && ableToStopInNextStep
}

func assertTraceAsNodeStateEqual(t *testing.T, actualTrace, expectedTrace []AStarNodeState) {
	if !reflect.DeepEqual(actualTrace, expectedTrace) {
		t.Errorf("Wrong trace. Got %v, expected %v.", actualTrace, expectedTrace)
	}
}

func toAStarNodeState(trace []Node) []AStarNodeState {
	traceWithVelocity := make([]AStarNodeState, 0)
	for _, node := range trace {
		traceWithVelocity = append(traceWithVelocity, AsTracePoint(node))
	}
	return traceWithVelocity
}

func AsTracePoint(node Node) AStarNodeState {
	return node.(AStarNode).state
}

func TracePoint(x, y, vx, vy int) AStarNodeState {
	return AStarNodeState{x: x, y: y, vx: vx, vy: vy}
}