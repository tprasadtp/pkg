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
		log.F("root-key-10", "root-value-10"),
		log.M("map-01", log.F("map-01-key-01", "map-01-value-01")),
		log.M("map-02", log.F("map-02-key-01", "map-02-value-01")),
		log.M("map-03", log.F("map-03-key-01", "map-03-value-01")),
		log.M("map-04",
			log.F("map-04-key-01", "map-04-value-01"),
			log.F("map-04-key-02", "map-04-value-02"),
			log.F("map-04-key-03", "map-04-value-03"),
			log.F("map-04-key-04", "map-04-value-04"),
			log.F("map-04-key-05", "map-04-value-05"),
			log.F("map-04-key-06", "map-04-value-06"),
			log.F("map-04-key-07", "map-04-value-07"),
			log.F("map-04-key-08", "map-04-value-08"),
			log.F("map-04-key-09", "map-04-value-09"),
			log.F("map-04-key-10", "map-04-value-10"),
			log.F("map-04-key-11", "map-04-value-11"),
			log.F("map-04-key-12", "map-04-value-12")),
	)
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
