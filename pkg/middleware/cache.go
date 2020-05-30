package middleware

import (
	"bufio"
	"bytes"
	"errors"
	"net/http"
	"net/http/httputil"

	"github.com/YouEclipse/rum/pkg/rum"
)

// CacheAdapter defines the cache adapter for cache middleware.
type CacheAdapter interface {
	Set(key string, value []byte, expire int) error
	Get(key string) ([]byte, error)
}

// ErrNil for not found.
var ErrNil = errors.New("result is nil")

// CacheMiddleware provides the cache middleware.
type CacheMiddleware struct {
	adapter CacheAdapter
	hash    rum.RequestHash
	expire  int
}

// CacheOption defines the options for cache middleware.
type CacheOption func(c *CacheMiddleware)

// NewCacheMiddleware returns a new cache middleware with options
func NewCache(adapter CacheAdapter, options ...CacheOption) *CacheMiddleware {
	c := &CacheMiddleware{
		adapter: adapter,
	}
	for _, optFunc := range options {
		optFunc(c)
	}
	if c.hash == nil {
		c.hash = rum.DefaultHash
	}
	return c
}

func CacheOptionExpire(expire int) CacheOption {
	return func(c *CacheMiddleware) {
		c.expire = expire
	}
}

func CacheOptionHash(h rum.RequestHash) CacheOption {
	return func(c *CacheMiddleware) {
		c.hash = h
	}
}

func (c *CacheMiddleware) Cache(next rum.RoundTripperFunc) rum.RoundTripperFunc {
	return func(r *http.Request) (*http.Response, error) {
		key := c.hash(r)
		if val, err := c.adapter.Get(key); err == nil && err != ErrNil {
			return http.ReadResponse(bufio.NewReader(bytes.NewBuffer(val)), r)
		}

		resp, err := next.RoundTrip(r)
		if err != nil {
			return nil, err
		}

		//dump Response for cache
		respDup, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, err
		}

		c.adapter.Set(key, respDup, c.expire)

		return resp, nil
	}
}
