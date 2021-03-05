// Package elena is deprecated Use see zlog instead (Internal package)
package elena

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/mgutz/ansi"
	"github.com/sirupsen/logrus"
)

func getColorByLevel(level logrus.Level) string {
	switch level {
	case logrus.TraceLevel:
		return ansi.LightGreen
	case logrus.DebugLevel:
		return ansi.LightGreen
	case logrus.InfoLevel:
		return ansi.LightCyan
	case logrus.WarnLevel:
		return ansi.LightYellow
	case logrus.ErrorLevel:
		return ansi.LightMagenta
	case logrus.FatalLevel, logrus.PanicLevel:
		return ansi.LightRed
	default:
		return ansi.LightBlue
	}
}

const resetColor = ansi.Reset

// Formatter - logrus formatter, implements logrus.Formatter
type Formatter struct {
	// Enable colors - Disable colors
	EnableColors bool
}

// Format an log entry
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	levelColor := getColorByLevel(entry.Level)
	// output buffer
	b := &bytes.Buffer{}
	if !f.EnableColors {
		fmt.Fprintf(b, "[%-7s] %-44s", strings.ToUpper(entry.Level.String()), strings.TrimSpace(entry.Message))
	} else {
		fmt.Fprintf(b, "%s[•]%s %-44s", levelColor, resetColor, strings.TrimSpace(entry.Message))
	}

	f.writeCaller(b, entry)

	f.writeFields(b, entry, levelColor)

	b.WriteByte('\n')

	return b.Bytes(), nil
}

func (f *Formatter) writeCaller(b *bytes.Buffer, entry *logrus.Entry) {
	if entry.HasCaller() {
		fmt.Fprintf(
			b,
			"(%s:%d) ",
			entry.Caller.Function,
			entry.Caller.Line,
		)
	}
}

func (f *Formatter) writeFields(b *bytes.Buffer, entry *logrus.Entry, levelColor string) {
	if len(entry.Data) != 0 {
		fields := make([]string, 0, len(entry.Data))
		for field := range entry.Data {
			fields = append(fields, field)
		}

		sort.Strings(fields)

		for _, field := range fields {
			f.writeField(b, entry, field, levelColor)
		}
	}
}

func (f *Formatter) writeField(b *bytes.Buffer, entry *logrus.Entry, field, levelColor string) {
	if !f.EnableColors {
		fmt.Fprintf(b, " %v=%v", field, entry.Data[field])
	} else {
		fmt.Fprintf(b, " %s%v%s=%v", levelColor, field, resetColor, entry.Data[field])
	}
}
