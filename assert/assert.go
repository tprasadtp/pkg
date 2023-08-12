package assert

import (
	"errors"
	"fmt"
	"testing"
)

// wrapper to handle format messages.
func msgf(fallback string, args ...any) string {
	if len(args) == 0 {
		return fallback
	}

	if s, ok := (args[0]).(string); ok {
		return fmt.Sprintf(s, args[1:])
	}
	return fallback
}

// Panics asserts that function f panics.
//
// args can be used to customize the error message.
//
//	parser := func(){panic("this function panics")}
//	assert.Panics(t, "%s => parser function should panic, it did not", t.Name())
func Panics(t testing.TB, f func(), args ...any) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Errorf(msgf("Expected panic, but did not", args...))
		}
	}()
	f()
}

// NotPanics asserts that the given function does not panic.
func NotPanics(t testing.TB, f func(), args ...any) {
	t.Helper()
	defer func() {
		if err := recover(); err != nil {
			fallback := fmt.Sprintf("Expcted not to panic: %v", err)
			t.Errorf(msgf(fallback, args...))
		}
	}()
	f()
}

// IsError asserts than any error in "err"'s tree matches "target".
//
//	_, err := someFunction()
//	assert.IsError(t, err, ExpectedErr)
func IsError(t testing.TB, err, target error, args ...any) {
	t.Helper()
	if errors.Is(err, target) {
		return
	}
	fallback := fmt.Sprintf("Error tree %q should contain error %q", err, target)
	t.Error(msgf(fallback, args...))
}

// IsError asserts than any error in "err"'s tree matches "target".
//
//	_, err := someFunction()
//	assert.NotIsError(t, err, NotExpectedErr)
func NotIsError(t testing.TB, err, target error, args ...any) {
	t.Helper()
	if !errors.Is(err, target) {
		return
	}
	fallback := fmt.Sprintf("Error tree %q should NOT contain error %q", err, target)
	t.Error(msgf(fallback, args...))
}

// Errors asserts that an error is not nil.
func Errors(t testing.TB, err error, args ...any) {
	t.Helper()
	if err != nil {
		return
	}
	t.Errorf(msgf("Expected an error, but got nil", args...))
}

// NoErrors asserts that an error is nil.
func NoErrors(t testing.TB, err error, args ...any) {
	if err == nil {
		return
	}
	t.Helper()
	fallback := fmt.Sprintf("Expected no error, but got: %s", err)
	t.Errorf(msgf(fallback, args...))
}

// True asserts that an expression is true.
func True(t testing.TB, ok bool, args ...any) {
	if ok {
		return
	}
	t.Helper()
	t.Fatal(msgf("Expected expression to be true", args...))
}

// False asserts that an expression is false.
func False(t testing.TB, ok bool, args ...any) {
	if !ok {
		return
	}
	t.Helper()
	t.Fatal(msgf("Expected expression to be false", args...))
}
