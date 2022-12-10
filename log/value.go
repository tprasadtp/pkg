package log

import (
	"fmt"
	"math"
	"strconv"
	"time"
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

// Value can store any value, but for most common cases it does not allocate.
type Value struct {
	store uint64
	s     string
	k     Kind
	any   any
}

func (v Value) Kind() Kind {
	return v.k
}

func (v Value) Int64() (int64, error) {
	if v.Kind() == Int64Kind {
		return int64(v.store), nil
	}
	return 0, ErrInvalidKind
}

func (v Value) String() (int64, error) {
	if v.Kind() == Int64Kind {
		return int64(v.store), nil
	}
	return 0, ErrInvalidKind
}

// Converts to Value
func ToValue(v any) Value {
	switch v := v.(type) {
	case bool:
		store := uint64(0)
		if v {
			store = 1
		}
		return Value{
			store: store,
			k:     BoolKind,
		}
	case *bool:
		store := uint64(0)
		if *v {
			store = 1
		}
		return Value{store: store, k: BoolKind}
	case string:
		return Value{
			any: StringKind,
		}
	case *string:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			str: *v,
			any: StringKind,
		}
	case int:
		return Value{
			store: uint64(v),
			any:   Int64Kind,
		}
	case *int:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			store: uint64(*v),
			any:   Int64Kind,
		}
	case int8:
		return Value{
			store: uint64(v),
			any:   Int64Kind,
		}
	case *int8:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			store: uint64(*v),
			any:   Int64Kind,
		}
	case int16:
		return Value{
			store: uint64(v),
			any:   Int64Kind,
		}
	case *int16:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			store: uint64(*v),
			any:   Int64Kind,
		}
	case int32:
		return Value{
			store: uint64(v),
			any:   Int64Kind,
		}
	case *int32:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			store: uint64(*v),
			any:   Int64Kind,
		}
	case int64:
		return Value{
			store: uint64(v),
			any:   Int64Kind,
		}
	case *int64:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			store: uint64(*v),
			any:   Int64Kind,
		}
	// Unsigned integers
	case uint:
		return Value{
			store: uint64(v),
			any:   Uint64Kind,
		}
	case *uint:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			store: uint64(*v),
			any:   Uint64Kind,
		}
	case uint8:
		return Value{
			store: uint64(v),
			any:   Uint64Kind,
		}
	case *uint8:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			store: uint64(*v),
			any:   Int64Kind,
		}
	case uint16:
		return Value{
			store: uint64(v),
			any:   Uint64Kind,
		}
	case *uint16:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			store: uint64(*v),
			any:   Uint64Kind,
		}
	case uint32:
		return Value{
			store: uint64(v),
			any:   Uint64Kind,
		}
	case *uint32:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			store: uint64(*v),
			any:   Uint64Kind,
		}
	case uint64:
		return Value{
			store: uint64(v),
			any:   Uint64Kind,
		}
	case *uint64:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			store: uint64(*v),
			any:   Int64Kind,
		}
	// Floats
	case float32:
		return Value{
			store: math.Float64bits(float64(v)),
			any:   Float64Kind,
		}
	case *float32:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			store: math.Float64bits(float64(*v)),
			any:   Int64Kind,
		}
	case float64:
		return Value{
			store: math.Float64bits(float64(v)),
			any:   Float64Kind,
		}
	case *float64:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			store: math.Float64bits(float64(*v)),
			any:   Int64Kind,
		}
	// Complex
	case complex64:
		return Value{
			str: strconv.FormatComplex(complex128(v), 4, 'g', 64),
			any: StringKind,
		}
	case *complex64:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			str: strconv.FormatComplex(complex128(*v), 4, 'g', 64),
			any: StringKind,
		}
	case complex128:
		return Value{
			str: strconv.FormatComplex(complex128(v), 4, 'g', 128),
			any: StringKind,
		}
	case *complex128:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			str: strconv.FormatComplex(complex128(*v), 4, 'g', 128),
			any: StringKind,
		}
	// Objects implementing stringer will be transformed to StringKind.
	// Not all stringer implementations have zero garbage. Many use
	// fmt.Sprintf under the hood, which allocates.
	case fmt.Stringer:
		return Value{
			str: v.String(),
			any: StringKind,
		}
	// time.Time
	case time.Duration:
		return Value{
			store: uint64(v.Nanoseconds()),
			any:   DurationKind,
		}
	case *time.Duration:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			store: uint64(v.Nanoseconds()),
			any:   DurationKind,
		}
	case time.Time:
		return Value{
			store: uint64(v.UnixNano()),
		}
	case *time.Time:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			store: uint64(v.UnixNano()),
			any:   DurationKind,
		}
	default:
		return Value{
			any: v,
		}
	}
}
