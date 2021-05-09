package lease_test

import (
	"testing"
	"time"

	"sync/atomic"

	"github.com/donutloop/toolkit/lease"
)

func TestLeaser_Lease(t *testing.T) {
	var counter int32
	leaser := lease.NewLeaser()
	leaser.Lease("cleanup-cache", 1*time.Second, func() {
		atomic.AddInt32(&counter, 1)
	})

	<-time.After(2 * time.Second)

	if counter != 1 {
		t.Errorf(`unexpected counter value (actual:"%d", expected: "%d")`, counter, 1)
	}
}

func TestLeaser_OverwriteLease(t *testing.T) {
	var counter int32
	leaser := lease.NewLeaser()
	leaser.Lease("cleanup-cache", 2*time.Second, func() {
		atomic.AddInt32(&counter, 1)
	})

	leaser.Lease("cleanup-cache", 1*time.Second, func() {
		atomic.AddInt32(&counter, 2)
	})

	<-time.After(3 * time.Second)

	if counter != 2 {
		t.Errorf(`unexpected counter value (actual:"%d", expected: "%d")`, counter, 2)
	}
}

func TestLeaser_Return(t *testing.T) {
	var counter int32
	leaser := lease.NewLeaser()
	leaser.Lease("cleanup-cache", 1*time.Second, func() {
		atomic.AddInt32(&counter, 1)
	})

	ok := leaser.Return("cleanup-cache")
	if !ok {
		t.Error("error couldn't return resource")
	}

	<-time.After(2 * time.Second)

	if counter != 0 {
		t.Errorf(`unexpected counter value (actual:"%d", expected: "%d")`, counter, 0)
	}
}
