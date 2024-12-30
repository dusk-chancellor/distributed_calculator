package zaplog

// zap logger init

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	dev = "dev"
	prod = "prod"
)

func New(env string) (*zap.Logger, error) {
	var cfg zap.Config

	switch env {
	case dev:
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // colorize
	case prod:
		cfg = zap.NewProductionConfig()
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // ISO8601 timestamps
	default:
		cfg = zap.NewDevelopmentConfig()
	}

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
