package middleware

import (
	"net/http"
	"testing"

	"github.com/YouEclipse/rum/pkg/rum"
)

func TestAuth(t *testing.T) {
	httpClient := http.Client{}
	transport := &rum.Transport{}
	m := NewAuthenticationMiddleware()
	transport.Use(m.BasicAuth)

	httpClient.Transport = transport
}
