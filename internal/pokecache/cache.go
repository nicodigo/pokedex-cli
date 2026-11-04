package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]cacheEntry
	mu       *sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		entries:  map[string]cacheEntry{},
		mu:       &sync.Mutex{},
		interval: interval,
	}

	go cache.reapLoop()

	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, found := c.entries[key]
	return entry.val, found
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.entries {
			if ellapsed := time.Since(v.createdAt); ellapsed > c.interval {
				delete(c.entries, k)
			}
		}
		c.mu.Unlock()
	}
}
