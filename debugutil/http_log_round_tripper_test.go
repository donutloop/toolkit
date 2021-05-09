package debugutil_test

import (
	"fmt"
	"github.com/donutloop/toolkit/debugutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
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

	response, err := httpClient.Get(server.URL)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatal("response is bad, got: $v, want: $v", response.StatusCode, http.StatusOK)
	}
}
