package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ZapLogger(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		rl := middleware.RequestLogger(NewZapLogFormatter(logger))
		return rl(next)
	}
}

type ZapLogEntry struct {
	logger *zap.Logger
	fields []zapcore.Field
}

func NewZapLogEntry(logger *zap.Logger, fields []zapcore.Field) *ZapLogEntry {
	return &ZapLogEntry{
		logger: logger,
		fields: fields,
	}
}

func (l *ZapLogEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	l.logger.With(l.fields...).Info(
		"http request",
		zap.Int("http.code", status),
		zap.Int("http.size", bytes),
		zap.Float32("http.time_ms", float32(elapsed.Nanoseconds()/1000)/1000),
	)
}

func (l *ZapLogEntry) Panic(v interface{}, stack []byte) {
	var msg = "panic occured"
	if e, ok := v.(error); ok {
		msg = e.Error()
	}
	entry := zapcore.Entry{
		Level:   zap.PanicLevel,
		Time:    time.Now(),
		Message: msg,
		Stack:   string(stack),
	}
	l.logger.Core().Write(entry, nil)
}

type ZapLogFormatter struct {
	logger *zap.Logger
}

func NewZapLogFormatter(logger *zap.Logger) *ZapLogFormatter {
	return &ZapLogFormatter{
		logger: logger,
	}
}

func (l *ZapLogFormatter) NewLogEntry(r *http.Request) middleware.LogEntry {
	fields := []zapcore.Field{
		zap.String("http.proto", r.Proto),
		zap.String("http.method", r.Method),
		zap.String("http.path", r.URL.Path),
		zap.String("access.client.ip", r.RemoteAddr),
		zap.String("access.user_agent", r.UserAgent()),
	}
	return &ZapLogEntry{
		logger: l.logger,
		fields: fields,
	}
}
