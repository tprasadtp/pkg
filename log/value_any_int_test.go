package log

import (
	"math"
	"reflect"
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
			name: "zero-value",
			expect: Value{
				k: KindInt64,
				x: 0,
			},
		},
		{
			name:  "positive-value",
			input: 10,
			expect: Value{
				k: KindInt64,
				x: 10,
			},
		},
		{
			name:  "negative-value",
			input: -10,
			expect: Value{
				k: KindInt64,
				x: 0xfffffffffffffff6,
			},
		},
		{
			name:  "max-value",
			input: math.MaxInt,
			expect: Value{
				k: KindInt64,
				x: math.MaxInt,
			},
		},
		{
			name:  "min-value",
			input: math.MinInt,
			expect: Value{
				k: KindInt64,
				x: 0x8000000000000000,
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

func TestToValueIntPtr(t *testing.T) {
	type testCase struct {
		name   string
		input  *int
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
			input: func() *int {
				i := new(int)
				*i = 10
				return i
			}(),
			expect: Value{
				k: KindInt64,
				x: 10,
			},
		},
		{
			name: "negative-value",
			input: func() *int {
				i := new(int)
				*i = -10
				return i
			}(),
			expect: Value{
				k: KindInt64,
				x: 0xfffffffffffffff6,
			},
		},
		{
			name: "max-value",
			input: func() *int {
				i := new(int)
				*i = math.MaxInt
				return i
			}(),
			expect: Value{
				k: KindInt64,
				x: math.MaxInt,
			},
		},
		{
			name: "min-value",
			input: func() *int {
				i := new(int)
				*i = math.MinInt
				return i
			}(),
			expect: Value{
				k: KindInt64,
				x: 0x8000000000000000,
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

func TestToValueInt8(t *testing.T) {
	type testCase struct {
		name   string
		input  int8
		expect Value
	}

	tt := []testCase{
		{
			name: "zero-value",
			expect: Value{
				k: KindInt64,
				x: 0,
			},
		},
		{
			name:  "positive-value",
			input: 10,
			expect: Value{
				k: KindInt64,
				x: 10,
			},
		},
		{
			name:  "negative-value",
			input: -10,
			expect: Value{
				k: KindInt64,
				x: 0xfffffffffffffff6,
			},
		},
		{
			name:  "max-value",
			input: math.MaxInt8,
			expect: Value{
				k: KindInt64,
				x: math.MaxInt8,
			},
		},
		{
			name:  "min-value",
			input: math.MinInt8,
			expect: Value{
				k: KindInt64,
				x: 0xffffffffffffff80,
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

func TestToValueInt8Ptr(t *testing.T) {
	type testCase struct {
		name   string
		input  *int8
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
			input: func() *int8 {
				i := new(int8)
				*i = 10
				return i
			}(),
			expect: Value{
				k: KindInt64,
				x: 10,
			},
		},
		{
			name: "negative-value",
			input: func() *int8 {
				i := new(int8)
				*i = -10
				return i
			}(),
			expect: Value{
				k: KindInt64,
				x: 0xfffffffffffffff6,
			},
		},
		{
			name: "max-value",
			input: func() *int8 {
				i := new(int8)
				*i = math.MaxInt8
				return i
			}(),
			expect: Value{
				k: KindInt64,
				x: math.MaxInt8,
			},
		},
		{
			name: "min-value",
			input: func() *int8 {
				i := new(int8)
				*i = math.MinInt8
				return i
			}(),
			expect: Value{
				k: KindInt64,
				x: 0xffffffffffffff80,
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

func TestToValueInt16(t *testing.T) {
	type testCase struct {
		name   string
		input  int16
		expect Value
	}

	tt := []testCase{
		{
			name: "zero-value",
			expect: Value{
				k: KindInt64,
				x: 0,
			},
		},
		{
			name:  "positive-value",
			input: 10,
			expect: Value{
				k: KindInt64,
				x: 10,
			},
		},
		{
			name:  "negative-value",
			input: -10,
			expect: Value{
				k: KindInt64,
				x: 0xfffffffffffffff6,
			},
		},
		{
			name:  "max-value",
			input: math.MaxInt16,
			expect: Value{
				k: KindInt64,
				x: math.MaxInt16,
			},
		},
		{
			name:  "min-value",
			input: math.MinInt16,
			expect: Value{
				k: KindInt64,
				x: 0xffffffffffff8000,
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

func TestToValueInt16Ptr(t *testing.T) {
	type testCase struct {
		name   string
		input  *int16
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
			input: func() *int16 {
				i := new(int16)
				*i = 10
				return i
			}(),
			expect: Value{
				k: KindInt64,
				x: 10,
			},
		},
		{
			name: "negative-value",
			input: func() *int16 {
				i := new(int16)
				*i = -10
				return i
			}(),
			expect: Value{
				k: KindInt64,
				x: 0xfffffffffffffff6,
			},
		},
		{
			name: "max-value",
			input: func() *int16 {
				i := new(int16)
				*i = math.MaxInt16
				return i
			}(),
			expect: Value{
				k: KindInt64,
				x: math.MaxInt16,
			},
		},
		{
			name: "min-value",
			input: func() *int16 {
				i := new(int16)
				*i = math.MinInt16
				return i
			}(),
			expect: Value{
				k: KindInt64,
				x: 0xffffffffffff8000,
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

func TestToValueInt32(t *testing.T) {
	type testCase struct {
		name   string
		input  int32
		expect Value
	}

	tt := []testCase{
		{
			name: "zero-value",
			expect: Value{
				k: KindInt64,
				x: 0,
			},
		},
		{
			name:  "positive-value",
			input: 10,
			expect: Value{
				k: KindInt64,
				x: 10,
			},
		},
		{
			name:  "negative-value",
			input: -10,
			expect: Value{
				k: KindInt64,
				x: 0xfffffffffffffff6,
			},
		},
		{
			name:  "max-value",
			input: math.MaxInt32,
			expect: Value{
				k: KindInt64,
				x: math.MaxInt32,
			},
		},
		{
			name:  "min-value",
			input: math.MinInt32,
			expect: Value{
				k: KindInt64,
				x: 0xffffffff80000000,
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

func TestToValueInt32Ptr(t *testing.T) {
	type testCase struct {
		name   string
		input  *int32
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
			input: func() *int32 {
				i := new(int32)
				*i = 10
				return i
			}(),
			expect: Value{
				k: KindInt64,
				x: 10,
			},
		},
		{
			name: "negative-value",
			input: func() *int32 {
				i := new(int32)
				*i = -10
				return i
			}(),
			expect: Value{
				k: KindInt64,
				x: 0xfffffffffffffff6,
			},
		},
		{
			name: "max-value",
			input: func() *int32 {
				i := new(int32)
				*i = math.MaxInt32
				return i
			}(),
			expect: Value{
				k: KindInt64,
				x: math.MaxInt32,
			},
		},
		{
			name: "min-value",
			input: func() *int32 {
				i := new(int32)
				*i = math.MinInt32
				return i
			}(),
			expect: Value{
				k: KindInt64,
				x: 0xffffffff80000000,
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

func TestToValueInt64(t *testing.T) {
	type testCase struct {
		name   string
		input  int64
		expect Value
	}

	tt := []testCase{
		{
			name: "zero-value",
			expect: Value{
				k: KindInt64,
				x: 0,
			},
		},
		{
			name:  "positive-value",
			input: 10,
			expect: Value{
				k: KindInt64,
				x: 10,
			},
		},
		{
			name:  "negative-value",
			input: -10,
			expect: Value{
				k: KindInt64,
				x: 0xfffffffffffffff6,
			},
		},
		{
			name:  "max-value",
			input: math.MaxInt64,
			expect: Value{
				k: KindInt64,
				x: math.MaxInt64,
			},
		},
		{
			name:  "min-value",
			input: math.MinInt64,
			expect: Value{
				k: KindInt64,
				x: 0x8000000000000000,
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

func TestToValueInt64Ptr(t *testing.T) {
	type testCase struct {
		name   string
		input  *int64
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
			input: func() *int64 {
				i := new(int64)
				*i = 10
				return i
			}(),
			expect: Value{
				k: KindInt64,
				x: 10,
			},
		},
		{
			name: "negative-value",
			input: func() *int64 {
				i := new(int64)
				*i = -10
				return i
			}(),
			expect: Value{
				k: KindInt64,
				x: 0xfffffffffffffff6,
			},
		},
		{
			name: "max-value",
			input: func() *int64 {
				i := new(int64)
				*i = math.MaxInt64
				return i
			}(),
			expect: Value{
				k: KindInt64,
				x: math.MaxInt64,
			},
		},
		{
			name: "min-value",
			input: func() *int64 {
				i := new(int64)
				*i = math.MinInt64
				return i
			}(),
			expect: Value{
				k: KindInt64,
				x: 0x8000000000000000,
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
