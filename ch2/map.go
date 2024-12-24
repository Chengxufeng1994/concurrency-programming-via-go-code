package ch2

import "sync"

type Map[K comparable, V any] struct {
	mu sync.Mutex
	m  map[K]V
}

func NewMap[K comparable, V any](size ...int) *Map[K, V] {
	if len(size) > 0 {
		return &Map[K, V]{
			m: make(map[K]V, size[0]),
		}
	}

	return &Map[K, V]{
		m: make(map[K]V),
	}
}

func (m *Map[K, V]) Get(k K) (V, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	value, ok := m.m[k]
	return value, ok
}

func (m *Map[K, V]) Set(k K, v V) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.m[k] = v
}
