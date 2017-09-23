package concurrent

import (
	"fmt"
	"sync"
)

// Run executes the provided functions in concurrent and collects any errors they return.
// Be careful about your resource consumption
func Run(fns ...func() error) []error {
	wg := sync.WaitGroup{}
	errc := make(chan error, len(fns))
	wg.Add(len(fns))
	for i := range fns {
		go func(do func() error, errc chan error) {
			defer func() {

				if v := recover(); v != nil {
					errc <- fmt.Errorf("do is panicked (%v)", v)
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
