package log

import (
	"reflect"
	"testing"
)

func TestAnyValueAnyKind(t *testing.T) {
	type testCase struct {
		name   string
		input  any
		expect Value
	}

	tt := []testCase{
		{
			name:   "zero-value",
			expect: Value{},
		},
		{
			name:  "string-slice",
			input: []string{"a", "b"},
			expect: Value{
				any: []string{"a", "b"},
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
	}
}
