package httpclient

import "net/http"

// Our client is defined as something that does a http request and gets a response and error
type Client interface {
	Do(*http.Request) (*http.Response, error)
}

// Singature of DoFunc
type ClientFunc func(*http.Request) (*http.Response, error)

// Add method to DoFunc type to satisfy Client interface
func (f ClientFunc) Do(r *http.Request) (*http.Response, error) {
	return f(r)
}
