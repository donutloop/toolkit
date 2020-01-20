package handlepanic

import (
	"bytes"
	"fmt"
	"os/signal"
	"runtime/pprof"
	"strings"
	"syscall"
)

// HandlePanic writes a message to the logger and causes a backtrace to be produced.
// If crashOnError is set a coredump will be produced and kill program else it continues.
func HandlePanic(Errorf func(format string, args ...interface{}), crashOnError bool) {
	r := recover()
	if r != nil {
		Errorf("capture panic infos")

		Errorf(fmt.Sprintf("panic: %s", r))
		Backtrace(Errorf)

		if crashOnError {
			signal.Reset(syscall.SIGABRT)
			Errorf("finished capturing of panic infos")
			syscall.Kill(0, syscall.SIGABRT)
		} else {
			Errorf("finished capturing of panic infos")
		}
	}
}

// Backtrace writes a multi-line backtrace to the logger.
func Backtrace(Errorf func(format string, args ...interface{})) {
	profiles := pprof.Profiles()
	buf := new(bytes.Buffer)

	for _, p := range profiles {
		// https://golang.org/pkg/runtime/pprof/#Profile.WriteTo.
		err := pprof.Lookup(p.Name()).WriteTo(buf, 2)
		if err != nil {
			Errorf("could not write profile: %v", err)
		}
	}

	for _, line := range strings.Split(buf.String(), "\n") {
		Errorf(line)
	}
}
