// Copyright 2017 The toolkit Authors. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package loop_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/donutloop/toolkit/loop"
)

var StubErr error = errors.New("stub error")

func TestLoop(t *testing.T) {
	var counter int
	l := loop.NewLooper(1*time.Millisecond, func() error {
		counter++
		return nil
	})

	<-time.After(10 * time.Millisecond)
	l.Stop()

	if !(counter > 0) {
		t.Fatalf(`unexpected counter value (actual: "%d", expected: "counter > 0")`, counter)
	}
}

var GoroutineError error = fmt.Errorf("check isolation of goroutine")

func TestLoopFail(t *testing.T) {
	l := loop.NewLooper(1*time.Millisecond, func() error {
		panic(GoroutineError)
	})

	err := <-l.Error()
	if err.Error() != "Do panicked: check isolation of goroutine" {
		t.Fatal(err)
	}

	l = loop.NewLooper(1*time.Millisecond, func() error {
		return StubErr
	})

	err = <-l.Error()
	if !errors.Is(err, StubErr) {
		t.Fatal(err)
	}
}

func BenchmarkLoop(b *testing.B) {
	for n := 0; n < b.N; n++ {
		loop.NewLooper(1*time.Millisecond, func() error { return nil })
	}
}
