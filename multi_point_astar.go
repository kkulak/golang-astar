package astar

type MultiPointAstarNode struct {
	points []AStarNode
}

func (multiPoint MultiPointAstarNode) Cost(other Node) float64 {
	thisPoints := multiPoint.points
	thatPoints := other.(MultiPointAstarNode).points

	maxCost := 0.0
	for i := range thisPoints {
		costForSinglePoint := thisPoints[i].Cost(thatPoints[i])
		if (costForSinglePoint > maxCost) {
			maxCost = costForSinglePoint
		}
	}
	return maxCost
}

func (multiPoint MultiPointAstarNode)EstimatedCost(other Node) float64 {
	//todo implement
	return 0
}

func (multiPoint MultiPointAstarNode) AdjacentNodes() []Node {
	allPermutations := possibleStatesOfAllPoints(multiPoint)
	withoutConflictingPositions := filterOutConflictingPositions(allPermutations)
	return asMultiPointNodes(withoutConflictingPositions)
}

func asMultiPointNodes(groupsOfPoints [][]AStarNode) []Node {
	multiPointNodes := make([]Node, 0)
	for _, multiplePoints := range groupsOfPoints {
		multiPointNodes = append(multiPointNodes, MultiPointAstarNode{multiplePoints})
	}
	return multiPointNodes
}

func filterOutConflictingPositions(allStates [][]AStarNode) [][]AStarNode {
	withoutDuplicatedPositions := make([][]AStarNode, 0)
	for _, groupOfPoints := range allStates {
		if !conflictingPositions(groupOfPoints) {
			withoutDuplicatedPositions = append(withoutDuplicatedPositions, groupOfPoints)
		}
	}
	return withoutDuplicatedPositions
}

func conflictingPositions(multiplePoints []AStarNode) bool {
	if (len(multiplePoints) < 2) {
		return false
	}

	pointToCheck := multiplePoints[0]
	otherPoints := multiplePoints[1:]

	for _, otherPoint := range otherPoints {
		if otherPoint.state.coordinates.x == pointToCheck.state.coordinates.x &&
			otherPoint.state.coordinates.y == pointToCheck.state.coordinates.y {
			return true
		}
	}

	return false
}

func possibleStatesOfAllPoints(multiPoint MultiPointAstarNode) [][]AStarNode {
	allStatesOfPoints := make([][]AStarNode, 0)
	for _, singlePoint := range multiPoint.points {
		allStatesOfPoints = append(allStatesOfPoints, toAstarNodes(singlePoint.AdjacentNodes()))
	}
	return permute(allStatesOfPoints)
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
