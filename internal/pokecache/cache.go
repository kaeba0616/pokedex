package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]CacheEntry
	mu      *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		entries: make(map[string]CacheEntry),
		mu:      &sync.Mutex{},
	}

	go c.ReapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = CacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[key]
	return entry.val, ok
}

func (c *Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.Reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) Reap(now time.Time, interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.entries {
		if v.createdAt.Before(now.Add(-interval)) {
			delete(c.entries, k)
		}
	}
}
