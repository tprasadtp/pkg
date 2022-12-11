package log

import (
	"fmt"
)

// ValueKind represents Kind of the value.
// This similar to reflect.Kinds, but does not attempt
// to preserve exact types.
//   - Uint64Kind represents uint, uint8, uint32 and uint64.
//   - Int64Kind represents int, int8, int32, and int64
//   - Float64Kind represents float32 and float64
//   - complex64 and complex128 are converted to their string representation.
//   - [time.Time] is saved as TimeKind, but loses time.Time
//   - Pointers to all the values are dereferenced unless they are nil.
//     In case of a nil pointer, type information is lost. This
//     may not seem optimal, but most logging solutions convert
//     all field values to json or string anyway so it does not affect much
//     in applications.
//
//go:generate stringer -type=Kind -output value_string.go
type Kind int

const (
	AnyKind Kind = iota
	BoolKind
	StringKind
	Int64Kind
	Uint64Kind
	Float64Kind
	DurationKind
	TimeKind

	NullKind = 255
)

const (
	StringSliceKind Kind = iota + 1<<12
)

// Value can store any value, but for most common cases,
// it does not allocate. Inspired by slog proposal.
type Value struct {
	num uint64
	s   string
	k   Kind
	val any
}

// Kind returns Kind of the Value.
func (v Value) Kind() Kind {
	return v.k
}

func (v Value) Int64() (int64, error) {
	if v.Kind() == Int64Kind {
		return int64(v.num), nil
	}
	return 0, fmt.Errorf("log: Kind is %s, not Int64Kind", v.Kind().String())
}

func (v Value) Uint64() (uint64, error) {
	if v.Kind() == Uint64Kind {
		return v.num, nil
	}
	return 0, fmt.Errorf("log: Kind is %s, not Int64Kind", v.Kind().String())
}
