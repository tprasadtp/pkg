package log_test

import (
	"fmt"
	"testing"

	"github.com/tprasadtp/pkg/log"
)

func TestLevelStringer(t *testing.T) {
	type testCase struct {
		lvl    log.Level
		expect string
	}

	tt := []testCase{
		{
			lvl:    log.Level(0),
			expect: "DEBUG-10",
		},
		{
			lvl:    log.Level(10),
			expect: "DEBUG",
		},
		{
			lvl:    log.INFO,
			expect: "INFO",
		},
		{
			lvl:    log.Level(100),
			expect: "CRITICAL+20",
		},
	}
	for _, tc := range tt {
		tn := fmt.Sprintf("level=%d", tc.lvl)
		t.Run(tn, func(t *testing.T) {
			val := tc.lvl.String()
			if tc.expect != val {
				t.Errorf("%s => got=%v, want=%v", tn, val, tc.expect)
			}
		})
	}
}

func TestLevelMarshalJSON(t *testing.T) {
	type testCase struct {
		lvl    log.Level
		expect []byte
	}

	tt := []testCase{
		{
			lvl:    log.Level(0),
			expect: []byte{34, 68, 69, 66, 85, 71, 45, 49, 48, 34},
		},
		{
			lvl:    log.Level(10),
			expect: []byte{34, 68, 69, 66, 85, 71, 34},
		},
		{
			lvl:    log.INFO,
			expect: []byte{34, 73, 78, 70, 79, 34},
		},
		{
			lvl:    log.Level(100),
			expect: []byte{34, 73, 78, 70, 79, 34},
		},
	}
	for _, tc := range tt {
		tn := fmt.Sprintf("level=%d", tc.lvl)
		t.Run(tn, func(t *testing.T) {
			val, err := tc.lvl.MarshalJSON()
			if tc.expect != val {
				t.Errorf("%s => got=%v, want=%v", tn, val, tc.expect)
			}
		})
	}
}
