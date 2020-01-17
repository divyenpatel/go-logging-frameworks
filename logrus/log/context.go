package log

import (
	"context"
	"github.com/sirupsen/logrus"
)

type (
	loggerKey struct{}
)

// RFC3339NanoFixed is time.RFC3339Nano with nanoseconds padded using zeros to
// ensure the formatted time is always the same number of characters.
const RFC3339NanoFixed = "2006-01-02T15:04:05.000000000Z07:00"

// WithLogger returns a new context with the provided logger. Use in
// combination with logger.WithField(s) for great effect.
func WithLogger(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// WithFields returns a new context derived from ctx
// that has a logger that always logs the given fields.
func WithFields(ctx context.Context, fields logrus.Fields) context.Context {
	return WithLogger(ctx, GetLogger(ctx).WithFields(fields))
}

// GetLogger retrieves the current logger from the context. If no logger is
// available, the default logger is returned.
func GetLogger(ctx context.Context) *logrus.Entry {
	logger := ctx.Value(loggerKey{})
	if logger == nil {
		return logrus.NewEntry(logrus.StandardLogger())
	}
	return logger.(*logrus.Entry)
}