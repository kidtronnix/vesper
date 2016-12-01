package vesper

import "net/http"

// RateLimit will make sure http calls are limited according to provided rate limiter.
// Client will block until limiter sends signal to let client make request.
func RateLimit(limiter <-chan struct{}) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			<-limiter
			return c.Do(r)
		})
	}
}
