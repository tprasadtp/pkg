package log

import (
	"net/netip"
	"reflect"
	"testing"
)

func TestAnyValueNetIPAddr(t *testing.T) {
	type testCase struct {
		name   string
		input  netip.Addr
		expect Value
	}

	tt := []testCase{
		{
			name: "zero-value",
			expect: Value{
				k: KindIPAddr,
				s: "invalid IP",
			},
		},
		{
			name:  "some-value",
			input: netip.MustParseAddr("192.0.2.1"),
			expect: Value{
				k: KindIPAddr,
				s: "192.0.2.1",
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

func TestAnyValueNetIPAddrPtr(t *testing.T) {
	type testCase struct {
		name   string
		input  *netip.Addr
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
			input: func() *netip.Addr {
				i := new(netip.Addr)
				*i = netip.MustParseAddr("192.0.2.1")
				return i
			}(),
			expect: Value{
				k: KindIPAddr,
				s: "192.0.2.1",
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

func TestAnyValueNetPrefix(t *testing.T) {
	type testCase struct {
		name   string
		input  netip.Prefix
		expect Value
	}

	tt := []testCase{
		{
			name: "zero-value",
			expect: Value{
				k: KindIPPrefix,
				s: "invalid Prefix",
			},
		},
		{
			name:  "some-value",
			input: netip.MustParsePrefix("192.0.2.0/24"),
			expect: Value{
				k: KindIPPrefix,
				s: "192.0.2.0/24",
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

func TestAnyValueNetIPPrefixPtr(t *testing.T) {
	type testCase struct {
		name   string
		input  *netip.Prefix
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
			input: func() *netip.Prefix {
				i := new(netip.Prefix)
				*i = netip.MustParsePrefix("192.0.2.0/24")
				return i
			}(),
			expect: Value{
				k: KindIPPrefix,
				s: "192.0.2.0/24",
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
