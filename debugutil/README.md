# Usage

PrettySprint creates a human readable representation of the value v.

#### Example 
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

#### Example 

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

PrettyRequestDump creates a human readable representation of the value http.Request.

#### Example 

```go 
package main 

import (
	"github.com/donutloop/toolkit/debugutil"
	"log"
	"net/http"
)

func main() {

    req := &http.Request{}
    s , err := debugutil.PrettySprintRequest(req)
    if err != nil {
        log.Fatal(err)
    }    
    log.Println(s)
}
```

LogRoundTripper which logs all requests (request and response dump)

#### Example 

```go 
package main 

import (
	"github.com/donutloop/toolkit/debugutil"
	"net/http"
)

type logger struct {}

func (l logger) Errorf(format string, v ...interface{}) {
	log.Println(fmt.Sprintf(format, v...))
}
func (l logger) Infof(format string, v ...interface{}) {
	log.Println(fmt.Sprintf(format, v...))
}

func main() {

	httpClient := new(http.Client)
	httpClient.Transport = debugutil.NewLogRoundTripper(http.DefaultTransport, logger{}, true)

	response, err := httpClient.Get(server.URL)
}
```