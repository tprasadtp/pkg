// SPDX-FileCopyrightText: Copyright 2023 Prasad Tengse
// SPDX-License-Identifier: MIT

package log

import (
	"context"
	"log/slog"
)

// context key for storing Logger in context.
type ctxLoggerKey struct{}

func WithContext(ctx context.Context, logger *slog.Logger) context.Context {
	v, _ := any(logger).(*slog.Logger)
	if v == nil {
		return context.WithValue(ctx, ctxLoggerKey{}, slog.New(NewDiscardHandler()))
	}
	return context.WithValue(ctx, ctxLoggerKey{}, logger)

}

// FromContext returns the Logger stored in ctx (using [WithContext]).
// If ctx has no logger, it returns a logger with [DiscardHandler]
// (which ignores everything written to it). It is responsibility of the
// caller to ensure that logger has a suitable handler. [IsDiscard] can be
// used to check if given handler is [DiscardHandler].
//
//	logger := log.FromContext(ctx)
//	if log.IsDiscard(logger) {
//	    panic("logger's Handler MUST NOT be log.DiscardHandler")
//	}
func FromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(ctxLoggerKey{}).(*slog.Logger); ok {
		return logger
	}
	return slog.New(NewDiscardHandler())
}
