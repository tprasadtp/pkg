package log

import (
	"math"
	"net/netip"
	"strconv"
	"time"
)

const complexStringFmt byte = 'G'
const complexPrecision = 4

// Constructor for Value.
//
//nolint:funlen,gocognit,gocyclo,cyclop // Avoiding this leads to function sprawl.
func AnyValue(v any) Value {
	switch v := v.(type) {
	case bool:
		store := uint64(0)
		if v {
			store = 1
		}
		return Value{
			x:   store,
			k:   KindBool,
			any: nil,
		}
	case *bool:
		if v == nil {
			return Value{
				k:   KindNull,
				any: nil,
			}
		}
		store := uint64(0)
		if *v {
			store = 1
		}
		return Value{
			x:   store,
			k:   KindBool,
			any: nil,
		}
	case string:
		return Value{
			k:   KindString,
			s:   v,
			any: nil,
		}
	case *string:
		if v == nil {
			return Value{
				k:   KindNull,
				any: nil,
			}
		}
		return Value{
			s:   *v,
			k:   KindString,
			any: nil,
		}
	case int:
		return Value{
			x: uint64(v),
			k: KindInt64,
		}
	case *int:
		if v == nil {
			return Value{
				k: KindNull,
			}
		}
		return Value{
			x: uint64(*v),
			k: KindInt64,
		}
	case int8:
		return Value{
			x: uint64(v),
			k: KindInt64,
		}
	case *int8:
		if v == nil {
			return Value{
				k: KindNull,
			}
		}
		return Value{
			x: uint64(*v),
			k: KindInt64,
		}
	case int16:
		return Value{
			x: uint64(v),
			k: KindInt64,
		}
	case *int16:
		if v == nil {
			return Value{
				k: KindNull,
			}
		}
		return Value{
			x: uint64(*v),
			k: KindInt64,
		}
	case int32:
		return Value{
			x: uint64(v),
			k: KindInt64,
		}
	case *int32:
		if v == nil {
			return Value{
				k: KindNull,
			}
		}
		return Value{
			x: uint64(*v),
			k: KindInt64,
		}
	case int64:
		return Value{
			x: uint64(v),
			k: KindInt64,
		}
	case *int64:
		if v == nil {
			return Value{
				k: KindNull,
			}
		}
		return Value{
			x: uint64(*v),
			k: KindInt64,
		}
	// Unsigned integers
	case uint:
		return Value{
			x: uint64(v),
			k: KindUint64,
		}
	case *uint:
		if v == nil {
			return Value{
				k: KindNull,
			}
		}
		return Value{
			x: uint64(*v),
			k: KindUint64,
		}
	case uint8:
		return Value{
			x: uint64(v),
			k: KindUint64,
		}
	case *uint8:
		if v == nil {
			return Value{
				k: KindNull,
			}
		}
		return Value{
			x: uint64(*v),
			k: KindUint64,
		}
	case uint16:
		return Value{
			x: uint64(v),
			k: KindUint64,
		}
	case *uint16:
		if v == nil {
			return Value{
				k: KindNull,
			}
		}
		return Value{
			x: uint64(*v),
			k: KindUint64,
		}
	case uint32:
		return Value{
			x: uint64(v),
			k: KindUint64,
		}
	case *uint32:
		if v == nil {
			return Value{
				k: KindNull,
			}
		}
		return Value{
			x: uint64(*v),
			k: KindUint64,
		}
	case uint64:
		return Value{
			x: v,
			k: KindUint64,
		}
	case *uint64:
		if v == nil {
			return Value{
				k: KindNull,
			}
		}
		return Value{
			x: *v,
			k: KindUint64,
		}
	// Floats
	case float32:
		return Value{
			x: math.Float64bits(float64(v)),
			k: KindFloat64,
		}
	case *float32:
		if v == nil {
			return Value{
				k: KindNull,
			}
		}
		return Value{
			x: math.Float64bits(float64(*v)),
			k: KindFloat64,
		}
	case float64:
		return Value{
			x: math.Float64bits(v),
			k: KindFloat64,
		}
	case *float64:
		if v == nil {
			return Value{
				k: KindNull,
			}
		}
		return Value{
			x: math.Float64bits(*v),
			k: KindFloat64,
		}
	// Complex
	case complex64:
		//nolint:gomnd // Linter, you are useless here.
		return Value{
			s: strconv.FormatComplex(
				complex128(v), complexStringFmt,
				complexPrecision, 64,
			),
			k: KindComplex128,
		}
	case *complex64:
		if v == nil {
			return Value{
				k: KindNull,
			}
		}
		//nolint:gomnd // Linter, you are useless here.
		return Value{
			s: strconv.FormatComplex(
				complex128(*v), complexStringFmt,
				complexPrecision, 64,
			),
			k: KindComplex128,
		}
	case complex128:
		//nolint:gomnd // Linter, you are useless here.
		return Value{
			s: strconv.FormatComplex(
				v, complexStringFmt,
				complexPrecision, 128,
			),
			k: KindComplex128,
		}
	case *complex128:
		if v == nil {
			return Value{
				k: KindNull,
			}
		}
		//nolint:gomnd // Linter, you are useless here.
		return Value{
			s: strconv.FormatComplex(
				*v, complexStringFmt,
				complexPrecision, 128),
			k: KindComplex128,
		}
	// time.Time
	case time.Duration:
		return Value{
			x: uint64(v.Nanoseconds()),
			k: KindDuration,
		}
	case *time.Duration:
		if v == nil {
			return Value{
				k: KindNull,
			}
		}
		return Value{
			x: uint64(v.Nanoseconds()),
			k: KindDuration,
		}
	case time.Time:
		return Value{
			x: uint64(v.UnixNano()),
			s: v.Location().String(),
			k: KindTime,
		}
	case *time.Time:
		if v == nil {
			return Value{
				k: KindNull,
			}
		}
		return Value{
			x: uint64(v.UnixNano()),
			s: v.Location().String(),
			k: KindTime,
		}
	// netip.Addr (alloc free)
	case netip.Addr:
		return Value{
			s: v.String(),
			k: KindIPAddr,
		}
	case *netip.Addr:
		if v == nil {
			return Value{
				k: KindNull,
			}
		}
		return Value{
			s: v.String(),
			k: KindIPAddr,
		}
	case netip.Prefix:
		return Value{
			s: v.String(),
			k: KindIPPrefix,
		}
	case *netip.Prefix:
		if v == nil {
			return Value{
				k: KindNull,
			}
		}
		return Value{
			s: v.String(),
			k: KindIPPrefix,
		}
	default:
		return Value{
			k:   KindAny,
			any: v,
		}
	}
}
