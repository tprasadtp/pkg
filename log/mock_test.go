package log_test

import (
	"errors"
	"testing"

	"github.com/tprasadtp/pkg/log"
)

func TestMockHandlerLevel(t *testing.T) {
	h := log.MockHandler{
		Level: log.InfoLevel,
	}

	type testCase struct {
		Name   string
		Level  log.Level
		Expect bool
	}

	tt := []testCase{
		{
			Name:   "DebugLevel",
			Level:  log.DebugLevel,
			Expect: false,
		},
		{
			Name:   "VerboseLevel",
			Level:  log.VerboseLevel,
			Expect: false,
		},
		{
			Name:   "InfoLevel",
			Level:  log.InfoLevel,
			Expect: true,
		},
		{
			Name:   "SuccessLevel",
			Level:  log.SuccessLevel,
			Expect: true,
		},
		{
			Name:   "NoticeLevel",
			Level:  log.NoticeLevel,
			Expect: true,
		},
		{
			Name:   "WarningLevel",
			Level:  log.WarningLevel,
			Expect: true,
		},
		{
			Name:   "ErrorLevel",
			Level:  log.ErrorLevel,
			Expect: true,
		},
		{
			Name:   "CriticalLevel",
			Level:  log.CriticalLevel,
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

func TestMockHandlerCustomLevelFunc(t *testing.T) {
	h := log.MockHandler{
		EnabledFunc: func(l log.Level) bool { return l == log.ErrorLevel },
	}

	type testCase struct {
		Name   string
		Level  log.Level
		Expect bool
	}

	tt := []testCase{
		{
			Name:   "DebugLevel",
			Level:  log.DebugLevel,
			Expect: false,
		},
		{
			Name:   "VerboseLevel",
			Level:  log.VerboseLevel,
			Expect: false,
		},
		{
			Name:   "InfoLevel",
			Level:  log.InfoLevel,
			Expect: false,
		},
		{
			Name:   "SuccessLevel",
			Level:  log.SuccessLevel,
			Expect: false,
		},
		{
			Name:   "NoticeLevel",
			Level:  log.NoticeLevel,
			Expect: false,
		},
		{
			Name:   "WarningLevel",
			Level:  log.WarningLevel,
			Expect: false,
		},
		{
			Name:   "ErrorLevel",
			Level:  log.ErrorLevel,
			Expect: true,
		},
		{
			Name:   "CriticalLevel",
			Level:  log.CriticalLevel,
			Expect: false,
		},
		{
			Name:   "CriticalLevel",
			Level:  log.CriticalLevel,
			Expect: false,
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

func TestMockHandlerHandle(t *testing.T) {
	t.Run("InfoLevel", func(t *testing.T) {
		h := log.MockHandler{
			Level: log.InfoLevel,
		}
		var err error

		for _, e := range events {
			if h.Enabled(e.Level) {
				err = h.Handle(e)
				if err != nil {
					t.Errorf("handler returned error(%e), event=%+v", err, e)
				}
			}
		}
		if h.EventCount != 12 {
			t.Errorf("handler incorrect Events. expected=12, got=%d", h.EventCount)
		}
		if h.HandleCount != 12 {
			t.Errorf("handler incorrect HandleCount. expected=12, got=%d", h.HandleCount)
		}

		err = h.Flush()
		if err != nil {
			t.Errorf("handler flush error(%e)", err)
		}
		if h.EventCount != 0 {
			t.Errorf("handler did not flush events")
		}
	})
}

func TestMockHandlerHandleAlwaysErr(t *testing.T) {
	h := log.MockHandler{
		Level:     log.InfoLevel,
		AlwaysErr: true,
	}

	for _, e := range events {
		if h.Enabled(e.Level) {
			if err := h.Handle(e); !errors.Is(err, log.ErrMockHandler) {
				t.Errorf("handler.Handle() AlwaysErr=true did not return %s error",
					log.ErrMockHandler)
			}
		}
	}
	if h.HandleCount != 12 {
		t.Errorf("handler incorrect HandleCount. expected=2, got=%d", h.HandleCount)
	}

	if h.EventCount != 0 {
		t.Errorf("handler incorrect Events. expected=0, got=%d", h.EventCount)
	}

	if err := h.Flush(); !errors.Is(err, log.ErrMockHandler) {
		t.Errorf("handler.Flush() AlwaysErr=true dif not return an error")
	}
}
