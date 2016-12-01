package httpclient

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRatelimit(t *testing.T) {
	assert := assert.New(t)

	var attempts int
	testclient := ClientFunc(func(r *http.Request) (*http.Response, error) {
		attempts++
		return nil, nil
	})

	limit := 5
	l := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		l <- struct{}{}
	}

	m := RateLimit(l)

	c := m(testclient)
	quit := time.After(200 * time.Millisecond)
	go func() {
		for {
			req, _ := http.NewRequest("GET", "http://example.com", nil)
			c.Do(req)
		}
	}()
	<-quit

	assert.Equal(limit, attempts)
}
