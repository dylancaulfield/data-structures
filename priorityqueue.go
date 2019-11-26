package structures

import (
	"sync"
)

type PriorityQueue struct {
	items []priorityQueueItem
	mutex sync.RWMutex
}

type priorityQueueItem struct {
	item     interface{}
	priority int
}

func (q *PriorityQueue) Enqueue(item interface{}, priority int) {

	q.mutex.Lock()
	defer q.mutex.Unlock()

	newItem := priorityQueueItem{item, priority}

	q.items = append(q.items, newItem)

}

func (q *PriorityQueue) GetHighest() interface{} {

	q.mutex.Lock()
	defer q.mutex.Unlock()

	if len(q.items) == 0 {
		return nil
	}

	highestPriority := q.items[0].priority
	highestPriorityIndex := 0

	for i, item := range q.items {

		if item.priority > highestPriority {
			highestPriorityIndex = i
		}

	}

	return q.items[highestPriorityIndex].item

}

func (q *PriorityQueue) DeleteHighest() {

	q.mutex.Lock()
	defer q.mutex.Unlock()

	if len(q.items) == 0 {
		return
	}

	highestPriority := q.items[0].priority
	highestPriorityIndex := 0

	for i, item := range q.items {

		if item.priority > highestPriority {
			highestPriorityIndex = i
		}

	}

	q.items = append(q.items[:highestPriorityIndex], q.items[highestPriorityIndex + 1:]...)

}