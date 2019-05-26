# Usage

PrettySprint creates a human readable representation of the value v.

## Example 
```go 
package main 

import (
	"github.com/donutloop/toolkit/debugutil"
	"log"
)

func main() {
    log.Println(debugutil.PrettySprint([]string{}))
}
```

PrettyResponseDump creates a human readable representation of the value http.Response.

## Example 

```go 
package main 

import (
	"github.com/donutloop/toolkit/debugutil"
	"log"
	"net/http"
)

func main() {

    resp := &http.Response{}
    s , err := debugutil.PrettySprintResponse(resp)
    if err != nil {
        log.Fatal(err)
    }    
    log.Println(s)
}
```