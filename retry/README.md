# Usage of Retrier

Retry supervised do funcs which automatically handle failures when they occur by performing retries.

## Example 
```go 
package main 

import (
	"github.com/donutloop/toolkit/retry"
	"log"
	"context"
)

func main() {

    r := retry.NewRetrier()
  
    err := r.Retry(context.Background(), func() (bool, error) {
   		// do things
   		return true, nil
    })
    
    if err != nil {
        log.Fatal(err)
    }
}
```

# Usage of Roundtripper Retrier

## Example 
```go 
package main 

import (
	"github.com/donutloop/toolkit/retry"
	"log"
	"context"
)

func main() {
        retryRoundTripper := retry.NewRoundTripper(http.DefaultTransport, .50 , .15 , 3, []int{http.StatusBadRequest})
	httpClient := new(http.Client)
	httpClient.Transport = retryRoundTripper

	req, err := http.NewRequest(http.MethodGet, "http://example.com", nil )
	if err != nil {
		//...
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		//...
	}
}
```