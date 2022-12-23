package log

import (
	"context"
	"math"
	"testing"
	"time"
)

func BenchmarkMini(b *testing.B) {
	logger := New(NewDiscardHandler(LevelTrace))
	t := time.Duration(1)
	// cpx := complex(1.0, 0.5)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		l2 := logger.With(
			F("string", "value"),
			F("uint", math.MaxUint/2),
			F("uint8", math.MaxUint8),
			F("uint16", math.MaxUint16),
			F("uint32", math.MaxUint32),
			F("uint64", math.MaxUint64/2),

			F("int", math.MaxInt),
			F("int8", math.MaxInt8),
			F("int16", math.MaxInt16),
			F("int32", math.MaxInt32),
			F("int64", math.MaxInt64),

			F("bool", false),

			F("float32", math.MaxFloat32),
			F("float64", math.MaxFloat64),
			F("time.Time", t),
		)
		l2.WithCtx(context.Background()).WithErr(ErrInvalidKind).Info("INFO L2")
	}
}
