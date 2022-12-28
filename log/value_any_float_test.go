package log

import (
	"fmt"
	"math"
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

func TestAnyValueFloat32(t *testing.T) {
	type testCase struct {
		name   string
		input  float32
		expect Value
	}

	tt := []testCase{
		{
			name: "zero-value",
			expect: Value{
				k: KindFloat64,
			},
		},
		{
			name:  "positive-value",
			input: 10.0,
			expect: Value{
				k: KindFloat64,
				x: 0x4024000000000000,
			},
		},
		{
			name:  "negative-value",
			input: -10.0,
			expect: Value{
				k: KindFloat64,
				x: 0xc024000000000000,
			},
		},
		{
			name:  "max-value",
			input: math.MaxFloat32,
			expect: Value{
				k: KindFloat64,
				x: 0x47efffffe0000000,
			},
		},
		{
			name:  "positive-inf-value",
			input: float32(math.Inf(1)),
			expect: Value{
				k: KindFloat64,
				x: 0x7ff0000000000000,
			},
		},
		{
			name:  "negative-inf-value",
			input: float32(math.Inf(-1)),
			expect: Value{
				k: KindFloat64,
				x: 0xfff0000000000000,
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

func TestAnyValueFloat32Ptr(t *testing.T) {
	type testCase struct {
		name   string
		input  *float32
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
			name: "positive-value",
			input: func() *float32 {
				i := new(float32)
				*i = 10.0
				return i
			}(),
			expect: Value{
				k: KindFloat64,
				x: 0x4024000000000000,
			},
		},
		{
			name: "negative-value",
			input: func() *float32 {
				i := new(float32)
				*i = -10.0
				return i
			}(),
			expect: Value{
				k: KindFloat64,
				x: 0xc024000000000000,
			},
		},
		// Inf
		{
			name: "positive-inf-value",
			input: func() *float32 {
				i := new(float32)
				*i = float32(math.Inf(1))
				return i
			}(),
			expect: Value{
				k: KindFloat64,
				x: 0x7ff0000000000000,
			},
		},
		{
			name: "negative-inf-value",
			input: func() *float32 {
				i := new(float32)
				*i = float32(math.Inf(-1))
				return i
			}(),
			expect: Value{
				k: KindFloat64,
				x: 0xfff0000000000000,
			},
		},
		// Max
		{
			name: "max-value",
			input: func() *float32 {
				i := new(float32)
				*i = math.MaxFloat32
				return i
			}(),
			expect: Value{
				k: KindFloat64,
				x: 0x47efffffe0000000,
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

func TestAnyValueFloat64(t *testing.T) {
	type testCase struct {
		name   string
		input  float64
		expect Value
	}

	tt := []testCase{
		{
			name: "zero-value",
			expect: Value{
				k: KindFloat64,
			},
		},
		{
			name:  "positive-value",
			input: 10.0,
			expect: Value{
				k: KindFloat64,
				x: 0x4024000000000000,
			},
		},
		{
			name:  "negative-value",
			input: -10.0,
			expect: Value{
				k: KindFloat64,
				x: 0xc024000000000000,
			},
		},
		{
			name:  "max-value",
			input: math.MaxFloat64,
			expect: Value{
				k: KindFloat64,
				x: 0x7fefffffffffffff,
			},
		},
		{
			name:  "positive-inf-value",
			input: math.Inf(1),
			expect: Value{
				k: KindFloat64,
				x: 0x7ff0000000000000,
			},
		},
		{
			name:  "negative-inf-value",
			input: math.Inf(-1),
			expect: Value{
				k: KindFloat64,
				x: 0xfff0000000000000,
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

func TestAnyValueFloat64Ptr(t *testing.T) {
	type testCase struct {
		name   string
		input  *float64
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
			name: "positive-value",
			input: func() *float64 {
				i := new(float64)
				*i = 10.0
				return i
			}(),
			expect: Value{
				k: KindFloat64,
				x: 0x4024000000000000,
			},
		},
		{
			name: "negative-value",
			input: func() *float64 {
				i := new(float64)
				*i = -10.0
				return i
			}(),
			expect: Value{
				k: KindFloat64,
				x: 0xc024000000000000,
			},
		},
		// Inf
		{
			name: "positive-inf-value",
			input: func() *float64 {
				i := new(float64)
				*i = math.Inf(1)
				return i
			}(),
			expect: Value{
				k: KindFloat64,
				x: 0x7ff0000000000000,
			},
		},
		{
			name: "negative-inf-value",
			input: func() *float64 {
				i := new(float64)
				*i = math.Inf(-1)
				return i
			}(),
			expect: Value{
				k: KindFloat64,
				x: 0xfff0000000000000,
			},
		},
		// Max
		{
			name: "max-value",
			input: func() *float64 {
				i := new(float64)
				*i = math.MaxFloat64
				return i
			}(),
			expect: Value{
				k: KindFloat64,
				x: 0x7fefffffffffffff,
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
