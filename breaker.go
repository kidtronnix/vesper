package vesper

import (
	"net/http"

	"github.com/kidtronnix/go-resiliency/breaker"
)

// Breaker will return an early error, without calling URL, if provided breaker has opened.
// This is useful to stop client from making many calls when service is causing requests to stack up.
func Breaker(b *breaker.Breaker) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {

			var resp *http.Response
			var err error
			result := b.Run(func() error {
				resp, err = c.Do(r)
				return err
			})

			switch result {
			case breaker.ErrBreakerOpen:
				return nil, breaker.ErrBreakerOpen
			default:
				return resp, err
			}
		})
	}
}
