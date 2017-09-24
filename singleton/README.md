# Usage

the singleton pattern is a software design pattern that restricts the instantiation of a type to one object

## Example 

```go 
package main 

import (
	"github.com/donutloop/toolkit/singleton"
)

type config struct {
	Addr string
	Port int
}

var configSingleton = NewSingleton(func() (interface{}, error) {
	return &config{Addr:"localhost", Port:80,}, nil
})

func Config() (*config, error) {
	s, err := configSingleton.Get()
	if err != nil {
		return nil, err
	}
	return s.(*config), nil
}

func main() {
	config, err := Config()
	// do things 
}
```