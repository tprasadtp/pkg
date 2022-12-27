package log

import (
	"reflect"
	"testing"
)

// Test ref values.
//
// max   float32 bits = 0x47efffffe0000000
// max   float64 bits = 0x7fefffffffffffff
// -inf  float64 bits = 0xfff0000000000000
// +inf  float64 bits = 0x7ff0000000000000
// -10.0 float64 bits = 0xc024000000000000
// +10.0 float64 bits = 0x4024000000000000

func TestAnyValueComplex64(t *testing.T) {
	type testCase struct {
		name   string
		input  complex64
		expect Value
	}

	tt := []testCase{
		{
			name: "zero-value",
			expect: Value{
				k: KindComplex128,
				s: "(0+0i)",
			},
		},
		{
			name: "some-value",
			input: func() complex64 {
				return complex64(complex(1.1, 1.2))
			}(),
			expect: Value{
				k: KindComplex128,
				s: "(1.1+1.2i)",
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

func TestAnyValueComplex64Ptr(t *testing.T) {
	type testCase struct {
		name   string
		input  *complex64
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
			name: "zero-value",
			input: func() *complex64 {
				i := new(complex64)
				return i
			}(),
			expect: Value{
				k: KindComplex128,
				s: "(0+0i)",
			},
		},
		{
			name: "some-value",
			input: func() *complex64 {
				i := new(complex64)
				*i = complex64(complex(1.1, 1.2))
				return i
			}(),
			expect: Value{
				k: KindComplex128,
				s: "(1.1+1.2i)",
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

func TestAnyValueComplex128(t *testing.T) {
	type testCase struct {
		name   string
		input  complex128
		expect Value
	}

	tt := []testCase{
		{
			name: "zero-value",
			expect: Value{
				k: KindComplex128,
				s: "(0+0i)",
			},
		},
		{
			name: "some-value",
			input: func() complex128 {
				return complex(1.1, 1.2)
			}(),
			expect: Value{
				k: KindComplex128,
				s: "(1.1+1.2i)",
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

func TestAnyValueComplex128Ptr(t *testing.T) {
	type testCase struct {
		name   string
		input  *complex128
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
			name: "zero-value",
			input: func() *complex128 {
				i := new(complex128)
				return i
			}(),
			expect: Value{
				k: KindComplex128,
				s: "(0+0i)",
			},
		},
		{
			name: "some-value",
			input: func() *complex128 {
				i := new(complex128)
				*i = complex(1.1, 1.2)
				return i
			}(),
			expect: Value{
				k: KindComplex128,
				s: "(1.1+1.2i)",
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
