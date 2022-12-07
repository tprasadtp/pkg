package scripts_test

import (
	"testing"

	"github.com/tprasadtp/pkg/scripts"
)

func BenchmarkPC(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		scripts.GetPC(1)
	}
}
