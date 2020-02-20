package rum

import (
	"bufio"
	"bytes"
	"net/http"
	"net/http/httputil"
	"time"
)

// CacheMiddleware cached the http.Response.
type CacheMiddleware struct {
	cache   CacheAdapter
	keyFunc UniqueKeyFunc
	expire  int
}

func NewCacheMiddleware(maxRequest, expire int) *CacheMiddleware {
	return &CacheMiddleware{}
}

var DefaultCacheMiddleware = &CacheMiddleware{
	keyFunc: DefaultUniqueKeyFunc,
}

type CacheMiddlewareConfig struct {
	Expire time.Duration
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

		c.cache.Set(key, respDup, c.expire)

		return resp, nil
	}
}
