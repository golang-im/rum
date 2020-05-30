package rum

import (
	"net/http"
)

// Transport implements the http.RoundTripper interface.
// The same as http.Transport,if is nil,will
// use http.DefaultTransport.
type Transport struct {
	http.RoundTripper
	middlewares MiddlewareChain
}

// RequestHash hashes the request for identifying a unique request.
type RequestHash func(r *http.Request) string

// RoundTrip executes  HTTP transaction with roundtriper middlewares.
func (t *Transport) RoundTrip(r *http.Request) (*http.Response, error) {
	if t.RoundTripper == nil {
		t.RoundTripper = http.DefaultTransport
	}

	return t.RoundTripper.RoundTrip(r)
}

// Use adds middleware to the t.middlewares,and wraps t.RoundTripper.
func (t *Transport) Use(wares ...Middleware) {
	if t.middlewares != nil {
		t.middlewares = make([]Middleware, 0)
	}

	if t.RoundTripper == nil {
		t.RoundTripper = http.DefaultTransport
	}

	// use middlewares in order
	t.middlewares = append(t.middlewares, wares...)
	for i := 0; i < t.middlewares.Len(); i++ {
		t.RoundTripper = t.middlewares[i](t.RoundTripper.RoundTrip)
	}
	return
}
