package bus_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/donutloop/toolkit/bus"
)

type msg struct {
	Id   int64
	body string
}

func TestHandlerReturnsError(t *testing.T) {
	b := bus.New()

	err := b.AddHandler(func(m *msg) error {
		return errors.New("handler error")
	})
	if err != nil {
		t.Fatal(err)
	}

	err = b.Dispatch(new(msg))
	if err == nil {
		t.Fatalf("dispatch msg failed (%v)", err)
	}
}

func TestHandlerReturn(t *testing.T) {
	b := bus.New()

	err := b.AddHandler(func(m *msg) error {
		m.body = "Hello, world!"
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	msg := new(msg)
	err = b.Dispatch(msg)
	if err != nil {
		t.Fatalf("dispatch msg failed (%s)", err.Error())
	}

	if msg.body != "Hello, world!" {
		t.Fatal("failed to get response from handler")
	}
}

func TestEventListeners(t *testing.T) {
	b := bus.New()
	count := 0

	b.AddEventListener(func(m *msg) error {
		count += 1
		return nil
	})

	b.AddEventListener(func(m *msg) error {
		count += 10
		return nil
	})

	err := b.Publish(new(msg))
	if err != nil {
		t.Fatalf("publish msg failed (%s)", err.Error())
	}

	if count != 11 {
		t.Fatal(fmt.Sprintf("publish msg failed, listeners called: %v, expected: %v", count, 11))
	}
}

func TestAddHandlerBadFunc(t *testing.T) {
	defer func() {
		if v := recover(); v != nil {
			_, ok := v.(bus.BadFuncError)
			if !ok {
				t.Fatalf("unexpected object (%v)", v)
			}
		}
	}()

	b := bus.New()
	err := b.AddHandler(func(m *msg, s string) error {
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddListenerBadFunc(t *testing.T) {
	defer func() {
		if v := recover(); v != nil {
			_, ok := v.(bus.BadFuncError)
			if !ok {
				t.Fatalf("unexpected object (%v)", v)
			}
		}
	}()

	b := bus.New()
	b.AddEventListener(func(m *msg, s string) error {
		return nil
	})
}

func BenchmarkHandlerRun(b *testing.B) {
	for n := 0; n < b.N; n++ {
		bs := bus.New()

		bs.AddHandler(func(m *msg) error {
			return nil
		})

		if err := bs.Dispatch(new(msg)); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkListenerRun(b *testing.B) {
	for n := 0; n < b.N; n++ {
		bs := bus.New()

		bs.AddEventListener(func(m *msg) error {
			return nil
		})

		bs.AddEventListener(func(m *msg) error {
			return nil
		})

		bs.AddEventListener(func(m *msg) error {
			return nil
		})

		bs.AddEventListener(func(m *msg) error {
			return nil
		})

		bs.AddEventListener(func(m *msg) error {
			return nil
		})

		if err := bs.Publish(new(msg)); err != nil {
			b.Fatal(err)
		}
	}
}
