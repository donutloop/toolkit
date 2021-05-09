package event

import (
	"fmt"
	"runtime/debug"
	"sync"
)

// Hooks holds a list of parameter-less functions to call whenever the set is
// triggered with Fire().
type Hooks struct {
	funcs []func()
	mu    sync.Mutex
	wg    sync.WaitGroup
}

func (h *Hooks) Add(f func()) {
	h.mu.Lock()
	h.funcs = append(h.funcs, f)
	h.mu.Unlock()
}

// Fire calls all the functions in a given Hooks list. It launches a goroutine
// for each function and then waits for all of them to finish before returning.
func (h *Hooks) Fire() []error {
	h.mu.Lock()
	defer h.mu.Unlock()

	errc := make(chan error, len(h.funcs))
	for _, hook := range h.funcs {
		h.wg.Add(1)
		go hookWrapper(&h.wg, hook, errc)
	}

	h.wg.Wait()
	close(errc)
	errs := make([]error, 0, len(h.funcs))
	for err := range errc {
		errs = append(errs, err)
	}
	return errs
}

func hookWrapper(wg *sync.WaitGroup, hook func(), errc chan error) {
	defer func() {
		if v := recover(); v != nil {
			errc <- &RecoverError{Err: v, Stack: debug.Stack()}
		}
		wg.Done()
	}()

	hook()
}

type RecoverError struct {
	Err   interface{}
	Stack []byte
}

func (e *RecoverError) Error() string { return fmt.Sprintf("Do panicked: %v", e.Err) }
