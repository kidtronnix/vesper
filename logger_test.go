package httpclient

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggerOK(t *testing.T) {
	assert := assert.New(t)

	testclient := ClientFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{Status: "200 OK"}, nil
	})

	buff := bytes.NewBufferString("")

	decorator := Logger(log.New(buff, "", 0))

	c := decorator(testclient)

	req, _ := http.NewRequest("GET", "http://example.com/boom", nil)
	fmt.Println("here")
	c.Do(req)

	assert.NotEqual(0, len(buff.String()))
}

func TestLoggerError(t *testing.T) {
	assert := assert.New(t)

	testclient := ClientFunc(func(r *http.Request) (*http.Response, error) {
		return nil, errors.New("Random")
	})

	buff := bytes.NewBufferString("")

	decorator := Logger(log.New(buff, "", 0))

	c := decorator(testclient)

	req, _ := http.NewRequest("GET", "http://example.com/boom", nil)
	fmt.Println("here")
	c.Do(req)

	assert.NotEqual(0, len(buff.String()))
}
