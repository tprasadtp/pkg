package log

import (
	"math"
)

type Level uint16

const (
	DEBUG   Level = 10
	VERBOSE Level = 15
	INFO    Level = 20
	SUCCESS Level = 21
	NOTICE  Level = 25
	WARNING Level = 30
	ERROR   Level = 40
	FATAL   Level = 50
	PANIC   Level = FATAL
)

const (
	ALL     Level = 0
	NONE    Level = math.MaxUint16
	UNKNOWN Level = math.MaxUint16 - 1
)
