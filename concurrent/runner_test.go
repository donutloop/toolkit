package concurrent_test

import (
	"errors"
	"github.com/donutloop/toolkit/concurrent"
	"sync/atomic"
	"testing"
)

func TestRun(t *testing.T) {
	counter := int32(0)
	errs := concurrent.Run(
		func() error {
			atomic.AddInt32(&counter, 1)
			return nil
		},
		func() error {
			atomic.AddInt32(&counter, 5)
			return nil
		},
	)
	if len(errs) != 0 {
		for err := range errs {
			t.Log(err)
		}
		t.Fatalf("unexpected errors")
	}

	expectedValue := int32(6)
	if counter != expectedValue {
		t.Errorf(`unexpected value (actual: %d, expected: %d)`, counter, expectedValue)
	}
}

func TestRunFail(t *testing.T) {

	counter := int32(0)
	errs := concurrent.Run(
		func() error {
			return errors.New("stub error")
		},
		func() error {
			panic("check isolation of goroutine")
		},
		func() error {
			atomic.AddInt32(&counter, 3)
			return nil
		},
	)

	expectedCountOfErrors := 2
	if len(errs) != expectedCountOfErrors {
		t.Fatalf(`unexpected count of errors (actual: %d, expected: %d)`, len(errs), expectedCountOfErrors)
	}

	expectedValue := int32(3)
	if counter != expectedValue {
		t.Errorf(`unexpected value (actual: %d, expected: %d)`, counter, expectedValue)
	}
}

func BenchmarkRun(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concurrent.Run(
			func() error {
				return errors.New("stub error")
			},
			func() error {
				panic("check isolation of goroutine")
			},
			func() error {
				return nil
			},
		)
	}
}
