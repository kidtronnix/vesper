package vesper

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/kidtronnix/go-resiliency/breaker"
	"github.com/stretchr/testify/assert"
)

func TestBreaker(t *testing.T) {
	assert := assert.New(t)

	errDown := errors.New("Service down!")

	testclient := ClientFunc(func(r *http.Request) (*http.Response, error) {
		return nil, errDown
	})

	m := Breaker(breaker.New(2, 1, time.Second))

	c := m(testclient)

	req, _ := http.NewRequest("GET", "http://example.com", nil)
	_, err := c.Do(req)
	assert.Equal(errDown, err)
	_, err = c.Do(req)
	assert.Equal(errDown, err)
	_, err = c.Do(req)
	assert.Equal(breaker.ErrBreakerOpen, err)

}
