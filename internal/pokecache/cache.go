package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		data:     make(map[string]cacheEntry),
		interval: interval,
		mu:       &sync.RWMutex{},
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(time.Millisecond)
	for {
		<-ticker.C
		c.mu.Lock()
		for k, cache := range c.data {
			if time.Since(cache.createdAt) > c.interval {
				delete(c.data, k)
			}
		}
		c.mu.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	cache, ok := c.data[key]
	defer c.mu.Unlock()
	if ok {
		return cache.val, true
	}
	return nil, false
}
