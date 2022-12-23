package logrotate

import (
	"io"
	"os"
	"sync"
)

// Compile time check ensures that File always
// implements Write Closer.
// This will fail if multi.Handler does not
// implement log.Handler interface.
var _ io.WriteCloser = &File{}

type File struct {
	// Filename is the file to write logs to.
	// Backup log files will be retained in the same directory.
	// If not specified empty, uses <processname>-log.log.
	// on Linux
	//  - if $LOGS_DIRECTORY is non empty and writable, it is used.
	//  - if $LOGS_DIRECTORY is not set, uses /var/log if it writable
	//  - if both above conditions are not specified, panics.
	// on Windows
	//  - $env:PROGRAMDATA is used if not specified or empty.
	FileName string
	// MaxBackups is the maximum number of old log files to retain.
	// The default is to retain all old log files
	// (though MaxAge may still cause them to get deleted.)
	MaxBackups uint
	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 MB.
	MaxSizeMB uint
	// underlying file
	file *os.File
	mu   sync.Mutex
}

// Implements io.Writer.
func (f *File) Write(b []byte) (int, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.file != nil {
		return 0, os.ErrInvalid
	}
	return f.file.Write(b)
}

// Implements io.Closer.
func (f *File) Close() error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.file == nil {
		return os.ErrInvalid
	}
	return f.file.Close()
}

func (f *File) Rotate() error {
	f.mu.Lock()
	defer f.mu.Unlock()
	return nil
}
