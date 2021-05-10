package xpanic

import (
	"bytes"
	"fmt"
	"os/signal"
	"runtime/pprof"
	"strings"
	"syscall"
)

// If crashOnError is set a coredump will be produced and kill program else it continues.
const (
	CrashOnErrorActivated   = true
	CrashOnErrorDeactivated = false
)

// BuildPanicHandler builds a panic handler and verifies a none nil logger got passed.
func BuildPanicHandler(errorf func(format string, args ...interface{}), crashOnError bool) func() {
	if errorf == nil {
		panic("errorf is not set")
	}
	// handlePanic writes a message to the logger and causes a backtrace to be produced.
	return func() {
		r := recover()
		if r != nil {
			errorf("capture panic infos")

			errorf(fmt.Sprintf("panic: %s", r))
			backtrace(errorf)

			if crashOnError {
				signal.Reset(syscall.SIGABRT)
				errorf("finished capturing of panic infos")

				err := syscall.Kill(0, syscall.SIGABRT)
				if err != nil {
					errorf("syscall.Kill failed: %v", err)
				}
			} else {
				errorf("finished capturing of panic infos")
			}
		}
	}
}

// Backtrace writes a multi-line backtrace to the logger.
func backtrace(errorf func(format string, args ...interface{})) {
	profiles := pprof.Profiles()
	buf := new(bytes.Buffer)

	for _, p := range profiles {
		// https://golang.org/pkg/runtime/pprof/#Profile.WriteTo.
		err := pprof.Lookup(p.Name()).WriteTo(buf, 2)
		if err != nil {
			errorf("could not write profile: %v", err)
		}
	}

	for _, line := range strings.Split(buf.String(), "\n") {
		errorf(line)
	}
}
