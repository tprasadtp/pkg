package log

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAnyValueString(t *testing.T) {
	type testCase struct {
		name   string
		input  string
		expect Value
	}

	tt := []testCase{
		{
			name: "zero-value",
			expect: Value{
				k: KindString,
				x: 0,
			},
		},
		{
			name:  "some-value",
			input: "a string",
			expect: Value{
				k: KindString,
				s: "a string",
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

func TestAnyValueStringStringPtr(t *testing.T) {
	type testCase struct {
		name   string
		input  *string
		expect Value
	}

	tt := []testCase{
		{
			name: "nil-value",
			expect: Value{
				k: KindNull,
			},
		},
		{
			name: "some-value",
			input: func() *string {
				i := new(string)
				*i = "a string"
				return i
			}(),
			expect: Value{
				k: KindString,
				s: "a string",
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
