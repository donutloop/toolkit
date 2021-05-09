package xhttp

import (
	"errors"
	"net/http"
)

type Middleware func(m http.RoundTripper) http.RoundTripper

// Use is wrapping up a RoundTripper with a set of middleware.
func Use(client *http.Client, middlewares ...Middleware) *http.Client {
	if client == nil {
		panic(errors.New("client is nil"))
	}
	if len(middlewares) == 0 {
		panic(errors.New("middlewares is nil"))
	}
	if client.Transport == nil {
		client.Transport = http.DefaultTransport
	}
	current := client.Transport
	for _, middleware := range middlewares {
		if middleware == nil {
			panic(errors.New("middleware is nil"))
		}
		current = middleware(current)
	}
	client.Transport = current
	return client
}
