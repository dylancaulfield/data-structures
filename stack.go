package structures

import (
	"sync"
)

type Stack struct {
	items []interface{}
	mutex  sync.RWMutex
}

func (s *Stack) Pop() (item interface{}) {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.items) == 0 {
		
		return nil
		
	}

	s.items, item = s.items[:len(s.items)-1], s.items[len(s.items)-1]

	return

}

func (s *Stack) Push(item interface{}) {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.items = append(s.items, item)

}
