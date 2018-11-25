package worker_test

import (
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
	workerHandler := func(parameter worker.GenericType) {
		v, ok := parameter.(string)
		if !ok {
			t.Errorf("value is not a string got=%v", parameter)
			return
		}

		if !containes([]string{"hello", "golang", "world"}, v)	{
			t.Errorf("value is bad got=%v", parameter)
		}

		t.Logf("value: %v", v)
		atomic.AddInt32(&counter, 1)
	}

	queue := worker.New(3, workerHandler, 10)

	queue <- "hello"
	queue <- "golang"
	queue <- "world"

	<- time.After(500 * time.Millisecond)

	if atomic.LoadInt32(&counter) != 3 {
		t.Errorf("counter is bad (want=3, got=%v)", atomic.LoadInt32(&counter))
	}

	t.Logf("counter value is %v", atomic.LoadInt32(&counter))
}
