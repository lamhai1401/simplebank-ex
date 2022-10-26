package caching

import (
	"github.com/bluele/gcache"
)

type localCache struct {
	gc gcache.Cache
}

// NewLocalLRUCache linter
func NewLocalLRUCache(capacity int) Caching {
	gc := gcache.New(capacity).
		LRU().
		Build()
	return &localCache{
		gc: gc,
	}
}

// Set linter
func (c *localCache) Set(key interface{}, value interface{}) error {
	return c.gc.SetWithExpire(key, value, DefaultExpireTime)
}

// Get linter
func (c *localCache) Get(key interface{}) (interface{}, error) {
	return c.gc.Get(key)
}

// Remove linter
func (c *localCache) Remove(key interface{}) bool {
	return c.gc.Remove(key)
}
