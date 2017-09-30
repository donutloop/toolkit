package bus_test

import (
	"fmt"
	"github.com/donutloop/toolkit/bus"
)

// Creates a bus and adds a listener to a message afterward it publishes a new message
func ExampleBusListener() {

	type msg struct {
		Id   int64
		body string
	}

	b := bus.New()

	b.AddEventListener(func(m *msg) error {
		fmt.Println("db insert listener")
		return nil
	})

	if err := b.Publish(new(msg)); err != nil {
		fmt.Println(fmt.Sprintf("bus: %v", err))
	}

	// Output: db insert listener
}

// Creates a bus and adds a handler for a message afterward it dispatch a new message
func ExampleBusHandler() {

	type msg struct {
		Id   int64
		body string
	}

	b := bus.New()

	b.AddHandler(func(m *msg) error {
		fmt.Println("db insert listener")
		return nil
	})

	if err := b.Dispatch(new(msg)); err != nil {
		fmt.Println(fmt.Sprintf("bus: %v", err))
	}

	// Output: db insert listener
}
