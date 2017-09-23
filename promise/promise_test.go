// Copyright 2017 The toolkit Authors. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package promise_test

import (
	"context"
	"fmt"
	"github.com/donutloop/toolkit/promise"
	"strings"
	"testing"
	"time"
)

func TestDoPanic(t *testing.T) {

	done, errc := promise.Do(context.Background(), func(ctx context.Context) error {
		panic("check isolation of goroutine")
		return nil
	})

	select {
	case <-done:
	case err := <-errc:
		if err == nil {
			t.Fatal("Unexpected nil error")
		}

		if !strings.Contains(err.Error(), "promise is panicked") {
			t.Fatalf(`Unexpected error message (Actual: %s, Expected: promise is panicked (*))`, err.Error())
		}
	}
}

func TestDoFail(t *testing.T) {

	done, errc := promise.Do(context.Background(), func(ctx context.Context) error {
		return fmt.Errorf("stub")
	})

	select {
	case <-done:
	case err := <-errc:
		if err == nil {
			t.Fatal("Unexpected nil error")
		}

		expectedMessage := "stub"
		if err.Error() != expectedMessage {
			t.Fatalf(`Unexpected error message (Actual: %s, Expected: %s)`, err.Error(), expectedMessage)
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
			t.Fatalf("Unexpected error (%v)", err)
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
				b.Fatalf("Unexpected error (%v)", err)
			}
		}
	}
}
