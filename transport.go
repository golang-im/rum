package rum

import "net/http"

type Transport struct {
	http.RoundTripper
	middlewares MiddlewareChain

	index int8
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
