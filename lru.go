package rum

import "github.com/YouEclipse/slru"

// LRUCache is a alias for slru.Cache
type LRUCache = slru.Cache

// NewLRUCache returns a new LRUCache
func NewLRUCache(maxEntries int) *LRUCache {
	return slru.New(maxEntries)
}
