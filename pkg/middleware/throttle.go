package middleware

import "time"

// ThrottleMiddleware behaviors like throttle function  in browser.
// the RoundTrip function will only be called after the TTL.
type ThrottleMiddleware struct {
	//TODO: implement
	TTL time.Duration
}
