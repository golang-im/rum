package middleware

import (
	"io"
	"net/http"

	"github.com/YouEclipse/rum/pkg/rum"
)

// LoggerMiddleware logging the request and response to the given io.Writer.
type LoggerMiddleware struct {
	requestFunc  func(req *http.Request) []byte
	responseFunc func(res *http.Response) []byte
	logger       io.Writer
}

// NewLoggerMiddleware returns a new LoggerMiddleware with options.
func NewLoggerMiddleware(options ...LoggerOption) *LoggerMiddleware {
	m := &LoggerMiddleware{}
	for _, optFunc := range options {
		optFunc(m)
	}
	return nil
}

// LoggerOption defines the option function for LoggerMiddleware.
type LoggerOption func(*LoggerMiddleware)

func LoggerOptionLogger(w io.Writer) LoggerOption {
	return func(m *LoggerMiddleware) {
		m.logger = w
	}
}

func LoggerOptionRequest(f func(req *http.Request) []byte) LoggerOption {
	return func(m *LoggerMiddleware) {
		m.requestFunc = f
	}
}

func LoggerOptionResponse(f func(res *http.Response) []byte) LoggerOption {
	return func(m *LoggerMiddleware) {
		m.responseFunc = f
	}
}

func (l *LoggerMiddleware) Log(next rum.RoundTripperFunc) rum.RoundTripperFunc {
	return func(r *http.Request) (*http.Response, error) {
		resp, err := next.RoundTrip(r)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
