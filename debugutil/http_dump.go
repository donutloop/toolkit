package debugutil

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"strings"
)


// PrettyDumpResponse is like DumpResponse but dump is pretty formatted.
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