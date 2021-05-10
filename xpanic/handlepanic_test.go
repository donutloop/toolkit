package xpanic_test

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
	"testing"

	"github.com/donutloop/toolkit/xpanic"
)

func TestHandlePanic(t *testing.T) {
	// if darwin then skip this test
	if runtime.GOOS == "darwin" {
		t.Skip()
	}

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
