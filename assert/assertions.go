// SPDX-FileCopyrightText: 2012 Mat Ryer
// SPDX-FileCopyrightText: 2012 Tyler Bunnell
// SPDX-FileCopyrightText: 2022 Prasad Tengse <tprasadtp@users.noreply.github.com>
// SPDX-License-Identifier: MIT OR Apache-2.0

package assert

import (
	"fmt"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"testing"

	"github.com/tprasadtp/pkg/assert/internal/diff"
)

var testNameRegexp = regexp.MustCompile(`\.(Test[\p{L}_\p{N}]*)`)

// Assertions provides assertion methods
// and helps avoiding passing [testing.T]
// every time.
type Assertions struct {
	t *testing.T
}

// New returns a new Assertions
func New(t *testing.T) *Assertions {
	return &Assertions{
		t: t,
	}
}

// callerInfo is weird, as assert can be used in helper libs which
// in turn use assert package, so we trace the call tree till we find
// test func, and print it and line number which called assert eventually.
func callerInfo(t *testing.T) string {
	var file string = "<unknown>"
	var funcName string = "<unknown>"
	var line int
	var pc uintptr
	var ok bool

	for i := 0; ; i++ {
		pc, file, line, ok = runtime.Caller(i)
		if !ok {
			break
		}

		f := runtime.FuncForPC(pc)
		if f == nil {
			break
		}

		funcName = f.Name()

		if testNameRegexp.FindStringSubmatch(funcName) != nil {
			break
		}
		t.Logf("Line=%d, File=%s, Func=%s Iterator=%d", line, file, funcName, i)
	}

	return fmt.Sprintf("[ %s:%d ] - ", file, line)
}

// Implements asserts that an object is implemented by the specified interface.
//
//	assert.Implements(t, (*MyInterface)(nil), new(MyObject), "MyObject")
func Implements(t *testing.T, interfaceObject any, object any, message ...string) bool {
	interfaceType := reflect.TypeOf(interfaceObject).Elem()
	return Equal(t, reflect.TypeOf(object).Implements(interfaceType), fmt.Sprintf("%sObject must implement %s. %s", callerInfo(t), interfaceType, message))
}

// IsType asserts that the specified objects are of the same type.
func IsType(t *testing.T, expectedType any, object any, message ...string) bool {
	return Equal(t, reflect.TypeOf(object), reflect.TypeOf(expectedType), fmt.Sprintf("Object expected to be of type %s, but was %s. %s", reflect.TypeOf(expectedType), reflect.TypeOf(object), message))
}

// Equal asserts that two objects are equal.
//
//	assert.Equal(t, 123, 123, "123 and 123 should be equal")
//
// Returns whether the assertion was successful (true) or not (false).
func Equal(t *testing.T, got, expect any, message ...string) bool {
	diff := diff.Diff(got, expect)
	if len(diff) > 0 {
		t.Errorf("%s must be equal: %s", callerInfo(t), diff)
		return false
	}
	return true
}

// NotNil asserts that the specified object is not nil.
//
//	assert.NotNil(t, err, "err should be something")
//
// Returns whether the assertion was successful (true) or not (false).
func NotNil(t *testing.T, object any, message ...string) bool {

	var success bool = true

	if object == nil {
		success = false
	} else if reflect.ValueOf(object).IsNil() {
		success = false
	}

	if !success {
		t.Errorf("%sExpected not to be nil. %s", callerInfo(t), message)
	}

	return success
}

// Nil asserts that the specified object is nil.
//
//	assert.Nil(t, err, "err should be nothing")
//
// Returns whether the assertion was successful (true) or not (false).
func Nil(t *testing.T, object any, message ...string) bool {

	if object == nil {
		return true
	} else if reflect.ValueOf(object).IsNil() {
		return true
	}

	t.Errorf("%sExpected to be nil but was %#v. %s", callerInfo(t), object, message)

	return false
}

// NotEqual asserts that the specified values are NOT equal.
//
//	assert.NotEqual(t, obj1, obj2, "two objects shouldn't be equal")
//
// Returns whether the assertion was successful (true) or not (false).
func NotEqual(t *testing.T, a, b any, message ...string) bool {

	diff := diff.Diff(a, b)
	if len(diff) == 0 {
		t.Errorf("%s diff=%s, %s", callerInfo(t), diff, message)
		return false
	}
	return true

}

// Contains asserts that the specified string contains the specified substring.
//
//	assert.Contains(t, "Hello World", "World", "But 'Hello World' does contain 'World'")
//
// Returns whether the assertion was successful (true) or not (false).
func Contains(t *testing.T, s, contains string, message ...string) bool {

	if !strings.Contains(s, contains) {
		t.Errorf("%s %s '%s' does not contain '%s'", callerInfo(t), message, s, contains)
		return false
	}

	return true
}

// NotContains asserts that the specified string does NOT contain the specified substring.
//
//	assert.NotContains(t, "Hello World", "Earth", "But 'Hello World' does NOT contain 'Earth'")
//
// Returns whether the assertion was successful (true) or not (false).
func NotContains(t *testing.T, s, contains string, message ...string) bool {

	if strings.Contains(s, contains) {
		t.Errorf("%s%s '%s' should not contain '%s'", callerInfo(t), message, s, contains)
		return false
	}

	return true

}

// didPanic returns true if the function passed to it panics. Otherwise, it returns false.
func didPanic(f func()) bool {
	var didPanic bool = false
	func() {

		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()

		// call the target function
		f()

	}()

	return didPanic
}

// Panics asserts that the code inside the specified PanicTestFunc panics.
//
//	assert.Panics(t, func(){
//	  GoCrazy()
//	}, "Calling GoCrazy() should panic")
//
// Returns whether the assertion was successful (true) or not (false).
func Panics(t *testing.T, f func(), message ...string) bool {
	if didPanic(f) {
		return true
	}
	t.Errorf("Func MUST panic but didn't. %s", message)
	return false
}

// NotPanics asserts that the code inside the specified PanicTestFunc does NOT panic.
//
//	assert.NotPanics(t, func(){
//	  RemainCalm()
//	}, "Calling RemainCalm() should NOT panic")
//
// Returns whether the assertion was successful (true) or not (false).
func NotPanics(t *testing.T, f func(), message ...string) bool {
	if !didPanic(f) {
		return true
	}
	t.Errorf("Func MUST NOT panic but did. %s", message)
	return false
}
