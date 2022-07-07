package ds

import "sync"

type Map[K comparable, V any] struct {
	m     map[K]V
	mutex *sync.RWMutex
}

func NewMap[K comparable, V any]() Map[K, V] {
	return Map[K, V]{
		m:     make(map[K]V),
		mutex: &sync.RWMutex{},
	}
}

func (m *Map[K, V]) Get(key K) (V, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	v, ok := m.m[key]
	return v, ok
}

func (m *Map[K, V]) Set(key K, value V) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.m[key] = value
}
