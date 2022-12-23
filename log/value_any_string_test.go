package log

import (
	"fmt"
	"net"
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
			name: "<string>-zero-value",
			expect: Value{
				k:   KindString,
				num: 0,
			},
		},
		{
			name:  "<string>-some-value",
			input: "a string",
			expect: Value{
				k: KindString,
				s: "a string",
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)

			if got.any != nil {
				t.Errorf("Value.any non nil for String(%s)", tc.input)
			}
			if got.num != 0 {
				t.Errorf("Value.num non 0 for String(%s)", tc.input)
			}
			if got.k != KindString {
				t.Errorf("Value.kind expected=%s got=%s", tc.expect.k.String(), got.k.String())
			}
			if got.s != tc.expect.s {
				t.Errorf("Value.num expect=%s got=%s", tc.expect.s, got.s)
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
			name: "<stringptr>-nil-value",
			expect: Value{
				k: KindNull,
			},
		},
		{
			name: "<stringptr>-some-value",
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
			got := ToValue(tc.input)
			if got.any != nil {
				t.Errorf("Value.any not nil for string(%s)", *tc.input)
			}
			if got.num != 0 {
				t.Errorf("Value.num not 0 for string(%s)", *tc.input)
			}

			if tc.input == nil {
				if got.k != KindNull {
					t.Errorf("Value.k expected=%s got=%s", tc.expect.k.String(), got.k.String())
				}
				if got.s != "" {
					t.Errorf("Value.num expect=\"\" got=%s", got.s)
				}
			} else {
				if got.s != tc.expect.s {
					t.Errorf("Value.num expect=%s got=%s", tc.expect.s, got.s)
				}
				if got.k != KindString {
					t.Errorf("Value.k expected=%s got=%s", tc.expect.k.String(), got.k.String())
				}
			}
		})
	}
}

func TestToValueStringer(t *testing.T) {
	type testCase struct {
		name   string
		input  fmt.Stringer
		expect Value
	}

	tt := []testCase{
		{
			name:  "<string>-stringer-ip",
			input: net.IPv6loopback,
			expect: Value{
				k: KindString,
				s: net.IPv6loopback.String(),
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)

			if got.any != nil {
				t.Errorf("Value.any non nil for String(%s)", tc.input)
			}
			if got.num != 0 {
				t.Errorf("Value.num non 0 for String(%s)", tc.input)
			}
			if got.k != KindString {
				t.Errorf("Value.kind expected=%s got=%s", tc.expect.k.String(), got.k.String())
			}
			if got.s != tc.expect.s {
				t.Errorf("Value.num expect=%s got=%s", tc.expect.s, got.s)
			}
		})
	}
}
