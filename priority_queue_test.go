package astar

import (
	"container/heap"
	"testing"
)

func Test__Pop_Takes_Item_With_Lowest_Priority_Value(t *testing.T) {
	// given
	pq := &PriorityQueue{}
	heap.Init(pq)
	pq.Push(&PriorityQueueItem{ value: -1, aPriority: 5, index: 0 })
	pq.Push(&PriorityQueueItem{ value: -1, aPriority: 100, index: 1 })

	// when
	poppedItem := heap.Pop(pq).(*PriorityQueueItem)

	// then
	if poppedItem.Priority() != 5 {
		t.Error("Expected item with priority = 5")
	}
}

func Test__Fix_Repairs_Heap_After_Manual_Priority_Modification(t *testing.T) {
	// given
	pq := &PriorityQueue{}
	heap.Init(pq)
	item := &PriorityQueueItem{ value: -1, aPriority: 5, index: 0 }
	pq.Push(item)
	pq.Push(&PriorityQueueItem{ value: -1, aPriority: 100, index: 1 })

	// when
	item.aPriority = 200
	heap.Fix(pq, item.index)
	poppedItem := heap.Pop(pq).(*PriorityQueueItem)

	// then
	if poppedItem.Priority() != 100 {
		t.Error("Expected item with priority = 100")
	}
}

func Test__Remove_Followed_By_Push_Reestablish_Heap_Ordering(t *testing.T) {
	// given
	pq := &PriorityQueue{}
	heap.Init(pq)
	pq.Push(&PriorityQueueItem{ value: -1, aPriority: 5, index: 0 })
	pq.Push(&PriorityQueueItem{ value: -1, aPriority: 100, index: 1 })

	// when
	item := heap.Remove(pq, 0).(*PriorityQueueItem)
	item.aPriority = 200
	heap.Push(pq, item)
	poppedItem := heap.Pop(pq).(*PriorityQueueItem)

	// then
	if poppedItem.Priority() != 100 {
		t.Error("Expected item with priority = 100")
	}
}
