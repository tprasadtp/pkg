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
			name: "<bool>-zero-value",
			expect: Value{
				k:   KindBool,
				num: 0,
			},
		},
		{
			name:  "<bool>-true-value",
			input: true,
			expect: Value{
				k:   KindBool,
				num: 1,
			},
		},
		{
			name:  "<bool>-false-value",
			input: false,
			expect: Value{
				k:   KindBool,
				num: 0,
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
			name: "<boolptr>-nil-value",
			expect: Value{
				k: KindNull,
			},
		},
		{
			name: "<boolptr>-true-value",
			input: func() *bool {
				i := new(bool)
				*i = true
				return i
			}(),
			expect: Value{
				k:   KindBool,
				num: 1,
			},
		},
		{
			name: "<boolptr>-false-value",
			input: func() *bool {
				i := new(bool)
				*i = false
				return i
			}(),
			expect: Value{
				k:   KindBool,
				num: 0,
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
