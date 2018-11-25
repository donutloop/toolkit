package worker

import (
	"container/list"
)

type worker struct {
	c chan interface{}
	done chan bool
	jobs chan interface{}
	fn    func(n GenericType)
	buf   *list.List
}

// NewWorker starts n*Workers goroutines running func on incoming
// parameters sent on the returned channel.
func New(nWorkers uint, fn func(gt GenericType), buffer uint) chan<- interface{} {
	retc := make(chan interface{}, buffer)
	w := &worker{
		c:     retc,
		jobs: make(chan interface{}, buffer),
		done: make(chan bool),
		fn:    fn,
		buf:   list.New(),
	}
	go w.listener()
	for i := uint(0); i < nWorkers; i++ {
		go w.work()
	}
	go func() {
		for i := uint(0); i < nWorkers; i++ {
			<-w.done
		}
	}()
	return retc
}

func (w *worker) listener() {
	inc := w.c
	for inc != nil || w.buf.Len() > 0 {
		outc := w.jobs
		var frontNode interface{}
		if e := w.buf.Front(); e != nil {
			frontNode = e.Value
		} else {
			outc = nil
		}
		select {
		case outc <- frontNode:
			w.buf.Remove(w.buf.Front())
		case el, ok := <-inc:
			if !ok {
				inc = nil
				continue
			}
			w.buf.PushBack(el)
		}
	}
	close(w.jobs)
}

func (w *worker) work() {
	for {
		select {
		case genericType, ok := <-w.jobs:
			if !ok {
				w.done <- true
				return
			}
			w.fn(genericType)
		}
	}
}
