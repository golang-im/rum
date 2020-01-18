package rum

import (
	"encoding/base64"
	"net/http"
)

type AuthenticationMiddleware struct {
}

type BasicAuthConfig struct {
	UserName string
	Password string
}

func (m *AuthenticationMiddleware) BasicAuth(next RoundTripperFunc, config *BasicAuthConfig) RoundTripperFunc {
	return func(r *http.Request) (*http.Response, error) {
		r.Header.Set("Authorization", "Basic "+base64.URLEncoding.EncodeToString([]byte(config.UserName+":"+config.Password)))
		return next.RoundTrip(r)
	}
}

type BearerTokenConfig struct {
	BearerToken string
}

func (m *AuthenticationMiddleware) BearerToken(next RoundTripperFunc, config *BearerTokenConfig) RoundTripperFunc {
	return func(r *http.Request) (*http.Response, error) {
		r.Header.Set("Authorization", "Bearer "+config.BearerToken)
		return next.RoundTrip(r)
	}
}

func (m *AuthenticationMiddleware) DigestAuth(next RoundTripperFunc) RoundTripperFunc {
	return func(r *http.Request) (*http.Response, error) {
		return next.RoundTrip(r)
	}
}

func (m *AuthenticationMiddleware) OAuth1(next RoundTripperFunc) RoundTripperFunc {
	return func(r *http.Request) (*http.Response, error) {
		return next.RoundTrip(r)
	}
}

func (m *AuthenticationMiddleware) OAuth2(next RoundTripperFunc) RoundTripperFunc {
	return func(r *http.Request) (*http.Response, error) {
		return next.RoundTrip(r)
	}
}
