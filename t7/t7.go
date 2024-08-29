package t7

import "sync"

//Реализовать конкурентную запись данных в map.

type SafeMap struct {
	mu sync.RWMutex
	m  map[string]struct{}
}

func New(size int) *SafeMap {
	return &SafeMap{
		m: make(map[string]struct{}, size),
	}
}

func (s *SafeMap) Write(val string) {
	s.mu.Lock()
	s.m[val] = struct{}{}
	s.mu.Unlock()
}
