// Package mapgoroutine as Map in Go but it allow handle cocurrent
package mapgoroutine

import (
	"sync"
)

// MapGoroutine is as map in Go but it consist RWMutex
type MapGoroutine struct {
	sync.RWMutex
	List map[interface{}]interface{}
}

// NewMapGoroutine buffer map
func NewMapGoroutine() *MapGoroutine {
	return &MapGoroutine{
		List: make(map[interface{}]interface{}),
	}
}

// Load is load value of key
func (rm *MapGoroutine) Load(key interface{}) (value interface{}, ok bool) {
//	rm.RLock()
//	defer rm.RUnlock()
	value, ok = rm.List[key]
	return
}

// Delete is remove key, value
func (rm *MapGoroutine) Delete(key interface{}) {
//	rm.Lock()
//	defer rm.Unlock()
	delete(rm.List, key)
}

// Store is add key, value
func (rm *MapGoroutine) Store(key interface{}, value interface{}) {
//	rm.Lock()
//	defer rm.Unlock()
	rm.List[key] = value
}

// Len is get length of map
func (rm *MapGoroutine) Len() uint32 {
//	rm.Lock()
//	defer rm.Unlock()
	return uint32(len(rm.List))
}
