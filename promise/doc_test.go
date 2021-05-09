package promise_test

import (
	"context"
	"fmt"

	"github.com/donutloop/toolkit/promise"
)

func Example() {

	done, errc := promise.Do(context.Background(), func(ctx context.Context) error {
		fmt.Println("do things")
		return nil
	})

	select {
	case <-done:
	case err := <-errc:
		fmt.Printf("error: %v \n", err)
	}

	// Output: do things
}
