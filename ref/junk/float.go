//go:build ignore

package main

import (
	"fmt"
	"math"
)

func main() {
	var stored uint64
	var value = func() float32 {
		return float32(math.Inf(-1))
	}()
	stored = math.Float64bits(float64(value))
	fmt.Printf("%#v", stored)
}
