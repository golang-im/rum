package rum

import "net/http"

// RoundTripperFunc defines the RoundTrip func used by http.RoundTripper.
type RoundTripperFunc func(*http.Request) (*http.Response, error)

// RoundTrip executes  HTTP transaction.
func (f RoundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

// Middleware defines the RoundTripperFunc Middleware.
type Middleware func(RoundTripperFunc) RoundTripperFunc

// MiddlewareChain defines a Middleware array.
type MiddlewareChain []Middleware

// Len returns the number of middlewares in the chain.
func (c MiddlewareChain) Len() int { return len(c) }
