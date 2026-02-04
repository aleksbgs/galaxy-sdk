package store

import "sync"

// MemStore is an in-memory KV store for tests and prototypes.
type MemStore struct {
	mu      sync.RWMutex
	data    map[string][]byte
	version int64
}

func NewMemStore() *MemStore {
	return &MemStore{data: make(map[string][]byte)}
}

func (m *MemStore) Get(key []byte) []byte {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value := m.data[string(key)]
	if value == nil {
		return nil
	}
	cpy := make([]byte, len(value))
	copy(cpy, value)
	return cpy
}

func (m *MemStore) Set(key, value []byte) {
	m.mu.Lock()
	defer m.mu.Unlock()
	cpy := make([]byte, len(value))
	copy(cpy, value)
	m.data[string(key)] = cpy
}

func (m *MemStore) Delete(key []byte) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, string(key))
}

func (m *MemStore) Commit() (int64, []byte) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.version++
	return m.version, nil
}

func (m *MemStore) Version() int64 {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.version
}
