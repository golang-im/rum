package rum

import "net/http"

import "net/http/httputil"

import "bufio"

import "bytes"

type CacheMiddleware struct {
	cache   CacheAdapter
	keyFunc UniqueKeyFunc
}

func NewCacheMiddleware(maxRequest int) *CacheMiddleware {
	return &CacheMiddleware{
		cache: NewLRUCache(maxRequest),
	}
}

var DefaultCacheMiddleware = &CacheMiddleware{
	cache:   NewLRUCache(200),
	keyFunc: DefaultUniqueKeyFunc,
}

func (c *CacheMiddleware) Default(next RoundTripperFunc) RoundTripperFunc {
	return func(r *http.Request) (*http.Response, error) {
		key := c.keyFunc(r)
		if val, err := c.cache.Get(key); err == nil && err != ErrNil {
			return http.ReadResponse(bufio.NewReader(bytes.NewBuffer(val.([]byte))), r)
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

		c.cache.Set(key, respDup)

		return resp, nil
	}
}
