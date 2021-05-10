package concurrent_test

import (
	"fmt"
	"sync/atomic"

	"github.com/donutloop/toolkit/concurrent"
)

// Run concurrently your func() error.
func ExampleRun() {
	counter := int32(0)
	errs := concurrent.Run(
		func() error {
			atomic.AddInt32(&counter, 40)
			return nil
		},
		func() error {
			atomic.AddInt32(&counter, 2)
			return nil
		},
	)

	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Printf("error: %v \n", err)
		}
	}

	fmt.Println(counter)
	// Output: 42
}
