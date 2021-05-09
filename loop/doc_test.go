package loop_test

import (
	"fmt"
	"time"

	"github.com/donutloop/toolkit/loop"
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
