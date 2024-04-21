// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: MIT

package testutils

import (
	"bytes"
	"io"
	"testing"
)

var _ io.WriteCloser = (*outputLogger)(nil)

// NewOutputLogger is used for redirecting i/o streams to t.Log in tests.
// Words are split on new lines. This does not handle windows style newlines.
func NewOutputLogger(t *testing.T, prefix string) io.WriteCloser {
	return &outputLogger{
		t:      t,
		prefix: prefix,
		buf:    make([]byte, 0, 1024),
	}
}

// Writes to t.Log when new lines are found.
type outputLogger struct {
	t      *testing.T
	buf    []byte
	prefix string
}

func (l *outputLogger) write(b []byte) {
	if len(b) == 0 {
		return
	}
	l.t.Helper()
	l.buf = append(l.buf, b...)
	var n int
	for {
		n = bytes.IndexByte(l.buf, '\n')
		if n < 0 {
			break
		}
		l.t.Logf("(%s) %s", l.prefix, l.buf[:n])
		if n+1 > len(l.buf) {
			l.buf = l.buf[0:]
		} else {
			l.buf = l.buf[n+1:]
		}
	}
}

func (l *outputLogger) Write(b []byte) (int, error) {
	l.t.Helper()
	l.write(b)
	return len(b), nil
}

func (l *outputLogger) Close() error {
	if len(l.buf) > 0 {
		l.t.Helper()
		l.t.Logf("(%s) %s", l.prefix, l.buf)
		l.buf = l.buf[0:]
	}
	return nil
}
