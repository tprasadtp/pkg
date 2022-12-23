//go:build ignore

package main

import (
	"fmt"
	"math"
)

func main() {
	var stored uint64
	var value = func() float64 {
		return 0
	}()
	stored = math.Float64bits(value)
	fmt.Printf("%#v", stored)
}
