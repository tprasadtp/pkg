package log

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

const complexStringFmt byte = 'G'
const complexPrecision = 4

// Converts to Value.
func ToValue(v any) Value {
	switch v := v.(type) {
	case bool:
		store := uint64(0)
		if v {
			store = 1
		}
		return Value{
			num: store,
			any: BoolKind,
		}
	case *bool:
		if v == nil {
			return Value{any: NullKind}
		}
		store := uint64(0)
		if *v {
			store = 1
		}
		return Value{num: store, any: BoolKind}
	case string:
		return Value{
			any: StringKind,
			s:   v,
		}
	case *string:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			s:   *v,
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
			num: v,
			any: Uint64Kind,
		}
	case *uint64:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			num: *v,
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
			num: math.Float64bits(v),
			any: Float64Kind,
		}
	case *float64:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			num: math.Float64bits(*v),
			any: Int64Kind,
		}
	// Complex
	case complex64:
		return Value{
			s:   strconv.FormatComplex(complex128(v), complexStringFmt, complexPrecision, 64),
			any: StringKind,
		}
	case *complex64:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			s:   strconv.FormatComplex(complex128(*v), complexStringFmt, complexPrecision, 64),
			any: StringKind,
		}
	case complex128:
		return Value{
			s:   strconv.FormatComplex(v, complexStringFmt, complexPrecision, 128),
			any: StringKind,
		}
	case *complex128:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			s:   strconv.FormatComplex(*v, complexStringFmt, complexPrecision, 128),
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
			s:   v.Location().String(),
			any: TimeKind,
		}
	case *time.Time:
		if v == nil {
			return Value{
				any: NullKind,
			}
		}
		return Value{
			num: uint64(v.UnixNano()),
			s:   v.Location().String(),
			any: TimeKind,
		}
	// Objects implementing stringer will be transformed to StringKind.
	// Not all stringer implementations have zero garbage. Many use
	// fmt.Sprintf under the hood, which allocates.
	case fmt.Stringer:
		return Value{
			s:   v.String(),
			any: StringKind,
		}
	default:
		return Value{
			any: v,
		}
	}
}
