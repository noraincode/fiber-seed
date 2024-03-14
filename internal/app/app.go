package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/noraincode/fiber-seed/internal/app/config"
)

func Init() (func(), error) {
	config.Init()

	injector, _, err := BuildInjector(context.Background())
	if err != nil {
		panic("Failed to build the injector, err: " + err.Error())
	}

	if err := injector.Srv.Run(); err != nil {
		panic("Failed to run the server, err: " + err.Error())
	}

	return func() {
		injector.Srv.Shutdown()
	}, nil
}

func Run() {
	cleanFunc, err := Init()
	if err != nil {
		panic(err)
	}
	shutdown := make(chan struct{})
	registerSignal(shutdown)
	<-shutdown
	cleanFunc()
}

func registerSignal(shutdown chan struct{}) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, []os.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1}...)
	go func() {
		for sig := range c {
			if handleSignals(sig) {
				close(shutdown)
				return
			}
		}
	}()
}

func handleSignals(signal os.Signal) (exitNow bool) {
	switch signal {
	case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
		return true
	case syscall.SIGUSR1:
		return false
	}
	return false
}
