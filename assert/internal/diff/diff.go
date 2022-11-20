// Copyright (c) 2015-2017 Daniel Nichter. All rights reserved.
// Copyright (c) 2022 Prasad Tengse. All rights reserved.
// SPDX-License-Identifier: MIT OR Apache-2.0

package diff

import (
	"fmt"
	"reflect"
	"strings"
)

// Config is diff configuration allowing to tweak diff output
type Config struct {
	// FloatPrecision is the number of decimal places to round float values
	// to when comparing.
	FloatPrecision uint
	// MaxDiff specifies the maximum number of differences to return.
	MaxDiff uint
	// CompareUnexportedFields causes un-exported struct fields, like s in
	// T{s int}, to be compared when true.
	CompareUnexportedFields bool
	// NilSlicesAreEmpty causes a nil slice to be Equal to an empty slice.
	NilSlicesAreEmpty bool
	// NilMapsAreEmpty causes a nil map to be Equal to an empty map.
	NilMapsAreEmpty bool
}

// DefaultConfig is default diff configuration.
// NilSlicesAreEmpty and NilMapsAreEmpty are both set to false
// to avoid issues with json encoding and APIs which cannot handle
// null vs empty.
var DefaultConfig Config = Config{
	FloatPrecision:          10,
	MaxDiff:                 10,
	CompareUnexportedFields: false,
	NilSlicesAreEmpty:       false,
	NilMapsAreEmpty:         false,
}

type cmp struct {
	Equal       []string
	buff        []string
	floatFormat string
	config      Config
}

var errorType = reflect.TypeOf((*error)(nil)).Elem()

// Diff compares variables a and b, recursing into their structure
// and returns a list of differences,
// or nil if there are none. Some differences may not be found if an error is
// also returned.
//
// If a type has an Diff method, like time.Equal, it is called.
//
// When comparing a struct, if a field has the tag `deep:"-"` then it will be
// ignored.
func Diff(got, want any) []string {
	return DiffWithConfig(got, want, DefaultConfig)
}

// DiffWithConfig is same as Diff, but accepts a custom
// [github.com/tprasadtp/pkg/assert/internal/diff.Config]
func DiffWithConfig(got, want any, config Config) []string {
	aVal := reflect.ValueOf(got)
	bVal := reflect.ValueOf(want)
	c := &cmp{
		Equal:       []string{},
		buff:        []string{},
		floatFormat: fmt.Sprintf("%%.%df", config.FloatPrecision),
		config:      config,
	}
	if got == nil && want == nil {
		return nil
	} else if got == nil && want != nil {
		c.saveEqual("<nil pointer>", want)
	} else if got != nil && want == nil {
		c.saveEqual(got, "<nil pointer>")
	}
	if len(c.Equal) > 0 {
		return c.Equal
	}

	c.Diff(aVal, bVal, 0)
	if len(c.Equal) > 0 {
		return c.Equal // Equals
	}
	return nil // no Equals
}

func (c *cmp) Diff(a, b reflect.Value, level int) {
	// Check if one value is nil, e.g. T{x: *X} and T.x is nil
	if !a.IsValid() || !b.IsValid() {
		if a.IsValid() && !b.IsValid() {
			c.saveEqual(a.Type(), "<nil pointer>")
		} else if !a.IsValid() && b.IsValid() {
			c.saveEqual("<nil pointer>", b.Type())
		}
		return
	}

	// If different types, they can't be Equal
	aType := a.Type()
	bType := b.Type()
	if aType != bType {
		// Built-in types don't have a name, so don't report [3]int != [2]int as " != "
		if aType.Name() == "" || aType.Name() != bType.Name() {
			c.saveEqual(aType, bType)
		} else {
			// Type names can be the same, e.g. pkg/v1.Error and pkg/v2.Error
			// are both exported as pkg, so unless we include the full pkg path
			// the Equal will be "pkg.Error != pkg.Error"
			// https://github.com/go-test/deep/issues/39
			aFullType := aType.PkgPath() + "." + aType.Name()
			bFullType := bType.PkgPath() + "." + bType.Name()
			c.saveEqual(aFullType, bFullType)
		}
		return
	}

	// Primitive https://golang.org/pkg/reflect/#Kind
	aKind := a.Kind()
	bKind := b.Kind()

	// Do a and b have underlying elements? Yes if they're ptr or interface.
	aElem := aKind == reflect.Ptr || aKind == reflect.Interface
	bElem := bKind == reflect.Ptr || bKind == reflect.Interface

	// If both types implement the error interface, compare the error strings.
	// This must be done before dereferencing because the interface is on a
	// pointer receiver. Re https://github.com/go-test/deep/issues/31, a/b might
	// be primitive kinds; see TestErrorPrimitiveKind.
	if aType.Implements(errorType) && bType.Implements(errorType) {
		if (!aElem || !a.IsNil()) && (!bElem || !b.IsNil()) {
			aString := a.MethodByName("Error").Call(nil)[0].String()
			bString := b.MethodByName("Error").Call(nil)[0].String()
			if aString != bString {
				c.saveEqual(aString, bString)
				return
			}
		}
	}

	// Dereference pointers and any
	if aElem || bElem {
		if aElem {
			a = a.Elem()
		}
		if bElem {
			b = b.Elem()
		}
		c.Diff(a, b, level+1)
		return
	}

	switch aKind {

	/////////////////////////////////////////////////////////////////////
	// Iterable kinds
	/////////////////////////////////////////////////////////////////////

	case reflect.Struct:
		/*
			The variables are structs like:
				type T struct {
					FirstName string
					LastName  string
				}
			Type = <pkg>.T, Kind = reflect.Struct

			Iterate through the fields (FirstName, LastName), recurse into their values.
		*/

		// Types with an Equal() method, like time.Time, only if struct field
		// is exported (CanInterface)
		if eqFunc := a.MethodByName("Equal"); eqFunc.IsValid() && eqFunc.CanInterface() {
			// Handle https://github.com/go-test/deep/issues/15:
			// Don't call T.Equal if the method is from an embedded struct, like:
			//   type Foo struct { time.Time }
			// First, we'll encounter Equal(Ttime, time.Time) but if we pass b
			// as the 2nd arg we'll panic: "Call using pkg.Foo as type time.Time"
			// As far as I can tell, there's no way to see that the method is from
			// time.Time not Foo. So we check the type of the 1st (0) arg and skip
			// unless it's b type. Later, we'll encounter the time.Time anonymous/
			// embedded field and then we'll have Equal(time.Time, time.Time).
			funcType := eqFunc.Type()
			if funcType.NumIn() == 1 && funcType.In(0) == bType {
				retVals := eqFunc.Call([]reflect.Value{b})
				if !retVals[0].Bool() {
					c.saveEqual(a, b)
				}
				return
			}
		}

		for i := 0; i < a.NumField(); i++ {
			if aType.Field(i).PkgPath != "" && !c.config.CompareUnexportedFields {
				continue // skip unexported field, e.g. s in type T struct {s string}
			}

			if aType.Field(i).Tag.Get("deep") == "-" {
				continue // field wants to be ignored
			}

			c.push(aType.Field(i).Name) // push field name to buff

			// Get the Value for each field, e.g. FirstName has Type = string,
			// Kind = reflect.String.
			af := a.Field(i)
			bf := b.Field(i)

			// Recurse to compare the field values
			c.Diff(af, bf, level+1)

			c.pop() // pop field name from buff

			if len(c.Equal) >= int(c.config.MaxDiff) {
				break
			}
		}
	case reflect.Map:
		/*
			The variables are maps like:
				map[string]int{
					"foo": 1,
					"bar": 2,
				}
			Type = map[string]int, Kind = reflect.Map

			Or:
				type T map[string]int{}
			Type = <pkg>.T, Kind = reflect.Map

			Iterate through the map keys (foo, bar), recurse into their values.
		*/

		if a.IsNil() || b.IsNil() {
			if c.config.NilMapsAreEmpty {
				if a.IsNil() && b.Len() != 0 {
					c.saveEqual("<nil map>", b)
					return
				} else if a.Len() != 0 && b.IsNil() {
					c.saveEqual(a, "<nil map>")
					return
				}
			} else {
				if a.IsNil() && !b.IsNil() {
					c.saveEqual("<nil map>", b)
				} else if !a.IsNil() && b.IsNil() {
					c.saveEqual(a, "<nil map>")
				}
			}
			return
		}

		if a.Pointer() == b.Pointer() {
			return
		}

		for _, key := range a.MapKeys() {
			c.push(fmt.Sprintf("map[%v]", key))

			aVal := a.MapIndex(key)
			bVal := b.MapIndex(key)
			if bVal.IsValid() {
				c.Diff(aVal, bVal, level+1)
			} else {
				c.saveEqual(aVal, "<does not have key>")
			}

			c.pop()

			if len(c.Equal) >= int(c.config.MaxDiff) {
				return
			}
		}

		for _, key := range b.MapKeys() {
			if aVal := a.MapIndex(key); aVal.IsValid() {
				continue
			}

			c.push(fmt.Sprintf("map[%v]", key))
			c.saveEqual("<does not have key>", b.MapIndex(key))
			c.pop()
			if len(c.Equal) >= int(c.config.MaxDiff) {
				return
			}
		}
	case reflect.Array:
		n := a.Len()
		for i := 0; i < n; i++ {
			c.push(fmt.Sprintf("array[%d]", i))
			c.Diff(a.Index(i), b.Index(i), level+1)
			c.pop()
			if len(c.Equal) >= int(c.config.MaxDiff) {
				break
			}
		}
	case reflect.Slice:
		if c.config.NilSlicesAreEmpty {
			if a.IsNil() && b.Len() != 0 {
				c.saveEqual("<nil slice>", b)
				return
			} else if a.Len() != 0 && b.IsNil() {
				c.saveEqual(a, "<nil slice>")
				return
			}
		} else {
			if a.IsNil() && !b.IsNil() {
				c.saveEqual("<nil slice>", b)
				return
			} else if !a.IsNil() && b.IsNil() {
				c.saveEqual(a, "<nil slice>")
				return
			}
		}

		aLen := a.Len()
		bLen := b.Len()

		if a.Pointer() == b.Pointer() && aLen == bLen {
			return
		}

		n := aLen
		if bLen > aLen {
			n = bLen
		}
		for i := 0; i < n; i++ {
			c.push(fmt.Sprintf("slice[%d]", i))
			if i < aLen && i < bLen {
				c.Diff(a.Index(i), b.Index(i), level+1)
			} else if i < aLen {
				c.saveEqual(a.Index(i), "<no value>")
			} else {
				c.saveEqual("<no value>", b.Index(i))
			}
			c.pop()
			if len(c.Equal) >= int(c.config.MaxDiff) {
				break
			}
		}

	/////////////////////////////////////////////////////////////////////
	// Primitive kinds
	/////////////////////////////////////////////////////////////////////

	case reflect.Float32, reflect.Float64:
		// Round floats to FloatPrecision decimal places to compare with
		// user-defined precision. As is commonly know, floats have "imprecision"
		// such that 0.1 becomes 0.100000001490116119384765625. This cannot
		// be avoided; it can only be handled. Issue 30 suggested that floats
		// be compared using an epsilon: Equal = |a-b| < epsilon.
		// In many cases the result is the same, but I think epsilon is a little
		// less clear for users to reason about. See issue 30 for details.
		aval := fmt.Sprintf(c.floatFormat, a.Float())
		bval := fmt.Sprintf(c.floatFormat, b.Float())
		if aval != bval {
			c.saveEqual(a.Float(), b.Float())
		}
	case reflect.Bool:
		if a.Bool() != b.Bool() {
			c.saveEqual(a.Bool(), b.Bool())
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if a.Int() != b.Int() {
			c.saveEqual(a.Int(), b.Int())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if a.Uint() != b.Uint() {
			c.saveEqual(a.Uint(), b.Uint())
		}
	case reflect.String:
		if a.String() != b.String() {
			c.saveEqual(a.String(), b.String())
		}
	default:
		c.saveEqual(fmt.Sprintf("<un-comparable-%s>", aKind.String()), fmt.Sprintf("<un-comparable-%s>", bKind.String()))
	}
}

func (c *cmp) push(name string) {
	c.buff = append(c.buff, name)
}

func (c *cmp) pop() {
	if len(c.buff) > 0 {
		c.buff = c.buff[0 : len(c.buff)-1]
	}
}

func (c *cmp) saveEqual(aval, bval any) {
	if len(c.buff) > 0 {
		varName := strings.Join(c.buff, ".")
		c.Equal = append(c.Equal, fmt.Sprintf("%s: %v != %v", varName, aval, bval))
	} else {
		c.Equal = append(c.Equal, fmt.Sprintf("%v != %v", aval, bval))
	}
}
