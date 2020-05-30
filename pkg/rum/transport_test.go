package rum

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddlewareUsed(t *testing.T) {
	httpClient := http.Client{}
	Transport := &Transport{}
	var before, after int

	for i := 1; i <= 10; i++ {
		Transport.Use(func(next RoundTripperFunc) RoundTripperFunc {
			return func(r *http.Request) (*http.Response, error) {
				before++
				resp, err := next.RoundTrip(r)
				after++
				return resp, err
			}
		})
	}
	httpClient.Transport = Transport

	httpClient.Get("http://xxx")

	assert.Equal(t, 10, before)
	assert.Equal(t, 10, after)
}

func TestMiddlewareOrder(t *testing.T) {
	httpClient := http.Client{}
	Transport := &Transport{}
	var expectOrderArray = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	var actualOrderArray = make([]int, 0)

	for i := 1; i <= 10; i++ {
		tmp := i
		Transport.Use(func(next RoundTripperFunc) RoundTripperFunc {
			return func(r *http.Request) (*http.Response, error) {
				actualOrderArray = append(actualOrderArray, tmp)
				resp, err := next.RoundTrip(r)
				return resp, err
			}
		})
	}
	httpClient.Transport = Transport

	httpClient.Get("http://xxx")

	assert.Equal(t, actualOrderArray, expectOrderArray)

}
