package event_test

import (
	"testing"

	"github.com/donutloop/toolkit/event"
)

func TestHooksMultiFire(t *testing.T) {
	hooks := new(event.Hooks)
	hooks.Add(func() {})
	hooks.Add(func() {})

	for i := 0; i < 10; i++ {
		errs := hooks.Fire()
		if len(errs) > 0 {
			for _, err := range errs {
				t.Error(err)
			}
		}
	}
}

func TestHooks(t *testing.T) {
	triggered1 := false
	triggered2 := false

	hooks := new(event.Hooks)
	hooks.Add(func() { triggered1 = true })
	hooks.Add(func() { triggered2 = true })

	hooks.Fire()

	if !triggered1 {
		t.Errorf("registered (first) hook function failed to trigger")
	}

	if !triggered2 {
		t.Errorf("registered (second) hook function failed to trigger")
	}
}

func TestHooksPanic(t *testing.T) {
	hooks := new(event.Hooks)
	hooks.Add(func() { panic("check isolation of goroutine") })
	errs := hooks.Fire()
	if len(errs) != 1 {
		t.Fatalf("error count is bad (%d)", len(errs))
	}

	expectedMessage := "hook is panicked (check isolation of goroutine)"
	if errs[0].Error() != expectedMessage {
		t.Fatalf(`unexpected error message (actual: "%s", expected: "%s")`, errs[0].Error(), expectedMessage)
	}
}

func BenchmarkHooks(b *testing.B) {
	hooks := new(event.Hooks)
	hooks.Add(func() {})
	for n := 0; n < b.N; n++ {
		hooks.Fire()
	}
}
