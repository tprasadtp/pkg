package log

import (
	"time"
)

// Field is logger fields. Optionally with a namespace.
// Namespace is distinct from Logger's namespace.
// Handler MUST consider both logger's and Field's namespaces.
type Field struct {
	Namespace string
	Key       string
	Value     Value
}

// Includes caller info if available.
type CallerInfo struct {
	// Defined represents whether caller info is defined.
	Defined bool
	// Line number of the caller
	// If not available, this is 0.
	Line uint
	// File containing the code
	// If not available this is empty string.
	File string
	// Function name of the caller.
	// this includes full path of the package.
	// except for main package.
	// This is empty if information is not available.
	Func string
}

// Event represents a single Log event.
// Event should be considered immutable.
type Event struct {
	// Namespace is namespace of the logger that generated this event.
	Namespace string

	// Time (Global)
	Time time.Time

	// Log Level (Global)
	Level Level

	// Log Message (Global)
	Message string

	// Error (Global)
	Error error

	// Caller
	Caller CallerInfo

	// Fields is preallocated fields.
	// This typically enough for all the field attributes
	// This is an allocation optimization.
	Fields [fieldsBucketSize]Field
	// FieldsOverflow is same as Fields, but unlike Fields is not of fixed length.
	FieldsOverflow []Field
}
