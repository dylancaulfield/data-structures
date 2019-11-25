package data_structures

import (
	"go/types"
	"sync"
)

type Stack struct {
	items []types.Type
	lock  sync.RWMutex
}

func (s *Stack) Pop() (item types.Type) {

	s.lock.Lock()

	s.items, item = s.items[:len(s.items)-1], s.items[len(s.items)-1]

	s.lock.Unlock()

}

func (s *Stack) Push(item types.Type) {

	s.lock.Lock()

	s.items = append(s.items, item)

	s.lock.Unlock()

}
