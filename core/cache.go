package core

import (
"sync"
"time"

"github.com/dgraph-io/ristretto"

)

// Cache provides an in-memory caching layer with TTL and tag-based invalidation.
type Cache struct {
store       *ristretto.Cache
tagIndex    map[string]map[string]bool // tag -> keys
tagIndexMux sync.RWMutex
}

// NewCache creates a new Cache instance with default settings.
func NewCache() *Cache {
store, _ := ristretto.NewCache(&ristretto.Config{
NumCounters: 1e7,     // 10 million keys to track frequency
MaxCost:     1 << 30, // 1 GB max
BufferItems: 64,      // Number of keys per Get buffer
})

return &Cache{
	store:    store,
	tagIndex: make(map[string]map[string]bool),
}

}

// Set stores a value in the cache with a specific TTL.
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
c.store.SetWithTTL(key, value, 1, ttl)
}

// Get retrieves a value from the cache.
func (c *Cache) Get(key string) (interface{}, bool) {
return c.store.Get(key)
}

// SetTags associates tags with a cached key.
func (c *Cache) SetTags(key string, tags []string) {
c.tagIndexMux.Lock()
defer c.tagIndexMux.Unlock()

for _, tag := range tags {
	if _, exists := c.tagIndex[tag]; !exists {
		c.tagIndex[tag] = make(map[string]bool)
	}
	c.tagIndex[tag][key] = true
}

}

// InvalidateTags removes all cache entries associated with the given tags.
func (c *Cache) InvalidateTags(tags ...string) {
c.tagIndexMux.Lock()
defer c.tagIndexMux.Unlock()

for _, tag := range tags {
	if keys, found := c.tagIndex[tag]; found {
		for key := range keys {
			c.store.Del(key)
		}
		delete(c.tagIndex, tag)
	}
}

}

