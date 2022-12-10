package multi_test

import (
	"errors"
	"testing"

	"github.com/tprasadtp/pkg/log"
	"github.com/tprasadtp/pkg/log/handlers/multi"
	"github.com/tprasadtp/pkg/log/internal/testdata"
)

func TestHandler(t *testing.T) {
	h1 := log.MockHandler{
		Level: log.LevelInfo,
	}

	h2 := log.MockHandler{
		Level: log.LevelError,
	}

	m := multi.New(&h1, &h2)

	for _, e := range testdata.GetEvents() {
		if m.Enabled(e.Level) {
			if err := m.Write(e); err != nil {
				t.Errorf("handler returned error(%e), event=%s", err, e.Message)
			}
		}
	}

	if h1.EventsWritten != testdata.I {
		t.Errorf("incorrect number of events on h1(@InfoLevel) expected=%d, got=%d",
			testdata.I,
			h1.EventsWritten,
		)
	}
	if h2.EventsWritten != testdata.E {
		t.Errorf("incorrect number of events on h2(@ErrorLevel) expected=%d, got=%d",
			testdata.E,
			h2.EventsWritten,
		)
	}

	if err := m.Flush(); err != nil {
		t.Errorf("handler flush returned error(%e)", err)
	}

	if h1.EventsWritten != 0 {
		t.Errorf("flush did not clear events on h1(@InfoLevel) expected=0, got=%d",
			h1.EventsWritten)
	}
	if h2.EventsWritten != 0 {
		t.Errorf("flush did not clear events on h2(@ErrorLevel) expected=0, got=%d",
			h2.EventsWritten)
	}
}

func TestHandlerClose(t *testing.T) {
	h1 := log.MockHandler{
		Level: log.LevelInfo,
	}

	h2 := log.MockHandler{
		Level: log.LevelError,
	}

	m := multi.New(&h1, &h2)

	e := log.Event{
		Level:   log.LevelInfo,
		Message: "TestEvent",
	}

	// Close the handler
	if err := m.Close(); err != nil {
		t.Errorf("first handler close returned error(%e)", err)
	}

	// Write to closed handler should error
	if err := m.Write(e); !errors.Is(err, log.ErrHandlerClosed) {
		t.Errorf("write on closed handler returned unexpected error (%v)", err)
	}

	// Flushing closed handler should error
	if err := m.Flush(); !errors.Is(err, log.ErrHandlerClosed) {
		t.Errorf("flush on closed handler returned unexpected error (%v)", err)
	}

	// Close to already closed handler should error
	if err := m.Close(); !errors.Is(err, log.ErrHandlerClosed) {
		t.Errorf("close on closed handler returned unexpected error(%v)", err)
	}
}

func TestOneAlreadyClosedHandler(t *testing.T) {
	h1 := log.MockHandler{
		Level: log.LevelInfo,
	}

	h2 := log.MockHandler{
		Level: log.LevelError,
	}

	h1.Close()

	m := multi.New(&h1, &h2)

	for _, e := range testdata.GetEvents() {
		if m.Enabled(e.Level) {
			if err := m.Write(e); !errors.Is(err, log.ErrHandlerClosed) {
				t.Errorf("h1(closed) invalid error => got=(%s), expected=(%s)",
					err,
					log.ErrHandlerClosed,
				)
			}
		}
	}

	if h1.EventsWritten != 0 {
		t.Errorf("incorrect number of events on h1[closed](@InfoLevel) expected=0 got=%d",
			h1.EventsWritten,
		)
	}
	if h2.EventsWritten != testdata.E {
		t.Errorf("incorrect number of events on h2(@ErrorLevel) expected=%d, got=%d",
			testdata.E,
			h2.EventsWritten,
		)
	}

	if err := m.Flush(); !errors.Is(err, log.ErrHandlerClosed) {
		t.Errorf("handler flush returned error(%e)", err)
	}

	if h1.EventsWritten != 0 {
		t.Errorf("flush did not clear events on h1(@InfoLevel) expected=0, got=%d",
			h1.EventsWritten)
	}
	if h2.EventsWritten != 0 {
		t.Errorf("flush did not clear events on h2(@ErrorLevel) expected=0, got=%d",
			h2.EventsWritten)
	}
}

func TestMultiHandlerWithError(t *testing.T) {
	h1 := log.MockHandler{
		Level: log.LevelInfo,
	}

	h2 := log.MockHandler{
		Level:     log.LevelError,
		AlwaysErr: true,
	}

	h3 := log.MockHandler{
		Level:     log.LevelVerbose,
		AlwaysErr: true,
	}

	m := multi.New(&h1, &h2, &h3)

	for _, e := range testdata.GetEvents() {
		if m.Enabled(e.Level) {
			if err := m.Write(e); !errors.Is(err, log.ErrHandlerWrite) {
				t.Errorf(
					"handle error mismatch (@%s) => expected=%s, got=%s",
					e.Message,
					log.ErrHandlerWrite,
					err)
			}
		}
	}

	if h2.EventsWritten != 0 {
		t.Errorf("incorrect number of events on h2(@ErrorLevel) expected=0, got=%d",
			h2.EventsWritten)
	}
	if h2.WriteCalls != 5 {
		t.Errorf("incorrect number of Handle() calls on h2(@ErrorLevel) expected=5, got=%d",
			h2.EventsWritten)
	}

	if h3.EventsWritten != 0 {
		t.Errorf("incorrect number of events on h1(@ErrorLevel) expected=0, got=%d",
			h3.EventsWritten)
	}
	if h3.WriteCalls != 14 {
		t.Errorf("incorrect number of events on h1(@ErrorLevel) expected=14, got=%d",
			h3.EventsWritten)
	}

	if err := m.Flush(); !errors.Is(err, log.ErrHandlerWrite) {
		t.Errorf(
			"flush error mismatch => expected=%s, got=%s",
			log.ErrHandlerWrite,
			err)
	}
}
