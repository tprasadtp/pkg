package slogjson_test

import (
	"bytes"
	"context"
	"fmt"
	"runtime"
	"testing"

	"go.opencensus.io/trace"

	"github.com/tprasadtp/pkg/ref/coder/slog"
	"github.com/tprasadtp/pkg/ref/coder/slog/internal/assert"
	"github.com/tprasadtp/pkg/ref/coder/slog/internal/entryjson"
	"github.com/tprasadtp/pkg/ref/coder/slog/sloggers/slogjson"
)

var _, slogjsonTestFile, _, _ = runtime.Caller(0)

var bg = context.Background()

func TestMake(t *testing.T) {
	t.Parallel()

	ctx, s := trace.StartSpan(bg, "meow")
	b := &bytes.Buffer{}
	l := slog.Make(slogjson.Sink(b))
	l = l.Named("named")
	l.Error(ctx, "line1\n\nline2", slog.F("wowow", "me\nyou"))

	j := entryjson.Filter(b.String(), "ts")
	exp := fmt.Sprintf(`{"level":"ERROR","msg":"line1\n\nline2","caller":"%v:29","func":"github.com/tprasadtp/pkg/ref/coder/slog/sloggers/slogjson_test.TestMake","logger_names":["named"],"trace":"%v","span":"%v","fields":{"wowow":"me\nyou"}}
`, slogjsonTestFile, s.SpanContext().TraceID, s.SpanContext().SpanID)
	assert.Equal(t, "entry", exp, j)
}
