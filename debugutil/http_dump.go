package debugutil

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"strings"
)

// PrettyPrintResponse is pretty printing a http response.
func PrettySprintResponse(resp *http.Response) (string, error) {
	dump, err := PrettyDumpResponse(resp, true)
	if err != nil {
		return "", err
	}

	return string(dump), nil
}

// PrettyDumpResponse is like httputil.DumpResponse but dump is pretty formatted.
func PrettyDumpResponse(resp *http.Response, body bool) ([]byte, error) {

	b, err := httputil.DumpResponse(resp, body)
	if err != nil {
		return nil, err
	}

	header := resp.Header.Get("Content-type")
	if body && strings.Contains(header, "application/json") && resp.ContentLength > 0 {
		buffer := new(bytes.Buffer)
		jsonRaw := b[int64(len(b))-resp.ContentLength:]
		b = b[:int64(len(b))-resp.ContentLength]
		buffer.Write(b)

		if err := json.Indent(buffer, jsonRaw, "", "\t"); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return b, nil
}

// PrettySprintRequest is pretty printing a http request.
func PrettySprintRequest(resp *http.Request) (string, error) {
	dump, err := PrettyDumpRequest(resp, true)
	if err != nil {
		return "", err
	}

	return string(dump), nil
}

// PrettyDumpRequest is like httputil.DumpRequest but dump is pretty formatted.
func PrettyDumpRequest(req *http.Request, body bool) ([]byte, error) {

	b, err := httputil.DumpRequest(req, body)
	if err != nil {
		return nil, err
	}

	header := req.Header.Get("Content-type")
	if body && strings.Contains(header, "application/json") && req.ContentLength > 0 {
		buffer := new(bytes.Buffer)
		jsonRaw := b[int64(len(b))-req.ContentLength:]
		b = b[:int64(len(b))-req.ContentLength]
		buffer.Write(b)

		if err := json.Indent(buffer, jsonRaw, "", "\t"); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return b, nil
}
