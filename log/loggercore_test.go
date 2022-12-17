package log

import (
	"testing"
)

func TestAllocsCaller(t *testing.T) {
	a := testing.AllocsPerRun(10, func() {
		getCallerInfo(1)
	})
	if a != float64(1) {
		t.Errorf("Allocs =%f, Expected=1", a)
	}
}
