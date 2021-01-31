package pkg

import (
	"testing"
)

func TestDummy(t *testing.T) {
	if !ShadowPackage {
		t.FailNow()
	}
}
