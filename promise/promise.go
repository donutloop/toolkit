// Copyright 2017 The toolkit Authors. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package promise

import (
	"context"
	"fmt"
	"runtime/debug"
)

type RecoverError struct {
	Err   interface{}
	Stack []byte
}

func (e *RecoverError) Error() string { return fmt.Sprintf("Do panicked: %v", e.Err) }

// Do is a basic promise implementation: it wraps calls a function in a goroutine.
func Do(ctx context.Context, f func(ctx context.Context) error) (<-chan struct{}, chan error) {
	done := make(chan struct{}, 1)
	errc := make(chan error)
	go func(done chan struct{}, error chan error, ctx context.Context) {
		defer func() {
			if v := recover(); v != nil {
				errc <- &RecoverError{Err: v, Stack: debug.Stack()}
			}
		}()

		if err := f(ctx); err != nil {
			error <- err
			return
		}

		done <- struct{}{}
	}(done, errc, ctx)

	return done, errc
}
