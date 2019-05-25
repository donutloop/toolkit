package debugutil_test

import (
	"bufio"
	"bytes"
	"github.com/donutloop/toolkit/debugutil"
	"net/http"
	"testing"
)

func TestPrettyDumpResponse(t *testing.T) {

	r := []byte(`HTTP/2.0 200 OK
Content-Length: 1288
Cache-Control: no-cache, no-store, must-revalidate, max-age=0
Content-Type: application/json; charset=utf-8
Date: Wed, 08 May 2019 16:14:10 GMT
Expires: Wed, 08 May 2019 16:14:10 GMT
Server: nginx/1.10.3 (Ubuntu)
Set-Cookie: isLoggedIn=True; Path=/
Set-Cookie: sessionid_access=701229f3-491c-46c9-bcd8-cbe07cee05da; expires=Thu, 07-May-2020 16:14:10 GMT; httponly; Max-Age=31536000; Path=/
Vary: Cookie
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Iwoca: 844475
X-Partner: fincompare
X-State-Key: 8e4db383-2ead-4a47-87df-220432714c47
X-Xss-Protection: 1; mode=block

{"data": {"partner": {"verified_data": {"people": [{"identity_document_checks": [{"status": "passed", "file_links": [{"file_type": "photo", "link": "https://static.iwoca.com/assets/iwoca.4c17fef7de62.png"}], "check_id": "REFERENCE_0001", "datetime": "2017-06-12T14:05:51.666Z", "identity_document_type": "passport", "document_issuing_country": "gb", "provider_name": "test"}], "uid": "6cf7319e-f9ec-4038-ba4f-3561a6097484"}]}}, "state_key": "8e4db383-2ead-4a47-87df-220432714c47", "schema_version": "v1", "application": {"company": {"last_12_months_turnover": {"amount": 700000, "datetime": "2016-10-12T14:05:51.666Z"}, "type": "gmbh", "company_number": "01111112", "bank_details": {"iban": "DE89370400440532013000"}}, "requested_products": {"credit_facility": {"approval": {"amount": 15000}}}, "people": [{"residential_addresses": [{"town": "Ely", "uid": "cf9aa203-4e0c-4d7f-b42b-90c7b3d193d3", "house_number": "286", "date_from": "2014-02-03", "street_line_1": "Idverifier St", "postcode": "CB62AG"}], "last_name": "Norton", "uid": "6cf7319e-f9ec-4038-ba4f-3561a6097484", "roles": ["applicant", "shareholder", "guarantor", "director"], "title": "herr", "first_name": "Ervin", "privacy_policy": {"agreed": true, "datetime": "2016-10-12T14:05:51.666Z"}, "date_of_birth": "1980-01-01"}]}}}
`)
	reader := bufio.NewReader(bytes.NewReader(r))
	req, err := http.NewRequest(http.MethodGet, "/api/resource", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.ReadResponse(reader, req)
	if err != nil {
		t.Fatal(err)
	}

	dump, err := debugutil.PrettyDumpResponse(resp, true)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(dump))
}
