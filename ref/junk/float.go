//go:build ignore

package main

import (
	"fmt"
	"math"
)

func main() {
	var stored uint64
	var value = func() float64 {
		return -10.123
	}()
	stored = math.Float64bits(float64(value))
	fmt.Printf("%#v", stored)
}
