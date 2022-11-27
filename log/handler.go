package log

// A Handler describes the logging backend.
// To log to multiple backends simply wrap all the handlers in
// another handler.
type Handler interface {
	// Checks if handler is enabled at the specified Level.
	//  - This does not necessarily mean logs above this level are logged.
	//  - Handler may chose to only log a single level, for example a handler for
	//    error tracking services like [Sentry] or [Cloud Trace] might only
	//    process events at ERROR level.
	//
	// [Sentry]: https://sentry.io
	// [Cloud Trace]: https://cloud.google.com/trace/docs/setup/go
	Enabled(l Level) bool

	// Writes a Log Entry.
	//  - Depending on implementation, writes may be buffered/batched.
	//  - Please note that this is ONLY called if Enabled returns true.
	//  - Implementations SHOULD return error to when handler is
	//    not initialized or closed.
	//  - It is responsibility of the implementation to be thread safe.
	//    Logger WILL NOT handle thread safety.
	Handle(e Event) error

	// Flushes pending entries in the buffer.
	//  - Its up to the handler to implement timeouts,
	//    However it is highly encouraged to do so.
	//  - Panic/Panicf and method on the Logger will call this automatically.
	//  - It is NOT an error if there are no pending entries to flush.
	Flush() error
}
