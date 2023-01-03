package log

import (
	"net/netip"
	"testing"
	"time"
)

var (
	tUint     uint    = 1
	tUint8    uint8   = 8
	tUint16   uint16  = 16
	tUint32   uint32  = 32
	tUint64   uint64  = 64
	tInt      int     = 1
	tInt8     int8    = -8
	tInt16    int16   = -16
	tInt32    int32   = -32
	tInt64    int64   = -64
	tFloat32  float32 = 1.32
	tFloat64  float64 = 1.64
	tString   string  = "a string with spaces"
	tTime, _          = time.Parse(time.RFC3339, time.StampNano)
	tDuration         = time.Second
	tIP               = netip.MustParseAddr("192.0.2.1")
	tPrefix           = netip.MustParsePrefix("192.0.2.0/24")
)

func BenchmarkNumbers(b *testing.B) {
	logger := New(NewDiscardHandler(LevelTrace)).WithoutCaller()
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		l2 := logger.With(
			// F("uint", tUint),
			// F("uint8", tUint8),
			// F("uint16", tUint16),
			// F("uint32", tUint32),
			// F("uint64", uint64(1)),
			// F("int", tInt),
			// F("int8", tInt8),
			// F("int16", tInt16),
			// F("int32", tInt32),
			F("int64", 0),
		)
		l2.Info("INFO L2")
	}
}
