package log

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogLevel holds an AtomicLevel that can be used to change the logging
// level of Default.
var LogLevel = zap.NewAtomicLevelAt(zap.DebugLevel)

// loggerKey holds the context key used for loggers.
type loggerKey struct{}

// WithLogger returns a new context derived from ctx that
// is associated with the given logger.
func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// WithFields returns a new context derived from ctx
// that has a logger that always logs the given fields.
func WithFields(ctx context.Context, fields ...zapcore.Field) context.Context {
	return WithLogger(ctx, GetLogger(ctx).With(fields...))
}

// GetLogger returns the logger associated with the given
// context. If there is no logger, it will return new logger.
func GetLogger(ctx context.Context) *zap.Logger {
	if ctx == nil {
		panic("nil context passed to Logger")
	}
	if logger, _ := ctx.Value(loggerKey{}).(*zap.Logger); logger != nil {
		return logger
	}
	logger, _ := zap.NewProduction()
	return logger
}
