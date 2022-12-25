//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	var stored uint64
	var value = func() int64 {
		return int64(-time.Second)
	}()
	stored = uint64(value)
	fmt.Printf("%#v", stored)
}
