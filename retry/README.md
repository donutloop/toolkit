# Usage

Retry supervised do funcs which automatically handle failures when they occur by performing retries.

## Example 
```go 
package main 

import (
	"github.com/donutloop/toolkit/retry"
	"log"
	"context"
)

func main() {
	r := retry.NewRetrier()
    
    err := r.Retry(context.Background(), func() (bool, error) {
   		// do things
   		return true, nil
    })
    
    if err != nil {
        log.Fatal(err)
    }
}
```