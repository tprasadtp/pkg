package multi_test

import (
	"errors"
	"testing"

	"github.com/tprasadtp/pkg/log"
	"github.com/tprasadtp/pkg/log/handlers/multi"
)

func TestMultiHandler(t *testing.T) {
	h1 := log.MockHandler{
		Level: log.InfoLevel,
	}

	h2 := log.MockHandler{
		Level: log.ErrorLevel,
	}

	m := multi.New(&h1, &h2)

	for _, e := range events {
		if m.Enabled(e.Level) {
			if err := m.Handle(e); err != nil {
				t.Errorf("handler returned error(%e), event=%s", err, e.Message)
			}
		}
	}

	if h1.EventCount != 12 {
		t.Errorf("incorrect number of events on h1(@InfoLevel) expected=12, got=%d",
			h1.EventCount)
	}
	if h2.EventCount != 5 {
		t.Errorf("incorrect number of events on h2(@ErrorLevel) expected=5, got=%d",
			h2.EventCount)
	}

	if err := m.Flush(); err != nil {
		t.Errorf("handler flush returned error(%e)", err)
	}

	if h1.EventCount != 0 {
		t.Errorf("flush did not clear events on h1(@InfoLevel) expected=0, got=%d",
			h1.EventCount)
	}
	if h2.EventCount != 0 {
		t.Errorf("flush did not clear events on h2(@ErrorLevel) expected=0, got=%d",
			h2.EventCount)
	}
}

func TestMultiHandlerWithError(t *testing.T) {
	h1 := log.MockHandler{
		Level: log.InfoLevel,
	}

	h2 := log.MockHandler{
		Level:     log.ErrorLevel,
		AlwaysErr: true,
	}

	h3 := log.MockHandler{
		Level:     log.VerboseLevel,
		AlwaysErr: true,
	}

	m := multi.New(&h1, &h2, &h3)

	for _, e := range events {
		if m.Enabled(e.Level) {
			if err := m.Handle(e); !errors.Is(err, log.ErrMockHandler) {
				t.Errorf(
					"handle error mismatch (@%s) => expected=%s, got=%s",
					e.Message,
					log.ErrMockHandler,
					err)
			}
		}
	}

	if h2.EventCount != 0 {
		t.Errorf("incorrect number of events on h2(@ErrorLevel) expected=0, got=%d",
			h2.EventCount)
	}
	if h2.HandleCount != 5 {
		t.Errorf("incorrect number of Handle() calls on h2(@ErrorLevel) expected=5, got=%d",
			h2.EventCount)
	}

	if h3.EventCount != 0 {
		t.Errorf("incorrect number of events on h1(@ErrorLevel) expected=0, got=%d",
			h3.EventCount)
	}
	if h3.HandleCount != 14 {
		t.Errorf("incorrect number of events on h1(@ErrorLevel) expected=14, got=%d",
			h3.EventCount)
	}

	if err := m.Flush(); !errors.Is(err, log.ErrMockHandler) {
		t.Errorf(
			"flush error mismatch => expected=%s, got=%s",
			log.ErrMockHandler,
			err)
	}
}
