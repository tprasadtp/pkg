// Copyright (c) 2012 Mat Ryer
// Copyright (c) 2012 Tyler Bunnell
// Copyright (c) 2022 Prasad Tengse <tprasadtp@users.noreply.github.com>
// SPDX-License-Identifier: MIT OR Apache-2.0

package assert

import (
	"testing"
)

// AssertionTesterInterface defines an interface to be used for testing assertion methods
type AssertionTesterInterface interface {
	TestMethod()
}

// AssertionTesterConformingObject is an object that conforms to the AssertionTesterInterface interface
type AssertionTesterConformingObject struct {
}

func (a *AssertionTesterConformingObject) TestMethod() {
}

// AssertionTesterNonConformingObject is an object that does not conform to the AssertionTesterInterface interface
type AssertionTesterNonConformingObject struct {
}

func TestImplements(t *testing.T) {

	mockT := new(testing.T)

	if !Implements(mockT, (*AssertionTesterInterface)(nil), new(AssertionTesterConformingObject)) {
		t.Error("Implements method should return true: AssertionTesterConformingObject implements AssertionTesterInterface")
	}
	if Implements(mockT, (*AssertionTesterInterface)(nil), new(AssertionTesterNonConformingObject)) {
		t.Error("Implements method should return false: AssertionTesterNonConformingObject does not implements AssertionTesterInterface")
	}

}

func TestIsType(t *testing.T) {

	mockT := new(testing.T)

	if !IsType(mockT, new(AssertionTesterConformingObject), new(AssertionTesterConformingObject)) {
		t.Error("IsType should return true: AssertionTesterConformingObject is the same type as AssertionTesterConformingObject")
	}
	if IsType(mockT, new(AssertionTesterConformingObject), new(AssertionTesterNonConformingObject)) {
		t.Error("IsType should return false: AssertionTesterConformingObject is not the same type as AssertionTesterNonConformingObject")
	}

}

func TestEqual(t *testing.T) {

	mockT := new(testing.T)

	if !Equal(mockT, "Hello World", "Hello World") {
		t.Error("Equal should return true")
	}
	if !Equal(mockT, 123, 123) {
		t.Error("Equal should return true")
	}
	if !Equal(mockT, 123.5, 123.5) {
		t.Error("Equal should return true")
	}
	if !Equal(mockT, []byte("Hello World"), []byte("Hello World")) {
		t.Error("Equal should return true")
	}
	if !Equal(mockT, nil, nil) {
		t.Error("Equal should return true")
	}

}

func TestNotNil(t *testing.T) {

	mockT := new(testing.T)

	if !NotNil(mockT, new(AssertionTesterConformingObject)) {
		t.Error("NotNil should return true: object is not nil")
	}
	if NotNil(mockT, nil) {
		t.Error("NotNil should return false: object is nil")
	}

}

func TestNil(t *testing.T) {

	mockT := new(testing.T)

	if !Nil(mockT, nil) {
		t.Error("Nil should return true: object is nil")
	}
	if Nil(mockT, new(AssertionTesterConformingObject)) {
		t.Error("Nil should return false: object is not nil")
	}

}

func TestNotEqual(t *testing.T) {

	mockT := new(testing.T)

	if !NotEqual(mockT, "Hello World", "Hello World!") {
		t.Error("NotEqual should return true")
	}
	if !NotEqual(mockT, 123, 1234) {
		t.Error("NotEqual should return true")
	}
	if !NotEqual(mockT, 123.5, 123.55) {
		t.Error("NotEqual should return true")
	}
	if !NotEqual(mockT, []byte("Hello World"), []byte("Hello World!")) {
		t.Error("NotEqual should return true")
	}
	if !NotEqual(mockT, nil, new(AssertionTesterConformingObject)) {
		t.Error("NotEqual should return true")
	}
}

func TestContains(t *testing.T) {

	mockT := new(testing.T)

	if !Contains(mockT, "Hello World", "Hello") {
		t.Error("Contains should return true: \"Hello World\" contains \"Hello\"")
	}
	if Contains(mockT, "Hello World", "Salut") {
		t.Error("Contains should return false: \"Hello World\" does not contain \"Salut\"")
	}

}

func TestNotContains(t *testing.T) {

	mockT := new(testing.T)

	if !NotContains(mockT, "Hello World", "Hello!") {
		t.Error("NotContains should return true: \"Hello World\" does not contain \"Hello!\"")
	}
	if NotContains(mockT, "Hello World", "Hello") {
		t.Error("NotContains should return false: \"Hello World\" contains \"Hello\"")
	}

}

func TestDidPanic(t *testing.T) {

	if !didPanic(func() {
		panic("Panic!")
	}) {
		t.Error("didPanic should return true")
	}

	if didPanic(func() {
	}) {
		t.Error("didPanic should return false")
	}

}

func TestPanics(t *testing.T) {

	mockT := new(testing.T)

	if !Panics(mockT, func() {
		panic("Panic!")
	}) {
		t.Error("Panics should return true")
	}

	if Panics(mockT, func() {
	}) {
		t.Error("Panics should return false")
	}

}

func TestNotPanics(t *testing.T) {

	mockT := new(testing.T)

	if !NotPanics(mockT, func() {
	}) {
		t.Error("NotPanics should return true")
	}

	if NotPanics(mockT, func() {
		panic("Panic!")
	}) {
		t.Error("NotPanics should return false")
	}

}

func TestEqual_Funcs(t *testing.T) {

	type f func() int
	var f1 f = func() int { return 1 }
	var f2 f = func() int { return 2 }

	var f1_copy f = f1

	Equal(t, f1_copy, f1, "Funcs are the same and should be considered equal")
	NotEqual(t, f1, f2, "f1 and f2 are different")

}
