package mock_test

import (
	"errors"
	"testing"

	"github.com/tprasadtp/pkg/log"
	"github.com/tprasadtp/pkg/log/internal/mock"
	"github.com/tprasadtp/pkg/log/internal/testdata"
)

func TestMockHandlerEnabled(t *testing.T) {
	h := mock.Handler{
		Level: log.LevelInfo,
	}

	type testCase struct {
		Name   string
		Level  log.Level
		Expect bool
	}

	tt := []testCase{
		{
			Name:   "DebugLevel",
			Level:  log.LevelDebug,
			Expect: false,
		},
		{
			Name:   "VerboseLevel",
			Level:  log.LevelVerbose,
			Expect: false,
		},
		{
			Name:   "InfoLevel",
			Level:  log.LevelInfo,
			Expect: true,
		},
		{
			Name:   "SuccessLevel",
			Level:  log.LevelSuccess,
			Expect: true,
		},
		{
			Name:   "NoticeLevel",
			Level:  log.LevelNotice,
			Expect: true,
		},
		{
			Name:   "WarningLevel",
			Level:  log.LevelWarning,
			Expect: true,
		},
		{
			Name:   "ErrorLevel",
			Level:  log.LevelError,
			Expect: true,
		},
		{
			Name:   "CriticalLevel",
			Level:  log.LevelCritical,
			Expect: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			value := h.Enabled(tc.Level)
			if tc.Expect != value {
				t.Errorf("at %s => got=%t want=%t", tc.Name, value, tc.Expect)
			}
		})
	}
}

func TestMockHandler(t *testing.T) {
	t.Run("InfoLevel", func(t *testing.T) {
		h := mock.Handler{
			Level: log.LevelInfo,
		}
		// Write to handler
		for _, e := range testdata.GetEvents() {
			if h.Enabled(e.Level) {
				if err := h.Write(&e); err != nil {
					t.Errorf("handler returned error(%e), event=%+v", err, e)
				}
			}
		}
		if len(h.Events) != testdata.I {
			t.Errorf("handler incorrect Events. expected=12, got=%d", len(h.Events))
		}
		// Flush Handler
		if err := h.Flush(); err != nil {
			t.Errorf("handler flush error(%e)", err)
		}
		if len(h.Events) != 0 {
			t.Errorf("handler did not flush events")
		}

		// Close Handler
		if err := h.Close(); err != nil {
			t.Errorf("handler close error(%e)", err)
		}
		if len(h.Events) != 0 {
			t.Errorf("handler close not flush events")
		}

		// Write to already closed handler
		for _, e := range testdata.GetEvents() {
			if h.Enabled(e.Level) {
				if err := h.Write(&e); !errors.Is(err, log.ErrHandlerClosed) {
					t.Errorf("handler(closed) invalid error => got=(%s), expected=(%s)",
						err,
						log.ErrHandlerClosed,
					)
				}
			}
		}
		if len(h.Events) != 0 {
			t.Errorf("handler(closed) incorrect events. expected=0, got=%d",
				len(h.Events))
		}
		// Events are not written but handler is invoked.
		if h.WriteCalls != 24 {
			t.Errorf("handler(closed) incorrect WriteCalls. expected=24, got=%d",
				h.WriteCalls)
		}

		// Flush already closed Handler
		if err := h.Flush(); !errors.Is(err, log.ErrHandlerClosed) {
			t.Errorf("flushing already closed handler => got=(%s), expected=(%s)",
				err,
				log.ErrHandlerClosed,
			)
		}

		// Close already closed Handler
		if err := h.Close(); !errors.Is(err, log.ErrHandlerClosed) {
			t.Errorf("closing already closed handler => got=(%s), expected=(%s)",
				err,
				log.ErrHandlerClosed,
			)
		}
	})
}

func TestMockHandlerHandleAlwaysErr(t *testing.T) {
	h := mock.Handler{
		Level:     log.LevelInfo,
		AlwaysErr: true,
	}

	for _, e := range testdata.GetEvents() {
		if h.Enabled(e.Level) {
			if err := h.Write(&e); !errors.Is(err, log.ErrHandlerWrite) {
				t.Errorf("handler.Handle() AlwaysErr=true did not return %s error",
					log.ErrHandlerWrite)
			}
		}
	}
	if h.WriteCalls != 12 {
		t.Errorf("handler incorrect WriteCalls. expected=2, got=%d", h.WriteCalls)
	}

	if len(h.Events) != 0 {
		t.Errorf("handler incorrect Events. expected=0, got=%d", len(h.Events))
	}

	if err := h.Flush(); !errors.Is(err, log.ErrHandlerWrite) {
		t.Errorf("handler.Flush() AlwaysErr=true dif not return an error")
	}
}