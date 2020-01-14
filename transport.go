package rum

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"net/http"
)

// Transport
// If Transport.KeyFunc is nil,will use DefaultUniqueKeyFunc.
type Transport struct {
	http.RoundTripper
	middlewares MiddlewareChain
	KeyFunc     UniqueKeyFunc

	index int8
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

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.RoundTripper == nil {
		t.RoundTripper = http.DefaultTransport
	}
	if t.middlewares != nil {

	}
	return nil, nil
}

func (t *Transport) reset() {
	t.RoundTripper = nil
	t.middlewares = nil
	t.index = -1
}

func (t *Transport) Next() {
	t.index++
	for t.index < int8(len(t.middlewares)) {
		t.middlewares[t.index](t.RoundTrip)
		t.index++
	}
}
