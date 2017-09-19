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

// NewFIFOScheduler returns a Scheduler that schedules jobs in FIFO
// order sequentially
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
					done := asyncDo(f.ctx, f.PanicHandler, job)
					<-done
				}
				return
			}
		} else {
			done := asyncDo(f.ctx, f.PanicHandler, job)
			<-done

			f.finishCond.L.Lock()
			f.finished++
			f.pendings = f.pendings[1:]
			f.finishCond.Broadcast()
			f.finishCond.L.Unlock()
		}
	}
}

// AsyncDo is a basic promise implementation: it wraps calls a function in a goroutine
func asyncDo(ctx context.Context, panicHandler func(stack DebugStack), f func(ctx context.Context)) <-chan struct{} {
	ch := make(chan struct{}, 1)
	go func(ch chan struct{}, ctx context.Context, panicHandler func(stack DebugStack)) {
		defer func() {
			if err := recover(); err != nil {
				panicHandler(DebugStack(debug.Stack()))
			}
		}()

		f(ctx)
		ch <- struct{}{}
	}(ch, ctx, panicHandler)
	return ch
}
