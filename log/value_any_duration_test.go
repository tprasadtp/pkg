package log

import (
	"math"
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
				k:   KindDuration,
				num: 0,
			},
		},
		{
			name:  "<time.Duration>-positive-value",
			input: time.Second,
			expect: Value{
				k:   KindDuration,
				num: 1000000000,
			},
		},
		{
			name:  "<time.Duration>-negative-value",
			input: -time.Second,
			expect: Value{
				k:   KindDuration,
				num: 0xffffffffc4653600,
			},
		},
		{
			name:  "<time.Duration>-max-value",
			input: time.Duration(math.MaxInt64),
			expect: Value{
				k:   KindDuration,
				num: math.MaxInt,
			},
		},
		{
			name:  "<time.Duration>-min-value",
			input: time.Duration(math.MinInt64),
			expect: Value{
				k:   KindDuration,
				num: 0x8000000000000000,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)

			if got.any != nil {
				t.Errorf("Value.any non nil for time.Duration(%d)", tc.input)
			}
			if got.k != KindDuration {
				t.Errorf("Value.kind expected=%s got=%s", tc.expect.k.String(), got.k.String())
			}
			if got.num != tc.expect.num {
				t.Errorf("Value.num expect=%d got=%d", tc.expect.num, got.num)
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
				*i = 10
				return i
			}(),
			expect: Value{
				k:   KindUint64,
				num: 10,
			},
		},
		{
			name: "<time.DurationPtr>-max-value",
			input: func() *time.Duration {
				i := new(time.Duration)
				*i = math.MaxInt64
				return i
			}(),
			expect: Value{
				k:   KindUint64,
				num: math.MaxUint,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)
			if got.any != nil {
				t.Errorf("Value.any not nil for time.Duration(%d)", tc.input)
			}

			if tc.input == nil {
				if got.k != KindNull {
					t.Errorf("Value.k expected=%s got=%s", tc.expect.k.String(), got.k.String())
				}
				if got.num != 0 {
					t.Errorf("Value.num expect=0 got=%d", got.num)
				}
			} else {
				if got.num != tc.expect.num {
					t.Errorf("Value.num expect=%d got=%d", tc.expect.num, got.num)
				}
				if got.k != KindDuration {
					t.Errorf("Value.k expected=%s got=%s", tc.expect.k.String(), got.k.String())
				}
			}
		})
	}
}
