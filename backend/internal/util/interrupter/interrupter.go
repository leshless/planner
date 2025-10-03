package interrupter

import (
	"os"
	"os/signal"
	"syscall"
)

type Interrupter interface {
	Ch() <-chan struct{}
}

type interrupter struct {
	ch chan struct{}
}

var _ Interrupter = (*interrupter)(nil)

func NewInterrupter() Interrupter {
	ch := make(chan struct{}, 1)
	signalCh := make(chan os.Signal, 1)

	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalCh
		ch <- struct{}{}
	}()

	return &interrupter{
		ch: ch,
	}
}

func (i *interrupter) Ch() <-chan struct{} {
	return i.ch
}
