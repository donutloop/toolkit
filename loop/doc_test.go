package loop_test

import (
	"github.com/donutloop/toolkit/loop"
	"time"
	"fmt"
)

func ExampleLoop() {

	loop.NewLooper(1*time.Second, func() error {
		// do after one second things
		return nil
	})

	// error and stop handling is missing for simplicity

	fmt.Println("successfully")
	// Output: successfully
}
