package multi_test

import (
	"errors"
	"testing"

	"github.com/tprasadtp/pkg/log"
	"github.com/tprasadtp/pkg/log/handlers/internal/testdata"
	"github.com/tprasadtp/pkg/log/handlers/multi"
)

func TestHandler(t *testing.T) {
	h1 := log.MockHandler{
		Level: log.InfoLevel,
	}

	h2 := log.MockHandler{
		Level: log.ErrorLevel,
	}

	m := multi.New(&h1, &h2)

	for _, e := range testdata.GetEvents() {
		if m.Enabled(e.Level) {
			if err := m.Write(e); err != nil {
				t.Errorf("handler returned error(%e), event=%s", err, e.Message)
			}
		}
	}

	if len(h1.Events) != testdata.I {
		t.Errorf("incorrect number of events on h1(@InfoLevel) expected=%d, got=%d",
			testdata.I,
			len(h1.Events),
		)
	}
	if len(h2.Events) != testdata.E {
		t.Errorf("incorrect number of events on h2(@ErrorLevel) expected=%d, got=%d",
			testdata.E,
			len(h2.Events),
		)
	}

	if err := m.Flush(); err != nil {
		t.Errorf("handler flush returned error(%e)", err)
	}

	if len(h1.Events) != 0 {
		t.Errorf("flush did not clear events on h1(@InfoLevel) expected=0, got=%d",
			len(h1.Events))
	}
	if len(h2.Events) != 0 {
		t.Errorf("flush did not clear events on h2(@ErrorLevel) expected=0, got=%d",
			len(h2.Events))
	}
}

func TestOneClosedHandler(t *testing.T) {
	h1 := log.MockHandler{
		Level: log.InfoLevel,
	}

	h2 := log.MockHandler{
		Level: log.ErrorLevel,
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

	if len(h1.Events) != 0 {
		t.Errorf("incorrect number of events on h1[closed](@InfoLevel) expected=0 got=%d",
			len(h1.Events),
		)
	}
	if len(h2.Events) != testdata.E {
		t.Errorf("incorrect number of events on h2(@ErrorLevel) expected=%d, got=%d",
			testdata.E,
			len(h2.Events),
		)
	}

	if err := m.Flush(); !errors.Is(err, log.ErrHandlerClosed) {
		t.Errorf("handler flush returned error(%e)", err)
	}

	if len(h1.Events) != 0 {
		t.Errorf("flush did not clear events on h1(@InfoLevel) expected=0, got=%d",
			len(h1.Events))
	}
	if len(h2.Events) != 0 {
		t.Errorf("flush did not clear events on h2(@ErrorLevel) expected=0, got=%d",
			len(h2.Events))
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

	if len(h2.Events) != 0 {
		t.Errorf("incorrect number of events on h2(@ErrorLevel) expected=0, got=%d",
			len(h2.Events))
	}
	if h2.HandleCount != 5 {
		t.Errorf("incorrect number of Handle() calls on h2(@ErrorLevel) expected=5, got=%d",
			len(h2.Events))
	}

	if len(h3.Events) != 0 {
		t.Errorf("incorrect number of events on h1(@ErrorLevel) expected=0, got=%d",
			len(h3.Events))
	}
	if h3.HandleCount != 14 {
		t.Errorf("incorrect number of events on h1(@ErrorLevel) expected=14, got=%d",
			len(h3.Events))
	}

	if err := m.Flush(); !errors.Is(err, log.ErrHandlerWrite) {
		t.Errorf(
			"flush error mismatch => expected=%s, got=%s",
			log.ErrHandlerWrite,
			err)
	}
}
