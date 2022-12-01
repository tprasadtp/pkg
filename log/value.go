package log

type Kind uint8

// AnyKind is 0 so that a zero Value represents nil.
const AnyKind Kind = 0

const ValueMarshalerKind Kind = 255

const (
	BoolKind = iota + 1

	Int64Kind
	Uint64Kind
	Float64Kind

	StringKind

	DurationKind
	TimeKind
)

func NewValue(a any) Value {
	return Value{}
}

// A Value is can represent any Go value, but unlike type any,
// it can represent most small values without an allocation.
type Value struct {
	// num holds the value for Kinds Int64, Uint64, Float64, Bool and Duration,
	// and nanoseconds since the epoch for TimeKind.
	num uint64
	// s holds the value for StringKind.
	s string
	// If a.(type) is of Kind, then the value
	//  - is in num or s as described above.
	// If any is of type *time.Location, then the Kind is Time and time.Time
	// value can be constructed from the Unix nano sec in num and the location
	// (monotonic time is not preserved).
	// Otherwise, the Kind is Any and any is the value.
	// (This implies that Values cannot store Kinds or *time.Locations.)
	a any
}

// ValueMarshaler.
type ValueMarshaler interface {
	Value() Value
}
