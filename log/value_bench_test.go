package log

import (
	"math"
	"testing"
)

func BenchmarkValue(b *testing.B) {
	var i64 int64 = math.MaxInt
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ToValue("string")
		ToValue(64)
		ToValue(math.MaxInt64)
		ToValue(&i64)
		ToValue(true)
		ToValue(false)
	}
}

var stringSlice = []string{"a", "b", "c"}
var s = "STRING"

func BenchmarkF(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		F("testing.go.key-01", &s)
		F("testing.go.key-02", uint8(1))
		F("testing.go.key-02", math.MaxInt64)
		F("testing.go.key-02", math.MaxFloat64)
		FN("testing.go.key-01", &s, "A")
		FN("testing.go.key-02", uint8(1), "A")
		FN("testing.go.key-02", math.MaxInt64, "A")
		FN("testing.go.key-02", math.MaxFloat64, "A")
	}
}
