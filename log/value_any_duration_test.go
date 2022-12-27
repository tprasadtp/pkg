package log

import (
	"math"
	"reflect"
	"testing"
	"time"
)

func TestToValueDuration(t *testing.T) {
	type testCase struct {
		name   string
		input  time.Duration
		expect Value
	}

	tt := []testCase{
		{
			name: "<time.Duration>-zero-value",
			expect: Value{
				k: KindDuration,
				x: 0,
			},
		},
		{
			name:  "<time.Duration>-positive-value",
			input: time.Second,
			expect: Value{
				k: KindDuration,
				x: 1000000000,
			},
		},
		{
			name:  "<time.Duration>-negative-value",
			input: -time.Second,
			expect: Value{
				k: KindDuration,
				x: 0xffffffffc4653600,
			},
		},
		{
			name:  "<time.Duration>-max-value",
			input: time.Duration(math.MaxInt64),
			expect: Value{
				k: KindDuration,
				x: math.MaxInt,
			},
		},
		{
			name:  "<time.Duration>-min-value",
			input: time.Duration(math.MinInt64),
			expect: Value{
				k: KindDuration,
				x: 0x8000000000000000,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := ToValue(tc.input)
			if !reflect.DeepEqual(tc.expect, actual) {
				t.Errorf("%s => \n(expected) => %#v \n(got) => %#v", tc.name, tc.expect, actual)
			}
		})
	}
}

func TestToValueDurationPtr(t *testing.T) {
	type testCase struct {
		name   string
		input  *time.Duration
		expect Value
	}

	tt := []testCase{
		{
			name: "<time.DurationPtr>-nil-value",
			expect: Value{
				k: KindNull,
			},
		},
		{
			name: "<time.DurationPtr>-positive-value",
			input: func() *time.Duration {
				i := new(time.Duration)
				*i = time.Second
				return i
			}(),
			expect: Value{
				k: KindDuration,
				x: 1000000000,
			},
		},
		{
			name: "<time.Duration>-negative-value",
			input: func() *time.Duration {
				i := new(time.Duration)
				*i = -time.Second
				return i
			}(),
			expect: Value{
				k: KindDuration,
				x: 0xffffffffc4653600,
			},
		},
		{
			name: "<time.DurationPtr>-max-value",
			input: func() *time.Duration {
				i := new(time.Duration)
				*i = time.Duration(math.MaxInt64)
				return i
			}(),
			expect: Value{
				k: KindDuration,
				x: math.MaxInt,
			},
		},
		{
			name: "<time.Duration>-min-value",
			input: func() *time.Duration {
				i := new(time.Duration)
				*i = math.MinInt64
				return i
			}(),
			expect: Value{
				k: KindDuration,
				x: 0x8000000000000000,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := ToValue(tc.input)
			if !reflect.DeepEqual(tc.expect, actual) {
				t.Errorf("%s => \n(expected) => %#v \n(got) => %#v", tc.name, tc.expect, actual)
			}
		})
	}
}
