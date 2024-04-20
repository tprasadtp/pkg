// SPDX-FileCopyrightText: Copyright 2023 Prasad Tengse
// SPDX-License-Identifier: MIT

package log

import (
	"context"
	"io"
	"log/slog"
	"reflect"
	"testing"
)

func TestContext(t *testing.T) {
	t.Run("ctx-has-no-logger", func(t *testing.T) {
		ctx := context.Background()
		logger := FromContext(ctx)
		discard := NewDiscardHandler()
		handler := logger.Handler()
		if !reflect.DeepEqual(handler, discard) {
			t.Errorf("invalid logger from context expected: %#v, got: %#v", NewDiscardHandler(), handler)
		}
	})

	t.Run("ctx-has-logger", func(t *testing.T) {
		originalLogger := slog.New(slog.NewTextHandler(io.Discard, nil))
		ctx := WithContext(context.Background(), originalLogger)
		ctxLogger := FromContext(ctx)
		if !reflect.DeepEqual(originalLogger, ctxLogger) {
			t.Errorf("invalid logger from context")
			t.Errorf("expected: %#v", originalLogger)
			t.Errorf("got     : %#v", ctxLogger)
		}
	})
	t.Run("ctx-has-slog-logger", func(t *testing.T) {
		handler := slog.NewTextHandler(io.Discard, nil)
		originalLogger := slog.New(handler)
		ctx := WithContext(context.Background(), originalLogger)
		ctxLogger := FromContext(ctx)
		if _, ok := ctxLogger.Handler().(DiscardHandler); ok {
			t.Errorf("expected non discard logger")
		}
		if !reflect.DeepEqual(handler, ctxLogger.Handler()) {
			t.Errorf("context logger has different handler")
		}
	})
}
