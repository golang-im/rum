package rum

import "net/http"

import "net/http/httputil"

type CacheMiddleware struct {
	cache   *LRUCache
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
		if val, ok := c.cache.Get(key); ok {
			return val.(*http.Response), nil
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

		c.cache.Add(key, respDup)

		return resp, nil
	}
}
