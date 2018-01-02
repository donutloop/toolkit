// Copyright 2017 The toolkit Authors. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package multierror_test

import (
	"errors"

	"testing"

	"github.com/donutloop/toolkit/multierror"
)

func TestMultiError_Error(t *testing.T) {
	errs := []error{
		nil,
		errors.New("error connect to db failed"),
		errors.New("error marschaling json"),
	}
	expectedValue := "multiple errors: error connect to db failed; error marschaling json"
	err := multierror.New(errs...)
	if err.Error() != expectedValue {
		t.Errorf(`unexpected error message (actual:"%s", expected: "%s")`, err.Error(), expectedValue)
	}
}
