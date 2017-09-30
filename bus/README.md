# Usage

Bus it is a simple but powerful publish-subscribe event system. It requires object to
register themselves with the event bus to receive events.

## Example 
```go 
package main 

import (
	"github.com/donutloop/toolkit/bus"
	"log"
)

type msg struct {
	Id      int64
	counter int
}

func main() {
	b := bus.New()

    b.AddEventListener(func(m *msg) error {
        m.counter++
        return nil
    })

    b.AddEventListener(func(m *msg) error {
        m.counter++
        return nil
    })

    b.Publish(new(msg))
}
```