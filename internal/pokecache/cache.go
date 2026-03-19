package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu   sync.RWMutex
	data map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		data: make(map[string]cacheEntry),
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.data[key] = cacheEntry{createdAt: time.Now(), val: val}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	tick := time.NewTicker(interval)
	var keys []string // keys to delete
	for {
		<-tick.C
		keys = nil // reset

		c.mu.RLock()
		for key, val := range c.data {
			if time.Since(val.createdAt) > interval {
				keys = append(keys, key)
			}
		}
		c.mu.RUnlock()

		c.mu.Lock()
		for _, key := range keys {
			delete(c.data, key)
		}
		c.mu.Unlock()
	}
}
