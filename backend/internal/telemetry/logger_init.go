package telemetry

import (
	"fmt"
	"io"
	"planner/backend/internal/config"
	"planner/backend/internal/util/graceful"
	"planner/backend/internal/util/stdio"

	xfs "github.com/hack-pad/hackpadfs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(
	fs xfs.FS,
	stdIO stdio.StdIO,
	configHolder config.Holder,
	gracefulRegistrator graceful.Registrator,
) (Logger, error) {
	config := configHolder.Config().Logger

	var (
		zapEncoderConfig zapcore.EncoderConfig
		zapLevel         zapcore.Level
	)
	if config.Development {
		zapEncoderConfig = zap.NewDevelopmentEncoderConfig()
		zapLevel = zap.DebugLevel
	} else {
		zapEncoderConfig = zap.NewProductionEncoderConfig()
		zapLevel = zap.InfoLevel
	}

	file, err := fs.Open(config.FileOutputPath)
	if err != nil {
		return nil, fmt.Errorf("opening output path %q: %w", config.FileOutputPath, err)
	}

	gracefulRegistrator.Register(file.Close)

	// writer, ok := file.(io.Writer)
	// if !ok {
	// 	return nil, errors.New("provided output path is not writeable file")
	// }

	writer := io.MultiWriter(stdIO.StdOut())

	zapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapEncoderConfig),
		zapcore.AddSync(writer),
		zap.NewAtomicLevelAt(zapLevel),
	)

	zapLogger := zap.New(zapCore)

	return NewLogger(zapLogger), nil
}
