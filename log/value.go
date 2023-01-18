package log

// ValueKind represents Kind of the value.
// This similar to reflect.Kinds, but does not attempt
// to preserve exact types.
//   - Uint64Kind represents uint, uint8, uint32 and uint64.
//   - Int64Kind represents int, int8, int32, and int64
//   - Float64Kind represents float32 and float64
//   - KindComplex128 represents complex64 and complex128.
//   - [time.Time] is saved as TimeKind, but loses time.Time monotonic value.
//   - [time.Duration] is saved as DurationKind.
//   - Pointers to all the values are dereferenced unless they are nil.
//     In case of a nil pointer, type information is lost. This
//     may not seem optimal, but most logging solutions convert
//     all field values to json or string anyway so it does not affect much
//     in applications.
type Kind uint8

const (
	KindAny Kind = iota
	KindBool
	KindString
	KindInt64
	KindUint64
	KindFloat64
	KindComplex128
	KindDuration
	KindTime
	KindIPAddr
	KindIPPrefix
	KindNull = 255
)

// Value can store any value, but for most common cases,
// it does not allocate. Internal representation of a value
// is not subject to compatibility guarantees.
type Value struct {
	x   uint64
	s   string
	k   Kind
	any any
}

// Int64 returns int64 value if [Kind] is [KindInt64].
// Panics otherwise.
func (v Value) Int64() int64 {
	if v.k != KindInt64 {
		panic("Value is not of Kind KindInt64")
	}
	return int64(v.x)
}
