package log

import (
	"io"
	stdlog "log"
	"sync"
)

// Compile time check for bridge.
// This ensures that bridge implements [io.Writer] interface.
var _ io.Writer = &bridge{}

// default bridge is an global bridge instance.
var defaultBridge = bridge{
	level:  LevelInfo,
	logger: nil,
}

// Bridges stdlib logger.
// this implements an io.Writer interface
// which transforms the bytes written by
// stdlib logger to an [Event] and let configured
// handler handle it.
type bridge struct {
	mu     sync.Mutex
	logger *Logger
	level  Level
}

// Write implements io.Writer.
func (br *bridge) Write(b []byte) (int, error) {
	br.mu.Lock()
	defer br.mu.Unlock()
	// var message string
	if br.logger == nil {
		return 0, ErrLoggerInvalid
	}
	// Split "a/b/c/file.go:23: message" into "d.go", "23", and "message".
	// if parts := bytes.SplitN(b, []byte{':'}, 3); len(parts) != 3 || len(parts[0]) < 1 || len(parts[2]) < 1 {
	// 	message = fmt.Sprintf("bad log format: %s", b)
	// } else {
	// 	file := string(parts[0])
	// 	message = string(parts[2][1:]) // skip leading space
	// 	line, err = strconv.Atoi(string(parts[1]))
	// 	if err != nil {
	// 		text = fmt.Sprintf("bad line number: %s", b)
	// 		line = 1
	// 	}
	// }
	return len(b), nil
}

// Bridges standard library's default logger to given
// Logger at specified Level.
//   - This can only map all events from stdlib logger to a single Level.
//   - This cannot prevent standard library logger from calling panic
//     on Panic/Panicf/Panicln events.
//   - This cannot prevent standard library logger from calling os.Exit
//     on Fatal/Fatalf/Fatalln events.
func Bridge(logger *Logger, level Level) error {
	if logger == nil {
		return ErrLoggerInvalid
	}
	defaultBridge.mu.Lock()
	defer defaultBridge.mu.Unlock()

	stdlog.Default().SetOutput(&defaultBridge)
	stdlog.Default().SetPrefix("")
	if logger.Caller() {
		stdlog.Default().SetFlags(stdlog.Llongfile)
	} else {
		stdlog.Default().SetFlags(0)
	}
	return nil
}
