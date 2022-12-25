package log

import (
	"math"
	"testing"
)

func TestToValueUint(t *testing.T) {
	type testCase struct {
		name   string
		input  uint
		expect Value
	}

	tt := []testCase{
		{
			name: "<uint>-zero-value",
			expect: Value{
				k:   KindUint64,
				num: 0,
			},
		},
		{
			name:  "<uint>-positive-value",
			input: 10,
			expect: Value{
				k:   KindUint64,
				num: 10,
			},
		},
		{
			name:  "<uint>-max-value",
			input: math.MaxUint,
			expect: Value{
				k:   KindUint64,
				num: math.MaxUint,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)

			if got.any != nil {
				t.Errorf("Value.any non nil for Uint(%d)", tc.input)
			}
			if got.k != KindUint64 {
				t.Errorf("Value.kind expected=%s got=%s", tc.expect.k.String(), got.k.String())
			}
			if got.num != tc.expect.num {
				t.Errorf("Value.num expect=%d got=%d", tc.expect.num, got.num)
			}
		})
	}
}

func TestToValueUintPtr(t *testing.T) {
	type testCase struct {
		name   string
		input  *uint
		expect Value
	}

	tt := []testCase{
		{
			name: "<uintptr>-nil-value",
			expect: Value{
				k: KindNull,
			},
		},
		{
			name: "<uintptr>-positive-value",
			input: func() *uint {
				i := new(uint)
				*i = 10
				return i
			}(),
			expect: Value{
				k:   KindUint64,
				num: 10,
			},
		},
		{
			name: "<uintptr>-max-value",
			input: func() *uint {
				i := new(uint)
				*i = math.MaxUint
				return i
			}(),
			expect: Value{
				k:   KindUint64,
				num: math.MaxUint,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)
			if got.any != nil {
				t.Errorf("Value.any not nil for uinteger(%d)", tc.input)
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
				if got.k != KindUint64 {
					t.Errorf("Value.k expected=%s got=%s", tc.expect.k.String(), got.k.String())
				}
			}
		})
	}
}

func TestToValueUint8(t *testing.T) {
	type testCase struct {
		name   string
		input  uint8
		expect Value
	}

	tt := []testCase{
		{
			name: "<uint8>-zero-value",
			expect: Value{
				k:   KindUint64,
				num: 0,
			},
		},
		{
			name:  "<uint8>-positive-value",
			input: 10,
			expect: Value{
				k:   KindUint64,
				num: 10,
			},
		},
		{
			name:  "<uint8>-max-value",
			input: math.MaxUint8,
			expect: Value{
				k:   KindUint64,
				num: math.MaxUint8,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)

			if got.any != nil {
				t.Errorf("Value.any non nil for Uint(%d)", tc.input)
			}
			if got.k != KindUint64 {
				t.Errorf("Value.kind expected=%s got=%s", tc.expect.k.String(), got.k.String())
			}
			if got.num != tc.expect.num {
				t.Errorf("Value.num expect=%d got=%d", tc.expect.num, got.num)
			}
		})
	}
}

func TestToValueUint8Ptr(t *testing.T) {
	type testCase struct {
		name   string
		input  *uint8
		expect Value
	}

	tt := []testCase{
		{
			name: "<uint8ptr>-nil-value",
			expect: Value{
				k: KindNull,
			},
		},
		{
			name: "<uint8ptr>-positive-value",
			input: func() *uint8 {
				i := new(uint8)
				*i = 10
				return i
			}(),
			expect: Value{
				k:   KindUint64,
				num: 10,
			},
		},
		{
			name: "<uint8ptr>-max-value",
			input: func() *uint8 {
				i := new(uint8)
				*i = math.MaxUint8
				return i
			}(),
			expect: Value{
				k:   KindUint64,
				num: math.MaxUint8,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)
			if got.any != nil {
				t.Errorf("Value.any not nil for uinteger(%d)", tc.input)
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
				if got.k != KindUint64 {
					t.Errorf("Value.k expected=%s got=%s", tc.expect.k.String(), got.k.String())
				}
			}
		})
	}
}

func TestToValueUint16(t *testing.T) {
	type testCase struct {
		name   string
		input  uint16
		expect Value
	}

	tt := []testCase{
		{
			name: "<uint16>-zero-value",
			expect: Value{
				k:   KindUint64,
				num: 0,
			},
		},
		{
			name:  "<uint16>-positive-value",
			input: 10,
			expect: Value{
				k:   KindUint64,
				num: 10,
			},
		},
		{
			name:  "<uint16>-max-value",
			input: math.MaxUint16,
			expect: Value{
				k:   KindUint64,
				num: math.MaxUint16,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)

			if got.any != nil {
				t.Errorf("Value.any non nil for Uint(%d)", tc.input)
			}
			if got.k != KindUint64 {
				t.Errorf("Value.kind expected=%s got=%s", tc.expect.k.String(), got.k.String())
			}
			if got.num != tc.expect.num {
				t.Errorf("Value.num expect=%d got=%d", tc.expect.num, got.num)
			}
		})
	}
}

func TestToValueUint16Ptr(t *testing.T) {
	type testCase struct {
		name   string
		input  *uint16
		expect Value
	}

	tt := []testCase{
		{
			name: "<uint16ptr>-nil-value",
			expect: Value{
				k: KindNull,
			},
		},
		{
			name: "<uint16ptr>-positive-value",
			input: func() *uint16 {
				i := new(uint16)
				*i = 10
				return i
			}(),
			expect: Value{
				k:   KindUint64,
				num: 10,
			},
		},
		{
			name: "<uint16ptr>-max-value",
			input: func() *uint16 {
				i := new(uint16)
				*i = math.MaxUint16
				return i
			}(),
			expect: Value{
				k:   KindUint64,
				num: math.MaxUint16,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)
			if got.any != nil {
				t.Errorf("Value.any not nil for uinteger(%d)", tc.input)
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
				if got.k != KindUint64 {
					t.Errorf("Value.k expected=%s got=%s", tc.expect.k.String(), got.k.String())
				}
			}
		})
	}
}

func TestToValueUint32(t *testing.T) {
	type testCase struct {
		name   string
		input  uint32
		expect Value
	}

	tt := []testCase{
		{
			name: "<uint32>-zero-value",
			expect: Value{
				k:   KindUint64,
				num: 0,
			},
		},
		{
			name:  "<uint32>-positive-value",
			input: 10,
			expect: Value{
				k:   KindUint64,
				num: 10,
			},
		},
		{
			name:  "<uint32>-max-value",
			input: math.MaxUint32,
			expect: Value{
				k:   KindUint64,
				num: math.MaxUint32,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)

			if got.any != nil {
				t.Errorf("Value.any non nil for Uint(%d)", tc.input)
			}
			if got.k != KindUint64 {
				t.Errorf("Value.kind expected=%s got=%s", tc.expect.k.String(), got.k.String())
			}
			if got.num != tc.expect.num {
				t.Errorf("Value.num expect=%d got=%d", tc.expect.num, got.num)
			}
		})
	}
}

func TestToValueUint32Ptr(t *testing.T) {
	type testCase struct {
		name   string
		input  *uint32
		expect Value
	}

	tt := []testCase{
		{
			name: "<uint32ptr>-nil-value",
			expect: Value{
				k: KindNull,
			},
		},
		{
			name: "<uint32ptr>-positive-value",
			input: func() *uint32 {
				i := new(uint32)
				*i = 10
				return i
			}(),
			expect: Value{
				k:   KindUint64,
				num: 10,
			},
		},
		{
			name: "<uint32ptr>-max-value",
			input: func() *uint32 {
				i := new(uint32)
				*i = math.MaxUint32
				return i
			}(),
			expect: Value{
				k:   KindUint64,
				num: math.MaxUint32,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)
			if got.any != nil {
				t.Errorf("Value.any not nil for uinteger(%d)", tc.input)
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
				if got.k != KindUint64 {
					t.Errorf("Value.k expected=%s got=%s", tc.expect.k.String(), got.k.String())
				}
			}
		})
	}
}

func TestToValueUint64(t *testing.T) {
	type testCase struct {
		name   string
		input  uint64
		expect Value
	}

	tt := []testCase{
		{
			name: "<uint64>-zero-value",
			expect: Value{
				k:   KindUint64,
				num: 0,
			},
		},
		{
			name:  "<uint64>-positive-value",
			input: 10,
			expect: Value{
				k:   KindUint64,
				num: 10,
			},
		},
		{
			name:  "<uint64>-max-value",
			input: math.MaxUint64,
			expect: Value{
				k:   KindUint64,
				num: math.MaxUint64,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)

			if got.any != nil {
				t.Errorf("Value.any non nil for Uint(%d)", tc.input)
			}
			if got.k != KindUint64 {
				t.Errorf("Value.kind expected=%s got=%s", tc.expect.k.String(), got.k.String())
			}
			if got.num != tc.expect.num {
				t.Errorf("Value.num expect=%d got=%d", tc.expect.num, got.num)
			}
		})
	}
}

func TestToValueUint64Ptr(t *testing.T) {
	type testCase struct {
		name   string
		input  *uint64
		expect Value
	}

	tt := []testCase{
		{
			name: "<uint64ptr>-nil-value",
			expect: Value{
				k: KindNull,
			},
		},
		{
			name: "<uint64ptr>-positive-value",
			input: func() *uint64 {
				i := new(uint64)
				*i = 10
				return i
			}(),
			expect: Value{
				k:   KindUint64,
				num: 10,
			},
		},
		{
			name: "<uint64ptr>-max-value",
			input: func() *uint64 {
				i := new(uint64)
				*i = math.MaxUint64
				return i
			}(),
			expect: Value{
				k:   KindUint64,
				num: math.MaxUint64,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ToValue(tc.input)
			if got.any != nil {
				t.Errorf("Value.any not nil for uinteger(%d)", tc.input)
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
				if got.k != KindUint64 {
					t.Errorf("Value.k expected=%s got=%s", tc.expect.k.String(), got.k.String())
				}
			}
		})
	}
}
