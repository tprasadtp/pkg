//go:build ignore

package main

import (
	"context"
	"fmt"

	"github.com/tprasadtp/pkg/log"
	"github.com/tprasadtp/pkg/scripts/internal/pkg01"
)

// Compile time check for handler.
// This will fail if discard.Handler does not
// implement Handler interface.
var _ log.Handler = &ConsoleHandler{}

func NewConsole() *ConsoleHandler {
	return &ConsoleHandler{
		level: log.DebugLevel,
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
	print("PRINT WITH HELPER --> LINE 19")
	pkg01.LogSomething(&logger)
	fmt.Println("-----------------------------------------------------------")

	logger.WithNamespace("namespace-1").WithCtx(context.Background()).WithError(log.ErrHandlerWrite).Error("NESTED ERROR")
	logger.WithNamespace("namespace-1").With(
		log.M("map-1", log.F("nested-field-1", "nested-value-1")),
		log.M("map-2", log.F("nested-field-2", "nested-value-2")),
		log.F("root-field-01", "root-value-01"),
	).Info("NESTED FIELDS")
	logger.Debug("DEBUG")
	logger.Verbose("VERBOSE")
	logger.Info("INFO")
	logger.Success("SUCCESS")
	logger.Warning("WARNING")
	logger.Error("ERROR")
	logger.Critical("CRITICAL")

	logger.Fatal("MUST EXIT")
}
