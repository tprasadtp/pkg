package color

import (
	"fmt"
	"testing"
)

func TestFlag(t *testing.T) {
	type testCase struct {
		flag   string
		tty    bool
		expect bool
	}

	var tt = []testCase{
		{
			flag:   "always",
			tty:    false,
			expect: true,
		},
		{
			flag:   "ALWAYS",
			tty:    false,
			expect: true,
		},
		{
			flag:   "alWaYs",
			tty:    false,
			expect: true,
		},
		{
			flag:   " alWaYs ",
			tty:    false,
			expect: true,
		},
		{
			flag:   "ALWAYS  ",
			tty:    false,
			expect: true,
		},
		// Force
		{
			flag:   "force",
			tty:    false,
			expect: true,
		},
		{
			flag:   "FORCE",
			tty:    false,
			expect: true,
		},
		{
			flag:   "FoRCE",
			tty:    false,
			expect: true,
		},
		{
			flag:   " FoRCE ",
			tty:    false,
			expect: true,
		},
		{
			flag:   "FORCE  ",
			tty:    false,
			expect: true,
		},
		// Never
		{
			flag:   "never",
			tty:    false,
			expect: false,
		},
		{
			flag:   "NEVER",
			tty:    false,
			expect: false,
		},
		{
			flag:   "NeVeR",
			tty:    false,
			expect: false,
		},
		{
			flag:   " NEvEr ",
			tty:    false,
			expect: false,
		},
		{
			flag:   "NEVER  ",
			tty:    false,
			expect: false,
		},
		// False
		{
			flag:   "false",
			tty:    false,
			expect: false,
		},
		{
			flag:   "False",
			tty:    false,
			expect: false,
		},
		{
			flag:   "FaLSE",
			tty:    false,
			expect: false,
		},
		{
			flag:   " FaLSE ",
			tty:    false,
			expect: false,
		},
		{
			flag:   "False  ",
			tty:    false,
			expect: false,
		},
		// no
		{
			flag:   "no",
			tty:    false,
			expect: false,
		},
		{
			flag:   "No",
			tty:    false,
			expect: false,
		},
		{
			flag:   "NO",
			tty:    false,
			expect: false,
		},
		{
			flag:   " no ",
			tty:    false,
			expect: false,
		},
		{
			flag:   "No  ",
			tty:    false,
			expect: false,
		},
	}

	for _, tc := range tt {
		tn := fmt.Sprintf("flag=%s,tty=%t", tc.flag, tc.tty)
		t.Run(tn, func(t *testing.T) {
			val := isColorable(tc.flag, tc.tty)
			if tc.expect != val {
				t.Errorf("%s => got=%v, want=%v", tn, val, tc.expect)
			}
		})
	}
}

func TestEnvVariables(t *testing.T) {
	// This test MUST not be parallel
	type testCase struct {
		CLICOLOR_FORCE string
		NO_COLOR       string
		CLICOLOR       string
		CI             string
		tty            bool
		expect         bool
	}

	var tt = []testCase{
		// CLICOLOR_FORCE (not defined or empty or zero)
		{
			CLICOLOR_FORCE: "",
			tty:            false,
			expect:         false,
		},
		{
			CLICOLOR_FORCE: "",
			tty:            true,
			expect:         true,
		},
		{
			CLICOLOR_FORCE: "0",
			tty:            false,
			expect:         false,
		},
		{
			CLICOLOR_FORCE: "0",
			tty:            true,
			expect:         true,
		},
		// CLICOLOR_FORCE (defined!=0)
		{
			CLICOLOR_FORCE: "1",
			tty:            true,
			expect:         true,
		},
		{
			CLICOLOR_FORCE: "1",
			tty:            false,
			expect:         true,
		},
		{
			CLICOLOR_FORCE: "non-zero",
			tty:            true,
			expect:         true,
		},
		{
			CLICOLOR_FORCE: "non-zero",
			tty:            false,
			expect:         true,
		},
		{
			CLICOLOR_FORCE: "non-zero-with space",
			tty:            false,
			expect:         true,
		},
		//NO_COLOR (not defined or empty)
		{
			NO_COLOR: "",
			tty:      false,
			expect:   false,
		},
		{
			NO_COLOR: "",
			tty:      true,
			expect:   true,
		},
		{
			NO_COLOR: "  ",
			tty:      false,
			expect:   false,
		},
		{
			NO_COLOR: "  ",
			tty:      true,
			expect:   false,
		},
		//NO_COLOR (defined or not empty)
		{
			NO_COLOR: "defined",
			tty:      false,
			expect:   false,
		},
		{
			NO_COLOR: "defined",
			tty:      true,
			expect:   false,
		},
		{
			NO_COLOR: "  ",
			tty:      false,
			expect:   false,
		},
		{
			NO_COLOR: "  ",
			tty:      true,
			expect:   false,
		},
		// CI
		{
			CI:     "true",
			tty:    true,
			expect: true,
		},
		{
			CI:     "true",
			tty:    false,
			expect: true,
		},
		{
			CI:     "none",
			tty:    true,
			expect: true,
		},
		{
			CI:     "none",
			tty:    false,
			expect: false,
		},
	}

	// Eliminate side effects from TERM
	// They only apply on linux, and are tested separately.
	t.Setenv("TERM", "")

	for _, tc := range tt {
		tn := fmt.Sprintf(
			"CLICOLOR_FORCE=%s,NO_COLOR=%s,CLICOLOR=%s,tty=%t",
			tc.CLICOLOR_FORCE,
			tc.NO_COLOR,
			tc.CLICOLOR,
			tc.tty,
		)
		t.Run(tn, func(t *testing.T) {
			if tc.CLICOLOR_FORCE != "" {
				t.Setenv("CLICOLOR_FORCE", tc.CLICOLOR_FORCE)
			} else {
				t.Setenv("CLICOLOR_FORCE", "")
			}
			if tc.CLICOLOR != "undefined" {
				t.Setenv("CLICOLOR", tc.CLICOLOR)
			} else {
				t.Setenv("CLICOLOR", "")
			}
			if tc.NO_COLOR != "" {
				t.Setenv("NO_COLOR", tc.NO_COLOR)
			} else {
				t.Setenv("NO_COLOR", "")
			}
			if tc.CI != "" {
				t.Setenv("CI", tc.CI)
			} else {
				t.Setenv("CI", "")
			}
			val := isColorable("auto", tc.tty)
			if tc.expect != val {
				t.Errorf("%s => got=%v, want=%v", tn, val, tc.expect)
			}
		})
	}
}
