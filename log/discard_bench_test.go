package log

import (
	"testing"
)

func BenchmarkMini(b *testing.B) {
	logger := New(NewNoOpHandler(TraceLevel))
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		// l2 := logger.With(
		// 	M("map-01", F("nested-field-01", "value-01")))
		l2 := logger.With(F("nested-field-01", "value-01"))
		l2.Info("INFO L2")
	}
}

func BenchmarkDiscardDisabledLevel(b *testing.B) {
	b.ReportAllocs()
	logger := New(NewNoOpHandler(ErrorLevel))
	for n := 0; n < b.N; n++ {
		l2 := logger.WithNamespace("namespace-01").WithError(ErrHandlerClosed)
		l2.With(
			F("root-key-01", "root-value-01"),
			F("root-key-02", "root-value-02"),
			F("root-key-03", "root-value-03"),
			M("map-01", F("map-01-key-01", "map-01-value-01")),
			M("map-02", F("map-02-key-01", "map-02-value-01")),
			M("map-03", F("map-03-key-01", "map-03-value-01")),
			M("map-04",
				F("map-04-key-01", "map-04-value-01"),
				F("map-04-key-02", "map-04-value-02"),
				F("map-04-key-03", "map-04-value-03"),
				F("map-04-key-04", "map-04-value-04"),
				F("map-04-key-05", "map-04-value-05"),
				F("map-04-key-06", "map-04-value-06"),
				F("map-04-key-07", "map-04-value-07"),
				F("map-04-key-08", "map-04-value-08"),
				F("map-04-key-09", "map-04-value-09"),
				F("map-04-key-10", "map-04-value-10"),
				F("map-04-key-11", "map-04-value-11"),
				F("map-04-key-12", "map-04-value-12")),
		).Info("INFO L2")
	}
}

func BenchmarkDiscardEnabled(b *testing.B) {
	b.ReportAllocs()
	logger := New(NewNoOpHandler(TraceLevel))
	for n := 0; n < b.N; n++ {
		l2 := logger.WithNamespace("namespace-01").WithError(ErrHandlerClosed)
		l2.With(
			F("root-key-01", "root-value-01"),
			F("root-key-02", "root-value-02"),
			F("root-key-03", "root-value-03"),
			F("root-key-04", "root-value-04"),
			F("root-key-05", "root-value-05"),
			F("root-key-06", "root-value-06"),
			F("root-key-07", "root-value-07"),
			F("root-key-08", "root-value-08"),
			F("root-key-09", "root-value-09"),
			F("root-key-10", "root-value-10"),
			F("root-key-01", "root-value-01"),
			F("root-key-02", "root-value-02"),
			F("root-key-03", "root-value-03"),
			F("root-key-04", "root-value-04"),
			F("root-key-05", "root-value-05"),
			F("root-key-06", "root-value-06"),
			F("root-key-07", "root-value-07"),
			F("root-key-08", "root-value-08"),
			F("root-key-09", "root-value-09"),
			F("root-key-10", "root-value-10"),
			F("root-key-01", "root-value-01"),
			F("root-key-02", "root-value-02"),
			F("root-key-03", "root-value-03"),
			F("root-key-04", "root-value-04"),
			F("root-key-05", "root-value-05"),
			F("root-key-06", "root-value-06"),
			F("root-key-07", "root-value-07"),
			F("root-key-08", "root-value-08"),
			F("root-key-09", "root-value-09"),
			F("root-key-10", "root-value-10"),
			M("map-01", F("map-01-key-01", "map-01-value-01")),
			M("map-02", F("map-02-key-01", "map-02-value-01")),
			M("map-03", F("map-03-key-01", "map-03-value-01")),
			M("map-04",
				F("map-04-key-01", "map-04-value-01"),
				F("map-04-key-02", "map-04-value-02"),
				F("map-04-key-03", "map-04-value-03"),
				F("map-04-key-04", "map-04-value-04"),
				F("map-04-key-05", "map-04-value-05"),
				F("map-04-key-06", "map-04-value-06"),
				F("map-04-key-07", "map-04-value-07"),
				F("map-04-key-08", "map-04-value-08"),
				F("map-04-key-09", "map-04-value-09"),
				F("map-04-key-10", "map-04-value-10"),
				F("map-04-key-11", "map-04-value-11"),
				F("map-04-key-12", "map-04-value-12")),
		).Info("INFO L2")
	}
}

// func BenchmarkDiscardEnabledF(b *testing.B) {
// 	b.ReportAllocs()
// 	logger := New(discard.New(TraceLevel))
// 	for n := 0; n < b.N; n++ {
// 		l2 := logger.WithNamespace("namespace-01").WithError(ErrHandlerClosed)
// 		l2.With(
// 			F("root-key-01", "root-value-01"),
// 			F("root-key-02", "root-value-02"),
// 			F("root-key-03", "root-value-03"),
// 			M("map-01", F("map-01-key-01", "map-01-value-01")),
// 			M("map-02", F("map-02-key-01", "map-02-value-01")),
// 			M("map-03", F("map-03-key-01", "map-03-value-01")),
// 			M("map-04",
// 				F("map-04-key-01", "map-04-value-01"),
// 				F("map-04-key-02", "map-04-value-02"),
// 				F("map-04-key-03", "map-04-value-03"),
// 				F("map-04-key-04", "map-04-value-04"),
// 				F("map-04-key-05", "map-04-value-05"),
// 				F("map-04-key-06", "map-04-value-06"),
// 				F("map-04-key-07", "map-04-value-07"),
// 				F("map-04-key-08", "map-04-value-08"),
// 				F("map-04-key-09", "map-04-value-09"),
// 				F("map-04-key-10", "map-04-value-10"),
// 				F("map-04-key-11", "map-04-value-11"),
// 				F("map-04-key-12", "map-04-value-12")),
// 		).Logf(InfoLevel, "INFO L2 %d", 1)
// 	}
// }
