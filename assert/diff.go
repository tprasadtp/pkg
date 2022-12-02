package assert

import (
	"reflect"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// Diff returns a diff between exp and act.
func Diff(exp, act any, opts ...cmp.Option) string {
	opts = append(opts, cmpopts.EquateErrors(), cmp.Exporter(func(r reflect.Type) bool {
		return true
	}))
	return cmp.Diff(exp, act, opts...)
}
