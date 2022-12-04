package slogstackdriver

import (
	logpbtype "google.golang.org/genproto/googleapis/logging/type"

	"github.com/tprasadtp/pkg/ref/coder/slog"
)

func Sev(level slog.Level) logpbtype.LogSeverity {
	return sev(level)
}
