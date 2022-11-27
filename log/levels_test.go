package log_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/tprasadtp/pkg/assert"
	"github.com/tprasadtp/pkg/log"
)

func TestLevelStringer(t *testing.T) {
	type testCase struct {
		lvl    log.Level
		expect string
	}

	const customLevelConst log.Level = 9

	tt := []testCase{
		{
			lvl:    log.Level(1),
			expect: "DEBUG-9",
		},
		{
			lvl:    customLevelConst,
			expect: "DEBUG-1",
		},
		{
			lvl:    log.Level(uint16(19)),
			expect: "VERBOSE-1",
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
			assert.Equal(t, tc.expect, val)
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
			expect: []byte{34, 67, 82, 73, 84, 73, 67, 65, 76, 43, 50, 48, 34},
		},
	}
	for _, tc := range tt {
		tn := fmt.Sprintf("level=%d", tc.lvl)
		t.Run(tn, func(t *testing.T) {
			got, err := tc.lvl.MarshalJSON()
			assert.Equal(t, tc.expect, got)
			assert.Nil(t, err)
		})
	}
}

func TestLevelInterfaces(t *testing.T) {
	assert.Implements(t, (*fmt.Stringer)(nil), new(log.Level), "log.Level Stringer")
	assert.Implements(t, (*json.Marshaler)(nil), new(log.Level), "log.Level MarshalJSON")
}
