package main

import (
	"fmt"
	"testing"
	"time"
)

var _a Attr
var _d time.Duration

func main() {
	d := 34353 * time.Hour // arbitrary int that won't fit in an interface without allocation.
	r := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_a = Mk("foo", d)
		}
	})
	fmt.Println(r, r.MemString())

	r = testing.Benchmark(func(b *testing.B) {
		a := Mk("foo", d)
		for i := 0; i < b.N; i++ {
			_d = Value[time.Duration](a)
		}
	})
	fmt.Println(r, r.MemString())
}

type Attr struct {
	kind Kind
	x    uint64
	s    string
	i    any
}

func Duration(key string, val time.Duration) Attr {
	return Attr{x: uint64(val)}
}

func Mk[T any](key string, val T) Attr {
	var a Attr
	switch any((*T)(nil)).(type) {
	case *time.Duration:
		a.kind = DurationKind
		a.x = uint64(any(val).(time.Duration))
	case *string:
		a.kind = StringKind
		a.s = any(val).(string)
	default:
		a.i = val
		a.kind = AnyKind
	}
	return a
}

type Kind int

const (
	NilKind Kind = iota
	BoolKind
	DurationKind
	Float64Kind
	Int64Kind
	StringKind
	TimeKind
	Uint64Kind
	AnyKind
)

func Value[T any](a Attr) T {
	var v T
	switch v := any(&v).(type) {
	case *time.Duration:
		if a.kind != DurationKind {
			panic("not duration")
		}
		*v = time.Duration(a.x)
	case *string:
		if a.kind != StringKind {
			panic("not string")
		}
		*v = a.s
	case *any:
		switch a.kind {
		case DurationKind:
			*v = time.Duration(a.x)
		case StringKind:
			*v = a.s
		case AnyKind:
			*v = a.i
		}
	default:
		panic("unexpected type")
	}
	return v
}
