// Copyright 2017 The toolkit Authors. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package retry

import (
	"context"
	"fmt"
	"math"
	"time"
)

type ExhaustedError struct{}

func (r *ExhaustedError) Error() string {
	return "function never succeeded in Retry"
}

type Retrier interface {
	Retry(ctx context.Context, do RetryableDo) error
}

type RetryableDo func() (bool, error)

func NewRetrier(InitialIntervalInSeconds, maxIntervalInSeconds float64, tries uint) Retrier {
	return &retrier{
		initialIntervalInSeconds: InitialIntervalInSeconds,
		maxIntervalInSeconds:     maxIntervalInSeconds,
		tries:           tries,
	}
}

// Retry supervised do funcs which automatically handle failures when they occur by performing retries.
type retrier struct {
	initialIntervalInSeconds float64
	maxIntervalInSeconds     float64
	tries           uint
}

func (r *retrier) Retry(ctx context.Context, do RetryableDo) error {
	if ctx == nil {
		ctx = context.Background()
	}

	if r.initialIntervalInSeconds > r.maxIntervalInSeconds {
		return fmt.Errorf("initial interval is greater than max (initial: %f, max: %f)", r.initialIntervalInSeconds, r.maxIntervalInSeconds)
	}

	var err error
	var done bool
	interval := r.initialIntervalInSeconds
	for i := uint(0); !done && i < r.tries; i++ {
		done, err = do()

		if ctx.Err() != nil {
			return ctx.Err()
		}

		if err != nil {
			return err
		}

		if !done {
			time.Sleep(time.Duration(interval) * time.Second)
			interval = math.Min(interval*2, r.maxIntervalInSeconds)
		}
	}

	if !done {
		return new(ExhaustedError)
	}
	return nil
}
