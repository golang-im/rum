package auth

import (
	"encoding/base64"
	"net/http"

	"github.com/YouEclipse/rum/pkg/rum"
)

// AuthenticationMiddleware provides authentication middlewares.
type AuthenticationMiddleware struct {
	config *BasicAuthConfig
}

// NewAuthenticationMiddleware returns a new authentication middleware.
func NewAuthenticationMiddleware() *AuthenticationMiddleware {
	return &AuthenticationMiddleware{}
}

// BasicAuthConfig is the config for basicAuth authorization.
type BasicAuthConfig struct {
	UserName string
	Password string
}

// BasicAuth adds the Authorization header for the basic authorization request.
func (m *AuthenticationMiddleware) BasicAuth(next rum.RoundTripperFunc) rum.RoundTripperFunc {
	return func(r *http.Request) (*http.Response, error) {
		r.Header.Set("Authorization", "Basic "+base64.URLEncoding.EncodeToString([]byte(m.config.UserName+":"+m.config.Password)))
		return next.RoundTrip(r)
	}
}

//BearerTokenConfig is the config for Bearer token authorization.
type BearerTokenConfig struct {
	BearerToken string
}

// BearerToken adds the Authorization header  for the bearer token authorization request.
func (m *AuthenticationMiddleware) BearerToken(next rum.RoundTripperFunc, config *BearerTokenConfig) rum.RoundTripperFunc {
	return func(r *http.Request) (*http.Response, error) {
		r.Header.Set("Authorization", "Bearer "+config.BearerToken)
		return next.RoundTrip(r)
	}
}
