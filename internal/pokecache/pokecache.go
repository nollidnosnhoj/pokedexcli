package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	data      []byte
}

type Cache struct {
	cache map[string]CacheEntry
	mutex *sync.RWMutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]CacheEntry),
		mutex: &sync.RWMutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Set(key string, data []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = CacheEntry{
		createdAt: time.Now(),
		data:      data,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return entry.data, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, interval time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-interval)) {
			delete(c.cache, k)
		}
	}
}