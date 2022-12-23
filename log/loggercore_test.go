package log

import (
	"testing"
)

func TestAllocs_getCallerInfo(t *testing.T) {
	a := testing.AllocsPerRun(10, func() {
		getCallerInfo(1)
	})
	if a != float64(2) {
		t.Errorf("Allocs =%f, Expected=2", a)
	}
}
