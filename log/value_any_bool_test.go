package log

import (
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
			got := ToValue(tc.input)

			if got.any != nil {
				t.Errorf("Value.any non nil for Bool(%t)", tc.input)
			}
			if got.k != KindBool {
				t.Errorf("Value.kind expected=%s got=%s", tc.expect.k.String(), got.k.String())
			}
			if got.num != tc.expect.num {
				t.Errorf("Value.num expect=%d got=%d", tc.expect.num, got.num)
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
			got := ToValue(tc.input)
			if got.any != nil {
				t.Errorf("Value.any not nil for bool(%t)", *tc.input)
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
				if got.k != KindBool {
					t.Errorf("Value.k expected=%s got=%s", tc.expect.k.String(), got.k.String())
				}
			}
		})
	}
}
