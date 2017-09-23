// Copyright 2017 The toolkit Authors. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package loop_test

import (
	"testing"
	"github.com/donutloop/toolkit/loop"
	"time"
	"fmt"
	"errors"
)

func TestLoop(t *testing.T) {
	var counter int
	l := loop.NewLooper(1*time.Millisecond, func() error {
		counter++
		return nil
	})

	<- time.After(10*time.Millisecond)
	l.Stop()

	expectedValue := 9
	if !(counter >= 9) {
		t.Fatalf(`unexpected counter value (actual: "%d", expected: "%d")`, counter, expectedValue)
	}
}

func TestLoopFail(t *testing.T) {
	l := loop.NewLooper(1*time.Millisecond, func() error {
		panic(fmt.Errorf("check isolation of goroutine"))
		return nil
	})


	err := <- l.Error()
	if err.Error() != "event is panicked (check isolation of goroutine)" {
		t.Fatal(err)
	}

	l = loop.NewLooper(1*time.Millisecond, func() error {
		return errors.New("stub error")
	})

	err = <- l.Error()
	if err.Error() != "stub error" {
		t.Fatal(err)
	}
}

func BenchmarkLoop(b *testing.B) {
	for n := 0; n < b.N; n++ {
		loop.NewLooper(1*time.Millisecond, func() error {return nil})
	}
}
