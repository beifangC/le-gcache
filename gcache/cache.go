package gcache

import (
	"gcache/lru"
	"sync"
)

type cache struct {
	mu  sync.Mutex
	lru *lru.Cache
	cap int64
}

func (c *cache) add(key string, value Byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = lru.NewCache(0, c.cap)
	}
	c.lru.Add(key, value)
}

func (c *cache) get(key string) (value Byte, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}

	if v, ok := c.lru.Get(key); ok {
		return v.(Byte), ok
	}
	return
}
