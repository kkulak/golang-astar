package astar

import "math"

func SomeStrangeHeuristics(from, to AStarNode) float64 {
	minStepsX := calculateMinNumberOfSteps(from.state.coordinates.x, from.state.velocity.x, to.state.coordinates.x)
	minStepsY := calculateMinNumberOfSteps(from.state.coordinates.y, from.state.velocity.y, to.state.coordinates.y)
	return math.Max(float64(minStepsX), float64(minStepsY))
}

func calculateMinNumberOfSteps(start, speed, end int) int {
	distance := int(math.Abs(float64(start - end)))
	var speedTowardsTarget = speed
	if (end < start) {
		speedTowardsTarget = -speedTowardsTarget
	}
	return calculateMinNumberOfStepsForDistanceAndSpeed(distance, speedTowardsTarget)
}

func calculateMinNumberOfStepsForDistanceAndSpeed(distance, speedTowardsTarget int) int {
	if (speedTowardsTarget < 0) {
		speedAwayFromTarget := -speedTowardsTarget
		return minStepsToGoThrough(distanceWhileStoppingFrom(speedAwayFromTarget) + distance) + speedAwayFromTarget
	} else {
		distanceWhileStoppingFromStartSpeed := distanceWhileStoppingFrom(speedTowardsTarget)
		if (distanceWhileStoppingFromStartSpeed > distance) {
			return minStepsToGoThrough(distanceWhileStoppingFromStartSpeed - distance) + speedTowardsTarget
		} else {
			return minStepsToGoThroughStartingAtVo(distance - distanceWhileStoppingFromStartSpeed, speedTowardsTarget) + speedTowardsTarget
		}
	}
}

func distanceWhileStoppingFrom(v int) int {
	return v2o(v)
}

func minStepsToGoThrough(x int) int {
	return int(math.Min(float64(o2v2oSteps(x)), float64(o2vv2oSteps(x))))
}

func minStepsToGoThroughStartingAtVo(x, v0 int) int {
	return int(math.Min(float64(vo2v2voSteps(x, v0)), float64(vo2vv2voSteps(x, v0))))
}

func v2o(v int) int {
	return v * (v - 1) / 2
}

func o2v2oSteps(x int) int {
	return int(math.Ceil(2 * math.Sqrt(float64(x))))
}

func o2vv2oSteps(x int) int {
	return int(math.Ceil(math.Sqrt(1 + 4 * float64(x))))
}

func vo2v2voSteps(x, v0 int) int {
	return int(math.Ceil(2 * (math.Sqrt(float64((x + v0 * v0) - v0 * v0)))))
}

func vo2vv2voSteps(x, v0 int) int {
	return int(math.Ceil(math.Sqrt(float64(1 + 4 * (x + v0 * v0))) - float64(2 * v0 * v0)))
}
