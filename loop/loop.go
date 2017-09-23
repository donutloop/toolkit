// Copyright 2017 The toolkit Authors. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package loop

import (
	"time"
	"fmt"
)

func NewLooper(rate time.Duration, event func() error) *looper {
	l := &looper{
		event: event,
		rate:     rate,
		shutdown: make(chan struct{}),
		err: make(chan error),
	}
	go l.doLoop()
	return l
}

type looper struct {
	event func() error
	rate time.Duration
	shutdown chan struct{}
	err chan error
}

func (l *looper) Stop() {
	close(l.shutdown)
	close(l.err)
}

func (l *looper) Error() <- chan error {
	return l.err
}

func (l *looper) doLoop() {
	ticker := time.NewTicker(l.rate)
	defer ticker.Stop()
	defer func() {
		if v := recover(); v != nil {
			l.err <- fmt.Errorf("event is panicked (%v)", v)
		}
	}()

	for {
		select {
		case <-ticker.C:
			if err := l.event(); err != nil {
				l.err <- err
				return
			}
		case <-l.shutdown:
			return
		}
	}
}

