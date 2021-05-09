package singleton_test

import (
	"testing"

	"github.com/donutloop/toolkit/singleton"
)

func TestNewSingleton(t *testing.T) {

	var counter int
	stubSingleton := singleton.NewSingleton(func() (interface{}, error) {
		counter++
		return counter, nil
	})

	object, err := stubSingleton.Get()
	if err != nil {
		t.Fatal(err)
	}

	expectedValue := 1

	if object.(int) != expectedValue {
		t.Fatalf(`unexpected error message (actual: "%d", expected: "%d")`, object.(int), expectedValue)
	}

	object, err = stubSingleton.Get()
	if err != nil {
		t.Fatal(err)
	}

	if object.(int) != expectedValue {
		t.Fatalf(`unexpected error message (actual: "%d", expected: "%d")`, object.(int), expectedValue)
	}
}

func TestSingletonReset(t *testing.T) {

	var counter int
	stubSingleton := singleton.NewSingleton(func() (interface{}, error) {
		counter++
		return counter, nil
	})

	object, err := stubSingleton.Get()
	if err != nil {
		t.Fatal(err)
	}

	expectedValue := 1

	if object.(int) != expectedValue {
		t.Fatalf(`unexpected error message (actual: "%d", expected: "%d")`, object.(int), expectedValue)
	}

	object, err = stubSingleton.Get()
	if err != nil {
		t.Fatal(err)
	}

	if object.(int) != expectedValue {
		t.Fatalf(`unexpected error message (actual: "%d", expected: "%d")`, object.(int), expectedValue)
	}

	expectedValue = 2

	stubSingleton.Reset()

	object, err = stubSingleton.Get()
	if err != nil {
		t.Fatal(err)
	}

	if object.(int) != expectedValue {
		t.Fatalf(`unexpected error message (actual: "%d", expected: "%d")`, object.(int), expectedValue)
	}

	object, err = stubSingleton.Get()
	if err != nil {
		t.Fatal(err)
	}

	if object.(int) != expectedValue {
		t.Fatalf(`unexpected error message (actual: "%d", expected: "%d")`, object.(int), expectedValue)
	}
}

func BenchmarkSingleton_Get(b *testing.B) {
	stubSingleton := singleton.NewSingleton(func() (interface{}, error) {
		return nil, nil
	})

	for n := 0; n < b.N; n++ {
		_, err := stubSingleton.Get()
		if err != nil {
			b.Fatal(err)
		}
	}
}
