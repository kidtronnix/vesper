package httpclient

import "net/http"

// RateLimit will make sure http calls are limited according to a provided rate limiter
func RateLimit(limiter <-chan struct{}) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			<-limiter
			return c.Do(r)
		})
	}
}
