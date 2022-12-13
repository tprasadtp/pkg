package log

import (
	"math"
	"strconv"
	"time"
)

const complexStringFmt byte = 'G'
const complexPrecision = 4

// Converts to Value.
//
//nolint:funlen,gocognit,gocyclo,cyclop // Avoiding this leads to function sprawl.
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
			any: nil,
		}
	case *bool:
		if v == nil {
			return Value{
				k:   NullKind,
				any: nil,
			}
		}
		store := uint64(0)
		if *v {
			store = 1
		}
		return Value{
			num: store,
			k:   BoolKind,
			any: nil,
		}
	case string:
		return Value{
			k:   StringKind,
			s:   v,
			any: nil,
		}
	case *string:
		if v == nil {
			return Value{
				k:   NullKind,
				any: nil,
			}
		}
		return Value{
			s:   *v,
			k:   StringKind,
			any: nil,
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
		//nolint:gomnd // Linter, you are useless here.
		return Value{
			s: strconv.FormatComplex(
				complex128(v), complexStringFmt,
				complexPrecision, 64,
			),
			k: StringKind,
		}
	case *complex64:
		if v == nil {
			return Value{
				k: NullKind,
			}
		}
		//nolint:gomnd // Linter, you are useless here.
		return Value{
			s: strconv.FormatComplex(
				complex128(*v), complexStringFmt,
				complexPrecision, 64,
			),
			k: StringKind,
		}
	case complex128:
		//nolint:gomnd // Linter, you are useless here.
		return Value{
			s: strconv.FormatComplex(
				v, complexStringFmt,
				complexPrecision, 128,
			),
			k: StringKind,
		}
	case *complex128:
		if v == nil {
			return Value{
				k: NullKind,
			}
		}
		//nolint:gomnd // Linter, you are useless here.
		return Value{
			s: strconv.FormatComplex(
				*v, complexStringFmt,
				complexPrecision, 128),
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
	default:
		return Value{
			k:   AnyKind,
			any: v,
		}
	}
}
