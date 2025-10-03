package app

import (
	"errors"
	"sync/atomic"
)

type App interface {
	Run() error
}

type app struct {
	Primitives
	dependencies

	isStarted atomic.Bool
}

var _ App = (*app)(nil)

func New(
	primitives Primitives,
	dependencies dependencies,
) *app {
	return &app{
		Primitives:   primitives,
		dependencies: dependencies,
	}
}

func (app *app) Run() error {
	if app.isStarted.Load() {
		return errors.New("app is already started")
	}
	app.isStarted.Store(true)

	app.telemetry.Logger.Info("app successfully started")

	return nil
}
