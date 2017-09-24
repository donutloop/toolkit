// Copyright 2017 The toolkit Authors. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package retry_test

import (
	"context"
	"errors"
	"github.com/donutloop/toolkit/retry"
	"testing"
)

func TestRetrierRetryContextDeadlineFail(t *testing.T) {
	r := retry.NewRetrier()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := r.Retry(ctx, func() (bool, error) {
		return true, nil
	})

	if err == nil {
		t.Fatal("unexpected nil error")
	}

	expectedErrorMessage := "context canceled"
	if err.Error() != expectedErrorMessage {
		t.Fatal(err)
	}
}

func TestRetrierRetry(t *testing.T) {
	r := retry.NewRetrier()
	err := r.Retry(context.Background(), func() (bool, error) {
		return true, nil
	})

	if err != nil {
		t.Fatalf("unexpected error (%v)", err)
	}
}

func TestRetrierRetryTriggerError(t *testing.T) {
	r := retry.NewRetrier()
	err := r.Retry(context.Background(), func() (bool, error) {
		return false, errors.New("stub error")
	})

	if err == nil {
		t.Fatal("unexpected nil error")
	}

	expectedErrorMessage := "stub error"
	if err.Error() != expectedErrorMessage {
		t.Fatal(err)
	}
}

func TestRetrierRetryFail(t *testing.T) {
	r := retry.NewRetrier()
	r.InitialInterval = 0.125
	r.Tries = 2
	r.MaxInterval = 0.25

	err := r.Retry(context.Background(), func() (bool, error) {
		return false, nil
	})

	if err == nil {
		t.Fatal("unexpected nil error")
	}

	expectedErrorMessage := "function never succeeded in Retry"
	if err.Error() != expectedErrorMessage {
		t.Fatal(err)
	}
}
