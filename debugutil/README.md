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
    b , err := debugutil.PrettyDumpResponse(resp, true)
    if err != nil {
        log.Fatal(err)
    }    
    log.Println(string(b))
}
```