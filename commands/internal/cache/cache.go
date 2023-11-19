package cache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mutex    sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	data      []byte
	createdAt time.Time
	interval  time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Print() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	println("Cache:")
	for key, entry := range c.entries {
		println(key, entry.createdAt.String())
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval * time.Second)
	for {
		select {
		case <-ticker.C:
			c.mutex.Lock()
			for key, entry := range c.entries {
				if time.Since(entry.createdAt) > c.interval*time.Second {
					delete(c.entries, key)
				}
			}
			c.mutex.Unlock()
		}
	}
}

func (c *Cache) Add(key string, data []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.entries[key] = cacheEntry{
		data:      data,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.data, true
}
