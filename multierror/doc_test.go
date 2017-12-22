package multierror_test

import (
	"fmt"

	"errors"

	"github.com/donutloop/toolkit/multierror"
)

func ExampleMultiError() {
	errs := []error{
		errors.New("error connect to db failed"),
		errors.New("error marschaling json"),
	}
	fmt.Println(multierror.New(errs...))
	// Output: multiple errors: error connect to db failed; error marschaling json
}
