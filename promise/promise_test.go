// Copyright 2017 The toolkit Authors. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package promise_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/donutloop/toolkit/promise"
)

var ErrStub error = errors.New("stub")

func TestDoPanic(t *testing.T) {
	done, errc := promise.Do(context.Background(), func(ctx context.Context) error {
		panic("check isolation of goroutine")
	})

	select {
	case <-done:
	case err := <-errc:
		if err == nil {
			t.Fatal("unexpected nil error")
		}

		expectedError := "Do panicked: check isolation of goroutine"
		if expectedError != err.Error() {
			t.Fatalf(`unexpected error message (actual: "%s", expected: "%v")`, err.Error(), expectedError)
		}
	}
}

func TestDoFail(t *testing.T) {
	done, errc := promise.Do(context.Background(), func(ctx context.Context) error {
		return ErrStub
	})

	select {
	case <-done:
	case err := <-errc:
		if err == nil {
			t.Fatal("unexpected nil error")
		}

		expectedMessage := "stub"
		if err.Error() != expectedMessage {
			t.Fatalf(`unexpected error message (actual: "%s", expected: "%s")`, err.Error(), expectedMessage)
		}
	}
}

func TestDo(t *testing.T) {
	done, errc := promise.Do(context.Background(), func(ctx context.Context) error {
		<-time.After(1 * time.Second)
		return nil
	})

	select {
	case <-done:
	case err := <-errc:
		if err != nil {
			t.Fatalf("unexpected error (%v)", err)
		}
	}
}

func BenchmarkDo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		done, errc := promise.Do(context.Background(), func(ctx context.Context) error { return nil })

		select {
		case <-done:
		case err := <-errc:
			if err != nil {
				b.Fatalf("unexpected error (%v)", err)
			}
		}
	}
}
