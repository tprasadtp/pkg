package log

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestAnyValueTime(t *testing.T) {
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
				k: KindTime,
				s: "UTC",
				x: uint64(tsc.UnixNano()),
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := AnyValue(tc.input)
			if !reflect.DeepEqual(tc.expect, actual) {
				t.Errorf("%s => \n(expected) => %#v \n(got) => %#v", tc.name, tc.expect, actual)
			}
		})
		t.Run(fmt.Sprintf("%s-<allocs>", tc.name), func(t *testing.T) {
			allocs := testing.AllocsPerRun(10, func() {
				_ = F("key", tc.input)
			})
			if allocs != 0 {
				t.Errorf("(expected-allocs)0 != (actual-allocs)%f", allocs)
			}
		})
	}
}

func TestAnyValueTimePtr(t *testing.T) {
	type testCase struct {
		name   string
		input  *time.Time
		expect Value
	}
	tsc, _ := time.Parse(time.RFC3339, time.StampNano)

	tt := []testCase{
		{
			name:  "<*time.Time>-Nil",
			input: nil,
			expect: Value{
				k: KindNull,
			},
		},
		{
			name: "<*time.Time>-UTC",
			input: func() *time.Time {
				i := new(time.Time)
				*i = tsc.UTC()
				return i
			}(),
			expect: Value{
				k: KindTime,
				s: "UTC",
				x: uint64(tsc.UnixNano()),
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := AnyValue(tc.input)
			if !reflect.DeepEqual(tc.expect, actual) {
				t.Errorf("%s => \n(expected) => %#v \n(got) => %#v", tc.name, tc.expect, actual)
			}
		})
		t.Run(fmt.Sprintf("%s-<allocs>", tc.name), func(t *testing.T) {
			allocs := testing.AllocsPerRun(10, func() {
				_ = F("key", tc.input)
			})
			if allocs != 0 {
				t.Errorf("(expected-allocs)0 != (actual-allocs)%f", allocs)
			}
		})
	}
}