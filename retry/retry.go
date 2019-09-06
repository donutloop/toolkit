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

func NewRetrier(InitialIntervalInSeconds, maxIntervalInSeconds float64, tries uint, strategy Strategy) Retrier {

	if strategy == nil {
		panic("strategy is missing")
	}

	if InitialIntervalInSeconds > maxIntervalInSeconds {
		panic(fmt.Sprintf("initial interval is greater than max (initial: %f, max: %f)", InitialIntervalInSeconds, maxIntervalInSeconds))
	}

	return &retrier{
		InitialIntervalInSeconds: InitialIntervalInSeconds,
		maxIntervalInSeconds: maxIntervalInSeconds,
		strategy:        strategy,
		tries:           tries,
	}
}

// Retry supervised do funcs which automatically handle failures when they occur by performing retries.
type retrier struct {
	InitialIntervalInSeconds, maxIntervalInSeconds float64
	strategy             Strategy
	tries           uint
}

func (r *retrier) Retry(ctx context.Context, do RetryableDo) error {
	if ctx == nil {
		ctx = context.Background()
	}


	var err error
	var done bool
	for i := uint(0); !done && i < r.tries; i++ {
		done, err = do()

		if ctx.Err() != nil {
			return ctx.Err()
		}

		if err != nil {
			return err
		}

		if !done {
			r.InitialIntervalInSeconds = r.strategy.Policy(r.InitialIntervalInSeconds, r.maxIntervalInSeconds)
		}
	}

	if !done {
		return new(ExhaustedError)
	}
	return nil
}

type Strategy interface {
	Policy(intervalInSeconds, maxIntervalInSeconds float64) float64
}

type Exp struct {}

func (e *Exp) Policy(intervalInSeconds, maxIntervalInSeconds float64) float64 {
	time.Sleep(time.Duration(intervalInSeconds) * time.Second)
	return math.Min(intervalInSeconds*2, maxIntervalInSeconds)
}