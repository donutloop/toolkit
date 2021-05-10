package lease_test

import (
	"fmt"
	"time"

	"github.com/donutloop/toolkit/lease"
)

func ExampleLeaser_Lease() {
	leaser := lease.NewLeaser()
	leaser.Lease("cleanup-cache", 1*time.Second, func() {
		fmt.Println("cleaned up cache")
	})

	<-time.After(2 * time.Second)

	// Output: cleaned up cache
}
