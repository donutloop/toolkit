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
