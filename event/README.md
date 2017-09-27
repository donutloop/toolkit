# Usage

Hooks holds a list of functions (func error) to call whenever the set is
triggered.

## Example 
```go 
package main 

import (
	"github.com/donutloop/toolkit/event"
)

func main() {
    hooks := new(event.Hooks)
    hooks.Add(func() { 
    	// do things
    })
    hooks.Add(func() { 
        // do things 
    })
    hooks.Fire()
}
```