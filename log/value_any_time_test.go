package log

import (
	"reflect"
	"testing"
	"time"
)

func TestToValueTime(t *testing.T) {
	type testCase struct {
		name   string
		input  time.Time
		expect Value
	}
	tsc, _ := time.Parse(time.RFC3339, time.StampNano)

	tt := []testCase{
		{
			name:  "<time.Time>-UTC",
			input: tsc.UTC(),
			expect: Value{
				k:   KindTime,
				s:   "UTC",
				num: uint64(tsc.UnixNano()),
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

// func TestToValueTimePtr(t *testing.T) {
// 	type testCase struct {
// 		name   string
// 		input  *time.Time
// 		expect Value
// 	}

// 	tt := []testCase{
// 		{
// 			name: "<time.TimePtr>-nil-value",
// 			expect: Value{
// 				k: KindNull,
// 			},
// 		},
// 		{
// 			name: "<time.TimePtr>-positive-value",
// 			input: func() *time.Time {
// 				i := new(time.Time)
// 				*i = time.Second
// 				return i
// 			}(),
// 			expect: Value{
// 				k:   KindTime,
// 				num: 1000000000,
// 			},
// 		},
// 		{
// 			name: "<time.Time>-negative-value",
// 			input: func() *time.Time {
// 				i := new(time.Time)
// 				*i = -time.Second
// 				return i
// 			}(),
// 			expect: Value{
// 				k:   KindTime,
// 				num: 0xffffffffc4653600,
// 			},
// 		},
// 		{
// 			name: "<time.TimePtr>-max-value",
// 			input: func() *time.Time {
// 				i := new(time.Time)
// 				*i = time.Time(math.MaxInt64)
// 				return i
// 			}(),
// 			expect: Value{
// 				k:   KindTime,
// 				num: math.MaxInt,
// 			},
// 		},
// 		{
// 			name: "<time.Time>-min-value",
// 			input: func() *time.Time {
// 				i := new(time.Time)
// 				*i = math.MinInt64
// 				return i
// 			}(),
// 			expect: Value{
// 				k:   KindTime,
// 				num: 0x8000000000000000,
// 			},
// 		},
// 	}
// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			actual := ToValue(tc.input)
// 			if !reflect.DeepEqual(tc.expect, actual) {
// 				t.Errorf("%s => \n(expected) => %#v \n(got) => %#v", tc.name, tc.expect, actual)
// 			}
// 		})
// 	}
// }
