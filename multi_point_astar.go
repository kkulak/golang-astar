package astar

type MultiPointAstarNode struct {
	points      []AStarNode
	// pozdrawiam pachola legutka
	destination []Coordinates
}

func (multiPoint MultiPointAstarNode) Cost(other Node) float64 {
	thisPoints := multiPoint.points
	thatPoints := other.(MultiPointAstarNode).points

	totalCost := 0.0
	for i := range thisPoints {
		totalCost += thisPoints[i].Cost(thatPoints[i])
	}
	return totalCost
}

func (multiPoint MultiPointAstarNode)EstimatedCost(other Node) float64 {
	thisPoints := multiPoint.points
	thatPoints := other.(MultiPointAstarNode).points

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
	withoutConflictingPositions := filterOutConflictingPositions(allPermutations, multiPoint.destination)
	return asMultiPointNodes(withoutConflictingPositions, multiPoint.destination)
}

func asMultiPointNodes(groupsOfPoints [][]AStarNode, destination []Coordinates) []Node {
	multiPointNodes := make([]Node, 0)
	for _, multiplePoints := range groupsOfPoints {
		multiPointNodes = append(multiPointNodes, MultiPointAstarNode{multiplePoints, destination})
	}
	return multiPointNodes
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
	for i, singlePoint := range multiPoint.points {
		movesOfAllPoints = append(
			movesOfAllPoints,
			toAstarNodes(possibleMovesOf(singlePoint, multiPoint.destination[i])),
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
		return listOfLists
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
