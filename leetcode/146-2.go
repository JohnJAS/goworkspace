package main

import (
	"container/list"
)

type Pair struct {
	key, value int
}
type LRUCache struct {
	queue *list.List
	m     map[int]*list.Element
	cap   int
}

func Constructor(cap int) LRUCache {
	return LRUCache{
		queue: new(list.List),
		m:     make(map[int]*list.Element),
		cap:   cap,
	}
}

func (lru *LRUCache) Put(key int, value int) {
	if v, ok := lru.m[key]; ok {
		v.Value = Pair{key, value}
		lru.queue.MoveToFront(v)
	} else {
		if lru.cap <= lru.queue.Len() {
			back := lru.queue.Back()
			delete(lru.m, back.Value.(Pair).key)
			lru.queue.Remove(back)
		}
		lru.queue.PushFront(Pair{key, value})
		lru.m[key] = lru.queue.Front()
	}
}

func (lru *LRUCache) Get(key int) int {
	if v, ok := lru.m[key]; ok {
		lru.queue.MoveToFront(v)
		return v.Value.(Pair).value
	} else {
		return -1
	}
}
