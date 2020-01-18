package rum

import (
	"errors"

	"github.com/YouEclipse/slru"
)

type CacheAdapter interface {
	Set(key, value interface{}) error
	Get(key interface{}) (interface{}, error)
}

// LRUCache is slru.Cache which is thread-safely.
type LRUCache slru.Cache

// NewLRUCache returns a new LRUCache
func NewLRUCache(maxEntries int) *LRUCache {
	return (*LRUCache)(slru.New(maxEntries))
}

var ErrNil = errors.New("cache: result is nil.")

func (cache *LRUCache) Get(key interface{}) (interface{}, error) {
	value, ok := (*slru.Cache)(cache).Get(key)
	if !ok {
		return nil, ErrNil
	}
	return value, nil
}

func (cache *LRUCache) Set(key interface{}, value interface{}) error {
	(*slru.Cache)(cache).Add(key, value)
	return nil
}
