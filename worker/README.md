# Usage

worker.New() starts n * Workers goroutines running func on incoming
parameters sent on the returned channel.

## Example 
```go 
package main 

import (
	"github.com/donutloop/toolkit/worker"
	"log"
)

func main() {
	workerHandler := func(parameter interface{}) {
		v := parameter.(string)
		log.Println(v)	
	}

	queue := worker.New(2, workerHandler, 10)

	queue <- "hello"
	queue <- "world"
}
```