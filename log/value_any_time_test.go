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
				k: KindTime,
				s: "UTC",
				x: uint64(tsc.UnixNano()),
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

func TestToValueTimePtr(t *testing.T) {
	type testCase struct {
		name   string
		input  *time.Time
		expect Value
	}
	tsc, _ := time.Parse(time.RFC3339, time.StampNano)

	tt := []testCase{
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
			actual := ToValue(tc.input)
			if !reflect.DeepEqual(tc.expect, actual) {
				t.Errorf("%s => \n(expected) => %#v \n(got) => %#v", tc.name, tc.expect, actual)
			}
		})
	}
}
