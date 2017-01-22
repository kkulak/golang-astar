package astar

type MultiPointAstarNode struct {
	points      [100]AStarNode
	// pozdrawiam pachola legutka
	destination [100]Coordinates
}

func (multiPoint MultiPointAstarNode) Points() []AStarNode {
	pointsAsSlice := make([]AStarNode, 0)
	for _, point := range multiPoint.points {
		if point == (AStarNode{}) {
			break
		}

		pointsAsSlice = append(pointsAsSlice, point)
	}
	return pointsAsSlice
}

func (multiPoint MultiPointAstarNode) Destination() []Coordinates {
	destinationAsSlice := make([]Coordinates, 0)
	for _, singleDest := range multiPoint.destination {
		if singleDest == (Coordinates{}) {
			break
		}

		destinationAsSlice = append(destinationAsSlice, singleDest)
	}
	return destinationAsSlice
}

func (multiPoint MultiPointAstarNode) Cost(other Node) float64 {
	thisPoints := multiPoint.Points()
	thatPoints := other.(MultiPointAstarNode).Points()

	totalCost := 0.0
	for i := range thisPoints {
		totalCost += thisPoints[i].Cost(thatPoints[i])
	}
	return totalCost
}

func (multiPoint MultiPointAstarNode)EstimatedCost(other Node) float64 {
	thisPoints := multiPoint.Points()
	thatPoints := other.(MultiPointAstarNode).Points()

	maxCost := 0.0
	for i := range thisPoints {
		costForSinglePoint := thisPoints[i].EstimatedCost(thatPoints[i])
		if (costForSinglePoint > maxCost) {
			maxCost = costForSinglePoint
		}
	}
	return maxCost
}

func (multiPoint MultiPointAstarNode) AdjacentNodes() []Node {
	allPermutations := possibleNextMovesOfAllPoints(multiPoint)
	withoutConflictingPositions := filterOutConflictingPositions(allPermutations, multiPoint.Destination())
	return asMultiPointNodes(withoutConflictingPositions, multiPoint.Destination())
}

func asMultiPointNodes(groupsOfPoints [][]AStarNode, destination []Coordinates) []Node {
	multiPointNodes := make([]Node, 0)
	for _, multiplePoints := range groupsOfPoints {
		multiPointNodes = append(multiPointNodes, MultiPointAstarNode{pointsAsArray(multiplePoints), destinationAsArray(destination)})
	}
	return multiPointNodes
}

func destinationAsArray(destination []Coordinates) [100]Coordinates {
	var destAsArray [100]Coordinates
	for i, coord := range destination {
		destAsArray[i] = coord
	}
	return destAsArray
}

func pointsAsArray(points []AStarNode) [100]AStarNode {
	var pointsAsArray [100]AStarNode
	for i, point := range points {
		pointsAsArray[i] = point
	}
	return pointsAsArray
}

func filterOutConflictingPositions(allStates [][]AStarNode, destination []Coordinates) [][]AStarNode {
	withoutDuplicatedPositions := make([][]AStarNode, 0)
	for _, groupOfPoints := range allStates {
		if !conflictingPositions(groupOfPoints, destination) {
			withoutDuplicatedPositions = append(withoutDuplicatedPositions, groupOfPoints)
		}
	}
	return withoutDuplicatedPositions
}

func conflictingPositions(multiplePoints []AStarNode, destination[] Coordinates) bool {
	if (len(multiplePoints) < 2) {
		return false
	}

	for i, point := range multiplePoints {
		if reachedHisDestination(point, destination[i]) {
			continue
		}

		if (i + 1 == len(multiplePoints)) {
			break
		}

		for _, otherPoint := range multiplePoints[i + 1:] {
			if areInTheSamePosition(point, otherPoint) {
				return true
			}
		}
	}

	return false
}

func reachedHisDestination(point AStarNode, destination Coordinates) bool {
	return equal(point.state.coordinates, destination)
}

func areInTheSamePosition(first, second AStarNode) bool {
	return equal(first.state.coordinates, second.state.coordinates)
}

func equal(first, second Coordinates) bool {
	return first.x == second.x && first.y == second.y;
}

func possibleNextMovesOfAllPoints(multiPoint MultiPointAstarNode) [][]AStarNode {
	movesOfAllPoints := make([][]AStarNode, 0)
	for i, singlePoint := range multiPoint.Points() {
		movesOfAllPoints = append(
			movesOfAllPoints,
			toAstarNodes(possibleMovesOf(singlePoint, multiPoint.Destination()[i])),
		)
	}
	return permute(movesOfAllPoints)
}

func possibleMovesOf(singlePoint AStarNode, destination Coordinates) []Node {
	if reachedHisDestination(singlePoint, destination) {
		return make([]Node, 0)
	}

	return singlePoint.AdjacentNodes()
}

func toAstarNodes(nodes []Node) []AStarNode {
	astarNodes := make([]AStarNode, 0)
	for _, node := range nodes {
		astarNodes = append(astarNodes, node.(AStarNode))
	}
	return astarNodes
}

func permute(listOfLists [][]AStarNode) [][]AStarNode {
	if (len(listOfLists) == 1) {
		result := make([][]AStarNode, 0)
		for _, element := range listOfLists[0] {
			result = append(result, []AStarNode{element})
		}
		return result
	}

	if (len(listOfLists) == 2) {
		result := make([][]AStarNode, 0)
		for _, firstElement := range listOfLists[0] {
			for _, secondElement := range listOfLists[1] {
				result = append(result, []AStarNode{firstElement, secondElement})
			}
		}
		return result
	}

	first := listOfLists[0]
	second := permute(listOfLists[1:])

	result := make([][]AStarNode, 0)
	for _, singleElement := range first {
		for _, array := range second {
			result = append(result, append([]AStarNode{singleElement}, array...))
		}
	}
	return result
}
