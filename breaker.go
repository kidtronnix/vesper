package httpclient

import (
	"net/http"

	"github.com/smaxwellstewart/go-resiliency/breaker"
)

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
