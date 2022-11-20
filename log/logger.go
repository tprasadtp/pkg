package log

import (
	"time"
)

type Event struct {
	Level     Level
	Timestamp time.Time
	Message   string
	Fields    []Field
}
