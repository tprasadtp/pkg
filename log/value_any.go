package log

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"time"
	"unsafe"
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
			k:   BoolKind,
		}
	case *bool:
		if v == nil {
			return Value{k: NullKind}
		}
		store := uint64(0)
		if *v {
			store = 1
		}
		return Value{num: store, k: BoolKind}
	case string:
		return Value{
			k: StringKind,
			s: v,
		}
	case *string:
		if v == nil {
			return Value{
				k: NullKind,
			}
		}
		return Value{
			s: *v,
			k: StringKind,
		}
	case int:
		return Value{
			num: uint64(v),
			k:   Int64Kind,
		}
	case *int:
		if v == nil {
			return Value{
				k: NullKind,
			}
		}
		return Value{
			num: uint64(*v),
			k:   Int64Kind,
		}
	case int8:
		return Value{
			num: uint64(v),
			k:   Int64Kind,
		}
	case *int8:
		if v == nil {
			return Value{
				k: NullKind,
			}
		}
		return Value{
			num: uint64(*v),
			k:   Int64Kind,
		}
	case int16:
		return Value{
			num: uint64(v),
			k:   Int64Kind,
		}
	case *int16:
		if v == nil {
			return Value{
				k: NullKind,
			}
		}
		return Value{
			num: uint64(*v),
			k:   Int64Kind,
		}
	case int32:
		return Value{
			num: uint64(v),
			k:   Int64Kind,
		}
	case *int32:
		if v == nil {
			return Value{
				k: NullKind,
			}
		}
		return Value{
			num: uint64(*v),
			k:   Int64Kind,
		}
	case int64:
		return Value{
			num: uint64(v),
			k:   Int64Kind,
		}
	case *int64:
		if v == nil {
			return Value{
				k: NullKind,
			}
		}
		return Value{
			num: uint64(*v),
			k:   Int64Kind,
		}
	// Unsigned integers
	case uint:
		return Value{
			num: uint64(v),
			k:   Uint64Kind,
		}
	case *uint:
		if v == nil {
			return Value{
				k: NullKind,
			}
		}
		return Value{
			num: uint64(*v),
			k:   Uint64Kind,
		}
	case uint8:
		return Value{
			num: uint64(v),
			k:   Uint64Kind,
		}
	case *uint8:
		if v == nil {
			return Value{
				k: NullKind,
			}
		}
		return Value{
			num: uint64(*v),
			k:   Int64Kind,
		}
	case uint16:
		return Value{
			num: uint64(v),
			k:   Uint64Kind,
		}
	case *uint16:
		if v == nil {
			return Value{
				k: NullKind,
			}
		}
		return Value{
			num: uint64(*v),
			k:   Uint64Kind,
		}
	case uint32:
		return Value{
			num: uint64(v),
			k:   Uint64Kind,
		}
	case *uint32:
		if v == nil {
			return Value{
				k: NullKind,
			}
		}
		return Value{
			num: uint64(*v),
			k:   Uint64Kind,
		}
	case uint64:
		return Value{
			num: v,
			k:   Uint64Kind,
		}
	case *uint64:
		if v == nil {
			return Value{
				k: NullKind,
			}
		}
		return Value{
			num: *v,
			k:   Int64Kind,
		}
	// Floats
	case float32:
		return Value{
			num: math.Float64bits(float64(v)),
			k:   Float64Kind,
		}
	case *float32:
		if v == nil {
			return Value{
				k: NullKind,
			}
		}
		return Value{
			num: math.Float64bits(float64(*v)),
			k:   Int64Kind,
		}
	case float64:
		return Value{
			num: math.Float64bits(v),
			k:   Float64Kind,
		}
	case *float64:
		if v == nil {
			return Value{
				k: NullKind,
			}
		}
		return Value{
			num: math.Float64bits(*v),
			k:   Int64Kind,
		}
	// Complex
	case complex64:
		return Value{
			s: strconv.FormatComplex(complex128(v), complexStringFmt, complexPrecision, 64),
			k: StringKind,
		}
	case *complex64:
		if v == nil {
			return Value{
				k: NullKind,
			}
		}
		return Value{
			s: strconv.FormatComplex(complex128(*v), complexStringFmt, complexPrecision, 64),
			k: StringKind,
		}
	case complex128:
		return Value{
			s: strconv.FormatComplex(v, complexStringFmt, complexPrecision, 128),
			k: StringKind,
		}
	case *complex128:
		if v == nil {
			return Value{
				k: NullKind,
			}
		}
		return Value{
			s: strconv.FormatComplex(*v, complexStringFmt, complexPrecision, 128),
			k: StringKind,
		}
	// time.Time
	case time.Duration:
		return Value{
			num: uint64(v.Nanoseconds()),
			k:   DurationKind,
		}
	case *time.Duration:
		if v == nil {
			return Value{
				k: NullKind,
			}
		}
		return Value{
			num: uint64(v.Nanoseconds()),
			k:   DurationKind,
		}
	case time.Time:
		return Value{
			num: uint64(v.UnixNano()),
			s:   v.Location().String(),
			k:   TimeKind,
		}
	case *time.Time:
		if v == nil {
			return Value{
				k: NullKind,
			}
		}
		return Value{
			num: uint64(v.UnixNano()),
			s:   v.Location().String(),
			k:   TimeKind,
		}
	// Experimental.
	// Only this has some significant usage (~5k hits on [GitHub])
	// [GitHub]: https://github.com/search?q=zap.Strings+language%3AGo+language%3AGo&type=Code
	case []string:
		h := (*reflect.SliceHeader)(unsafe.Pointer(&v))
		return Value{
			num: uint64(h.Len),
			k:   StringSliceKind,
			val: unsafe.Pointer(h.Data),
		}
	// Objects implementing stringer will be transformed to StringKind.
	// Not all stringer implementations have zero garbage. Many use
	// fmt.Sprintf under the hood, which allocates.
	case fmt.Stringer:
		return Value{
			s: v.String(),
			k: StringKind,
		}
	default:
		return Value{
			k:   AnyKind,
			val: v,
		}
	}
}
