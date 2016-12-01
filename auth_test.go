package httpclient

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	assert := assert.New(t)

	testclient := ClientFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal("Bearer 123", r.Header.Get("Authorization"))
		return nil, nil
	})

	m := Auth("123")

	c := m(testclient)

	req, _ := http.NewRequest("GET", "http://example.com", nil)
	c.Do(req)
}
