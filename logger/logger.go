package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	cfg := zap.NewProductionConfig()
	// cfg := zap.NewDevelopmentConfig()
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.StacktraceKey = ""
	cfg.EncoderConfig = encoderCfg
	// Customize logger settings if needed, e.g., cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	var err error
	log, err = cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}
func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}
func Fatal(msg string, fields ...zap.Field) {
	log.Fatal(msg, fields...)
}
