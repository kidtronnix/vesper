package httpclient

import "net/http"

// Auth will make sure to add auth token to request header
func Auth(token string) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			r.Header.Set("Authorization", "Bearer "+token)
			r.Header.Set("Content-Type", "application/json")
			return c.Do(r)
		})
	}
}
