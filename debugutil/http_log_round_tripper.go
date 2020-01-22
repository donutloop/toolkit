package debugutil

import (
	"net/http"
	"net/http/httputil"
)

type logger interface {
	Errorf(format string, v ...interface{})
	Infof(format string, v ...interface{})
}

type LogRoundTripper struct {
	http.RoundTripper
	logger logger
	dumpBody bool
}

// RoundTripper returns a new http.RoundTripper which logs all requests (request and response dump)
// Should only be used for none production envs
func NewLogRoundTripper(roundTripper http.RoundTripper, logger logger, dumpBody bool) http.RoundTripper {
	return LogRoundTripper{roundTripper, logger, dumpBody}
}

func (tr LogRoundTripper) RoundTrip(req *http.Request) (res *http.Response, err error) {
	requestDump, err := httputil.DumpRequestOut(req, tr.dumpBody)
	if err != nil {
		tr.logger.Errorf("could not dump request: %v", err)
	} else {
		tr.logger.Infof("------------  HTTP REQUEST -----------\n%s", requestDump)
	}

	res, err = tr.RoundTripper.RoundTrip(req)
	if err != nil {
		return res, err
	}

	if res != nil {
		responseDump, err := httputil.DumpResponse(res, tr.dumpBody)
		if err != nil {
			tr.logger.Errorf("could not dump response: %v", err)
		} else {
			tr.logger.Infof( "------------  HTTP RESPONSE ----------\n%s", responseDump)
		}
	}

	return res, err
}
