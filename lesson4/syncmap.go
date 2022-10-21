package lesson4

import "sync"

type SyncMap struct {
	data map[int]int
	mu   *sync.RWMutex
}

func NewSyncMap() *SyncMap {
	mu := sync.RWMutex{}
	data := map[int]int{}
	return &SyncMap{data, &mu}
}

func (sm *SyncMap) Set(key int, value int) {
	sm.mu.Lock()
	sm.data[key] = value
	sm.mu.Unlock()
}

func (sm *SyncMap) Get(key int) int {
	sm.mu.RLock()
	value := sm.data[key]
	sm.mu.RUnlock()
	return value
}

func (sm *SyncMap) Delete(key int) {
	sm.mu.Lock()
	delete(sm.data, key)
	sm.mu.Unlock()
}
