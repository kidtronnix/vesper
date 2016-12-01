package httpclient

import (
	"errors"
	"net/http"
	"testing"

	"github.com/smaxwellstewart/go-resiliency/retrier"
	"github.com/stretchr/testify/assert"
)

func TestRetry(t *testing.T) {
	assert := assert.New(t)

	var errTransient = errors.New("Transient error")

	var retries = 2

	var attempts int
	testclient := ClientFunc(func(r *http.Request) (*http.Response, error) {
		var err error
		if attempts < retries {
			err = errTransient
		}
		attempts++
		return nil, err
	})

	m := Retry(retrier.New(retrier.ConstantBackoff(retries, 0), nil))

	c := m(testclient)

	req, _ := http.NewRequest("GET", "http://example.com", nil)
	c.Do(req)

	assert.Equal(retries+1, attempts)
}
