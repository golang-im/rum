package rum

import (
	"errors"

	"github.com/YouEclipse/slru"
)

// CacheAdapter defines the cache adapter for CacheMiddleware.
type CacheAdapter interface {
	Set(key, value interface{}) error
	Get(key interface{}) (interface{}, error)
}

// LRUCache is slru.Cache which is thread-safely.
// it warps slrr.Cache for implements CacheAdapter.
type LRUCache slru.Cache

// NewLRUCache returns a new LRUCache.
func NewLRUCache(maxEntries int) *LRUCache {
	return (*LRUCache)(slru.New(maxEntries))
}

// Error for not found.
var ErrNil = errors.New("cache: result is nil.")

// Get returns the value from the cache,if not found , returns a err.
func (cache *LRUCache) Get(key interface{}) (interface{}, error) {
	value, ok := (*slru.Cache)(cache).Get(key)
	if !ok {
		return nil, ErrNil
	}
	return value, nil
}

// Set sets the value to the cache.
func (cache *LRUCache) Set(key interface{}, value interface{}) error {
	(*slru.Cache)(cache).Add(key, value)
	return nil
}
