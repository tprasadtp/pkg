package log_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/tprasadtp/pkg/log"
)

func TestLoggerNamespace(t *testing.T) {
	type testCase struct {
		name      string
		logger    log.Logger
		namespace string
		expect    string
	}
	l := log.NewLogger(log.NewDiscardHandler(log.LevelTrace))
	tt := []testCase{
		{
			name:   "no-existing-namespace-with-empty-input",
			logger: l,
		},
		{
			name:      "no-existing-namespace-with-some-string",
			logger:    l,
			namespace: "some-value",
			expect:    "some-value",
		},
		{
			name:      "no-existing-namespace-with-dotted-string",
			logger:    l,
			namespace: "space.atomic.rockets",
			expect:    "space.atomic.rockets",
		},
		{
			name:      "no-existing-namespace-with-space-string",
			logger:    l,
			namespace: "some value with spaces which is a bad idea",
			expect:    "some value with spaces which is a bad idea",
		},
		// Existing namespace
		{
			name:      "with-existing-namespace-with-empty-input",
			namespace: "",
			logger:    l.WithNamespace("consul.service"),
			expect:    "consul.service",
		},
		{
			name:      "with-existing-namespace-with-some-string",
			namespace: "cynthia",
			logger:    l.WithNamespace("consul.service"),
			expect:    "consul.service.cynthia",
		},
		{
			name:      "with-existing-namespace-with-dotted-string",
			namespace: "cynthia.gateway",
			logger:    l.WithNamespace("consul.service"),
			expect:    "consul.service.cynthia.gateway",
		},
		{
			name:      "with-existing-namespace-with-space-string",
			logger:    l.WithNamespace("consul.service"),
			namespace: "foo bar",
			expect:    "consul.service.foo bar",
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.logger.WithNamespace(tc.namespace).Namespace()
			if tc.expect != actual {
				t.Errorf("(expected-namespace)%s, != (actual-namespace)%s", tc.expect, actual)
			}
		})
		t.Run(fmt.Sprintf("%s-<allocs>", tc.name), func(t *testing.T) {
			allocs := testing.AllocsPerRun(10, func() {
				_ = tc.logger.WithNamespace(tc.namespace)
			})
			if allocs != 0 {
				t.Errorf("(expected-allocs)0 != (actual-allocs)%f", allocs)
			}
		})
	}
}

func TestLoggerWithErr(t *testing.T) {
	type testCase struct {
		name   string
		logger log.Logger
		input  error
		expect error
	}
	l := log.NewLogger(log.NewDiscardHandler(log.LevelTrace))
	tt := []testCase{
		{
			name:   "no-existing-error-with-nil",
			logger: l,
			input:  nil,
			expect: nil,
		},
		{
			name:   "no-existing-err-with-some-error",
			logger: l,
			input:  io.EOF,
			expect: io.EOF,
		},
		{
			name:   "existing-err-with-nil",
			logger: l.WithErr(io.EOF),
			input:  nil,
			expect: nil,
		},
		{
			name:   "existing-err-with-some-error",
			logger: l.WithErr(io.EOF),
			input:  os.ErrClosed,
			expect: os.ErrClosed,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.logger.WithErr(tc.input).Err()
			if errors.Is(actual, tc.expect) {
				t.Errorf("(expected-err)%s, != (actual-err)%s", tc.expect, actual)
			}
		})
		t.Run(fmt.Sprintf("%s-<allocs>", tc.name), func(t *testing.T) {
			allocs := testing.AllocsPerRun(10, func() {
				_ = tc.logger.WithErr(tc.input)
			})
			if allocs != 0 {
				t.Errorf("(expected-allocs)0 != (actual-allocs)%f", allocs)
			}
		})
	}
}

func TestLoggerCtx(t *testing.T) {
	type testCase struct {
		name   string
		logger log.Logger
		input  context.Context
		expect any
	}
	l := log.NewLogger(log.NewDiscardHandler(log.LevelTrace))
	tt := []testCase{
		{
			name:   "no-existing-err-ctx",
			input:  context.WithValue(context.Background(), "key", "value-1"),
			expect: "value-1",
		},
		{
			name:   "existing-ctx",
			logger: l.WithCtx(context.WithValue(context.Background(), "key", "value-1")),
			input:  context.WithValue(context.Background(), "key", "value-2"),
			expect: "value-2",
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctx := tc.logger.WithCtx(tc.input).Ctx()
			actual := ctx.Value("key")
			if tc.expect != actual {
				t.Errorf("(expected-context-value)%s, != (actual-context-value)%s", tc.expect, actual)
			}
		})
		t.Run(fmt.Sprintf("%s-<allocs>", tc.name), func(t *testing.T) {
			allocs := testing.AllocsPerRun(10, func() {
				_ = tc.logger.WithCtx(tc.input)
			})
			if allocs != 0 {
				t.Errorf("(expected-allocs)0 != (actual-allocs)%f", allocs)
			}
		})
	}
}

func TestLoggerWithCaller(t *testing.T) {
	type testCase struct {
		name   string
		logger log.Logger
		expect bool
	}
	l := log.NewLogger(log.NewDiscardHandler(log.LevelTrace))
	tt := []testCase{
		{
			name:   "not-enabled-<zero-value>",
			logger: l,
			expect: true,
		},
		{
			name:   "already-enabled",
			logger: l.WithCaller(),
			expect: true,
		},
		{
			name:   "already-disabled",
			logger: l.WithoutCaller(),
			expect: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.logger.WithCaller().Caller()
			if tc.expect != actual {
				t.Errorf("(expected-caller)%#v != (actual-caller)%#v", tc.expect, actual)
			}
		})
		t.Run(fmt.Sprintf("%s-<allocs>", tc.name), func(t *testing.T) {
			allocs := testing.AllocsPerRun(10, func() {
				_ = tc.logger.WithCaller()
			})
			if allocs != 0 {
				t.Errorf("(expected-allocs)0 != (actual-allocs)%f", allocs)
			}
		})
	}
}

func TestLoggerWithoutCaller(t *testing.T) {
	type testCase struct {
		name   string
		logger log.Logger
		expect bool
	}
	l := log.NewLogger(log.NewDiscardHandler(log.LevelTrace))
	tt := []testCase{
		{
			name:   "not-already-enabled",
			logger: l,
			expect: false,
		},
		{
			name:   "already-enabled",
			logger: l.WithCaller(),
			expect: false,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.logger.WithoutCaller().Caller()
			if tc.expect != actual {
				t.Errorf("(expected-caller)%#v != (actual-caller)%#v", tc.expect, actual)
			}
		})
		t.Run(fmt.Sprintf("%s-<allocs>", tc.name), func(t *testing.T) {
			allocs := testing.AllocsPerRun(10, func() {
				_ = tc.logger.WithoutCaller()
			})
			if allocs != 0 {
				t.Errorf("(expected-allocs)0 != (actual-allocs)%f", allocs)
			}
		})
	}
}

func TestLoggerHandler(t *testing.T) {
	type testCase struct {
		name   string
		logger log.Logger
		expect log.Handler
	}
	h := log.NewDiscardHandler(log.LevelTrace)
	l := log.NewLogger(h)

	tt := []testCase{
		{
			name:   "nil-value",
			logger: log.Logger{},
			expect: nil,
		},
		{
			name:   "already-enabled",
			logger: l,
			expect: h,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.logger.Handler()
			if tc.expect != actual {
				t.Errorf("(expected-caller)%#v, != (actual-caller)%#v", tc.expect, actual)
			}
		})
	}
}
