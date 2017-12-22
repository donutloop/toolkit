package lease

import "time"

type Leaser interface {
	// lease the resource r for duration. When the lease expires, invoke func call.
	// revoke a lease can be refreshed by calling Lease() again on the same resource
	Lease(r string, d time.Duration, revoke func())
	// if resource exists then cancel the old timer and delete the entry
	Return(r string) bool
}

// NewLeaser creates a instance of leaser
func NewLeaser() Leaser {
	return &leaser{
		timers: make(map[string]*time.Timer),
	}
}

type leaser struct {
	timers map[string]*time.Timer
}

func (l *leaser) Lease(r string, d time.Duration, f func()) {
	timer := time.AfterFunc(d, f)
	// if resource exists then cancel the old timer and overwrite the old one
	if t, ok := l.timers[r]; ok {
		t.Stop()
	}
	l.timers[r] = timer
}

// if resource exists then cancel the old timer and delete the entry
func (l *leaser) Return(r string) bool {
	if t, ok := l.timers[r]; ok {
		t.Stop()
		return true
	}
	return false
}
