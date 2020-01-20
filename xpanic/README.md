# Usage

It's a powerful panic handler to simplify reasoning of panics

## Example 
```go 
package main 

import (
	"github.com/donutloop/toolkit/xpanic"
	"log"
)

func main() {
    logF := func(format string, args ...interface{}) { log.Println(fmt.Sprintf(format, args...)) }
    panicHandler := BuildPanicHandler(logF, xpanic.CrashOnErrorDeactivated)
    defer panicHandler
    panic("hello world")
}
```