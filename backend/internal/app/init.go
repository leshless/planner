package app

import (
	"fmt"
	"planner/backend/internal/config"
	"planner/backend/internal/telemetry"
	"planner/backend/internal/util/graceful"
	"planner/backend/internal/util/stupid"

	"go.uber.org/dig"
)

func Init(primitives Primitives) (App, error) {
	container := dig.New()

	providePrimitives(container, primitives)
	provideDependencies(container)

	app, err := resolveApp(container)
	if err != nil {
		return nil, fmt.Errorf("resolving app from DI container: %w", err)
	}

	go func() {
		<-app.interrupter.Ch()

		app.telemetry.Logger.Info("app shutdown initiated")
		app.gracefulTerminator.Terminate()
	}()

	return app, nil
}

func providePrimitives(container *dig.Container, primitives Primitives) {
	container.Provide(stupid.NewReflect(primitives.fs))
	container.Provide(stupid.NewReflect(primitives.clock))
	container.Provide(stupid.NewReflect(primitives.interrupter))
	container.Provide(stupid.NewReflect(primitives.stdIO))

	container.Provide(stupid.NewReflect(primitives))
}

func provideDependencies(container *dig.Container) {
	container.Provide(graceful.NewManager, dig.As(new(graceful.Registrator)), dig.As(new(graceful.Terminator)))
	container.Provide(config.InitHolder, dig.As(new(config.Holder)))
	container.Provide(telemetry.InitLogger, dig.As(new(telemetry.Logger)))
	container.Provide(telemetry.NewTelemetry)

	container.Provide(newDependencies)
}

func resolveApp(container *dig.Container) (*app, error) {
	var app *app
	err := container.Invoke(func(
		primitives Primitives,
		dependencies dependencies,
	) {
		app = New(
			primitives,
			dependencies,
		)
	})

	return app, err
}
