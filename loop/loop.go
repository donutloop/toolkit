// Copyright 2017 The toolkit Authors. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package loop

import (
	"fmt"
	"runtime/debug"
	"time"
)

func NewLooper(rate time.Duration, event func() error) *Looper {
	l := &Looper{
		event:    event,
		rate:     rate,
		shutdown: make(chan struct{}),
		err:      make(chan error),
	}

	go l.doLoop()

	return l
}

type Looper struct {
	event    func() error
	rate     time.Duration
	shutdown chan struct{}
	err      chan error
}

func (l *Looper) Stop() {
	close(l.shutdown)
	close(l.err)
}

func (l *Looper) Error() <-chan error {
	return l.err
}

func (l *Looper) doLoop() {
	ticker := time.NewTicker(l.rate)
	defer ticker.Stop()
	defer func() {
		if v := recover(); v != nil {
			l.err <- &RecoverError{Err: v, Stack: debug.Stack()}
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

type RecoverError struct {
	Err   interface{}
	Stack []byte
}

func (e *RecoverError) Error() string { return fmt.Sprintf("Do panicked: %v", e.Err) }
