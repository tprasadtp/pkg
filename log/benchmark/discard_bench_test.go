package log_test

import (
	"context"
	"testing"

	"github.com/tprasadtp/pkg/log"
	"github.com/tprasadtp/pkg/log/handlers/discard"
)

func BenchmarkDiscardDisabledLevel(b *testing.B) {
	b.ReportAllocs()
	logger := log.New(discard.New(log.ErrorLevel))
	for n := 0; n < b.N; n++ {
		l2 := logger.WithNamespace("namespace-01").WithError(log.ErrHandlerClosed).WithCtx(context.Background())
		l2.With(
			log.F("root-key-01", "root-value-01"),
			log.F("root-key-02", "root-value-02"),
			log.F("root-key-03", "root-value-03"),
			log.M("map-01", log.F("map-01-key-01", "map-01-value-01")),
			log.M("map-02", log.F("map-02-key-01", "map-02-value-01")),
			log.M("map-03", log.F("map-03-key-01", "map-03-value-01")),
			log.M("map-04",
				log.F("map-04-key-01", "map-04-value-01"),
				log.F("map-04-key-02", "map-04-value-02"),
				log.F("map-04-key-03", "map-04-value-03"),
				log.F("map-04-key-04", "map-04-value-04"),
				log.F("map-04-key-05", "map-04-value-05"),
				log.F("map-04-key-06", "map-04-value-06"),
				log.F("map-04-key-07", "map-04-value-07"),
				log.F("map-04-key-08", "map-04-value-08"),
				log.F("map-04-key-09", "map-04-value-09"),
				log.F("map-04-key-10", "map-04-value-10"),
				log.F("map-04-key-11", "map-04-value-11"),
				log.F("map-04-key-12", "map-04-value-12")),
		).Info("INFO L2")
	}
}

func BenchmarkDiscardEnabled(b *testing.B) {
	b.ReportAllocs()
	logger := log.New(discard.New(log.TraceLevel))
	for n := 0; n < b.N; n++ {
		l2 := logger.WithNamespace("namespace-01").WithError(log.ErrHandlerClosed).WithCtx(context.Background())
		l2.With(
			log.F("root-key-01", "root-value-01"),
			log.F("root-key-02", "root-value-02"),
			log.F("root-key-03", "root-value-03"),
			log.M("map-01", log.F("map-01-key-01", "map-01-value-01")),
			log.M("map-02", log.F("map-02-key-01", "map-02-value-01")),
			log.M("map-03", log.F("map-03-key-01", "map-03-value-01")),
			log.M("map-04",
				log.F("map-04-key-01", "map-04-value-01"),
				log.F("map-04-key-02", "map-04-value-02"),
				log.F("map-04-key-03", "map-04-value-03"),
				log.F("map-04-key-04", "map-04-value-04"),
				log.F("map-04-key-05", "map-04-value-05"),
				log.F("map-04-key-06", "map-04-value-06"),
				log.F("map-04-key-07", "map-04-value-07"),
				log.F("map-04-key-08", "map-04-value-08"),
				log.F("map-04-key-09", "map-04-value-09"),
				log.F("map-04-key-10", "map-04-value-10"),
				log.F("map-04-key-11", "map-04-value-11"),
				log.F("map-04-key-12", "map-04-value-12")),
		).Info("INFO L2")
	}
}

func BenchmarkDiscardEnabledF(b *testing.B) {
	b.ReportAllocs()
	logger := log.New(discard.New(log.TraceLevel))
	for n := 0; n < b.N; n++ {
		l2 := logger.WithNamespace("namespace-01").WithError(log.ErrHandlerClosed).WithCtx(context.Background())
		l2.With(
			log.F("root-key-01", "root-value-01"),
			log.F("root-key-02", "root-value-02"),
			log.F("root-key-03", "root-value-03"),
			log.M("map-01", log.F("map-01-key-01", "map-01-value-01")),
			log.M("map-02", log.F("map-02-key-01", "map-02-value-01")),
			log.M("map-03", log.F("map-03-key-01", "map-03-value-01")),
			log.M("map-04",
				log.F("map-04-key-01", "map-04-value-01"),
				log.F("map-04-key-02", "map-04-value-02"),
				log.F("map-04-key-03", "map-04-value-03"),
				log.F("map-04-key-04", "map-04-value-04"),
				log.F("map-04-key-05", "map-04-value-05"),
				log.F("map-04-key-06", "map-04-value-06"),
				log.F("map-04-key-07", "map-04-value-07"),
				log.F("map-04-key-08", "map-04-value-08"),
				log.F("map-04-key-09", "map-04-value-09"),
				log.F("map-04-key-10", "map-04-value-10"),
				log.F("map-04-key-11", "map-04-value-11"),
				log.F("map-04-key-12", "map-04-value-12")),
		).Logf(log.InfoLevel, "INFO L2 %d", 1)
	}
}
