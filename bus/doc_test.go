package bus_test

import (
	"fmt"

	"github.com/donutloop/toolkit/bus"
)

// Creates a bus and adds a listener to a message afterward it publishes a new message.
func ExampleInProcBus_AddEventListener() {
	type msg struct {
		ID   int64
		body string
	}

	b := bus.New()

	b.AddEventListener(func(m *msg) error {
		fmt.Println("db insert listener id", m.ID)
		return nil
	})

	m := new(msg)
	m.ID = 1
	m.body = "test"

	if err := b.Publish(m); err != nil {
		fmt.Printf("error: (%v) \n", err)
	}

	// Output: db insert listener id 1
}

// Creates a bus and adds a handler for a message afterward it dispatch a new message.
func ExampleInProcBus_AddHandler() {
	type msg struct {
		ID   int64
		body string
	}

	b := bus.New()

	err := b.AddHandler(func(m *msg) error {
		fmt.Println("db insert listener id", m.ID)
		return nil
	})

	if err != nil {
		fmt.Printf("error: (%v) \n", err)
		return
	}

	m := new(msg)
	m.ID = 1
	m.body = "test"

	if err := b.Dispatch(m); err != nil {
		fmt.Printf("error: (%v) \n", err)
	}

	// Output: db insert listener id 1
}
