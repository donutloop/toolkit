package concurrent

import (
	"fmt"
	"runtime/debug"
	"sync"
)

type RecoverError struct {
	Err   interface{}
	Stack []byte
}

func (e *RecoverError) Error() string { return fmt.Sprintf("Do panicked: %v", e.Err) }

// Run executes the provided functions in concurrent and collects any errors they return.
// Be careful about your resource consumption.
func Run(fns ...func() error) []error {
	wg := sync.WaitGroup{}
	errc := make(chan error, len(fns))
	wg.Add(len(fns))
	for i := range fns {
		go func(do func() error, errc chan error) {
			defer func() {

				if v := recover(); v != nil {
					errc <- &RecoverError{Err: v, Stack: debug.Stack()}
				}

				wg.Done()
			}()

			if err := do(); err != nil {
				errc <- err
			}
		}(fns[i], errc)
	}
	wg.Wait()
	close(errc)
	errs := make([]error, 0, len(fns))
	for err := range errc {
		errs = append(errs, err)
	}
	return errs
}
