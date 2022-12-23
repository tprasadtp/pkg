package log

import (
	"math"
	"testing"
)

// max   float32 bits = 0x47efffffe0000000
// max   float64 bits = 0x7fefffffffffffff
// -inf  float64 bits = 0xfff0000000000000
// +inf  float64 bits = 0x7ff0000000000000
// -10.0 float64 bits = 0xc024000000000000
// +10.0 float64 bits = 0x4024000000000000

func TestToValueFloat32(t *testing.T) {
	type testCase struct {
		name   string
		input  float32
		expect Value
	}

	tt := []testCase{
		{
			name: "<float32>-zero-value",
			expect: Value{
				k: KindFloat64,
			},
		},
		{
			name:  "<float32>-positive-value",
			input: 10.0,
			expect: Value{
				k:   KindFloat64,
				num: 0x4024000000000000,
			},
		},
		{
			name:  "<float32>-negative-value",
			input: -10.0,
			expect: Value{
				k:   KindFloat64,
				num: 0xc024000000000000,
			},
		},
		{
			name:  "<float32>-max-value",
			input: math.MaxFloat32,
			expect: Value{
				k:   KindFloat64,
				num: 0x47efffffe0000000,
			},
		},
		{
			name:  "<float32>-positive-inf-value",
			input: float32(math.Inf(1)),
			expect: Value{
				k:   KindFloat64,
				num: 0x7ff0000000000000,
			},
		},
		{
			name:  "<float32>-negative-inf-value",
			input: float32(math.Inf(-1)),
			expect: Value{
				k:   KindFloat64,
				num: 0xfff0000000000000,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)

			if got.any != nil {
				t.Errorf("Value.any non nil for Float32(%f)", tc.input)
			}
			if got.k != KindFloat64 {
				t.Errorf("Value.kind expected=%s got=%s", tc.expect.k.String(), got.k.String())
			}
			if got.num != tc.expect.num {
				t.Errorf("Value.num expect=%#v got=%#v", tc.expect.num, got.num)
			}
		})
	}
}

func TestToValueFloat32Ptr(t *testing.T) {
	type testCase struct {
		name   string
		input  *float32
		expect Value
	}

	tt := []testCase{
		{
			name: "<float32ptr>-nil-value",
			expect: Value{
				k: KindNull,
			},
		},
		{
			name: "<float32ptr>-positive-value",
			input: func() *float32 {
				i := new(float32)
				*i = 10.0
				return i
			}(),
			expect: Value{
				k:   KindFloat64,
				num: 0x4024000000000000,
			},
		},
		{
			name: "<float32ptr>-negative-value",
			input: func() *float32 {
				i := new(float32)
				*i = -10.0
				return i
			}(),
			expect: Value{
				k:   KindFloat64,
				num: 0xc024000000000000,
			},
		},
		// Inf
		{
			name: "<float32ptr>-positive-inf-value",
			input: func() *float32 {
				i := new(float32)
				*i = float32(math.Inf(1))
				return i
			}(),
			expect: Value{
				k:   KindFloat64,
				num: 0x7ff0000000000000,
			},
		},
		{
			name: "<float32ptr>-negative-inf-value",
			input: func() *float32 {
				i := new(float32)
				*i = float32(math.Inf(-1))
				return i
			}(),
			expect: Value{
				k:   KindFloat64,
				num: 0xfff0000000000000,
			},
		},
		// Max
		{
			name: "<float32ptr>-max-value",
			input: func() *float32 {
				i := new(float32)
				*i = math.MaxFloat32
				return i
			}(),
			expect: Value{
				k:   KindFloat64,
				num: 0x47efffffe0000000,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)
			if got.any != nil {
				t.Errorf("Value.any not nil for Float32(%d)", tc.input)
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
					t.Errorf("Value.num expect=%#v got=%#v", tc.expect.num, got.num)
				}
				if got.k != KindFloat64 {
					t.Errorf("Value.k expected=%s got=%s", tc.expect.k.String(), got.k.String())
				}
			}
		})
	}
}

func TestToValueFloat64(t *testing.T) {
	type testCase struct {
		name   string
		input  float64
		expect Value
	}

	tt := []testCase{
		{
			name: "<float64>-zero-value",
			expect: Value{
				k: KindFloat64,
			},
		},
		{
			name:  "<float64>-positive-value",
			input: 10.0,
			expect: Value{
				k:   KindFloat64,
				num: 0x4024000000000000,
			},
		},
		{
			name:  "<float64>-negative-value",
			input: -10.0,
			expect: Value{
				k:   KindFloat64,
				num: 0xc024000000000000,
			},
		},
		{
			name:  "<float64>-max-value",
			input: math.MaxFloat64,
			expect: Value{
				k:   KindFloat64,
				num: 0x7fefffffffffffff,
			},
		},
		{
			name:  "<float64>-positive-inf-value",
			input: math.Inf(1),
			expect: Value{
				k:   KindFloat64,
				num: 0x7ff0000000000000,
			},
		},
		{
			name:  "<float64>-negative-inf-value",
			input: math.Inf(-1),
			expect: Value{
				k:   KindFloat64,
				num: 0xfff0000000000000,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)

			if got.any != nil {
				t.Errorf("Value.any non nil for Float64(%f)", tc.input)
			}
			if got.k != KindFloat64 {
				t.Errorf("Value.kind expected=%s got=%s", tc.expect.k.String(), got.k.String())
			}
			if got.num != tc.expect.num {
				t.Errorf("Value.num expect=%#v got=%#v", tc.expect.num, got.num)
			}
		})
	}
}

func TestToValueFloat64Ptr(t *testing.T) {
	type testCase struct {
		name   string
		input  *float64
		expect Value
	}

	tt := []testCase{
		{
			name: "<float64ptr>-nil-value",
			expect: Value{
				k: KindNull,
			},
		},
		{
			name: "<float64ptr>-positive-value",
			input: func() *float64 {
				i := new(float64)
				*i = 10.0
				return i
			}(),
			expect: Value{
				k:   KindFloat64,
				num: 0x4024000000000000,
			},
		},
		{
			name: "<float64ptr>-negative-value",
			input: func() *float64 {
				i := new(float64)
				*i = -10.0
				return i
			}(),
			expect: Value{
				k:   KindFloat64,
				num: 0xc024000000000000,
			},
		},
		// Inf
		{
			name: "<float64ptr>-positive-inf-value",
			input: func() *float64 {
				i := new(float64)
				*i = math.Inf(1)
				return i
			}(),
			expect: Value{
				k:   KindFloat64,
				num: 0x7ff0000000000000,
			},
		},
		{
			name: "<float64ptr>-negative-inf-value",
			input: func() *float64 {
				i := new(float64)
				*i = math.Inf(-1)
				return i
			}(),
			expect: Value{
				k:   KindFloat64,
				num: 0xfff0000000000000,
			},
		},
		// Max
		{
			name: "<float64ptr>-max-value",
			input: func() *float64 {
				i := new(float64)
				*i = math.MaxFloat64
				return i
			}(),
			expect: Value{
				k:   KindFloat64,
				num: 0x7fefffffffffffff,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)
			if got.any != nil {
				t.Errorf("Value.any not nil for Float64(%d)", tc.input)
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
					t.Errorf("Value.num expect=%#v got=%#v", tc.expect.num, got.num)
				}
				if got.k != KindFloat64 {
					t.Errorf("Value.k expected=%s got=%s", tc.expect.k.String(), got.k.String())
				}
			}
		})
	}
}
