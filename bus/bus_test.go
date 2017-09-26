package bus_test

import (
	"errors"
	"fmt"
	"github.com/donutloop/toolkit/bus"
	"testing"
)

type msg struct {
	Id   int64
	body string
}

func TestHandlerReturnsError(t *testing.T) {
	b := bus.New()

	b.AddHandler(func(m *msg) error {
		return errors.New("handler error")
	})

	err := b.Dispatch(new(msg))
	if err == nil {
		t.Fatal("Send query failed " + err.Error())
	}
}

func TestHandlerReturn(t *testing.T) {
	b := bus.New()

	b.AddHandler(func(q *msg) error {
		q.body = "Hello, world!"
		return nil
	})

	msg := new(msg)
	err := b.Dispatch(msg)

	if err != nil {
		t.Fatal("Send msg failed " + err.Error())
	}

	if msg.body != "Hello, world!" {
		t.Fatal("Failed to get response from handler")
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
		t.Fatal("Publish msg failed " + err.Error())
	}

	if count != 11 {
		t.Fatal(fmt.Sprintf("Publish msg failed, listeners called: %v, expected: %v", count, 11))
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
	b.AddHandler(func(q *msg, s string) error {
		return nil
	})
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
	b.AddEventListener(func(q *msg, s string) error {
		return nil
	})
}

func BenchmarkRun(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b := bus.New()

		b.AddEventListener(func(m *msg) error {
			return nil
		})

		b.AddEventListener(func(m *msg) error {
			return nil
		})

		b.AddEventListener(func(m *msg) error {
			return nil
		})

		b.AddEventListener(func(m *msg) error {
			return nil
		})

		b.AddEventListener(func(m *msg) error {
			return nil
		})

		b.AddHandler(func(q *msg) error {
			return nil
		})

		b.AddHandler(func(q *msg) error {
			return nil
		})

		b.AddHandler(func(q *msg) error {
			return nil
		})

		b.AddHandler(func(q *msg) error {
			return nil
		})

		b.AddHandler(func(q *msg) error {
			return nil
		})

		b.Dispatch(new(msg))
	}
}
