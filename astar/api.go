package astar

// A* API
type Node interface {
	AdjacentNodes() 	  []Node
	Cost(other Node) 	  float64
	EstimatedCost(other Node) float64
}

// Priority Queue API
type PriorityQueueAStarItem struct {
	node               Node
	open               bool
	closed             bool
	cost               float64
	estimatedTotalCost float64
	parent             *PriorityQueueAStarItem
	index              int
}

func (item PriorityQueueAStarItem) Priority() float64 {
	return item.estimatedTotalCost
}