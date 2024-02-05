package typeutl

import "sync"

type SyncMap[K comparable, V any] struct {
	m    map[K]V
	lock sync.RWMutex
}

func NewSyncMap[K comparable, V any]() SyncMap[K, V] {
	return SyncMap[K, V]{
		m: map[K]V{},
	}
}

func (it *SyncMap[K, V]) Put(k K, v V) {
	it.lock.Lock()
	defer it.lock.Unlock()

	it.m[k] = v
}

func (it *SyncMap[K, V]) Get(k K) (V, bool) {
	it.lock.RLock()
	defer it.lock.RUnlock()

	v, ok := it.m[k]

	return v, ok
}

func (it *SyncMap[K, V]) Delete(k K) {
	it.lock.Lock()
	defer it.lock.Unlock()

	delete(it.m, k)
}

func (it *SyncMap[K, V]) Values() []V {
	it.lock.RLock()
	defer it.lock.RUnlock()

	values := make([]V, 0, len(it.m))
	for _, v := range it.m {
		values = append(values, v)
	}

	return values
}
