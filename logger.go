package httpclient

import (
	"log"
	"net/http"
	"time"
)

// Auth will make sure to add auth token to request header
func Logger(l *log.Logger) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			start := time.Now()
			resp, err := c.Do(r)
			var msg string
			if err != nil {
				msg = "Error! " + err.Error()
			} else {
				msg = resp.Status
			}
			l.Printf("POST %s%s - %s (%v)", r.URL.Host, r.URL.Path, msg, time.Since(start))
			return resp, err
		})
	}
}
