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
	mockData         map[string][]byte
}

func NewMockTransport() *MockTransport {
	return &MockTransport{
		defaultTransport: http.DefaultTransport,
		mockData:         make(map[string][]byte),
	}
}

func (m *MockTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	path := request.URL.Path
	if data, ok := m.mockData[path]; ok {
		buf := bytes.NewBuffer(data)
		return http.ReadResponse(bufio.NewReader(buf), request)
	}

	return m.defaultTransport.RoundTrip(request)

	return nil, nil
}

func (m *MockTransport) LoadMockData(path string, data []byte) {
	m.mockData[path] = data
}
