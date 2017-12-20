// Copyright 2017 The toolkit Authors. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package debugutil_test

import (
	"testing"

	"github.com/donutloop/toolkit/debugutil"
)

func Test(t *testing.T) {

	strings := "dummy"

	tests := []struct {
		name   string
		input  interface{}
		output string
	}{
		{
			name:  "pretty print slice",
			input: make([]string, 0),
			output: `[]string{
}`,
		},
		{
			name:  "pretty print map",
			input: make(map[string]string),
			output: `map[string]string{
}`,
		},
		{
			name:   "pretty print pointer",
			input:  &strings,
			output: `&"dummy"`,
		},
		{
			name:   "pretty print value",
			input:  3,
			output: "3",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := debugutil.PrettySprint(test.input)
			if output != test.output {
				t.Errorf(`unepxected value (actual: "%s", exepected: "%s")`, output, test.output)
			}
		})
	}
}
