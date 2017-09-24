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

const standardTriesCount uint = 10

type ExhaustedError struct{}

func (r *ExhaustedError) Error() string {
	return "function never succeeded in Retry"
}

type RetryableDo func() (bool, error)

func NewRetrier() *Retrier {
	return &Retrier{
		InitialInterval: 1,
		MaxInterval:     3,
		Tries:           standardTriesCount,
	}
}

// Retry supervised do funcs which automatically handle failures when they occur by performing retries.
type Retrier struct {
	InitialInterval float64
	MaxInterval     float64
	Tries           uint
}

func (r *Retrier) Retry(ctx context.Context, do RetryableDo) error {
	if ctx == nil {
		ctx = context.Background()
	}

	if r.InitialInterval <= 0 {
		r.InitialInterval = 1
	}

	if r.Tries == 0 {
		r.Tries = standardTriesCount
	}

	if r.InitialInterval > r.MaxInterval {
		return fmt.Errorf("initial interval is greater than max (initial: %f, max: %f)", r.InitialInterval, r.MaxInterval)
	}

	var err error
	var done bool
	interval := r.InitialInterval
	for i := uint(0); !done && i < r.Tries; i++ {
		done, err = do()

		if ctx.Err() != nil {
			return ctx.Err()
		}

		if err != nil {
			return err
		}

		if !done {
			time.Sleep(time.Duration(interval) * time.Second)
			interval = math.Min(interval*2, r.MaxInterval)
		}
	}

	if !done {
		return new(ExhaustedError)
	}
	return nil
}
