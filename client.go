package vesper

import "net/http"

// Client is the interface all our http client decorators must use.
type Client interface {
	Do(*http.Request) (*http.Response, error)
}

// The ClientFunc type is an adapter to allow the use of ordinary functions as http clients.
// If f is a function with the appropriate signature, ClientFunc(f) is a Handler that calls f.
type ClientFunc func(*http.Request) (*http.Response, error)

// Do calls f(r).
func (f ClientFunc) Do(r *http.Request) (*http.Response, error) {
	return f(r)
}
