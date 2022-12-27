//go:build ignore

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var value = func() complex128 {
		return complex(math.Inf(1), 1.1)
	}()
	s := strconv.FormatComplex(
		complex128(value), 'G',
		4, 64,
	)
	fmt.Printf("%#v", s)
}
