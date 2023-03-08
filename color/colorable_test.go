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
		//nolint:revive,stylecheck // This is an env variable.
		envCLICOLOR_FORCE string
		//nolint:revive,stylecheck // This is an env variable.
		envNO_COLOR string
		envCLICOLOR string
		envCI       string
		tty         bool
		expect      bool
	}

	var tt = []testCase{
		// CLICOLOR_FORCE (not defined or empty or zero)
		{
			envCLICOLOR_FORCE: "",
			tty:               false,
			expect:            false,
		},
		{
			envCLICOLOR_FORCE: "",
			tty:               true,
			expect:            true,
		},
		// CLICOLOR_FORCE (defined=0)
		{
			envCLICOLOR_FORCE: "0",
			tty:               false,
			expect:            false,
		},
		{
			envCLICOLOR_FORCE: "0",
			tty:               true,
			expect:            true,
		},
		// CLICOLOR_FORCE (defined!=0)
		{
			envCLICOLOR_FORCE: "1",
			tty:               true,
			expect:            true,
		},
		{
			envCLICOLOR_FORCE: "1",
			tty:               false,
			expect:            true,
		},
		{
			envCLICOLOR_FORCE: "non-zero",
			tty:               true,
			expect:            true,
		},
		{
			envCLICOLOR_FORCE: "non-zero",
			tty:               false,
			expect:            true,
		},
		{
			envCLICOLOR_FORCE: "  ",
			tty:               false,
			expect:            true,
		},
		{
			envCLICOLOR_FORCE: "non-zero",
			tty:               false,
			expect:            true,
		},
		{
			envCLICOLOR_FORCE: "  ",
			tty:               false,
			expect:            true,
		},
		// NO_COLOR (not defined or empty)
		{
			envNO_COLOR: "",
			tty:         false,
			expect:      false,
		},
		{
			envNO_COLOR: "",
			tty:         true,
			expect:      true,
		},
		// NO_COLOR (defined whitespace only)
		{
			envNO_COLOR: "  ",
			tty:         false,
			expect:      false,
		},
		{
			envNO_COLOR: "  ",
			tty:         true,
			expect:      false,
		},
		// NO_COLOR defined and not empty
		{
			envNO_COLOR: "defined",
			tty:         false,
			expect:      false,
		},
		{
			envNO_COLOR: "defined",
			tty:         true,
			expect:      false,
		},
		// CI=true, enables colors even when TTY is not attached
		{
			envCI:  "true",
			tty:    true,
			expect: true,
		},
		{
			envCI:  "true",
			tty:    false,
			expect: true,
		},
		// CI != true
		{
			envCI:  "none",
			tty:    true,
			expect: true,
		},
		{
			envCI:  "none",
			tty:    false,
			expect: false,
		},
	}

	// Eliminate side effects from TERM
	// They only apply on linux, and are tested separately.
	t.Setenv("TERM", "")

	for _, tc := range tt {
		tn := fmt.Sprintf(
			"CI=%s,CLICOLOR_FORCE=%s,NO_COLOR=%s,CLICOLOR=%s,tty=%t",
			tc.envCI,
			tc.envCLICOLOR_FORCE,
			tc.envNO_COLOR,
			tc.envCLICOLOR,
			tc.tty,
		)
		t.Run(tn, func(t *testing.T) {
			if tc.envCLICOLOR_FORCE != "" {
				t.Setenv("CLICOLOR_FORCE", tc.envCLICOLOR_FORCE)
			} else {
				t.Setenv("CLICOLOR_FORCE", "")
			}
			if tc.envCLICOLOR != "undefined" {
				t.Setenv("CLICOLOR", tc.envCLICOLOR)
			} else {
				t.Setenv("CLICOLOR", "")
			}
			if tc.envNO_COLOR != "" {
				t.Setenv("NO_COLOR", tc.envNO_COLOR)
			} else {
				t.Setenv("NO_COLOR", "")
			}
			if tc.envCI != "" {
				t.Setenv("CI", tc.envCI)
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
