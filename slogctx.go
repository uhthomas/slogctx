package slogctx

import (
	"context"
	"log/slog"
)

type key struct{}

// From will return the logger from the context if present, or the default
// logger.
func From(ctx context.Context) *slog.Logger {
	if l, ok := ctx.Value(key{}).(*slog.Logger); ok {
		return l
	}
	return slog.Default()
}

// With returns a copy of the parent context with the logger.
func With(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, key{}, l)
}

// WithArgs returns a copy of the parent context with the logger which includes
// the given attributes in each output operation.
func WithArgs(ctx context.Context, args ...any) context.Context {
	return With(ctx, From(ctx).With(args...))
}

// WithArgs returns a copy of the parent context with the logger that
// starts a group, if name is non-empty.
func WithGroup(ctx context.Context, name string) context.Context {
	return With(ctx, From(ctx).WithGroup(name))
}

// Log gets the logger from the context and passes the given arguments to Log.
func Log(ctx context.Context, level slog.Level, msg string, args ...any) {
	From(ctx).Log(ctx, level, msg, args...)
}

// Debug gets the logger from the context and passes the given arguments to
// DebugContext.
func Debug(ctx context.Context, msg string, args ...any) {
	From(ctx).DebugContext(ctx, msg, args...)
}

// Info gets the logger from the context and passes the given arguments to
// InfoContext.
func Info(ctx context.Context, msg string, args ...any) {
	From(ctx).InfoContext(ctx, msg, args...)
}

// Warn gets the logger from the context and passes the given arguments to
// WarnContext.
func Warn(ctx context.Context, msg string, args ...any) {
	From(ctx).WarnContext(ctx, msg, args...)
}

// Error gets the logger from the context and passes the given arguments to
// ErrorContext.
func Error(ctx context.Context, msg string, args ...any) {
	From(ctx).ErrorContext(ctx, msg, args...)
}

// Enabled reports whether the logger from the context emits log records at the
// given level.
func Enabled(ctx context.Context, level slog.Level) bool {
	return From(ctx).Enabled(ctx, level)
}
