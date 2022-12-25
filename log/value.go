package log

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
//go:generate stringer -type=Kind -output value_kind_string.go -trimprefix=Kind
type Kind int

const (
	KindAny Kind = iota
	KindBool
	KindString
	KindInt64
	KindUint64
	KindFloat64
	KindDuration
	KindTime
	KindNull = 255
)

// Value can store any value, but for most common cases,
// it does not allocate. Inspired by slog proposal.
type Value struct {
	num uint64
	s   string
	k   Kind
	any any
}
