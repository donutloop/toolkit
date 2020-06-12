package xhttp

import (
	"net/http"
)

type Middleware func(m http.RoundTripper) http.RoundTripper

// Use is wrapping up a RoundTripper with a set of middleware
func Use(client *http.Client, middlewares ...Middleware) *http.Client {
	if client.Transport == nil {
		client.Transport = http.DefaultTransport
	}
	current := client.Transport
	for _, middleware := range middlewares {
		current = middleware(current)
	}
	client.Transport = current
	return client
}
