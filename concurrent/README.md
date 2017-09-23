# Usage

Runner executes the provided functions in concurrent and collects any errors they return.

## Example 
```go 
package main 

import (
	"github.com/donutloop/toolkit/concurrent"
	"log"
)

func main() {
	concurrent.Run(
		func() error {
	        // do things
		},
		func() error {
			// do things
			return nil
		},
	)
}
```