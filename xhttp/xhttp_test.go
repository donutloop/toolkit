package xhttp_test

import (
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/donutloop/toolkit/xhttp"
)

type TestMiddleware struct {
	roundtripper http.RoundTripper
	Log          func(v ...interface{})
	ID           int
}

func (m *TestMiddleware) RoundTrip(req *http.Request) (*http.Response, error) {
	m.Log("hit middleware ", m.ID)

	resp, err := m.roundtripper.RoundTrip(req)
	if err != nil {
		return resp, nil
	}

	return resp, nil
}

func TestInjectMiddleware(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		log.Println("hit handler")
	}

	s := httptest.NewServer(http.HandlerFunc(handler))
	defer s.Close()

	m1 := func(m http.RoundTripper) http.RoundTripper {
		return &TestMiddleware{m, log.Println, 1}
	}

	m2 := func(m http.RoundTripper) http.RoundTripper {
		return &TestMiddleware{m, log.Println, 2}
	}

	m3 := func(m http.RoundTripper) http.RoundTripper {
		return &TestMiddleware{m, log.Println, 3}
	}

	httpClient := new(http.Client)
	httpClient = xhttp.Use(httpClient, m1, m2)
	httpClient = xhttp.Use(httpClient, m3)

	resp, err := httpClient.Get(s.URL)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(err)
	}
}

func TestPanicNilClient(t *testing.T) {
	defer func() {
		v := recover()
		err := v.(error)

		if !errors.Is(err, xhttp.ClientNilError) {
			t.Errorf("error message is bad (%v)", v)
		}
	}()

	xhttp.Use(nil, nil)
}

func TestPanicNilMiddleware(t *testing.T) {
	defer func() {
		v := recover()
		err := v.(error)

		if !errors.Is(err, xhttp.MiddlewareNilError)  {
			t.Errorf("error message is bad (%v)", v)
		}
	}()

	xhttp.Use(new(http.Client), nil)
}

func TestPanicNilMiddlewares(t *testing.T) {
	defer func() {

		v := recover()
		err := v.(error)

		if !errors.Is(err, xhttp.MiddlewaresNilError) {
			t.Errorf("error message is bad (%v)", v)
		}
	}()

	xhttp.Use(new(http.Client))
}
