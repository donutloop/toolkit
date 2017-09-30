package loop_test

import (
	"fmt"
	"github.com/donutloop/toolkit/loop"
	"time"
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
