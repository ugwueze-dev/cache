package store

import (
	"time"

	"github.com/maypok86/otter"
)

type MemoryStore struct {
	store otter.CacheWithVariableTTL[string, any]
}

func NewMemoryStore(numItems int) (Store, error) {
	store, err := otter.MustBuilder[string, any](numItems).
		Cost(func(key string, value any) uint32 {
			return 1
		}).WithVariableTTL().Build()

	if err != nil {
		return nil, err
	}

	return &MemoryStore{
		store: store,
	}, nil
}

func (m *MemoryStore) Put(key string, value any, duration int) {
	m.store.Set(key, value, time.Duration(duration)*time.Second)
}

func (m *MemoryStore) Has(key string) bool {
	return m.store.Has(key)
}

func (m *MemoryStore) Get(key string) any {
	val, _ := m.store.Get(key)
	return val
}

func (m *MemoryStore) Remove(key string) {
	m.store.Delete(key)
}

func (m *MemoryStore) Flush() {
	m.store.Clear()
}

func (m *MemoryStore) Close() {
	m.store.Close()
}
