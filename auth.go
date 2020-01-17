package rum


type AuthenticationMiddleware struct{
	token string
}



func(m *AuthenticationMiddleware) BasicAuth(next RoundTripperFunc) RoundTripperFunc {
	return func(r *http.Request) (*http.Response, error) {
		return next.RoundTripper.RoundTrip(r)
	}
}

func(m *AuthenticationMiddleware) BearerToken(next RoundTripperFunc) RoundTripperFunc {
	return func(r *http.Request) (*http.Response, error) {
		return next.RoundTripper.RoundTrip(r)
	}
}

func(m *AuthenticationMiddleware) DigestAuth(next RoundTripperFunc) RoundTripperFunc {
	return func(r *http.Request) (*http.Response, error) {
		return next.RoundTripper.RoundTrip(r)
	}
}

func(m *AuthenticationMiddleware) OAuth1(next RoundTripperFunc) RoundTripperFunc {
	return func(r *http.Request) (*http.Response, error) {
		return next.RoundTripper.RoundTrip(r)
	}
}

func(m *AuthenticationMiddleware) OAuth2(next RoundTripperFunc) RoundTripperFunc {
	return func(r *http.Request) (*http.Response, error) {
		return next.RoundTripper.RoundTrip(r)
	}
}

