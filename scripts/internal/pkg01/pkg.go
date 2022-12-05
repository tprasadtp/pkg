package pkg01

import "github.com/tprasadtp/pkg/log"

func logHelperForLoggingSomething(logger *log.Logger) {
	log.Helper()
	logger.Critical("INSIDE HELPER IN A PKG")
}

func LogSomething(logger *log.Logger) {
	logHelperForLoggingSomething(logger)
	logger.Info("INSIDE A PKG")
}
