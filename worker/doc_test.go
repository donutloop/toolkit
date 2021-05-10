package worker_test

import (
	"fmt"
	"time"

	"github.com/donutloop/toolkit/worker"
)

func Example() {
	workerHandler := func(parameter interface{}) (interface{}, error) {
		v := parameter.(string)
		return v + " world", nil
	}

	request, response, _ := worker.New(2, workerHandler, 10)

	request <- "hello"

	<-time.After(time.Millisecond * 250)

	fmt.Println(<-response)

	// Output: hello world
}
