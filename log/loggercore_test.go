package log

import (
	"testing"
)

func TestCaller(t *testing.T) {
	a := testing.AllocsPerRun(10, func() {
		getCallerInfo(1)
	})
	if a != 0 {
		t.Errorf("Allocs =%f, Expected=0", a)
	}
}
