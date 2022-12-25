package log_test

import (
	"testing"
	"time"
)

func TestAllocs(t *testing.T) {
	type testCase struct {
		name   string
		f      func()
		allocs float64
	}

	tt := []testCase{
		{
			name: "<time.Duration>",
			f: func() {
				val := time.Second.String()
				_ = val
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			allocs := testing.AllocsPerRun(10, tc.f)
			if allocs != tc.allocs {
				t.Errorf("%s => alloc mismatch expected=%f, got=%f", tc.name, tc.allocs, allocs)
			}
		})
	}
}
