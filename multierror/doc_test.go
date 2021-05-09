package multierror_test

import (
	"fmt"

	"github.com/donutloop/toolkit/multierror"
)

func Example() {
	errs := []error{
		connectionError,
		marshalError,
	}
	fmt.Println(multierror.New(errs...))
	// Output: multiple errors: error connect to db failed; error marshal json
}
