package scripts

import "runtime"

func Caller(depth uint) {
	pc := make([]uintptr, 10)

	// Skip two extra frames to account for this function
	// and runtime.Callers itself.
	//nolint:gomnd // ignore this magic number.
	n := runtime.Callers(int(depth+2), pc)
	frames := runtime.CallersFrames(pc[:n])
	for i := 0; i < 10; i++ {
		_, more := frames.Next()
		// 	// We ran out of frames (This implies bug in log package)
		if !more {
			break
		}
	}
}

func GetPC(depth int) (string, int) {
	var pcs [10]uintptr
	runtime.Callers(depth, pcs[:])
	frames := runtime.CallersFrames(pcs[:])
	f, _ := frames.Next()
	return f.File, f.Line
}
