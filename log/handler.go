package log

// A Handler describes the logging backend.
// To log to multiple backends simply wrap all the handlers in
// another handler.
type Handler interface {
	// Checks if handler is enabled at the specified Level.
	//  - This does not necessarily mean logs above this level are logged.
	//    Handler may chose to only log a single level,
	//	- Enabled is called early, before any arguments are processed,
	//    to save effort if the log event should be discarded.
	Enabled(l Level) bool

	// Writes a Log Event.
	//  - Depending on implementation, writes may be buffered/batched.
	//  - Please note that this is ONLY called if Enabled(e.Level) returns true.
	//  - Implementations SHOULD return error to when handler is
	//    not initialized or closed.
	//  - It is responsibility of the implementation to be concurrent safe.
	Write(e Event) error

	// Writes pending entries in the buffer to disk/network.
	//  - If handler is a network handler it MUST write all pending entries to
	//    network endpoint.
	//  - If handler is a file handle, it MUST flush entries to disk AND call fsync.
	//  - Its up to the handler to implement timeouts,
	//    However it is highly encouraged to do so.
	//  - Panic/Panicf and method on the Logger will call this automatically.
	Flush() error

	// Closes the underlying file or network connection or socket after
	// writing pending entries in the buffer to file/network.
	//  - Logger will not invoke Flush method on the handler,
	//    it is up to the handler to flush its buffers before closing stream.
	//  - It is up to the handler to implement timeouts,
	//    However it is highly encouraged to do so.
	//  - Any calls to Handle() or Flush() after calling this MUST result in error.
	//  - Multiple calls to Close() MUST return an error.
	Close() error
}
