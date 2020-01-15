package rum

import "net/http"

type CacheMiddleware struct {
	cache *LRUCache
}

func NewCacheMiddleware(maxRequest int) *CacheMiddleware {
	return &CacheMiddleware{
		cache: NewLRUCache(maxRequest),
	}
}

var DefaultCacheMiddleware = &CacheMiddleware{
	cache: NewLRUCache(200),
}

func (c *CacheMiddleware) Default(next RoundTripperFunc) RoundTripperFunc {
	return func(r *http.Request) (*http.Response, error) {
		//TODO
		resp, err := next.RoundTrip(r)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
