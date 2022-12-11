// Code generated by "stringer -type=Kind -output value_string.go"; DO NOT EDIT.

package log

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AnyKind-0]
	_ = x[BoolKind-1]
	_ = x[StringKind-2]
	_ = x[Int64Kind-3]
	_ = x[Uint64Kind-4]
	_ = x[Float64Kind-5]
	_ = x[DurationKind-6]
	_ = x[TimeKind-7]
	_ = x[StringSliceKind-4096]
}

const (
	_Kind_name_0 = "AnyKindBoolKindStringKindInt64KindUint64KindFloat64KindDurationKindTimeKind"
	_Kind_name_1 = "StringSliceKind"
)

var (
	_Kind_index_0 = [...]uint8{0, 7, 15, 25, 34, 44, 55, 67, 75}
)

func (i Kind) String() string {
	switch {
	case 0 <= i && i <= 7:
		return _Kind_name_0[_Kind_index_0[i]:_Kind_index_0[i+1]]
	case i == 4096:
		return _Kind_name_1
	default:
		return "Kind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}