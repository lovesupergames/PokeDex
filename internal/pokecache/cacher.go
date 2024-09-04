package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Cache    map[string]CacheEntry
	Mu       *sync.RWMutex
	Interval time.Duration
}

type CacheEntry struct {
	CreatedAt time.Time
	Value     []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		Cache:    make(map[string]CacheEntry),
		Mu:       &sync.RWMutex{},
		Interval: interval,
	}
	go func() {
		ticker := time.NewTicker(c.Interval)
		defer ticker.Stop()
		for range ticker.C {
			c.ReapLoop()
		}
	}()

	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Cache[key] = CacheEntry{
		CreatedAt: time.Now(),
		Value:     value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.RLock()
	defer c.Mu.RUnlock()
	if value, ok := c.Cache[key]; ok {
		return value.Value, true
	}
	return nil, false
}

func (c *Cache) ReapLoop() {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	for key, value := range c.Cache {
		if time.Since(value.CreatedAt) > c.Interval {
			delete(c.Cache, key)
		}
	}
}
