package log

import (
	"math"
	"testing"
)

func TestToValueInt(t *testing.T) {
	type testCase struct {
		name   string
		input  int
		expect Value
	}

	tt := []testCase{
		{
			name: "<int>-zero-value",
			expect: Value{
				k:   KindInt64,
				num: 0,
			},
		},
		{
			name:  "<int>-positive-value",
			input: 10,
			expect: Value{
				k:   KindInt64,
				num: 10,
			},
		},
		{
			name:  "<int>-negative-value",
			input: -10,
			expect: Value{
				k:   KindInt64,
				num: 0xfffffffffffffff6,
			},
		},
		{
			name:  "<int>-max-value",
			input: math.MaxInt,
			expect: Value{
				k:   KindInt64,
				num: math.MaxInt,
			},
		},
		{
			name:  "<int>-min-value",
			input: math.MinInt,
			expect: Value{
				k:   KindInt64,
				num: 0x8000000000000000,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)

			if got.any != nil {
				t.Errorf("Value.any non nil for Int(%d)", tc.input)
			}
			if got.k != KindInt64 {
				t.Errorf("Value.kind expected=%s got=%s", tc.expect.k.String(), got.k.String())
			}
			if got.num != tc.expect.num {
				t.Errorf("Value.num expect=%d got=%d", tc.expect.num, got.num)
			}
		})
	}
}

func TestToValueIntPtr(t *testing.T) {
	type testCase struct {
		name   string
		input  *int
		expect Value
	}

	tt := []testCase{
		{
			name: "<intptr>-nil-value",
			expect: Value{
				k: KindNull,
			},
		},
		{
			name: "<intptr>-positive-value",
			input: func() *int {
				i := new(int)
				*i = 10
				return i
			}(),
			expect: Value{
				k:   KindInt64,
				num: 10,
			},
		},
		{
			name: "<intptr>-negative-value",
			input: func() *int {
				i := new(int)
				*i = -10
				return i
			}(),
			expect: Value{
				k:   KindInt64,
				num: 0xfffffffffffffff6,
			},
		},
		{
			name: "<intptr>-max-value",
			input: func() *int {
				i := new(int)
				*i = math.MaxInt
				return i
			}(),
			expect: Value{
				k:   KindInt64,
				num: math.MaxInt,
			},
		},
		{
			name: "<intptr>-min-value",
			input: func() *int {
				i := new(int)
				*i = math.MinInt
				return i
			}(),
			expect: Value{
				k:   KindInt64,
				num: 0x8000000000000000,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)
			if got.any != nil {
				t.Errorf("Value.any not nil for integer(%d)", tc.input)
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
				if got.k != KindInt64 {
					t.Errorf("Value.k expected=%s got=%s", tc.expect.k.String(), got.k.String())
				}
			}
		})
	}
}

func TestToValueInt8(t *testing.T) {
	type testCase struct {
		name   string
		input  int8
		expect Value
	}

	tt := []testCase{
		{
			name: "<int8>-zero-value",
			expect: Value{
				k:   KindInt64,
				num: 0,
			},
		},
		{
			name:  "<int8>-positive-value",
			input: 10,
			expect: Value{
				k:   KindInt64,
				num: 10,
			},
		},
		{
			name:  "<int8>-negative-value",
			input: -10,
			expect: Value{
				k:   KindInt64,
				num: 0xfffffffffffffff6,
			},
		},
		{
			name:  "<int8>-max-value",
			input: math.MaxInt8,
			expect: Value{
				k:   KindInt64,
				num: math.MaxInt8,
			},
		},
		{
			name:  "<int8>-min-value",
			input: math.MinInt8,
			expect: Value{
				k:   KindInt64,
				num: 0xffffffffffffff80,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)

			if got.any != nil {
				t.Errorf("Value.any non nil for Int(%d)", tc.input)
			}
			if got.k != KindInt64 {
				t.Errorf("Value.kind expected=%s got=%s", tc.expect.k.String(), got.k.String())
			}
			if got.num != tc.expect.num {
				t.Errorf("Value.num expect=%d got=%d", tc.expect.num, got.num)
			}
		})
	}
}

func TestToValueInt8Ptr(t *testing.T) {
	type testCase struct {
		name   string
		input  *int8
		expect Value
	}

	tt := []testCase{
		{
			name: "<int8ptr>-nil-value",
			expect: Value{
				k: KindNull,
			},
		},
		{
			name: "<int8ptr>-positive-value",
			input: func() *int8 {
				i := new(int8)
				*i = 10
				return i
			}(),
			expect: Value{
				k:   KindInt64,
				num: 10,
			},
		},
		{
			name: "<int8ptr>-negative-value",
			input: func() *int8 {
				i := new(int8)
				*i = -10
				return i
			}(),
			expect: Value{
				k:   KindInt64,
				num: 0xfffffffffffffff6,
			},
		},
		{
			name: "<int8ptr>-max-value",
			input: func() *int8 {
				i := new(int8)
				*i = math.MaxInt8
				return i
			}(),
			expect: Value{
				k:   KindInt64,
				num: math.MaxInt8,
			},
		},
		{
			name: "<int8ptr>-min-value",
			input: func() *int8 {
				i := new(int8)
				*i = math.MinInt8
				return i
			}(),
			expect: Value{
				k:   KindInt64,
				num: 0xffffffffffffff80,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)
			if got.any != nil {
				t.Errorf("Value.any not nil for integer(%d)", tc.input)
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
				if got.k != KindInt64 {
					t.Errorf("Value.k expected=%s got=%s", tc.expect.k.String(), got.k.String())
				}
			}
		})
	}
}

func TestToValueInt16(t *testing.T) {
	type testCase struct {
		name   string
		input  int16
		expect Value
	}

	tt := []testCase{
		{
			name: "<int16>-zero-value",
			expect: Value{
				k:   KindInt64,
				num: 0,
			},
		},
		{
			name:  "<int16>-positive-value",
			input: 10,
			expect: Value{
				k:   KindInt64,
				num: 10,
			},
		},
		{
			name:  "<int16>-negative-value",
			input: -10,
			expect: Value{
				k:   KindInt64,
				num: 0xfffffffffffffff6,
			},
		},
		{
			name:  "<int16>-max-value",
			input: math.MaxInt16,
			expect: Value{
				k:   KindInt64,
				num: math.MaxInt16,
			},
		},
		{
			name:  "<int16>-min-value",
			input: math.MinInt16,
			expect: Value{
				k:   KindInt64,
				num: 0xffffffffffff8000,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)

			if got.any != nil {
				t.Errorf("Value.any non nil for Int(%d)", tc.input)
			}
			if got.k != KindInt64 {
				t.Errorf("Value.kind expected=%s got=%s", tc.expect.k.String(), got.k.String())
			}
			if got.num != tc.expect.num {
				t.Errorf("Value.num expect=%d got=%d", tc.expect.num, got.num)
			}
		})
	}
}

func TestToValueInt16Ptr(t *testing.T) {
	type testCase struct {
		name   string
		input  *int16
		expect Value
	}

	tt := []testCase{
		{
			name: "<int16ptr>-nil-value",
			expect: Value{
				k: KindNull,
			},
		},
		{
			name: "<int16ptr>-positive-value",
			input: func() *int16 {
				i := new(int16)
				*i = 10
				return i
			}(),
			expect: Value{
				k:   KindInt64,
				num: 10,
			},
		},
		{
			name: "<int16ptr>-negative-value",
			input: func() *int16 {
				i := new(int16)
				*i = -10
				return i
			}(),
			expect: Value{
				k:   KindInt64,
				num: 0xfffffffffffffff6,
			},
		},
		{
			name: "<int16ptr>-max-value",
			input: func() *int16 {
				i := new(int16)
				*i = math.MaxInt16
				return i
			}(),
			expect: Value{
				k:   KindInt64,
				num: math.MaxInt16,
			},
		},
		{
			name: "<int16ptr>-min-value",
			input: func() *int16 {
				i := new(int16)
				*i = math.MinInt16
				return i
			}(),
			expect: Value{
				k:   KindInt64,
				num: 0xffffffffffff8000,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)
			if got.any != nil {
				t.Errorf("Value.any not nil for integer(%d)", tc.input)
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
				if got.k != KindInt64 {
					t.Errorf("Value.k expected=%s got=%s", tc.expect.k.String(), got.k.String())
				}
			}
		})
	}
}

func TestToValueInt32(t *testing.T) {
	type testCase struct {
		name   string
		input  int32
		expect Value
	}

	tt := []testCase{
		{
			name: "<int32>-zero-value",
			expect: Value{
				k:   KindInt64,
				num: 0,
			},
		},
		{
			name:  "<int32>-positive-value",
			input: 10,
			expect: Value{
				k:   KindInt64,
				num: 10,
			},
		},
		{
			name:  "<int32>-negative-value",
			input: -10,
			expect: Value{
				k:   KindInt64,
				num: 0xfffffffffffffff6,
			},
		},
		{
			name:  "<int32>-max-value",
			input: math.MaxInt32,
			expect: Value{
				k:   KindInt64,
				num: math.MaxInt32,
			},
		},
		{
			name:  "<int32>-min-value",
			input: math.MinInt32,
			expect: Value{
				k:   KindInt64,
				num: 0xffffffff80000000,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)

			if got.any != nil {
				t.Errorf("Value.any non nil for Int(%d)", tc.input)
			}
			if got.k != KindInt64 {
				t.Errorf("Value.kind expected=%s got=%s", tc.expect.k.String(), got.k.String())
			}
			if got.num != tc.expect.num {
				t.Errorf("Value.num expect=%d got=%d", tc.expect.num, got.num)
			}
		})
	}
}

func TestToValueInt32Ptr(t *testing.T) {
	type testCase struct {
		name   string
		input  *int32
		expect Value
	}

	tt := []testCase{
		{
			name: "<int32ptr>-nil-value",
			expect: Value{
				k: KindNull,
			},
		},
		{
			name: "<int32ptr>-positive-value",
			input: func() *int32 {
				i := new(int32)
				*i = 10
				return i
			}(),
			expect: Value{
				k:   KindInt64,
				num: 10,
			},
		},
		{
			name: "<int32ptr>-negative-value",
			input: func() *int32 {
				i := new(int32)
				*i = -10
				return i
			}(),
			expect: Value{
				k:   KindInt64,
				num: 0xfffffffffffffff6,
			},
		},
		{
			name: "<int32ptr>-max-value",
			input: func() *int32 {
				i := new(int32)
				*i = math.MaxInt32
				return i
			}(),
			expect: Value{
				k:   KindInt64,
				num: math.MaxInt32,
			},
		},
		{
			name: "<int32ptr>-min-value",
			input: func() *int32 {
				i := new(int32)
				*i = math.MinInt32
				return i
			}(),
			expect: Value{
				k:   KindInt64,
				num: 0xffffffff80000000,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)
			if got.any != nil {
				t.Errorf("Value.any not nil for integer(%d)", tc.input)
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
				if got.k != KindInt64 {
					t.Errorf("Value.k expected=%s got=%s", tc.expect.k.String(), got.k.String())
				}
			}
		})
	}
}

func TestToValueInt64(t *testing.T) {
	type testCase struct {
		name   string
		input  int64
		expect Value
	}

	tt := []testCase{
		{
			name: "<int64>-zero-value",
			expect: Value{
				k:   KindInt64,
				num: 0,
			},
		},
		{
			name:  "<int64>-positive-value",
			input: 10,
			expect: Value{
				k:   KindInt64,
				num: 10,
			},
		},
		{
			name:  "<int64>-negative-value",
			input: -10,
			expect: Value{
				k:   KindInt64,
				num: 0xfffffffffffffff6,
			},
		},
		{
			name:  "<int64>-max-value",
			input: math.MaxInt64,
			expect: Value{
				k:   KindInt64,
				num: math.MaxInt64,
			},
		},
		{
			name:  "<int64>-min-value",
			input: math.MinInt64,
			expect: Value{
				k:   KindInt64,
				num: 0x8000000000000000,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)

			if got.any != nil {
				t.Errorf("Value.any non nil for Int(%d)", tc.input)
			}
			if got.k != KindInt64 {
				t.Errorf("Value.kind expected=%s got=%s", tc.expect.k.String(), got.k.String())
			}
			if got.num != tc.expect.num {
				t.Errorf("Value.num expect=%d got=%d", tc.expect.num, got.num)
			}
		})
	}
}

func TestToValueInt64Ptr(t *testing.T) {
	type testCase struct {
		name   string
		input  *int64
		expect Value
	}

	tt := []testCase{
		{
			name: "<int64ptr>-nil-value",
			expect: Value{
				k: KindNull,
			},
		},
		{
			name: "<int64ptr>-positive-value",
			input: func() *int64 {
				i := new(int64)
				*i = 10
				return i
			}(),
			expect: Value{
				k:   KindInt64,
				num: 10,
			},
		},
		{
			name: "<int64ptr>-negative-value",
			input: func() *int64 {
				i := new(int64)
				*i = -10
				return i
			}(),
			expect: Value{
				k:   KindInt64,
				num: 0xfffffffffffffff6,
			},
		},
		{
			name: "<int64ptr>-max-value",
			input: func() *int64 {
				i := new(int64)
				*i = math.MaxInt64
				return i
			}(),
			expect: Value{
				k:   KindInt64,
				num: math.MaxInt64,
			},
		},
		{
			name: "<int64ptr>-min-value",
			input: func() *int64 {
				i := new(int64)
				*i = math.MinInt64
				return i
			}(),
			expect: Value{
				k:   KindInt64,
				num: 0x8000000000000000,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)
			if got.any != nil {
				t.Errorf("Value.any not nil for integer(%d)", tc.input)
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
				if got.k != KindInt64 {
					t.Errorf("Value.k expected=%s got=%s", tc.expect.k.String(), got.k.String())
				}
			}
		})
	}
}
