// Copyright 2017 The toolkit Authors. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package schedule

import (
	"context"
	"errors"
	"log"
	"runtime/debug"
	"sync"
)

type Job func(context.Context)

type DebugStack []byte

type Fifo struct {
	mu sync.Mutex

	resume       chan struct{}
	scheduled    int
	finished     int
	PanicHandler func(DebugStack)
	pendings     []Job

	ctx    context.Context
	cancel context.CancelFunc

	finishCond *sync.Cond
	done       chan struct{}
}

// NewFIFOScheduler returns a Scheduler that schedules jobs in FIFO order sequentially.
func NewFIFOScheduler() *Fifo {
	f := &Fifo{
		resume:       make(chan struct{}, 1),
		done:         make(chan struct{}, 1),
		PanicHandler: defaultPanicHandler,
	}
	f.finishCond = sync.NewCond(&f.mu)
	f.ctx, f.cancel = context.WithCancel(context.Background())
	go f.run()
	return f
}

func defaultPanicHandler(stack DebugStack) {
	log.Println(string(stack))
}

// Schedule schedules a job that will be ran in FIFO order sequentially.
func (f *Fifo) Schedule(j Job) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.cancel == nil {
		return errors.New("schedule: schedule to stopped scheduler")
	}

	if len(f.pendings) == 0 {
		select {
		case f.resume <- struct{}{}:
		default:
		}
	}
	f.pendings = append(f.pendings, j)
	return nil
}

func (f *Fifo) Pending() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return len(f.pendings)
}

func (f *Fifo) Scheduled() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.scheduled
}

func (f *Fifo) Finished() int {
	f.finishCond.L.Lock()
	defer f.finishCond.L.Unlock()
	return f.finished
}

func (f *Fifo) WaitFinish(n int) {
	f.finishCond.L.Lock()
	for f.finished < n || len(f.pendings) != 0 {
		f.finishCond.Wait()
	}
	f.finishCond.L.Unlock()
}

// Stop stops the scheduler and cancels all pending jobs.
func (f *Fifo) Stop() {
	f.mu.Lock()
	f.cancel()
	f.cancel = nil
	f.mu.Unlock()
	<-f.done
}

func (f *Fifo) run() {
	defer func() {
		close(f.done)
		close(f.resume)
	}()

	async := newAsync(f.ctx, f.PanicHandler)

	for {
		var job Job
		f.mu.Lock()
		if len(f.pendings) != 0 {
			f.scheduled++
			job = f.pendings[0]
		}
		f.mu.Unlock()
		if job == nil {
			select {
			case <-f.resume:
			case <-f.ctx.Done():
				f.mu.Lock()
				pendings := f.pendings
				f.pendings = nil
				f.mu.Unlock()
				// clean up pending jobs
				for _, job := range pendings {
					async.Do(job)
				}
				async.Close()
				return
			}
		} else {
			async.Do(job)
			f.finishCond.L.Lock()
			f.finished++
			f.pendings = f.pendings[1:]
			f.finishCond.Broadcast()
			f.finishCond.L.Unlock()
		}
	}
}

func newAsync(ctx context.Context, panicHandler func(stack DebugStack)) *async {
	a := &async{
		Ctx:          ctx,
		PanicHandler: panicHandler,
		Jobs:         make(chan func(ctx context.Context)),
	}
	a.init()
	return a
}

type async struct {
	Ctx          context.Context
	Jobs         chan func(ctx context.Context)
	PanicHandler func(stack DebugStack)
}

func (a *async) init() {
	go func(jobs chan func(ctx context.Context), ctx context.Context, panicHandler func(stack DebugStack)) {
		for job := range jobs {
			do(job, ctx, panicHandler)
		}
	}(a.Jobs, a.Ctx, a.PanicHandler)
}

func (a *async) Do(f func(ctx context.Context)) {
	a.Jobs <- f
}

func (a *async) Close() {
	close(a.Jobs)
}

func do(job func(ctx context.Context), ctx context.Context, panicHandler func(stack DebugStack)) {

	defer func() {
		if err := recover(); err != nil {
			panicHandler(DebugStack(debug.Stack()))
		}
	}()

	job(ctx)
}
