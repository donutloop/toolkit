# Usage

lease the resource for duration. When the lease expires, invoke func

## Example
```go 
package main 

import (
	"github.com/donutloop/toolkit/lease"
)

func main() {
    leaser := lease.NewLeaser()
    leaser.Lease("cleanup-cache", time.Duration(1*time.Second), func() {
    	fmt.Println("cleaned up cache")
    })
}
```
