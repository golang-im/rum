package rum

import (
	"encoding/base64"
	"net/http"
)

type AuthenticationMiddleware struct {
	used bool
}

// BasicAuthConfig is the config for basicAuth authorization.
type BasicAuthConfig struct {
	UserName string
	Password string
}

// BasicAuth adds the Authorization header for the basic authorization request.
func (m *AuthenticationMiddleware) BasicAuth(next RoundTripperFunc, config *BasicAuthConfig) RoundTripperFunc {
	return func(r *http.Request) (*http.Response, error) {
		r.Header.Set("Authorization", "Basic "+base64.URLEncoding.EncodeToString([]byte(config.UserName+":"+config.Password)))
		return next.RoundTrip(r)
	}
}

//BearerTokenConfig is the config for Bearer token authorization.
type BearerTokenConfig struct {
	BearerToken string
}

// BearerToken adds the Authorization header  for the bearer token authorization request.
func (m *AuthenticationMiddleware) BearerToken(next RoundTripperFunc, config *BearerTokenConfig) RoundTripperFunc {
	return func(r *http.Request) (*http.Response, error) {
		r.Header.Set("Authorization", "Bearer "+config.BearerToken)
		return next.RoundTrip(r)
	}
}
