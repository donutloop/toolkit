package retry_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/donutloop/toolkit/retry"
)

func TestRoundTripper_InternalServer(t *testing.T) {

	var counter int32
	testserver := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		atomic.AddInt32(&counter, 1)
		t.Log("hit endpoint")
		w.WriteHeader(http.StatusInternalServerError)
	}))

	retryRoundTripper := retry.NewRoundTripper(http.DefaultTransport, .50, .15, 3, nil, new(retry.Exp))
	httpClient := new(http.Client)
	httpClient.Transport = retryRoundTripper

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, testserver.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("response is bad, got=%v", resp.StatusCode)
	}

	if counter != 3 {
		t.Errorf("counter is bad, got=%v, want=%v", counter, 3)
	}
}

func TestRoundTripper_InternalServerBlacklisted(t *testing.T) {

	var counter int32
	testserver := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		atomic.AddInt32(&counter, 1)
		t.Log("hit endpoint")
		w.WriteHeader(http.StatusInternalServerError)
	}))

	retryRoundTripper := retry.NewRoundTripper(http.DefaultTransport, .50, .15, 3, []int{http.StatusInternalServerError}, new(retry.Exp))
	httpClient := new(http.Client)
	httpClient.Transport = retryRoundTripper

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, testserver.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("response is bad, got=%v", resp.StatusCode)
	}

	if counter != 1 {
		t.Errorf("counter is bad, got=%v, want=%v", counter, 1)
	}
}

func TestRoundTripper_StatusOk(t *testing.T) {

	var counter int32
	testserver := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		atomic.AddInt32(&counter, 1)
		t.Log("hit endpoint")
		w.WriteHeader(http.StatusOK)
	}))

	retryRoundTripper := retry.NewRoundTripper(http.DefaultTransport, .50, .15, 3, nil, new(retry.Exp))
	httpClient := new(http.Client)
	httpClient.Transport = retryRoundTripper

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, testserver.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("response is bad, got=%v", resp.StatusCode)
	}

	if counter != 1 {
		t.Errorf("counter is bad, got=%v, want=%v", counter, 1)
	}
}

func TestRoundTripper_JsonStatusOk(t *testing.T) {

	json := `{"hello":"world"}`

	var counter int32
	testserver := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		atomic.AddInt32(&counter, 1)
		t.Log("hit endpoint")

		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		t.Log(string(b))

		count := atomic.LoadInt32(&counter)
		if count == 1 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if string(b) != json {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(json))
		if err != nil {
			t.Fatal(err)
		}
	}))

	retryRoundTripper := retry.NewRoundTripper(http.DefaultTransport, .50, .15, 3, nil, new(retry.Exp))
	httpClient := new(http.Client)
	httpClient.Transport = retryRoundTripper

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, testserver.URL, bytes.NewBuffer([]byte(json)))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("response is bad, got=%v", resp.StatusCode)
	}

	if counter != 2 {
		t.Errorf("counter is bad, got=%v, want=%v", counter, 1)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("response is bad, got=%v", err)
	}

	if string(b) != json {
		t.Fatalf("response body is bad, got=%v", string(b))
	}
}
