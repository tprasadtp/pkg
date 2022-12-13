//go:build ignore

package main

import (
	"fmt"

	"github.com/tprasadtp/pkg/log"
)

// Compile time check for handler.
// This will fail if discard.Handler does not
// implement Handler interface.
var _ log.Handler = &ConsoleHandler{}

func NewConsole() *ConsoleHandler {
	return &ConsoleHandler{
		level: log.LevelTrace,
	}
}

type ConsoleHandler struct {
	level  log.Level
	closed bool
}

// Enabled Checks if given level is enabled.
func (h *ConsoleHandler) Enabled(level log.Level) bool {
	return true
}

// Write the Event.
func (h *ConsoleHandler) Write(event log.Event) error {
	if h.closed {
		return log.ErrHandlerClosed
	}
	fmt.Printf("%+v\n\n", event)
	return nil
}

// Flushes the handler.
func (h *ConsoleHandler) Flush() error {
	if h.closed {
		return log.ErrHandlerClosed
	}
	return nil
}

// Closes the handler.
func (h *ConsoleHandler) Close() error {
	if h.closed {
		return log.ErrHandlerClosed
	}
	h.closed = true
	return nil
}

var logger = log.New(&ConsoleHandler{})

func print(message string) {
	log.Helper()
	logger.Info(message)
}

func main() {
	// print("PRINT WITH HELPER --> LINE 19")
	// pkg01.LogSomething(&logger)

	// logger.WithNamespace("namespace-1").WithError(log.ErrHandlerWrite).Error("NESTED ERROR")
	l1 := logger.With(
		log.F("root-key-01", "root-value-01"),
		log.F("root-key-02", "root-value-02"))
	l2 := l1.With(
		log.F("root-key-03", "root-value-03"),
		log.F("root-key-04", "root-value-04"),
		log.F("root-key-05", "root-value-05"),
		log.F("root-key-06", "root-value-06"),
		log.F("root-key-07", "root-value-07"),
		log.F("root-key-08", "root-value-08"),
		log.F("root-key-09", "root-value-09"),
		log.F("root-key-10", "root-value-10"),
		log.F("root-key-01", "root-value-01"),
		log.F("root-key-02", "root-value-02"),
		log.F("root-key-03", "root-value-03"),
		log.F("root-key-04", "root-value-04"),
		log.F("root-key-05", "root-value-05"),
		log.F("root-key-06", "root-value-06"),
		log.F("root-key-07", "root-value-07"),
		log.F("root-key-08", "root-value-08"),
		log.F("root-key-09", "root-value-09"),
		log.F("root-key-10", "root-value-10"),
		log.F("root-key-01", "root-value-01"),
		log.F("root-key-02", "root-value-02"),
		log.F("root-key-03", "root-value-03"),
		log.F("root-key-04", "root-value-04"),
		log.F("root-key-05", "root-value-05"),
		log.F("root-key-06", "root-value-06"),
		log.F("root-key-07", "root-value-07"),
		log.F("root-key-08", "root-value-08"),
		log.F("root-key-09", "root-value-09"),
		log.F("root-key-10", "root-value-10"))
	l3 := l2.With(
		log.F("must-fit-in-buf-01", 1),
		log.F("must-fit-in-buf-02", "root-value-02"))
	l4 := l3.With(
		log.F("must-fit-in-buf-01", 1),
		log.F("must-fit-in-buf-02", "root-value-02"))
	l5 := l4.With(
		log.F("must-fit-in-buf-01", 1),
		log.F("must-fit-in-buf-02", "root-value-02"))
	l6 := l5.With(
		log.F("must-fit-in-buf-01", 1),
		log.F("must-fit-in-buf-02", "root-value-02"))
	l7 := l6.With(
		log.F("must-fit-in-buf-01", 1),
		log.F("must-fit-in-buf-02", "root-value-02"))

	// logger.Debug("DEBUG")
	// logger.Verbose("VERBOSE")
	// logger.Info("INFO")
	// logger.Success("SUCCESS")
	// logger.Warning("WARNING")
	// logger.Error("ERROR")
	// logger.Critical("CRITICAL")

	l7.Fatal("MUST EXIT")
}
