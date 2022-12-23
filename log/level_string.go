// Code generated by "stringer -type=Level -output level_string.go -trimprefix=Level"; DO NOT EDIT.

package log

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LevelTrace - -30]
	_ = x[LevelDebug - -20]
	_ = x[LevelVerbose - -10]
	_ = x[LevelInfo-0]
	_ = x[LevelSuccess-10]
	_ = x[LevelNotice-20]
	_ = x[LevelWarning-30]
	_ = x[LevelError-40]
	_ = x[LevelCritical-50]
	_ = x[LevelFatal-90]
}

const (
	_Level_name_0 = "Trace"
	_Level_name_1 = "Debug"
	_Level_name_2 = "Verbose"
	_Level_name_3 = "Info"
	_Level_name_4 = "Success"
	_Level_name_5 = "Notice"
	_Level_name_6 = "Warning"
	_Level_name_7 = "Error"
	_Level_name_8 = "Critical"
	_Level_name_9 = "Fatal"
)

func (i Level) String() string {
	switch {
	case i == -30:
		return _Level_name_0
	case i == -20:
		return _Level_name_1
	case i == -10:
		return _Level_name_2
	case i == 0:
		return _Level_name_3
	case i == 10:
		return _Level_name_4
	case i == 20:
		return _Level_name_5
	case i == 30:
		return _Level_name_6
	case i == 40:
		return _Level_name_7
	case i == 50:
		return _Level_name_8
	case i == 90:
		return _Level_name_9
	default:
		return "Level(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
