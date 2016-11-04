package astar

import "container/heap"

type NodeToPriorityItemMap map[Node]*PriorityQueueAStarItem

func AStar(start, end Node) (distance float64, path []Node) {

	// set up open/closed sets
	nodeToPriorityQueueItem := NodeToPriorityItemMap{}
	priorityQueue := &PriorityQueue{}
	heap.Init(priorityQueue)

	// add start node to open set
	startNode := nodeToPriorityQueueItem.get(start)
	startNode.open = true
	startNode.estimatedTotalCost = startNode.node.EstimatedCost(end)
	heap.Push(priorityQueue, startNode)

	for priorityQueue.Len() > 0 {
		current := heap.Pop(priorityQueue).(*PriorityQueueAStarItem)
		current.open = false
		current.closed = true

		if current == nodeToPriorityQueueItem.get(end) {
			return costAndPathToGoal(current)
		}

		for _, adjacent := range current.node.AdjacentNodes() {
			adjacentNode := nodeToPriorityQueueItem.get(adjacent)
			cost := current.cost + current.node.Cost(adjacent)
			if (adjacentNode.closed) { continue }

			// TODO: remove duplicated code
			if (!adjacentNode.open) {
				adjacentNode.cost = cost
				adjacentNode.open = true
				adjacentNode.estimatedTotalCost = cost + adjacent.EstimatedCost(end)
				adjacentNode.parent = current
				// node never visited before, should be PUSHED to heap
				heap.Push(priorityQueue, adjacentNode)
			} else if (cost < adjacentNode.cost) {
				adjacentNode.cost = cost
				adjacentNode.open = true
				adjacentNode.estimatedTotalCost = cost + adjacent.EstimatedCost(end)
				adjacentNode.parent = current
				// node already visited, heap should be fixed, as node value has changed
				heap.Fix(priorityQueue, adjacentNode.index)
			}

		}
	}

	return -1, nil
}

func costAndPathToGoal(from *PriorityQueueAStarItem) (distance float64, path []Node) {
	p := []Node{}
	curr := from
	for curr != nil {
		p = append(p, curr.node)
		curr = curr.parent
	}
	return from.cost, p
}

func (nodeToPqItem NodeToPriorityItemMap) get(node Node) *PriorityQueueAStarItem {
	n, ok := nodeToPqItem[node]
	if !ok {
		n = &PriorityQueueAStarItem{ node: node }
		nodeToPqItem[node] = n
	}
	return n
}