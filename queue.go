package structures

import "sync"

type Queue struct {
	items []interface{}
	mutex sync.RWMutex
}

func (q *Queue) Pop() (item interface{}) {

	q.mutex.Lock()
	defer q.mutex.Unlock()

	if len(q.items) == 0 {

		return nil

	}

	q.items, item = q.items[1:], q.items[0]

	return

}

func (q *Queue) Push(item interface{}) {

	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.items = append(q.items, item)

}
