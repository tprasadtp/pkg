package log

import (
	"crypto/rand"
	"encoding/hex"
)

// Get a new shiny SpanID
func NewSpanID() (SpanID, error) {
	var err error
	s := SpanID{}
	_, err = rand.Read(s.id[:])
	return s, err
}

// SpanID is trace ID used for open tracing
type SpanID struct {
	id [16]byte
}

// Hexadecimal String representation of SpanID
func (s SpanID) String() string {
	return hex.EncodeToString(s.id[:])
}
