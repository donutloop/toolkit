package xhttp

import (
	"errors"
	"net/http"
)

type Middleware func(m http.RoundTripper) http.RoundTripper

var ErrClientNil error = errors.New("client is nil")
var ErrMiddlewaresNil error = errors.New("middlewares is nil")
var ErrMiddlewareNil error = errors.New("middleware is nil")

// Use is wrapping up a RoundTripper with a set of middleware.
func Use(client *http.Client, middlewares ...Middleware) *http.Client {
	if client == nil {
		panic(ErrClientNil)
	}

	if len(middlewares) == 0 {
		panic(ErrMiddlewaresNil)
	}

	if client.Transport == nil {
		client.Transport = http.DefaultTransport
	}

	current := client.Transport

	for _, middleware := range middlewares {
		if middleware == nil {
			panic(ErrMiddlewareNil)
		}

		current = middleware(current)
	}

	client.Transport = current

	return client
}
