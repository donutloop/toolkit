package promise_test

import (
	"context"
	"fmt"
	"github.com/donutloop/toolkit/promise"
)

func ExamplePromise() {

	done, errc := promise.Do(context.Background(), func(ctx context.Context) error {
		fmt.Println("do things")
		return nil
	})

	select {
	case <-done:
	case err := <-errc:
		fmt.Println(fmt.Sprintf("error: %v", err))
	}

	// Output: do things
}
