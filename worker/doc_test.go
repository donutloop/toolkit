package worker_test

import (
	"fmt"
	"github.com/donutloop/toolkit/worker"
	"time"
)

func ExampleWorker() {
	workerHandler := func(parameter interface{}) {
		v := parameter.(string)
		fmt.Println(v)
	}

	queue := worker.New(2, workerHandler, 10)

	queue <- "hello"
	<- time.After(time.Millisecond * 250)

	// Output: hello
}
