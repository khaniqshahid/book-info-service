package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func init() {
	cfg := zap.NewProductionConfig()
	// cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// Customize logger settings if needed, e.g., cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	Logger, _ = cfg.Build()
}
