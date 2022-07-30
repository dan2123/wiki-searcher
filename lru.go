package main

import (
	"container/list"
	"errors"
)

type LRUCache struct {
	size      int
	evictList *list.List
	items     map[string]*list.Element
}

func NewLRUCache(size int) (*LRUCache, error) {
	if size <= 0 {
		return nil, errors.New("must provide error")
	}
	c := &LRUCache{
		size:      size,
		evictList: list.New(),
		items:     make(map[string]*list.Element),
	}
	return c, nil
}

func (c *LRUCache) Add(val string) (evicted bool) {
	// check if item exists
	if ent, ok := c.items[val]; ok {
		c.evictList.MoveToFront(ent)
		ent.Value = val
		return false
	}

	// add new item
	node := c.evictList.PushFront(val)
	c.items[val] = node

	// check if size has not been exceeded
	evict := c.evictList.Len() > c.size
	if evict {
		if oldNode := c.evictList.Back(); oldNode != nil {
			k := c.evictList.Remove(oldNode)
			delete(c.items, k.(string))
		}
	}
	return evict
}

func (c *LRUCache) GetItems() []string {
	out := make([]string, 0, len(c.items))

	// oldest to newest
	for curr := c.evictList.Back(); curr != nil; curr = curr.Prev() {
		out = append(out, curr.Value.(string))
	}

	return out
}
