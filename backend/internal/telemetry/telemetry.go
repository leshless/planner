package telemetry

type Telemetry struct {
	Logger Logger
}

func NewTelemetry(
	logger Logger,
) Telemetry {
	return Telemetry{
		Logger: logger,
	}
}
