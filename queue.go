package structures

import "sync"

type Queue struct {
	items []interface{}
	mutex sync.RWMutex
}

func (q *Queue) DeQueue() (item interface{}) {

	q.mutex.Lock()
	defer q.mutex.Unlock()

	if len(q.items) == 0 {

		return nil

	}

	q.items, item = q.items[1:], q.items[0]

	return

}

func (q *Queue) EnQueue(item interface{}) {

	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.items = append(q.items, item)

}
