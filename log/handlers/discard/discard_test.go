package discard_test

import (
	"testing"

	"github.com/tprasadtp/pkg/log"
	"github.com/tprasadtp/pkg/log/handlers/discard"
)

func TestDiscardHandler(t *testing.T) {
	h := discard.New(log.InfoLevel)
	var handleInvokeCount int
	for _, e := range events {
		if h.Enabled(e.Level) {
			handleInvokeCount++
			if err := h.Handle(e); err != nil {
				t.Errorf("handler returned error(%e), event=%s", err, e.Message)
			}
		}
	}

	if handleInvokeCount != 12 {
		t.Errorf("incorrect Enabled(), Handle() should be invoked=12 times, but got=%d", handleInvokeCount)
	}

	if err := h.Flush(); err != nil {
		t.Errorf("handler flush returned error(%e)", err)
	}
}
