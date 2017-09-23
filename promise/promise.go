// Copyright 2017 The toolkit Authors. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package promise

import (
	"context"
	"fmt"
)

//Do is a basic promise implementation: it wraps calls a function in a goroutine
func Do(ctx context.Context, f func(ctx context.Context) error) (<-chan struct{}, chan error) {
	done := make(chan struct{}, 1)
	errc := make(chan error)
	go func(done chan struct{}, error chan error, ctx context.Context) {
		defer func() {
			if v := recover(); v != nil {
				errc <- fmt.Errorf("promise is panicked (%v)", v)
			}
		}()

		if err := f(ctx); err != nil {
			errc <- err
			return
		}

		done <- struct{}{}
	}(done, errc, ctx)


	return done, errc
}
