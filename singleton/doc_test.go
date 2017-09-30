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

	configFunc()

	c, err := configFunc()
	if err != nil {
		fmt.Println(fmt.Sprintf("error: (%v)", err))
	}

	fmt.Println(fmt.Sprintf("%#v", c))
	// Output: &singleton_test.config{Addr:"localhost", Port:80}
}
