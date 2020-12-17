package osutil

import (
	"os"
	"os/signal"
	"syscall"
)

func RegisterExistSignal() <-chan os.Signal {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	return sc
}