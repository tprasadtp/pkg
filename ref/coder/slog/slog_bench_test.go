package slog_test

import (
	"testing"

	"github.com/tprasadtp/pkg/ref/coder/slog"
)

var logger = slog.Make(&fakeSink{})

func BenchmarkCoderSlog(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		l2 := logger.With(
			slog.F("a", "value"),
			slog.F("b", "value"),
			slog.F("b", "value"),
			slog.F("b", "value"),
			slog.F("b", "value"),
			slog.F("b", "value"),
			slog.F("b", "value"),
			slog.F("b", "value"),
			slog.F("b", "value"),
			slog.F("b", "value"),
			slog.F("b", "value"),
		)
		l2.With(
			slog.F("a", "value"),
			slog.F("b", "value"),
			slog.F("b", "value"),
			slog.F("b", "value"),
			slog.F("b", "value"),
			slog.F("b", "value"),
			slog.F("b", "value"),
			slog.F("b", "value"),
			slog.F("b", "value"),
			slog.F("b", "value"),
			slog.F("b", "value"),
		)
	}
}
