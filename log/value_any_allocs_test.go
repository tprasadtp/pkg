package log_test

import (
	"math"
	"testing"
	"time"

	"github.com/tprasadtp/pkg/log"
)

func TestToValueAllocs(t *testing.T) {
	type testCase struct {
		name   string
		f      func()
		allocs float64
	}

	tt := []testCase{
		{
			name: "<time.Duration>",
			f: func() {
				log.ToValue(time.Second)
			},
		},
		//Int
		{
			name: "<int>",
			f: func() {
				log.ToValue(math.MaxInt)
			},
		},
		{
			name: "<int8>",
			f: func() {
				log.ToValue(math.MaxInt8)
			},
		},
		{
			name: "<int16>",
			f: func() {
				log.ToValue(math.MaxInt16)
			},
		},
		{
			name: "<int32>",
			f: func() {
				log.ToValue(math.MaxInt32)
			},
		},
		{
			name: "<int64>",
			f: func() {
				log.ToValue(math.MaxInt64)
			},
		},
		// Integer Pointers
		{
			name: "<intptr>",
			f: func() {
				p := new(int)
				*p = math.MaxInt
				log.ToValue(p)
			},
		},
		{
			name: "<int8ptr>",
			f: func() {
				p := new(int8)
				*p = math.MaxInt8
				log.ToValue(p)
			},
		},
		{
			name: "<int16ptr>",
			f: func() {
				p := new(int16)
				*p = math.MaxInt16
				log.ToValue(p)
			},
		},
		{
			name: "<int32ptr>",
			f: func() {
				p := new(int32)
				*p = math.MaxInt32
				log.ToValue(p)
			},
		},
		{
			name: "<int64ptr>",
			f: func() {
				p := new(int64)
				*p = math.MaxInt64
				log.ToValue(p)
			},
		},
		// Uint
		{
			name: "<uint>",
			f: func() {
				log.ToValue(math.MaxUint32)
			},
		},
		{
			name: "<uint8>",
			f: func() {
				log.ToValue(math.MaxUint8)
			},
		},
		{
			name: "<uint16>",
			f: func() {
				log.ToValue(math.MaxUint16)
			},
		},
		{
			name: "<uint32>",
			f: func() {
				log.ToValue(math.MaxUint32)
			},
		},
		{
			name: "<uint64>",
			f: func() {
				log.ToValue(uint64(math.MaxUint))
			},
		},
		// Uint Pointers
		{
			name: "<uintptr>",
			f: func() {
				p := new(uint)
				*p = math.MaxUint
				log.ToValue(p)
			},
		},
		{
			name: "<uint8ptr>",
			f: func() {
				p := new(uint8)
				*p = math.MaxUint8
				log.ToValue(p)
			},
		},
		{
			name: "<uint16ptr>",
			f: func() {
				p := new(uint16)
				*p = math.MaxUint16
				log.ToValue(p)
			},
		},
		{
			name: "<uint32ptr>",
			f: func() {
				p := new(uint32)
				*p = math.MaxUint32
				log.ToValue(p)
			},
		},
		{
			name: "<uint64ptr>",
			f: func() {
				p := new(uint64)
				*p = math.MaxUint64
				log.ToValue(p)
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			allocs := testing.AllocsPerRun(10, tc.f)
			if allocs != tc.allocs {
				t.Errorf("%s => alloc mismatch expected=%f, actual=%f", tc.name, tc.allocs, allocs)
			}
		})
	}
}
