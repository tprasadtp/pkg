//go:build ignore

package main

import (
	"context"

	"github.com/tprasadtp/pkg/log"
)

var logger = log.New(&log.ConsoleHandler{})

func print(message string) {
	log.Helper()
	logger.Info(message)
}

func main() {
	print("PRINT WITH HELPER --> LINE 15")
	logger.WithNamespace("namespace-1").WithCtx(context.Background()).WithError(log.ErrHandlerWrite).Error("NESTED ERROR")
	logger.WithNamespace("namespace-1").WithFields(
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
