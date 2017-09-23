# Usage

promise.Do comes with a very primitive way of callback dispatch. It
immediately executes the callback, instead of scheduling it for execution

## Example 
```go 
package main 

import (
	"github.com/donutloop/toolkit/promise"
	"log"
)

func main() {
	done, errc := promise.Do(context.Background(), func(ctx context.Context) error {
		// do things
		return nil
	})

	select {
	case <-done:
	case err := <-errc:
		if err == nil {
			log.Fatal(err)
		}
	}
}
```