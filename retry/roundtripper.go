package retry

import (
	"context"
	"net/http"
)

type RoundTripper struct {
	retrier              Retrier
	next                 http.RoundTripper
	blacklistStatusCodes []int
}

// NewRoundTripper is constructing a new retry RoundTripper with given default values.
func NewRoundTripper(next http.RoundTripper, maxInterval, initialInterval float64, tries uint, blacklistStatusCodes []int, strategy Strategy) *RoundTripper {
	retrier := NewRetrier(initialInterval, maxInterval, tries, strategy)
	return &RoundTripper{
		retrier:              retrier,
		next:                 next,
		blacklistStatusCodes: blacklistStatusCodes,
	}
}

// RoundTrip is retrying a outgoing request in case of bad status code and not blacklisted status codes.
// if rt.next.RoundTrip(req) is return an error then it will abort the process retrying a request.
func (rt *RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	err := rt.retrier.Retry(context.Background(), func() (b bool, e error) {
		var err error
		resp, err = rt.next.RoundTrip(req)
		if err != nil {
			return false, err
		}

		// handle all 4xx and 5xx status codes
		if resp.StatusCode > http.StatusPermanentRedirect {
			if rt.isStatusCode(resp.StatusCode) {
				return true, nil
			}
			return false, nil
		}

		return true, nil
	})

	if _, ok := err.(*ExhaustedError); ok {
		return resp, nil
	}

	return resp, err
}

// isStatusCode iterates over list of black listed status code that it could abort the process of retrying a request
func (rt *RoundTripper) isStatusCode(statusCode int) bool {
	if rt.blacklistStatusCodes == nil {
		return false
	}
	for _, sc := range rt.blacklistStatusCodes {
		if statusCode == sc {
			return true
		}
	}
	return false
}
