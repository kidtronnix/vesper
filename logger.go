package vesper

import (
	"log"
	"net/http"
	"time"
)

// Logger will print a useful log message when a http request finishes.
func Logger(l *log.Logger) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			start := time.Now()
			resp, err := c.Do(r)
			var msg string
			if err != nil {
				msg = "Error: " + err.Error()
			} else {
				msg = resp.Status
			}
			l.Printf("POST %s%s - %s (%v)", r.URL.Host, r.URL.Path, msg, time.Since(start))
			return resp, err
		})
	}
}
