// Copyright (c) 2015-2017 Daniel Nichter. All rights reserved.
// Copyright (c) 2022 Prasad Tengse. All rights reserved.
// SPDX-License-Identifier: MIT OR Apache-2.0

package diff

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	v1 "github.com/tprasadtp/pkg/assert/internal/diff/test/v1"
	v2 "github.com/tprasadtp/pkg/assert/internal/diff/test/v2"
)

func TestStringNoDiff(t *testing.T) {
	diff := Diff("foo", "foo")
	if len(diff) > 0 {
		t.Error("should be no diffs, but got:", diff)
	}
}

func TestStringWithDiff(t *testing.T) {
	diff := Diff("foo", "bar")
	if diff == nil {
		t.Fatal("expected diff, but got none")
	}
	if len(diff) != 1 {
		t.Errorf("expected 1 diff, got %d : %s", len(diff), diff)
	}
	if diff[0] != "foo != bar" {
		t.Error("wrong diff:", diff[0])
	}
}

func TestFloat(t *testing.T) {
	diff := Diff(1.1, 1.1)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	diff = Diff(1.1234561, 1.1234562)
	if diff == nil {
		t.Error("no diff")
	}

	config := Config{
		FloatPrecision: 6,
	}

	diff = DiffWithConfig(1.1234561, 1.1234562, config)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	diff = DiffWithConfig(1.123456, 1.123457, config)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "1.123456 != 1.123457" {
		t.Error("wrong diff:", diff[0])
	}
}

func TestFloatCustomPrecision(t *testing.T) {
	config := Config{
		FloatPrecision: 6,
	}

	diff := DiffWithConfig(1.1234561, 1.1234562, config)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	diff = Diff(1.123456, 1.123457)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "1.123456 != 1.123457" {
		t.Error("wrong diff:", diff[0])
	}
}

func TestInt(t *testing.T) {
	diff := Diff(1, 1)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	diff = Diff(1, 2)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "1 != 2" {
		t.Error("wrong diff:", diff[0])
	}
}

func TestUint(t *testing.T) {
	diff := Diff(uint(2), uint(2))
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	diff = Diff(uint(2), uint(3))
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "2 != 3" {
		t.Error("wrong diff:", diff[0])
	}
}

func TestBool(t *testing.T) {
	diff := Diff(true, true)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	diff = Diff(false, false)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	diff = Diff(true, false)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "true != false" { // unless you're fipar
		t.Error("wrong diff:", diff[0])
	}
}

func TestTypeMismatch(t *testing.T) {
	type T1 int // same type kind (int)
	type T2 int // but different type
	var t1 T1 = 1
	var t2 T2 = 1
	diff := Diff(t1, t2)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "diff.T1 != diff.T2" {
		t.Error("wrong diff:", diff[0])
	}

	// Same pkg name but different full paths
	// https://github.com/go-test/diff/issues/39
	err1 := v1.Error{}
	err2 := v2.Error{}
	diff = Diff(err1, err2)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "github.com/tprasadtp/pkg/assert/internal/diff/test/v1.Error != github.com/tprasadtp/pkg/assert/internal/diff/test/v2.Error" {
		t.Error("wrong diff:", diff[0])
	}
}

func TestKindMismatch(t *testing.T) {
	var x int = 100
	var y float64 = 100
	diff := Diff(x, y)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "int != float64" {
		t.Error("wrong diff:", diff[0])
	}
}

func TestMaxDiff(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := []int{0, 0, 0, 0, 0, 0, 0}
	config := Config{
		MaxDiff:                 3,
		CompareUnexportedFields: true,
	}

	diff := DiffWithConfig(a, b, config)
	if diff == nil {
		t.Fatal("no diffs")
	}
	if len(diff) != int(config.MaxDiff) {
		t.Errorf("got %d diffs, expected %d", len(diff), int(config.MaxDiff))
	}
}

func TestStructWihtMaxDiff(t *testing.T) {
	config := Config{
		MaxDiff:                 3,
		CompareUnexportedFields: true,
	}

	type fiveFields struct {
		a int // unexported fields require ^
		b int
		c int
		d int
		e int
	}
	t1 := fiveFields{1, 2, 3, 4, 5}
	t2 := fiveFields{0, 0, 0, 0, 0}
	diff := DiffWithConfig(t1, t2, config)
	if diff == nil {
		t.Fatal("no diffs")
	}
	if len(diff) != int(config.MaxDiff) {
		t.Errorf("got %d diffs, expected %d", len(diff), int(config.MaxDiff))
	}
}

func TestMaxDiffMapsWithSameKeys(t *testing.T) {
	config := Config{
		MaxDiff:                 3,
		CompareUnexportedFields: true,
	}

	// Same keys, too many diffs
	m1 := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
	}
	m2 := map[int]int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
	}
	diff := DiffWithConfig(m1, m2, config)
	if diff == nil {
		t.Fatal("no diffs")
	}
	if len(diff) != int(config.MaxDiff) {
		t.Log(diff)
		t.Errorf("got %d diffs, expected %d", len(diff), int(config.MaxDiff))
	}
}

func TestMaxDiffMapsMissingKeys(t *testing.T) {
	config := Config{
		MaxDiff:                 3,
		CompareUnexportedFields: true,
	}

	m1 := map[int]int{
		1: 1,
		2: 2,
	}
	m2 := map[int]int{
		1: 1,
		2: 2,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
	}
	diff := DiffWithConfig(m1, m2, config)
	if diff == nil {
		t.Fatal("no diffs")
	}
	if len(diff) != int(config.MaxDiff) {
		t.Log(diff)
		t.Errorf("got %d diffs, expected %d", len(diff), int(config.MaxDiff))
	}
}

func TestUncomparable(t *testing.T) {
	a := func(int) {}
	b := func(int) {}
	diff := Diff(a, b)

	if len(diff) != 1 {
		t.Errorf("got %d diffs, expected 1", len(diff))
	}

	if diff[0] != "<un-comparable-func> != <un-comparable-func>" {
		t.Error("got diffs:", diff)
	}
}

func TestStruct(t *testing.T) {
	type s1 struct {
		id     int
		Name   string
		Number int
	}
	sa := s1{
		id:     1,
		Name:   "foo",
		Number: 2,
	}
	sb := sa
	diff := Diff(sa, sb)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	sb.Name = "bar"
	diff = Diff(sa, sb)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "Name: foo != bar" {
		t.Error("wrong diff:", diff[0])
	}

	sb.Number = 22
	diff = Diff(sa, sb)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 2 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "Name: foo != bar" {
		t.Error("wrong diff:", diff[0])
	}
	if diff[1] != "Number: 2 != 22" {
		t.Error("wrong diff:", diff[1])
	}

	sb.id = 11
	diff = Diff(sa, sb)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 2 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "Name: foo != bar" {
		t.Error("wrong diff:", diff[0])
	}
	if diff[1] != "Number: 2 != 22" {
		t.Error("wrong diff:", diff[1])
	}
}

func TestStructWithTags(t *testing.T) {
	config := Config{
		CompareUnexportedFields: true,
	}

	type s1 struct {
		same                    int
		modified                int
		sameIgnored             int `assert:"-"`
		modifiedIgnored         int `assert:"-"`
		ExportedSame            int
		ExportedModified        int
		ExportedSameIgnored     int `assert:"-"`
		ExportedModifiedIgnored int `assert:"-"`
	}
	type s2 struct {
		s1
		same                    int
		modified                int
		sameIgnored             int `assert:"-"`
		modifiedIgnored         int `assert:"-"`
		ExportedSame            int
		ExportedModified        int
		ExportedSameIgnored     int `assert:"-"`
		ExportedModifiedIgnored int `assert:"-"`
		recurseInline           s1
		recursePtr              *s2
	}
	sa := s2{
		s1: s1{
			same:                    0,
			modified:                1,
			sameIgnored:             2,
			modifiedIgnored:         3,
			ExportedSame:            4,
			ExportedModified:        5,
			ExportedSameIgnored:     6,
			ExportedModifiedIgnored: 7,
		},
		same:                    0,
		modified:                1,
		sameIgnored:             2,
		modifiedIgnored:         3,
		ExportedSame:            4,
		ExportedModified:        5,
		ExportedSameIgnored:     6,
		ExportedModifiedIgnored: 7,
		recurseInline: s1{
			same:                    0,
			modified:                1,
			sameIgnored:             2,
			modifiedIgnored:         3,
			ExportedSame:            4,
			ExportedModified:        5,
			ExportedSameIgnored:     6,
			ExportedModifiedIgnored: 7,
		},
		recursePtr: &s2{
			same:                    0,
			modified:                1,
			sameIgnored:             2,
			modifiedIgnored:         3,
			ExportedSame:            4,
			ExportedModified:        5,
			ExportedSameIgnored:     6,
			ExportedModifiedIgnored: 7,
		},
	}
	sb := s2{
		s1: s1{
			same:                    0,
			modified:                10,
			sameIgnored:             2,
			modifiedIgnored:         30,
			ExportedSame:            4,
			ExportedModified:        50,
			ExportedSameIgnored:     6,
			ExportedModifiedIgnored: 70,
		},
		same:                    0,
		modified:                10,
		sameIgnored:             2,
		modifiedIgnored:         30,
		ExportedSame:            4,
		ExportedModified:        50,
		ExportedSameIgnored:     6,
		ExportedModifiedIgnored: 70,
		recurseInline: s1{
			same:                    0,
			modified:                10,
			sameIgnored:             2,
			modifiedIgnored:         30,
			ExportedSame:            4,
			ExportedModified:        50,
			ExportedSameIgnored:     6,
			ExportedModifiedIgnored: 70,
		},
		recursePtr: &s2{
			same:                    0,
			modified:                10,
			sameIgnored:             2,
			modifiedIgnored:         30,
			ExportedSame:            4,
			ExportedModified:        50,
			ExportedSameIgnored:     6,
			ExportedModifiedIgnored: 70,
		},
	}

	diff := DiffWithConfig(sa, sb, config)

	want := []string{
		"s1.modified: 1 != 10",
		"s1.ExportedModified: 5 != 50",
		"modified: 1 != 10",
		"ExportedModified: 5 != 50",
		"recurseInline.modified: 1 != 10",
		"recurseInline.ExportedModified: 5 != 50",
		"recursePtr.modified: 1 != 10",
		"recursePtr.ExportedModified: 5 != 50",
	}
	if !reflect.DeepEqual(want, diff) {
		t.Errorf("got=%s, want=%s", strings.Join(diff, "\n"), strings.Join(want, "\n"))
	}
}

func TestNestedStruct(t *testing.T) {
	type s2 struct {
		Nickname string
	}
	type s1 struct {
		Name  string
		Alias s2
	}
	sa := s1{
		Name:  "Robert",
		Alias: s2{Nickname: "Bob"},
	}
	sb := sa
	diff := Diff(sa, sb)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	sb.Alias.Nickname = "Bobby"
	diff = Diff(sa, sb)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "Alias.Nickname: Bob != Bobby" {
		t.Error("wrong diff:", diff[0])
	}
}

func TestMap(t *testing.T) {
	ma := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	mb := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	diff := Diff(ma, mb)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	diff = Diff(ma, ma)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	mb["foo"] = 111
	diff = Diff(ma, mb)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "map[foo]: 1 != 111" {
		t.Error("wrong diff:", diff[0])
	}

	delete(mb, "foo")
	diff = Diff(ma, mb)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "map[foo]: 1 != <does not have key>" {
		t.Error("wrong diff:", diff[0])
	}

	diff = Diff(mb, ma)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "map[foo]: <does not have key> != 1" {
		t.Error("wrong diff:", diff[0])
	}

	var mc map[string]int
	diff = Diff(ma, mc)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	// handle hash order randomness
	if diff[0] != "map[foo:1 bar:2] != <nil map>" && diff[0] != "map[bar:2 foo:1] != <nil map>" {
		t.Error("wrong diff:", diff[0])
	}

	diff = Diff(mc, ma)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "<nil map> != map[foo:1 bar:2]" && diff[0] != "<nil map> != map[bar:2 foo:1]" {
		t.Error("wrong diff:", diff[0])
	}
}

func TestArray(t *testing.T) {
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}

	diff := Diff(a, b)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	diff = Diff(a, a)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	b[2] = 333
	diff = Diff(a, b)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "array[2]: 3 != 333" {
		t.Error("wrong diff:", diff[0])
	}

	c := [3]int{1, 2, 2}
	diff = Diff(a, c)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "array[2]: 3 != 2" {
		t.Error("wrong diff:", diff[0])
	}

	var d [2]int
	diff = Diff(a, d)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "[3]int != [2]int" {
		t.Error("wrong diff:", diff[0])
	}

	e := [12]int{}
	f := [12]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	diff = Diff(e, f)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != int(DefaultConfig.MaxDiff) {
		t.Error("not enough diffs:", diff)
	}
	for i := 0; i < int(DefaultConfig.MaxDiff); i++ {
		if diff[i] != fmt.Sprintf("array[%d]: 0 != %d", i+1, i+1) {
			t.Error("wrong diff:", diff[i])
		}
	}
}

func TestSlice(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}

	diff := Diff(a, b)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	diff = Diff(a, a)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	b[2] = 333
	diff = Diff(a, b)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "slice[2]: 3 != 333" {
		t.Error("wrong diff:", diff[0])
	}

	b = b[0:2]
	diff = Diff(a, b)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "slice[2]: 3 != <no value>" {
		t.Error("wrong diff:", diff[0])
	}

	diff = Diff(b, a)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "slice[2]: <no value> != 3" {
		t.Error("wrong diff:", diff[0])
	}

	var c []int
	diff = Diff(a, c)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "[1 2 3] != <nil slice>" {
		t.Error("wrong diff:", diff[0])
	}

	diff = Diff(c, a)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "<nil slice> != [1 2 3]" {
		t.Error("wrong diff:", diff[0])
	}
}

func TestSiblingSlices(t *testing.T) {
	father := []int{1, 2, 3, 4}
	a := father[0:3]
	b := father[0:3]

	diff := Diff(a, b)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}
	diff = Diff(b, a)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	a = father[0:3]
	b = father[0:2]
	diff = Diff(a, b)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "slice[2]: 3 != <no value>" {
		t.Error("wrong diff:", diff[0])
	}

	a = father[0:2]
	b = father[0:3]

	diff = Diff(a, b)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "slice[2]: <no value> != 3" {
		t.Error("wrong diff:", diff[0])
	}

	a = father[0:2]
	b = father[2:4]

	diff = Diff(a, b)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 2 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "slice[0]: 1 != 3" {
		t.Error("wrong diff:", diff[0])
	}
	if diff[1] != "slice[1]: 2 != 4" {
		t.Error("wrong diff:", diff[1])
	}

	a = father[0:0]
	b = father[1:1]

	diff = Diff(a, b)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}
	diff = Diff(b, a)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}
}

func TestEmptySlice(t *testing.T) {
	a := []int{1}
	b := []int{}
	var c []int

	// Non-empty is not Equal to empty.
	diff := Diff(a, b)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "slice[0]: 1 != <no value>" {
		t.Error("wrong diff:", diff[0])
	}

	// Empty is not Equal to non-empty.
	diff = Diff(b, a)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "slice[0]: <no value> != 1" {
		t.Error("wrong diff:", diff[0])
	}

	// Empty is not Equal to nil.
	diff = Diff(b, c)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "[] != <nil slice>" {
		t.Error("wrong diff:", diff[0])
	}

	// Nil is not Equal to empty.
	diff = Diff(c, b)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "<nil slice> != []" {
		t.Error("wrong diff:", diff[0])
	}
}

func TestNilSlicesAreEmpty(t *testing.T) {
	a := []int{1}
	b := []int{}
	var c []int

	config := Config{NilSlicesAreEmpty: true}

	// Empty is Equal to nil.
	diff := DiffWithConfig(b, c, config)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	// Nil is Equal to empty.
	diff = DiffWithConfig(c, b, config)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	// Non-empty is not Equal to nil.
	diff = DiffWithConfig(a, c, config)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "[1] != <nil slice>" {
		t.Error("wrong diff:", diff[0])
	}

	// Nil is not Equal to non-empty.
	diff = DiffWithConfig(c, a, config)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "<nil slice> != [1]" {
		t.Error("wrong diff:", diff[0])
	}

	// Non-empty is not Equal to empty.
	diff = DiffWithConfig(a, b, config)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "slice[0]: 1 != <no value>" {
		t.Error("wrong diff:", diff[0])
	}

	// Empty is not Equal to non-empty.
	diff = DiffWithConfig(b, a, config)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "slice[0]: <no value> != 1" {
		t.Error("wrong diff:", diff[0])
	}
}

func TestNilMapsAreEmpty(t *testing.T) {
	a := map[int]int{1: 1}
	b := map[int]int{}
	var c map[int]int

	config := Config{NilMapsAreEmpty: true}

	// Empty is Equal to nil.
	diff := DiffWithConfig(b, c, config)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	// Nil is Equal to empty.
	diff = DiffWithConfig(c, b, config)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	// Non-empty is not Equal to nil.
	diff = DiffWithConfig(a, c, config)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "map[1:1] != <nil map>" {
		t.Error("wrong diff:", diff[0])
	}

	// Nil is not Equal to non-empty.
	diff = DiffWithConfig(c, a, config)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "<nil map> != map[1:1]" {
		t.Error("wrong diff:", diff[0])
	}

	// Non-empty is not Equal to empty.
	diff = DiffWithConfig(a, b, config)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "map[1]: 1 != <does not have key>" {
		t.Error("wrong diff:", diff[0])
	}

	// Empty is not Equal to non-empty.
	diff = DiffWithConfig(b, a, config)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "map[1]: <does not have key> != 1" {
		t.Error("wrong diff:", diff[0])
	}
}

func TestNilInterface(t *testing.T) {
	type T struct{ i int }

	a := &T{i: 1}
	diff := Diff(nil, a)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "<nil pointer> != &{1}" {
		t.Error("wrong diff:", diff[0])
	}

	diff = Diff(a, nil)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "&{1} != <nil pointer>" {
		t.Error("wrong diff:", diff[0])
	}

	diff = Diff(nil, nil)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}
}

func TestPointer(t *testing.T) {
	type T struct{ i int }

	a, b := &T{i: 1}, &T{i: 1}
	diff := Diff(a, b)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	a, b = nil, &T{}
	diff = Diff(a, b)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "<nil pointer> != diff.T" {
		t.Error("wrong diff:", diff[0])
	}

	a, b = &T{}, nil
	diff = Diff(a, b)
	if diff == nil {
		t.Fatal("no diff")
	}
	if len(diff) != 1 {
		t.Error("too many diff:", diff)
	}
	if diff[0] != "diff.T != <nil pointer>" {
		t.Error("wrong diff:", diff[0])
	}

	a, b = nil, nil
	diff = Diff(a, b)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}
}

func TestTime(t *testing.T) {
	// In an interable kind (i.e. a struct)
	type sTime struct {
		T time.Time
	}
	now := time.Now()
	got := sTime{T: now}
	expect := sTime{T: now.Add(1 * time.Second)}
	diff := Diff(got, expect)
	if len(diff) != 1 {
		t.Error("expected 1 diff:", diff)
	}

	// Directly
	a := now
	b := now
	diff = Diff(a, b)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	// https://github.com/go-test/diff/issues/15
	type Time15 struct {
		time.Time
	}
	a15 := Time15{now}
	b15 := Time15{now}
	diff = Diff(a15, b15)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	later := now.Add(1 * time.Second)
	b15 = Time15{later}
	diff = Diff(a15, b15)
	if len(diff) != 1 {
		t.Errorf("got %d diffs, expected 1: %s", len(diff), diff)
	}

	// No diff in Equal should not affect diff of other fields (Foo)
	type Time17 struct {
		time.Time
		Foo int
	}
	a17 := Time17{Time: now, Foo: 1}
	b17 := Time17{Time: now, Foo: 2}
	diff = Diff(a17, b17)
	if len(diff) != 1 {
		t.Errorf("got %d diffs, expected 1: %s", len(diff), diff)
	}
}

func TestTimeUnexported(t *testing.T) {
	// https://github.com/go-test/diff/issues/18
	// Can't call Call() on exported Value func

	now := time.Now()
	type hiddenTime struct {
		t time.Time
	}
	htA := &hiddenTime{t: now}
	htB := &hiddenTime{t: now}
	diff := Diff(htA, htB)
	if len(diff) > 0 {
		t.Error("should be Equal:", diff)
	}

	// This doesn't call time.Time.Diff(), it Equals the un-exported fields
	// in time.Time, causing a diff like:
	// [t.wall: 13740788835924462040 != 13740788836998203864 t.ext: 1447549 != 1001447549]
	later := now.Add(1 * time.Second)
	htC := &hiddenTime{t: later}
	diff = DiffWithConfig(htA, htC, Config{CompareUnexportedFields: true})

	expected := 1
	if _, ok := reflect.TypeOf(htA.t).FieldByName("ext"); ok {
		expected = 2
	}
	if len(diff) != expected {
		t.Errorf("got %d diffs, expected %d: %s", len(diff), expected, diff)
	}
}

func TestInterface(t *testing.T) {
	a := map[string]any{
		"foo": map[string]string{
			"bar": "a",
		},
	}
	b := map[string]any{
		"foo": map[string]string{
			"bar": "b",
		},
	}
	diff := Diff(a, b)
	if len(diff) == 0 {
		t.Fatalf("expected 1 diff, got zero")
	}
	if len(diff) != 1 {
		t.Errorf("expected 1 diff, got %d: %s", len(diff), diff)
	}
}

func TestInterface2(t *testing.T) {
	defer func() {
		if val := recover(); val != nil {
			t.Fatalf("panic: %v", val)
		}
	}()

	a := map[string]any{
		"bar": 1,
	}
	b := map[string]any{
		"bar": 1.23,
	}
	diff := Diff(a, b)
	if len(diff) == 0 {
		t.Fatalf("expected 1 diff, got zero")
	}
	if len(diff) != 1 {
		t.Errorf("expected 1 diff, got %d: %s", len(diff), diff)
	}
}

func TestInterface3(t *testing.T) {
	//lint:ignore U1000 Required for diff tests
	type Value struct{ int }
	a := map[string]any{
		"foo": &Value{},
	}
	b := map[string]any{
		"foo": 1.23,
	}
	diff := Diff(a, b)
	if len(diff) == 0 {
		t.Fatalf("expected 1 diff, got zero")
	}

	if len(diff) != 1 {
		t.Errorf("expected 1 diff, got: %s", diff)
	}
}

func TestError(t *testing.T) {
	a := errors.New("it broke")
	b := errors.New("it broke")

	diff := Diff(a, b)
	if len(diff) != 0 {
		t.Fatalf("expected zero diffs, got %d: %s", len(diff), diff)
	}

	b = errors.New("it fell apart")
	diff = Diff(a, b)
	if len(diff) != 1 {
		t.Fatalf("expected 1 diff, got %d: %s", len(diff), diff)
	}
	if diff[0] != "it broke != it fell apart" {
		t.Errorf("got '%s', expected 'it broke != it fell apart'", diff[0])
	}

	// Both errors set
	type tWithError struct {
		Error error
	}
	t1 := tWithError{
		Error: a,
	}
	t2 := tWithError{
		Error: b,
	}
	diff = Diff(t1, t2)
	if len(diff) != 1 {
		t.Fatalf("expected 1 diff, got %d: %s", len(diff), diff)
	}
	if diff[0] != "Error: it broke != it fell apart" {
		t.Errorf("got '%s', expected 'Error: it broke != it fell apart'", diff[0])
	}

	// Both errors nil
	t1 = tWithError{
		Error: nil,
	}
	t2 = tWithError{
		Error: nil,
	}
	diff = Diff(t1, t2)
	if len(diff) != 0 {
		t.Log(diff)
		t.Fatalf("expected 0 diff, got %d: %s", len(diff), diff)
	}

	// One error is nil
	t1 = tWithError{
		Error: errors.New("foo"),
	}
	t2 = tWithError{
		Error: nil,
	}
	diff = Diff(t1, t2)
	if len(diff) != 1 {
		t.Log(diff)
		t.Fatalf("expected 1 diff, got %d: %s", len(diff), diff)
	}
	if diff[0] != "Error: *errors.errorString != <nil pointer>" {
		t.Errorf("got '%s', expected 'Error: *errors.errorString != <nil pointer>'", diff[0])
	}
}

func TestErrorWithOtherFields(t *testing.T) {
	a := errors.New("it broke")
	b := errors.New("it broke")

	diff := Diff(a, b)
	if len(diff) != 0 {
		t.Fatalf("expected zero diffs, got %d: %s", len(diff), diff)
	}

	b = errors.New("it fell apart")
	diff = Diff(a, b)
	if len(diff) != 1 {
		t.Fatalf("expected 1 diff, got %d: %s", len(diff), diff)
	}
	if diff[0] != "it broke != it fell apart" {
		t.Errorf("got '%s', expected 'it broke != it fell apart'", diff[0])
	}

	// Both errors set
	type tWithError struct {
		Error error
		Other string
	}
	t1 := tWithError{
		Error: a,
		Other: "ok",
	}
	t2 := tWithError{
		Error: b,
		Other: "ok",
	}
	diff = Diff(t1, t2)
	if len(diff) != 1 {
		t.Fatalf("expected 1 diff, got %d: %s", len(diff), diff)
	}
	if diff[0] != "Error: it broke != it fell apart" {
		t.Errorf("got '%s', expected 'Error: it broke != it fell apart'", diff[0])
	}

	// Both errors nil
	t1 = tWithError{
		Error: nil,
		Other: "ok",
	}
	t2 = tWithError{
		Error: nil,
		Other: "ok",
	}
	diff = Diff(t1, t2)
	if len(diff) != 0 {
		t.Log(diff)
		t.Fatalf("expected 0 diff, got %d: %s", len(diff), diff)
	}

	// Different Other value
	t1 = tWithError{
		Error: nil,
		Other: "ok",
	}
	t2 = tWithError{
		Error: nil,
		Other: "nope",
	}
	diff = Diff(t1, t2)
	if len(diff) != 1 {
		t.Fatalf("expected 1 diff, got %d: %s", len(diff), diff)
	}
	if diff[0] != "Other: ok != nope" {
		t.Errorf("got '%s', expected 'Other: ok != nope'", diff[0])
	}

	// Different Other value, same error
	t1 = tWithError{
		Error: a,
		Other: "ok",
	}
	t2 = tWithError{
		Error: a,
		Other: "nope",
	}
	diff = Diff(t1, t2)
	if len(diff) != 1 {
		t.Fatalf("expected 1 diff, got %d: %s", len(diff), diff)
	}
	if diff[0] != "Other: ok != nope" {
		t.Errorf("got '%s', expected 'Other: ok != nope'", diff[0])
	}
}

type primKindError string

func (e primKindError) Error() string {
	return string(e)
}

func TestErrorPrimitiveKind(t *testing.T) {
	// The primKindError type above is valid and used by Go, e.g.
	// url.EscapeError and url.InvalidHostError. Before fixing this bug
	// (https://github.com/go-test/diff/issues/31), we presumed a and b
	// were ptr or interface (and not nil), so a.Elem() worked. But when
	// a/b are primitive kinds, Elem() causes a panic.
	var err1 primKindError = "abc"
	var err2 primKindError = "abc"
	diff := Diff(err1, err2)
	if len(diff) != 0 {
		t.Fatalf("expected zero diffs, got %d: %s", len(diff), diff)
	}
}

func TestNil(t *testing.T) {
	type student struct {
		name string
		age  int
	}

	mark := student{"mark", 10}
	var someNilThing any = nil
	diff := Diff(someNilThing, mark)
	if diff == nil {
		t.Error("Nil value to comparison should not be Equal")
	}
	diff = Diff(mark, someNilThing)
	if diff == nil {
		t.Error("Nil value to comparison should not be Equal")
	}
	diff = Diff(someNilThing, someNilThing)
	if diff != nil {
		t.Error("Nil value to comparison should not be Equal")
	}
}
