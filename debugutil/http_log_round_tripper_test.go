package debugutil_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/donutloop/toolkit/debugutil"
)

type logger struct{}

func (l logger) Errorf(format string, v ...interface{}) {
	log.Println(fmt.Sprintf(format, v...))
}
func (l logger) Infof(format string, v ...interface{}) {
	log.Println(fmt.Sprintf(format, v...))
}

func TestLogRoundTripper_RoundTrip(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	testHandler := http.HandlerFunc(handler)

	server := httptest.NewServer(testHandler)
	defer server.Close()

	httpClient := new(http.Client)
	httpClient.Transport = debugutil.NewLogRoundTripper(http.DefaultTransport, logger{}, true)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	response, err := httpClient.Do(req)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatal("response is bad, got: $v, want: $v", response.StatusCode, http.StatusOK)
	}
}
