// stdlib plugin transforms events logged using standard library's
// log package to log.Events.
//
// Due to limitation of standard library package,
//   - This can only map all events from stdlib logger to a single level.
//   - This cannot prevent standard library logger from calling panic
//     on Panic/Panicf/Panicln events.
//   - This cannot prevent standard library logger from calling os.Exit
//     on Fatal/Fatalf/Fatalln events.
//   - This can only work on default stdlib logger [log.Default].
package stdlog

import (
	"io"
	stdliblog "log"
	"sync"

	"github.com/tprasadtp/pkg/log"
)

// Compile time check for bridge.
// this ensures that bridge implements io.Writer.
var _ io.Writer = &bridge{}

// default bridge.
var defaultBridge = bridge{
	level:  log.InfoLevel,
	logger: nil,
}

// Bridge implements bridge interfaces.
type bridge struct {
	mu     sync.Mutex
	logger *log.Logger
	level  log.Level
}

// Write implements io.Writer.
func (br *bridge) Write(b []byte) (int, error) {
	if br.logger == nil {
		return len(b), nil
	}
	// Split "a/b/c/file.go:23: message" into "d.go", "23", and "message".
	// if parts := bytes.SplitN(b, []byte{':'}, 3); len(parts) != 3 || len(parts[0]) < 1 || len(parts[2]) < 1 {
	// 	text = fmt.Sprintf("bad log format: %s", b)
	// } else {
	// 	file = string(parts[0])
	// 	text = string(parts[2][1:]) // skip leading space
	// 	line, err = strconv.Atoi(string(parts[1]))
	// 	if err != nil {
	// 		text = fmt.Sprintf("bad line number: %s", b)
	// 		line = 1
	// 	}
	// }
	return 0, nil
}

// Bridges standard library's default logger to given
// Logger at specified log Level.
//   - This can only map all events from stdlib logger to a single level.
//   - This cannot prevent standard library logger from calling panic
//     on Panic/Panicf/Panicln events.
//   - This cannot prevent standard library logger from calling os.Exit
//     on Fatal/Fatalf/Fatalln events.
func SetupBridge(logger *log.Logger, level log.Level) error {
	if logger == nil {
		return log.ErrLoggerInvalid
	}
	defaultBridge.mu.Lock()
	defer defaultBridge.mu.Unlock()

	stdliblog.Default().SetOutput(&defaultBridge)
	return nil
}
