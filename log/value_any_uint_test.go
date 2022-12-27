package log

import (
	"math"
	"reflect"
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
			name: "zero-value",
			expect: Value{
				k: KindUint64,
				x: 0,
			},
		},
		{
			name:  "positive-value",
			input: 10,
			expect: Value{
				k: KindUint64,
				x: 10,
			},
		},
		{
			name:  "max-value",
			input: math.MaxUint,
			expect: Value{
				k: KindUint64,
				x: math.MaxUint,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := ToValue(tc.input)
			if !reflect.DeepEqual(tc.expect, actual) {
				t.Errorf("%s => \n(expected) => %#v \n(got) => %#v", tc.name, tc.expect, actual)
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
			name: "nil-value",
			expect: Value{
				k: KindNull,
			},
		},
		{
			name: "positive-value",
			input: func() *uint {
				i := new(uint)
				*i = 10
				return i
			}(),
			expect: Value{
				k: KindUint64,
				x: 10,
			},
		},
		{
			name: "max-value",
			input: func() *uint {
				i := new(uint)
				*i = math.MaxUint
				return i
			}(),
			expect: Value{
				k: KindUint64,
				x: math.MaxUint,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := ToValue(tc.input)
			if !reflect.DeepEqual(tc.expect, actual) {
				t.Errorf("%s => \n(expected) => %#v \n(got) => %#v", tc.name, tc.expect, actual)
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
			name: "zero-value",
			expect: Value{
				k: KindUint64,
				x: 0,
			},
		},
		{
			name:  "positive-value",
			input: 10,
			expect: Value{
				k: KindUint64,
				x: 10,
			},
		},
		{
			name:  "max-value",
			input: math.MaxUint8,
			expect: Value{
				k: KindUint64,
				x: math.MaxUint8,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := ToValue(tc.input)
			if !reflect.DeepEqual(tc.expect, actual) {
				t.Errorf("%s => \n(expected) => %#v \n(got) => %#v", tc.name, tc.expect, actual)
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
			name: "nil-value",
			expect: Value{
				k: KindNull,
			},
		},
		{
			name: "positive-value",
			input: func() *uint8 {
				i := new(uint8)
				*i = 10
				return i
			}(),
			expect: Value{
				k: KindUint64,
				x: 10,
			},
		},
		{
			name: "max-value",
			input: func() *uint8 {
				i := new(uint8)
				*i = math.MaxUint8
				return i
			}(),
			expect: Value{
				k: KindUint64,
				x: math.MaxUint8,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := ToValue(tc.input)
			if !reflect.DeepEqual(tc.expect, actual) {
				t.Errorf("%s => \n(expected) => %#v \n(got) => %#v", tc.name, tc.expect, actual)
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
			name: "zero-value",
			expect: Value{
				k: KindUint64,
				x: 0,
			},
		},
		{
			name:  "positive-value",
			input: 10,
			expect: Value{
				k: KindUint64,
				x: 10,
			},
		},
		{
			name:  "max-value",
			input: math.MaxUint16,
			expect: Value{
				k: KindUint64,
				x: math.MaxUint16,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := ToValue(tc.input)
			if !reflect.DeepEqual(tc.expect, actual) {
				t.Errorf("%s => \n(expected) => %#v \n(got) => %#v", tc.name, tc.expect, actual)
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
			name: "nil-value",
			expect: Value{
				k: KindNull,
			},
		},
		{
			name: "positive-value",
			input: func() *uint16 {
				i := new(uint16)
				*i = 10
				return i
			}(),
			expect: Value{
				k: KindUint64,
				x: 10,
			},
		},
		{
			name: "max-value",
			input: func() *uint16 {
				i := new(uint16)
				*i = math.MaxUint16
				return i
			}(),
			expect: Value{
				k: KindUint64,
				x: math.MaxUint16,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := ToValue(tc.input)
			if !reflect.DeepEqual(tc.expect, actual) {
				t.Errorf("%s => \n(expected) => %#v \n(got) => %#v", tc.name, tc.expect, actual)
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
			name: "zero-value",
			expect: Value{
				k: KindUint64,
				x: 0,
			},
		},
		{
			name:  "positive-value",
			input: 10,
			expect: Value{
				k: KindUint64,
				x: 10,
			},
		},
		{
			name:  "max-value",
			input: math.MaxUint32,
			expect: Value{
				k: KindUint64,
				x: math.MaxUint32,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := ToValue(tc.input)
			if !reflect.DeepEqual(tc.expect, actual) {
				t.Errorf("%s => \n(expected) => %#v \n(got) => %#v", tc.name, tc.expect, actual)
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
			name: "nil-value",
			expect: Value{
				k: KindNull,
			},
		},
		{
			name: "positive-value",
			input: func() *uint32 {
				i := new(uint32)
				*i = 10
				return i
			}(),
			expect: Value{
				k: KindUint64,
				x: 10,
			},
		},
		{
			name: "max-value",
			input: func() *uint32 {
				i := new(uint32)
				*i = math.MaxUint32
				return i
			}(),
			expect: Value{
				k: KindUint64,
				x: math.MaxUint32,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := ToValue(tc.input)
			if !reflect.DeepEqual(tc.expect, actual) {
				t.Errorf("%s => \n(expected) => %#v \n(got) => %#v", tc.name, tc.expect, actual)
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
			name: "zero-value",
			expect: Value{
				k: KindUint64,
				x: 0,
			},
		},
		{
			name:  "positive-value",
			input: 10,
			expect: Value{
				k: KindUint64,
				x: 10,
			},
		},
		{
			name:  "max-value",
			input: math.MaxUint64,
			expect: Value{
				k: KindUint64,
				x: math.MaxUint64,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := ToValue(tc.input)
			if !reflect.DeepEqual(tc.expect, actual) {
				t.Errorf("%s => \n(expected) => %#v \n(got) => %#v", tc.name, tc.expect, actual)
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
			name: "nil-value",
			expect: Value{
				k: KindNull,
			},
		},
		{
			name: "positive-value",
			input: func() *uint64 {
				i := new(uint64)
				*i = 10
				return i
			}(),
			expect: Value{
				k: KindUint64,
				x: 10,
			},
		},
		{
			name: "max-value",
			input: func() *uint64 {
				i := new(uint64)
				*i = math.MaxUint64
				return i
			}(),
			expect: Value{
				k: KindUint64,
				x: math.MaxUint64,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := ToValue(tc.input)
			if !reflect.DeepEqual(tc.expect, actual) {
				t.Errorf("%s => \n(expected) => %#v \n(got) => %#v", tc.name, tc.expect, actual)
			}
		})
	}
}
