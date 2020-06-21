package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func Init() {
	switch os.Getenv("GO_ENV") {
	case "production":
		Logger = NewMyLoggerProduction()
	default:
		Logger = NewMyLoggerDevelopment()
	}
}

func NewLoggerConfig() *zap.Config {
	return &zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: true,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "severity",
			TimeKey:        "timestamp",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
			EncodeDuration: zapcore.MillisDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}
}

func NewCore(cfg *zap.Config) zapcore.Core {
	return zapcore.NewTee(
		newStdoutCore(cfg.EncoderConfig),
		newStderrCore(cfg.EncoderConfig),
	)
}

func newStdoutCore(encoderConfig zapcore.EncoderConfig) zapcore.Core {
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.Lock(os.Stdout),
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl <= zapcore.ErrorLevel
		}),
	)
}

func newStderrCore(encoderConfig zapcore.EncoderConfig) zapcore.Core {
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.Lock(os.Stderr),
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl > zapcore.ErrorLevel
		}),
	)
}

func NewMyLoggerProduction() *zap.Logger {
	cfg := NewLoggerConfig()
	cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	cfg.Development = false
	core := NewCore(cfg)
	return zap.New(core)
}

func NewMyLoggerDevelopment() *zap.Logger {
	cfg := NewLoggerConfig()
	core := NewCore(cfg)
	return zap.New(core)
}
