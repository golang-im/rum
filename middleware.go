package rum

import "net/http"

// RoundTripperFunc defines the RoundTrip func used by http.RoundTripper.
type RoundTripperFunc func(*http.Request) (*http.Response, error)

// RoundTrip
func (f RoundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

// Middleware defines the RoundTripperFunc Middleware.
type Middleware func(RoundTripperFunc) RoundTripperFunc

// MiddlewareChain defines a Middleware array.
type MiddlewareChain []Middleware
