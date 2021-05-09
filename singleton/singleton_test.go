package singleton

import (
	"testing"
)

func TestNewSingleton(t *testing.T) {

	var counter int
	stubSingleton := NewSingleton(func() (interface{}, error) {
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
	stubSingleton := NewSingleton(func() (interface{}, error) {
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
	stubSingleton := NewSingleton(func() (interface{}, error) {
		return nil, nil
	})

	for n := 0; n < b.N; n++ {
		stubSingleton.Get()
	}
}
