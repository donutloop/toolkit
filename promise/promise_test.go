package promise_test

import (
	"context"
	"github.com/donutloop/toolkit/promise"
	"strings"
	"testing"
	"fmt"
	"time"
)

func TestDoPanic(t *testing.T) {

	done, errc := promise.Do(context.Background(), func(ctx context.Context) error {
		panic("check isolation of process")
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
		<- time.After(1 * time.Second)
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


