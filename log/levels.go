package log

import (
	"math"
)

// Level Level represents log level.
type Level uint16

const (
	DEBUG   Level = 10
	VERBOSE Level = 20
	INFO    Level = 30
	SUCCESS Level = 40
	NOTICE  Level = 50
	WARNING Level = 60
	ERROR   Level = 70
	PANIC   Level = 80
	EXIT    Level = math.MaxUint16 - 1
)

const (
	ALL     Level = 0
	NONE    Level = math.MaxUint16
	UNKNOWN Level = math.MaxUint16 - 2
)
