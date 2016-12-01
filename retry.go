package httpclient

import (
	"net/http"

	"github.com/smaxwellstewart/go-resiliency/retrier"
)

// Retry will attempt to make http call multiple times accoring to provided retrier.
func Retry(retry *retrier.Retrier) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			var resp *http.Response
			var err error
			e := retry.Run(func() error {
				resp, err = c.Do(r)
				return err
			})
			return resp, e
		})
	}
}
