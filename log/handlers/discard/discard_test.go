package discard_test

import (
	"testing"

	"github.com/tprasadtp/pkg/log"
	"github.com/tprasadtp/pkg/log/handlers/discard"
	"github.com/tprasadtp/pkg/log/handlers/internal/testdata"
)

func TestDiscardHandler(t *testing.T) {
	h := discard.New(log.InfoLevel)
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

	if err := h.Flush(); err != nil {
		t.Errorf("handler flush returned error(%e)", err)
	}
}
