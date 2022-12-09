package log

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

type Value struct {
	num uint64
	str string
	any any
}

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

	// Avoid using this, this can be null
	NullKind = 255
)

// Converts any to Value
func AnyValue(v any) Value {
	switch v := v.(type) {
	case bool:
		store := uint64(0)
		if v {
			store = 1
		}
		return Value{num: store, any: BoolKind}
	case *bool:
		store := uint64(0)
		if *v {
			store = 1
		}
		return Value{num: store, any: BoolKind}
	case string:
		return Value{
			str: v,
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
			num: uint64(v),
			any: Int64Kind,
		}
	case *int:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			num: uint64(*v),
			any: Int64Kind,
		}
	case int8:
		return Value{
			num: uint64(v),
			any: Int64Kind,
		}
	case *int8:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			num: uint64(*v),
			any: Int64Kind,
		}
	case int16:
		return Value{
			num: uint64(v),
			any: Int64Kind,
		}
	case *int16:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			num: uint64(*v),
			any: Int64Kind,
		}
	case int32:
		return Value{
			num: uint64(v),
			any: Int64Kind,
		}
	case *int32:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			num: uint64(*v),
			any: Int64Kind,
		}
	case int64:
		return Value{
			num: uint64(v),
			any: Int64Kind,
		}
	case *int64:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			num: uint64(*v),
			any: Int64Kind,
		}
		// Unsigned integers
	case uint:
		return Value{
			num: uint64(v),
			any: Uint64Kind,
		}
	case *uint:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			num: uint64(*v),
			any: Uint64Kind,
		}
	case uint8:
		return Value{
			num: uint64(v),
			any: Uint64Kind,
		}
	case *uint8:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			num: uint64(*v),
			any: Int64Kind,
		}
	case uint16:
		return Value{
			num: uint64(v),
			any: Uint64Kind,
		}
	case *uint16:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			num: uint64(*v),
			any: Uint64Kind,
		}
	case uint32:
		return Value{
			num: uint64(v),
			any: Uint64Kind,
		}
	case *uint32:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			num: uint64(*v),
			any: Uint64Kind,
		}
	case uint64:
		return Value{
			num: uint64(v),
			any: Uint64Kind,
		}
	case *uint64:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			num: uint64(*v),
			any: Int64Kind,
		}
	// Floats
	case float32:
		return Value{
			num: math.Float64bits(float64(v)),
			any: Float64Kind,
		}
	case *float32:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			num: math.Float64bits(float64(*v)),
			any: Int64Kind,
		}
	case float64:
		return Value{
			num: math.Float64bits(float64(v)),
			any: Float64Kind,
		}
	case *float64:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			num: math.Float64bits(float64(*v)),
			any: Int64Kind,
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
	// objects implementing stringer will be transformed to StringKind
	// not all stringer implementations have zero garbage.
	case fmt.Stringer:
		return Value{
			str: v.String(),
			any: StringKind,
		}
	// time.Time
	case time.Duration:
		return Value{
			num: uint64(v.Nanoseconds()),
			any: DurationKind,
		}
	case *time.Duration:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			num: uint64(v.Nanoseconds()),
			any: DurationKind,
		}
	case time.Time:
		return Value{
			num: uint64(v.UnixNano()),
		}
	case *time.Time:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			num: uint64(v.UnixNano()),
			any: DurationKind,
		}
	default:
		return Value{
			any: v,
		}
	}
}
