package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val      []byte
}

type Cache struct {
	entries map[string]CacheEntry
	mu sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]CacheEntry),
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = CacheEntry{
		createdAt: time.Now(),
		val:      val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.entries[key]
	if !exists {
		return []byte{}, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration){
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > interval {
				c.mu.Lock()
				delete(c.entries, key)
				c.mu.Unlock()
			}
		}
	}	
}