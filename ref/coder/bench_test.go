package coder

import (
	"context"
	"io"
	"os"
	"testing"
	"time"

	"github.com/tprasadtp/pkg/ref/coder/slog"
	"github.com/tprasadtp/pkg/ref/coder/slog/sloggers/sloghuman"
	"golang.org/x/xerrors"
)

func BenchmarkXxx(b *testing.B) {
	log := slog.Make(sloghuman.Sink(os.Stdout))
	for n := 0; n < b.N; n++ {
		log.Info(context.Background(), "my message here",
			slog.F("field_name", "something or the other"),
			slog.F("some_map", slog.M(
				slog.F("nested_fields", time.Date(2000, time.February, 5, 4, 4, 4, 0, time.UTC)),
			)),
			slog.Error(
				xerrors.Errorf("wrap1: %w",
					xerrors.Errorf("wrap2: %w",
						io.EOF,
					),
				),
			),
		)
	}
}
