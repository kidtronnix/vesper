# vesper
[![GoDoc](https://godoc.org/github.com/smaxwellstewart/vesper?status.svg)](http://godoc.org/github.com/smaxwellstewart/vesper)

vesper is an idiomatic approach to building resilient http clients in go.

## Usage

```go
package main

import (
  "github.com/smaxwellstewart/go-resiliency/breaker"
  "github.com/smaxwellstewart/go-resiliency/limiter"
  "github.com/smaxwellstewart/go-resiliency/retrier"
)


func main() {
  // make rate limiter
	l := limiter.New(0, 5.0)
	defer l.Close()

	// make a resiliient http client
	client = vesper.Decorate(
		&http.Client{Timeout: time.Second}, // your http client
		vesper.Logger(log.New(os.Stdout, "[vesper]", 0)),
		vesper.Retry(retrier.New(retrier.ConstantBackoff(*retries, *backoff), nil)),
		vesper.Breaker(breaker.New(10, 1, 5*time.Second)),
		vesper.RateLimit(l.Limiter()),
	)

  // use client just like normal.
  req, _ := http.NewRequest("GET", "http://example.com", nil)
  resp, err := client.Do(req)
  // ...
}

```

## Decorators

Vesper provides a set of decorators for http clients, and framework for adding your own. It only decorates the `Do` method on a provided http client.

The following is the out the box decorators...

 - `Breaker` -  will error early without making http requests when an error threshold is met.
 - `Logger` - will print useful messages on each request.
 - `RateLimit` - will limit rate of making http reques.
 - `Retry` - will retry failed http requests.
