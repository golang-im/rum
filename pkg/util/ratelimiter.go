package util

import (
	"sync"

	"golang.org/x/time/rate"
)

type RateLimiter interface {
	Allow(item interface{}) bool
}

var _ RateLimiter = (*BucketRateLimiter)(nil)

// BucketRateLimiter implements the RateLimiter interface using standard rate.Limiter.
type BucketRateLimiter struct {
	r     rate.Limit
	burst int

	limitersLock sync.Mutex
	limiters     map[interface{}]*rate.Limiter
}

// NewBucketRateLimiter creates a new BucketRateLimiter instance.
func NewBucketRateLimiter(r rate.Limit, burst int) *BucketRateLimiter {
	return &BucketRateLimiter{
		r:        r,
		burst:    burst,
		limiters: make(map[interface{}]*rate.Limiter),
	}
}

// Allow returns if drops the given item.
func (l *BucketRateLimiter) Allow(item interface{}) bool {
	l.limitersLock.Lock()
	defer l.limitersLock.Unlock()

	limiter, ok := l.limiters[item]
	if !ok {
		limiter = rate.NewLimiter(l.r, l.burst)
		l.limiters[item] = limiter
	}

	return limiter.Allow()
}
