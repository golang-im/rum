package rum

import (
	"errors"
)

// CacheAdapter defines the cache adapter for CacheMiddleware.
type CacheAdapter interface {
	Set(key, value interface{}, expire int) error
	Get(key interface{}) (interface{}, error)
}

// ErrNil for not found.
var ErrNil = errors.New("result is nil")
