package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	interval time.Duration
	mutex    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{}
	cache.entries = make(map[string]cacheEntry)
	cache.interval = interval
	go cache.reapLoop()
	return &cache
}

func (c *Cache) Add(key string, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	c.entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mutex.Lock()
		defer c.mutex.Unlock()
		for key, entry := range c.entries {
			age := time.Since(entry.createdAt)
			if age > c.interval {
				delete(c.entries, key)
			}
		}
	}
}
