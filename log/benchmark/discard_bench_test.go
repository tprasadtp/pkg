package log_test

import (
	"testing"

	"github.com/tprasadtp/pkg/log"
	"github.com/tprasadtp/pkg/log/handlers/discard"
)

func BenchmarkDiscardHandler(b *testing.B) {
	b.ReportAllocs()
	logger := log.New(discard.New(log.CriticalLevel))
	logger.NoCallerTracing = true
	for n := 0; n < b.N; n++ {
		logger.Critical("IGN DEBUG")
	}
}
