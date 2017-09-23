# Usage

Looper executes the provided function once in while and collects any errors.

## Example 
```go 
package main 

import (
	"github.com/donutloop/toolkit/looper"
	"log"
)


func main() {
	l := loop.NewLooper(1*time.Millisecond, func() error {
		// do things
		return nil
	})

	for err := range l.Error() {
		log.Println(err)
	}
	
	// stop call is missing
}
```


