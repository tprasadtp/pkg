package log

import (
	"reflect"
	"testing"
)

func TestToValueString(t *testing.T) {
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
			actual := ToValue(tc.input)
			if !reflect.DeepEqual(tc.expect, actual) {
				t.Errorf("%s => \n(expected) => %#v \n(got) => %#v", tc.name, tc.expect, actual)
			}
		})
	}
}

func TestToValueStringStringPtr(t *testing.T) {
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
			actual := ToValue(tc.input)
			if !reflect.DeepEqual(tc.expect, actual) {
				t.Errorf("%s => \n(expected) => %#v \n(got) => %#v", tc.name, tc.expect, actual)
			}
		})
	}
}
