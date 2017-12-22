# Usage

MultiError concatenate errors into one error.

## Example
```go 
package main 

import (
	"github.com/donutloop/toolkit/multierror"
	"fmt"
)

func main() {
    errs := []error{
		errors.New("error connect to db failed"),
		errors.New("error marschaling json"),
    }
    fmt.Println(multierror.New(errs...))
}
```