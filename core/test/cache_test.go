package core_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gopher/core"
)

func TestCache_SetAndGet(t *testing.T) {
	t.Run("stores and retrieves a value", func(t *testing.T) {
		cache := core.NewCache()
		cache.Set("test:key", "value123", 1*time.Minute)

		val, ok := cache.Get("test:key")
		assert.True(t, ok)
		assert.Equal(t, "value123", val)
	})

	t.Run("respects TTL", func(t *testing.T) {
		cache := core.NewCache()
		cache.Set("short", "soon gone", 10*time.Millisecond)
		time.Sleep(20 * time.Millisecond)
		_, ok := cache.Get("short")
		assert.False(t, ok)
	})
}

func TestCache_Tags(t *testing.T) {
	t.Run("can invalidate by tag", func(t *testing.T) {
		cache := core.NewCache()
		cache.Set("item:1", "val1", time.Minute)
		cache.SetTags("item:1", []string{"group:one"})

		cache.InvalidateTags("group:one")
		_, ok := cache.Get("item:1")
		assert.False(t, ok)
	})

	t.Run("multiple keys for same tag", func(t *testing.T) {
		cache := core.NewCache()
		cache.Set("item:1", "one", time.Minute)
		cache.Set("item:2", "two", time.Minute)
		cache.SetTags("item:1", []string{"multi"})
		cache.SetTags("item:2", []string{"multi"})

		cache.InvalidateTags("multi")
		_, ok1 := cache.Get("item:1")
		_, ok2 := cache.Get("item:2")
		assert.False(t, ok1)
		assert.False(t, ok2)
	})
}
