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

func TestResponseCodes(t *testing.T) {
	tests := []struct {
		name         string
		responseCode int
		blacklisted  []int
		counter      uint
	}{
		{
			name:         "StatusCode",
			responseCode: http.StatusOK,
			counter:      1,
		},
		{
			name:         "StatusCode",
			responseCode: http.StatusInternalServerError,
			counter:      3,
		},
		{
			name:         "blacklisted",
			responseCode: http.StatusInternalServerError,
			blacklisted:  []int{http.StatusInternalServerError},
			counter:      1,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			var counter int32

			testsServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				atomic.AddInt32(&counter, 1)
				t.Log("hit endpoint")
				w.WriteHeader(test.responseCode)
			}))

			retryRoundTripper := retry.NewRoundTripper(http.DefaultTransport, .50, .15, test.counter, test.blacklisted, new(retry.Exp))
			httpClient := new(http.Client)
			httpClient.Transport = retryRoundTripper

			req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, testsServer.URL, nil)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := httpClient.Do(req)
			if err != nil {
				t.Fatal(err)
			}

			defer resp.Body.Close()

			if resp.StatusCode != test.responseCode {
				t.Errorf("response is bad, got=%v", resp.StatusCode)
			}

			if counter != int32(test.counter) {
				t.Errorf("counter is bad, got=%v, want=%v", counter, int32(test.counter))
			}
		})
	}
}

func TestRT_JsonStatusOK(t *testing.T) {

	json := `{"hello":"world"}`

	var counter int32

	testsServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
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

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, testsServer.URL, bytes.NewBuffer([]byte(json)))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

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
