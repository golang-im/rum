package rum

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"net/http"
)

// Transport
// if KeyFunc is nil,will use DefaultUniqueKeyFunc.
type Transport struct {
	http.RoundTripper
	middlewares MiddlewareChain
	KeyFunc     UniqueKeyFunc
}

// UniqueKeyFunc defines the unique key generator function of request.
type UniqueKeyFunc func(r *http.Request) string

// DefaultUniqueKeyFunc is the default unique key generator function
var DefaultUniqueKeyFunc = func(r *http.Request) string {
	h := sha1.New()
	io.WriteString(h, r.Method)
	io.WriteString(h, r.RequestURI)
	return hex.EncodeToString(h.Sum(nil))
}

func (t *Transport) RoundTrip(r *http.Request) (*http.Response, error) {
	if t.RoundTripper == nil {
		t.RoundTripper = http.DefaultTransport
	}

	return nil, nil
}

func (t *Transport) Use(wares ...Middleware) {
	if t.middlewares != nil {
		t.middlewares = make([]Middleware, 0)
	}
	t.middlewares = append(t.middlewares, wares...)
	for i := 0; i < t.middlewares.Len(); i++ {
		t.RoundTripper = t.middlewares[i](t.RoundTrip)
	}
	return
}
