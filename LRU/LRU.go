// Package lrucache implements an LRU cache.
// SET or GET will bring element to first cache
package lrucache

import (
	"errors"
	"fmt"

	mapgor "github.com/trongnhann/learning_go/MapGoroutine"
)

// Node is element's in LRU Cache
type Node struct {
	Key   interface{}
	Value interface{}

	PrevNode *Node
	NextNode *Node
}

// LRUCache consist of Size, Map element
type LRUCache struct {
	Size uint32
	//	Map       map[interface{}]*Node
	Map       *mapgor.MapGoroutine
	FirstNode *Node
	LastNode  *Node
}

// NewLRUCache is buffer a new LRUCache
func NewLRUCache(size uint32) (*LRUCache, error) {
	if size < 1 {
		return nil, errors.New("Size must bigger than 0")
	}
	return &LRUCache{
		Size:      size,
		Map:       mapgor.NewMapGoroutine(),
		FirstNode: nil,
		LastNode:  nil,
	}, nil
}

// Set is function set new key, value for node
// This node will be the first node of Cache
// If it existed it will be remove before set
func (lru *LRUCache) Set(key interface{}, value interface{}) {

	lru.Remove(key)
	lru.AddFirst(key, value)
	// Neu vuot qua size thi xoa LastNode
	if lru.Map.Len() > lru.Size {
		lru.Remove(lru.LastNode.Key)
	}
}

// Get will return value of key
func (lru *LRUCache) Get(key interface{}) (interface{}, bool) {
	tmpnode, err := lru.Map.Load(key)

	if tmpnode == nil {
		return nil, false
	}

	node := tmpnode.(*Node)

	if err != true {
		return nil, false
	} else {
		if node == nil {
			return nil, true
		}

		lru.Remove(key)
		lru.AddFirst(key, node.Value)
		return node.Value, true
	}
}

// Remove will delete value,key
func (lru *LRUCache) Remove(key interface{}) {
	tmpnode, err := lru.Map.Load(key)

	if tmpnode == nil {
		return
	} else {
		node := tmpnode.(*Node)

		if err != false {
			if node == nil {
				fmt.Println(node)
				return
			}

			if lru.Map.Len() == 1 {
				lru.FirstNode = nil
				lru.LastNode = nil
				lru.Map.Delete(key)
				return
			}

			if node.NextNode != nil && node.PrevNode != nil {
				node.PrevNode.NextNode = node.NextNode
			}

			if node.PrevNode != nil && node.NextNode != nil {
				node.NextNode.PrevNode = node.PrevNode
			}

			if key == lru.FirstNode.Key {
				lru.FirstNode = lru.FirstNode.NextNode
			}

			if key == lru.LastNode.Key {
				lru.LastNode = lru.LastNode.PrevNode
			}

			lru.Map.Delete(key)
		}
	}
}

// AddFirst add new key,value to the first of LRUCache
func (lru *LRUCache) AddFirst(key interface{}, value interface{}) {
	node := &Node{
		Key:      key,
		Value:    value,
		NextNode: lru.FirstNode,
		PrevNode: nil,
	}
	if lru.FirstNode == nil {
		lru.FirstNode = node
		lru.LastNode = node
	} else {
		if lru.LastNode == nil {
			lru.FirstNode.PrevNode = node
			lru.LastNode = lru.FirstNode
			lru.FirstNode = node
		} else {
			lru.FirstNode.PrevNode = node
			lru.FirstNode = node
		}
	}
	lru.Map.Store(key, node)

	if lru.Map.Len() > lru.Size {
		lru.Remove(lru.LastNode.Key)
	}
}
