package app

import (
	"planner/backend/internal/config"
	"planner/backend/internal/telemetry"
	"planner/backend/internal/util/graceful"
)

type dependencies struct {
	gracefulTerminator graceful.Terminator
	configHolder       config.Holder
	telemetry          telemetry.Telemetry
}

func newDependencies(
	gracefulTerminator graceful.Terminator,
	configHolder config.Holder,
	telemetry telemetry.Telemetry,
) dependencies {
	return dependencies{
		gracefulTerminator: gracefulTerminator,
		configHolder:       configHolder,
		telemetry:          telemetry,
	}
}
