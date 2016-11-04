package astar

// Priority Queue API

type PriorityQueueItem struct {
	value     interface{}
	aPriority float64
	index     int
}

func (item PriorityQueueItem) Priority() float64 {
	return item.aPriority
}