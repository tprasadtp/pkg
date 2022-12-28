package log

import (
	"context"
	"io"
	"testing"
)

func TestLoggerNamespace(t *testing.T) {
	type testCase struct {
		name      string
		input     Logger
		namespace string
		expect    string
	}
	tt := []testCase{
		{
			name: "no-existing-namespace-with-empty-input",
		},
		{
			name:      "no-existing-namespace-with-some-string",
			namespace: "some-value",
			expect:    "some-value",
		},
		{
			name:      "no-existing-namespace-with-dotted-string",
			namespace: "space.atomic.rockets",
			expect:    "space.atomic.rockets",
		},
		{
			name:      "no-existing-namespace-with-space-string",
			namespace: "some value with spaces which is a bad idea",
			expect:    "some value with spaces which is a bad idea",
		},
		// Existing namespace
		{
			name:      "with-existing-namespace-with-empty-input",
			namespace: "",
			input: Logger{
				namespace: "service.consul",
			},
			expect: "service.consul",
		},
		{
			name:      "with-existing-namespace-with-some-string",
			namespace: "cynthia",
			input: Logger{
				namespace: "lan.service",
			},
			expect: "lan.service.cynthia",
		},
		{
			name:      "with-existing-namespace-with-dotted-string",
			namespace: "cynthia.gateway",
			input: Logger{
				namespace: "lan.service",
			},
			expect: "lan.service.cynthia.gateway",
		},
		{
			name: "with-existing-namespace-with-space-string",
			input: Logger{
				namespace: "lan.service",
			},
			namespace: "foo bar",
			expect:    "lan.service.foo bar",
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.input.WithNamespace(tc.namespace).Namespace()
			if tc.expect != actual {
				t.Errorf("(expected-namespace)%s, != (actual-namespace)%s", tc.expect, actual)
			}
		})
	}
}

func TestLoggerError(t *testing.T) {
	type testCase struct {
		name   string
		logger Logger
		input  error
		expect error
	}
	tt := []testCase{
		{
			name:   "no-existing-error-with-nil",
			input:  nil,
			expect: nil,
		},
		{
			name:   "no-existing-err-with-some-error",
			input:  io.EOF,
			expect: io.EOF,
		},
		{
			name: "existing-err-with-nil",
			logger: Logger{
				err: io.EOF,
			},
			input:  nil,
			expect: nil,
		},
		{
			name: "existing-err-with-some-error",
			logger: Logger{
				err: io.EOF,
			},
			input:  ErrLoggerInvalid,
			expect: ErrLoggerInvalid,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.logger.WithErr(tc.input).Err()
			if tc.expect != actual {
				t.Errorf("(expected-err)%s, != (actual-err)%s", tc.expect, actual)
			}
		})
	}
}

func TestLoggerCtx(t *testing.T) {
	type testCase struct {
		name   string
		logger Logger
		input  context.Context
		expect any
	}
	tt := []testCase{
		{
			name:   "no-existing-err-ctx",
			input:  context.WithValue(context.Background(), "key", "value-1"),
			expect: "value-1",
		},
		{
			name: "existing-ctx",
			logger: Logger{
				ctx: context.WithValue(context.Background(), "key", "value-1"),
			},
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
	}
}

func TestLoggerWithCaller(t *testing.T) {
	type testCase struct {
		name   string
		logger Logger
		expect bool
	}
	tt := []testCase{
		{
			name:   "not-enabled-<zero-value>",
			expect: true,
		},
		{
			name:   "already-enabled",
			logger: Logger{}.WithCaller(),
			expect: true,
		},
		{
			name:   "already-disabled",
			logger: Logger{}.WithoutCaller(),
			expect: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.logger.WithCaller().Caller()
			if tc.expect != actual {
				t.Errorf("(expected-caller)%t, != (actual-caller)%t", tc.expect, actual)
			}
		})
	}
}

func TestLoggerWithoutCaller(t *testing.T) {
	type testCase struct {
		name   string
		logger Logger
		expect bool
	}
	tt := []testCase{
		{
			name:   "not-already-enabled",
			expect: false,
		},
		{
			name:   "already-enabled",
			logger: Logger{}.WithCaller(),
			expect: false,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.logger.WithoutCaller().Caller()
			if tc.expect != actual {
				t.Errorf("(expected-caller)%t, != (actual-caller)%t", tc.expect, actual)
			}
		})
	}
}
