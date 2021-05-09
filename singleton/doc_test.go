package singleton_test

import (
	"fmt"

	"github.com/donutloop/toolkit/singleton"
)

func ExampleSingleton() {

	type config struct {
		Addr string
		Port int
	}

	configSingleton := singleton.NewSingleton(func() (interface{}, error) {
		return &config{Addr: "localhost", Port: 80}, nil
	})

	configFunc := func() (*config, error) {
		s, err := configSingleton.Get()
		if err != nil {
			return nil, err
		}
		return s.(*config), nil
	}
	
	c, err := configFunc()
	if err != nil {
		fmt.Printf("error: (%v) \n", err)
	}

	fmt.Printf("%#v \n", c)
	// Output: &singleton_test.config{Addr:"localhost", Port:80}
}
