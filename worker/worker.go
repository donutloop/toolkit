package worker

import (
	"container/list"
)

type Request chan<- interface{}

type Response <-chan interface{}

type worker struct {
	request  chan interface{}
	response chan interface{}
	errs     chan error
	done     chan bool
	jobs     chan interface{}
	fn       func(n interface{}) (interface{}, error)
	buf      *list.List
}

// NewWorker starts n*Workers goroutines running func on incoming
// parameters sent on the returned channel.
func New(nWorkers uint, fn func(gt interface{}) (interface{}, error), buffer uint) (Request, Response, <-chan error) {

	request := make(chan interface{}, buffer)
	response := make(chan interface{}, buffer)
	errs := make(chan error, buffer)

	w := &worker{
		errs:     errs,
		request:  request,
		response: response,
		jobs:     make(chan interface{}, buffer),
		done:     make(chan bool),
		fn:       fn,
		buf:      list.New(),
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

	return request, response, errs
}

func (w *worker) listener() {
	inc := w.request

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

			v, err := w.fn(genericType)
			if err != nil {
				w.errs <- err
				continue
			}

			w.response <- v
		}
	}
}
