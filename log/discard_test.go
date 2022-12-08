package log_test

import (
	"errors"
	"testing"

	"github.com/tprasadtp/pkg/log"
	"github.com/tprasadtp/pkg/log/internal/testdata"
)

func TestDiscardHandler(t *testing.T) {
	h := log.NewNoOpHandler(log.InfoLevel)
	var handleInvokeCount int
	for _, e := range testdata.GetEvents() {
		if h.Enabled(e.Level) {
			handleInvokeCount++
			if err := h.Write(e); err != nil {
				t.Errorf("handler returned error(%e), event=%s", err, e.Message)
			}
		}
	}

	if handleInvokeCount != testdata.I {
		t.Errorf("incorrect Enabled(), Handle() should be invoked=%d times, but got=%d",
			testdata.I,
			handleInvokeCount)
	}
	// flush and close handler
	if err := h.Flush(); err != nil {
		t.Errorf("handler flush returned error(%e)", err)
	}
	if err := h.Close(); err != nil {
		t.Errorf("handler flush returned error(%e)", err)
	}

	// write to closed handler must error
	if err := h.Write(log.Event{Level: log.InfoLevel}); !errors.Is(err, log.ErrHandlerClosed) {
		t.Errorf("write on closed handler returned unexpected error (%v)", err)
	}

	// flush on closed handler must error
	if err := h.Flush(); !errors.Is(err, log.ErrHandlerClosed) {
		t.Errorf("flush on closed handler returned unexpected error (%v)", err)
	}

	// close on already closed handler must error
	if err := h.Close(); !errors.Is(err, log.ErrHandlerClosed) {
		t.Errorf("close on closed handler returned unexpected error (%v)", err)
	}
}
