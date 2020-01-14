package worker_test

import (
	"fmt"
	"github.com/donutloop/toolkit/worker"
	"sync/atomic"
	"testing"
	"time"
)

func TestWorker(t *testing.T) {

	containes := func(ls []string, s string) bool {
		for _, ss := range ls {
			if ss == s {
				return true
			}
		}
		return false
	}

	counter := int32(0)
	workerHandler := func(parameter interface{}) (interface{}, error) {
		v, ok := parameter.(string)
		if !ok {
			return false, fmt.Errorf("value is not a string got=%v", parameter)
		}

		if !containes([]string{"hello", "golang", "world"}, v) {
			t.Errorf("value is bad got=%v", parameter)
		}

		t.Logf("value: %v", v)
		atomic.AddInt32(&counter, 1)
		return true, nil
	}

	request, response, errs := worker.New(3, workerHandler, 10)

	request <- "hello"
	request <- "golang"
	request <- "world"

	go func() {
		for err := range errs {
			t.Error(err)
		}
	}()

	go func() {
		for v := range response {
			if !v.(bool) {
				t.Error("bad type")
			}
		}
	}()

	<-time.After(500 * time.Millisecond)

	if atomic.LoadInt32(&counter) != 3 {
		t.Errorf("counter is bad (want=3, got=%v)", atomic.LoadInt32(&counter))
	}

	t.Logf("counter value is %v", atomic.LoadInt32(&counter))
}
