package httptransport

import (
	"bufio"
	"bytes"
	"net/http"
)

var _ http.RoundTripper = &MockTransport{}

// MockTransport transport for mock
type MockTransport struct {
	defaultTransport http.RoundTripper
	//TODO cache Method/header/....
	mockData map[string][]byte
}

func NewMockTransport() *MockTransport {
	return &MockTransport{
		defaultTransport: http.DefaultTransport,
		mockData:         make(map[string][]byte),
	}
}

func (m *MockTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	u, err := url.ParseRequestURI(request.URL.RequestURI())
	if err != nil {
		return nil, err
	}
	path := u.Path

	resp := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Header:        http.Header{},
		Request:       request,
		Close:         true,
		ContentLength: -1,
	}

	if data, ok := m.mockData[path]; ok {
		resp.Header.Set("Content-Type", "application/json")
		buf := bytes.NewReader(data)
		resp.Body = ioutil.NopCloser(buf)
		return resp, nil
	}

	return m.defaultTransport.RoundTrip(request)
}

func (m *MockTransport) LoadMockData(path string, data []byte) {
	m.mockData[path] = data
}