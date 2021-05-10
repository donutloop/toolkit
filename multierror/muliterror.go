// Copyright 2017 The toolkit Authors. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package multierror

import "bytes"

const half int = 2

// New concatenate errors into one.
// If all errors are nil then it will returns nil
// otherwise the return value is a MultiError containing all the non-nil error.
func New(errs ...error) error {
	if len(errs) == 0 {
		return nil
	}

	errBucket := make([]error, 0, len(errs)/half)

	for _, err := range errs {
		if err != nil {
			errBucket = append(errBucket, err)
		}
	}

	if len(errBucket) == 0 {
		return nil
	}

	return multiError{errBucket}
}

// MultiError concatenate errors into one error.
type multiError struct {
	Errors []error
}

func (es multiError) Error() string {
	switch len(es.Errors) {
	case 1:
		return es.Errors[0].Error()
	default:
		var buf bytes.Buffer

		buf.WriteString("multiple errors: ")

		for i, e := range es.Errors {
			if i > 0 {
				buf.WriteString("; ")
			}

			buf.WriteString(e.Error())
		}

		return buf.String()
	}
}
