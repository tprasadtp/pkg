package log

import (
	"reflect"
	"testing"
)

func TestToValueBool(t *testing.T) {
	type testCase struct {
		name   string
		input  bool
		expect Value
	}

	tt := []testCase{
		{
			name: "zero-value",
			expect: Value{
				k: KindBool,
				x: 0,
			},
		},
		{
			name:  "true-value",
			input: true,
			expect: Value{
				k: KindBool,
				x: 1,
			},
		},
		{
			name:  "false-value",
			input: false,
			expect: Value{
				k: KindBool,
				x: 0,
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

func TestToValueBoolPtr(t *testing.T) {
	type testCase struct {
		name   string
		input  *bool
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
			name: "true-value",
			input: func() *bool {
				i := new(bool)
				*i = true
				return i
			}(),
			expect: Value{
				k: KindBool,
				x: 1,
			},
		},
		{
			name: "false-value",
			input: func() *bool {
				i := new(bool)
				*i = false
				return i
			}(),
			expect: Value{
				k: KindBool,
				x: 0,
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
