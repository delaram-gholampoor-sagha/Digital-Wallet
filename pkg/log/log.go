package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	OutputPaths       []string
	ErrorOutputPaths  []string
	DisableStacktrace bool
	Level             zapcore.Level
}

func New(service string, cfg Config) (*zap.SugaredLogger, error) {
	zapConfig := zap.NewProductionConfig()
	zapConfig.OutputPaths = cfg.OutputPaths
	zapConfig.ErrorOutputPaths = cfg.ErrorOutputPaths
	zapConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	zapConfig.DisableStacktrace = cfg.DisableStacktrace
	zapConfig.Level.SetLevel(cfg.Level)
	zapConfig.InitialFields = map[string]interface{}{
		"service": service,
	}

	logger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}
