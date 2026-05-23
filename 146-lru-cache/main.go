// 146. LRU Cache
//
// https://leetcode.com/problems/lru-cache/description/
//
// Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.
// Implement the LRUCache class:
//   - `LRUCache(int capacity)` Initialize the LRU cache with positive size capacity.
//   - `int get(int key)` Return the value of the key if the key exists, otherwise return -1.
//   - `void put(int key, int value)` Update the value of the key if the key exists.
//     Otherwise, add the key-value pair to the cache.
//     If the number of keys exceeds the capacity from this operation, evict the least recently used key.
//
// The functions get and put must each run in O(1) average time complexity.
package main

type CacheItem[K comparable] struct {
	Key   K
	Value any
	Prev  *CacheItem[K]
	Next  *CacheItem[K]
}

type LRUCache[K comparable] struct {
	capacity int
	head     *CacheItem[K]
	tail     *CacheItem[K]
	cache    map[K]*CacheItem[K]
}

func NewLRUCache[K comparable](capacity int) *LRUCache[K] {
	return &LRUCache[K]{
		capacity: max(1, capacity),
		cache:    make(map[K]*CacheItem[K], capacity),
	}
}

func (c *LRUCache[K]) Get(key K) (any, bool) {
	item, ok := c.cache[key]
	if !ok {
		return nil, false
	}

	c.moveToFront(item)
	return item.Value, true
}

func (c *LRUCache[K]) Put(key K, value any) {
	if item, ok := c.cache[key]; ok {
		item.Value = value
		c.moveToFront(item)
		return
	}

	if len(c.cache) >= c.capacity {
		c.removeLRU()
	}

	item := &CacheItem[K]{Key: key, Value: value}
	c.cache[key] = item
	c.insertAtFront(item)
}

func (c *LRUCache[K]) moveToFront(item *CacheItem[K]) {
	if item == c.head {
		return
	}
	c.detach(item)
	c.insertAtFront(item)
}

func (c *LRUCache[K]) insertAtFront(item *CacheItem[K]) {
	item.Prev = nil
	item.Next = c.head

	if c.head != nil {
		c.head.Prev = item
	}
	c.head = item

	if c.tail == nil {
		c.tail = item
	}
}

func (c *LRUCache[K]) detach(item *CacheItem[K]) {
	if item.Prev != nil {
		item.Prev.Next = item.Next
	} else {
		c.head = item.Next
	}

	if item.Next != nil {
		item.Next.Prev = item.Prev
	} else {
		c.tail = item.Prev
	}
}

func (c *LRUCache[K]) removeLRU() {
	if c.tail == nil {
		return
	}
	delete(c.cache, c.tail.Key)
	c.detach(c.tail)
}
