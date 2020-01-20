package xpanic_test

import (
	"bytes"
	"fmt"
	"github.com/donutloop/toolkit/xpanic"
	"sync"
	"testing"
)

func TestHandlePanic(t *testing.T) {

	var buff bytes.Buffer
	panicHandler := xpanic.BuildPanicHandler(func(format string, args ...interface{}) { buff.WriteString(fmt.Sprintf(format, args...)) }, xpanic.CrashOnErrorDeactivated)

	var wait sync.WaitGroup
	wait.Add(1)
	go func() {
		defer panicHandler()
		defer wait.Done()
		panic("hello world")
	}()

	wait.Wait()

	if buff.Len() == 0 {
		t.Fatal("buff is empty")
	}
}
