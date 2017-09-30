package event_test

import (
	"fmt"
	"github.com/donutloop/toolkit/event"
)

func ExampleHooks() {

	hooks := new(event.Hooks)
	hooks.Add(func() { fmt.Println("kernel request") })

	errs := hooks.Fire()
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(fmt.Sprintf("error: %v", err))
		}
	}

	// Output: kernel request
}
