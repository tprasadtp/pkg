package log

import (
	"crypto/rand"
	"encoding/hex"
)

// Get a new shiny TraceID
func NewTraceID() (TraceID, error) {
	var err error
	t := TraceID{}
	_, err = rand.Read(t.id[:])
	return t, err
}

// TraceID is trace ID used for open tracing
type TraceID struct {
	id [16]byte
}

// Hexadecimal String representation of TraceID
func (t TraceID) String() string {
	return hex.EncodeToString(t.id[:])
}
