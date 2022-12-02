package logrotate

import (
	"os"
	"sync"
)

// Compile time check ensures that File always
// implements Write Closer.
// This will fail if multi.Handler does not
// implement log.Handler interface.
// var _ io.WriteCloser = &File{}

type File struct {
	// Filename is the file to write logs to.  Backup log files will be retained
	// in the same directory.  It uses <processname>-lumberjack.log in
	// os.TempDir() if empty.
	Name string
	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackups uint
	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSizeMB uint
	// underlying file
	osFile *os.File
	mu     sync.Mutex
}

// Implements io.Writer.
func (f *File) Write(b []byte) (int, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.osFile != nil {
		return f.osFile.Write(b)
	}
	return 0, os.ErrNotExist
}
