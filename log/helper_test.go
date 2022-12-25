package log_test

import (
	"testing"

	"github.com/tprasadtp/pkg/log"
	"github.com/tprasadtp/pkg/log/internal/helpers"
)

func helperFoo() {
	log.Helper()
}

func helperBar() {
	log.Helper()
}

// This test is sensitive to package name changes as it uses runtime.Caller.
// If you change/move the package Please change below tests values accordingly.
const (
	helper01FuncName = "github.com/tprasadtp/pkg/log_test.helperFoo"
	helper02FuncName = "github.com/tprasadtp/pkg/log_test.helperBar"
)

func TestHelper(t *testing.T) {
	// reset helpers map
	// to avoid shared state between tests
	t.Cleanup(func() {
		helpers.Map.Range(func(key, value any) bool {
			helpers.Map.Delete(key)
			return true
		})
	})

	// call first helper func
	helperFoo()
	h1, ok1 := helpers.Map.Load(helper01FuncName)
	if !ok1 {
		t.Errorf("%s(key) should be in helpers stack", helper01FuncName)
	}
	if h1 != nil {
		t.Errorf("%s(value) should return nil as its being used as list", helper01FuncName)
	}

	// call second helper func
	helperBar()
	h2, ok2 := helpers.Map.Load(helper02FuncName)
	if !ok2 {
		t.Errorf("%s(key) should be in helpers stack", helper02FuncName)
	}
	if h2 != nil {
		t.Errorf("%s(value) should return nil as its being used as list", helper02FuncName)
	}

	// ensure caller 01 is still marked as  helper.
	h3, ok3 := helpers.Map.Load(helper01FuncName)
	if !ok3 {
		t.Errorf("%s(key) should still be in helpers stack after calling bar", helper01FuncName)
	}
	if h3 != nil {
		t.Errorf("%s(value) should still return nil as its being used as list", helper01FuncName)
	}
}

func BenchmarkHelper(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		log.Helper()
	}
}
