// Copyright 2017 The toolkit Authors. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package multierror_test

import (
	"errors"

	"testing"

	"github.com/donutloop/toolkit/multierror"
)

var marshalError error = errors.New("error marshal json")
var connectionError error = errors.New("error connect to db failed")

func TestMultiError_Error(t *testing.T) {
	errs := []error{
		nil,
		connectionError,
		marshalError,
	}
	expectedValue := "multiple errors: error connect to db failed; error marshal json"
	err := multierror.New(errs...)
	if err.Error() != expectedValue {
		t.Errorf(`unexpected error message (actual:"%v", expected: "%s")`, err, expectedValue)
	}
}
