package color

import (
	"fmt"
	"testing"

	"golang.org/x/sys/windows"
)

func TestWinBuild(t *testing.T) {
	type testCase struct {
		name         string
		majorVersion uint32
		minorVersion uint32
		buildNumber  uint32
		tty          bool
		expect       bool
	}

	var tt = []testCase{
		{
			name:         "Windows-Server-2022",
			majorVersion: 10,
			buildNumber:  20348,
			tty:          true,
			expect:       true,
		},
		{
			name:         "Windows-Server-2022",
			majorVersion: 10,
			buildNumber:  20348,
			tty:          false,
			expect:       false,
		},
		{
			name:         "Windows-Server-2019",
			majorVersion: 10,
			buildNumber:  17763,
			tty:          true,
			expect:       true,
		},
		{
			name:         "Windows-Server-2019",
			majorVersion: 10,
			buildNumber:  17763,
			tty:          false,
			expect:       false,
		},
		{
			name:         "Windows-Server-2016",
			majorVersion: 10,
			buildNumber:  14393,
			tty:          true,
			expect:       false,
		},
		{
			name:         "Windows-Server-2016",
			majorVersion: 10,
			buildNumber:  14393,
			tty:          false,
			expect:       false,
		},
		// Windows 11
		{
			name:         "Windows-11-21H2",
			majorVersion: 10,
			buildNumber:  19045,
			tty:          true,
			expect:       true,
		},
		{
			name:         "Windows-11-21H2",
			majorVersion: 10,
			buildNumber:  22000,
			tty:          false,
			expect:       false,
		},
		// Windows 10
		{
			name:         "Windows-10-22H2",
			majorVersion: 10,
			buildNumber:  22000,
			tty:          true,
			expect:       true,
		},
		{
			name:         "Windows-10-22H2",
			majorVersion: 10,
			buildNumber:  19045,
			tty:          false,
			expect:       false,
		},
		{
			name:         "Windows-10-21H2",
			majorVersion: 10,
			buildNumber:  19044,
			tty:          true,
			expect:       true,
		},
		{
			name:         "Windows-10-21H2",
			majorVersion: 10,
			buildNumber:  19044,
			tty:          false,
			expect:       false,
		},
		// Windows 10 RS3
		{
			name:         "Windows-10-1709",
			majorVersion: 10,
			buildNumber:  16299,
			tty:          true,
			expect:       true,
		},
		{
			name:         "Windows-10-1709",
			majorVersion: 10,
			buildNumber:  16299,
			tty:          false,
			expect:       false,
		},
		// Windows 7
		{
			name:         "Windows-7",
			majorVersion: 6,
			minorVersion: 1,
			buildNumber:  7601,
			tty:          false,
			expect:       false,
		},
		{
			name:         "Windows-8.1",
			majorVersion: 6,
			minorVersion: 3,
			buildNumber:  9600,
			tty:          false,
			expect:       false,
		},
	}

	for _, tc := range tt {
		tn := fmt.Sprintf("%s/terminal=%t", tc.name, tc.tty)
		t.Run(tn, func(t *testing.T) {

			// Simply call windows api again to reset everything.
			t.Cleanup(func() {
				osVersion = windows.RtlGetVersion()
			})

			if tc.buildNumber != 0 {
				osVersion.BuildNumber = tc.buildNumber
			}
			if tc.majorVersion != 0 {
				osVersion.MajorVersion = tc.majorVersion
			}
			if tc.minorVersion != 0 {
				osVersion.MinorVersion = tc.minorVersion
			}
			val := isColorable("auto", tc.tty)
			if tc.expect != val {
				t.Errorf("%s => got=%v, want=%v", tn, val, tc.expect)
			}
		})
	}
}

func TestWinBuildEnvOverride(t *testing.T) {
	type testCase struct {
		CI           string
		name         string
		majorVersion uint32
		minorVersion uint32
		buildNumber  uint32
		tty          bool
		expect       bool
	}

	var tt = []testCase{
		// Windows 10 RS3
		{
			name:         "Windows-10-1709",
			CI:           "true",
			majorVersion: 10,
			buildNumber:  16299,
			tty:          true,
			expect:       true,
		},
		{
			name:         "Windows-10-1709",
			CI:           "true",
			majorVersion: 10,
			buildNumber:  16299,
			tty:          false,
			expect:       true,
		},
		// Windows 7
		{
			name:         "Windows-7",
			CI:           "true",
			majorVersion: 6,
			minorVersion: 1,
			buildNumber:  7601,
			tty:          false,
			expect:       false,
		},
		{
			name:         "Windows-7",
			CI:           "true",
			majorVersion: 6,
			minorVersion: 1,
			buildNumber:  7601,
			tty:          true,
			expect:       false,
		},
	}

	for _, tc := range tt {
		tn := fmt.Sprintf("%s/terminal=%t", tc.name, tc.tty)
		t.Run(tn, func(t *testing.T) {
			// Simply call windows api again to reset everything.
			t.Cleanup(func() {
				osVersion = windows.RtlGetVersion()
			})

			t.Setenv("CI", "true")

			if tc.buildNumber != 0 {
				osVersion.BuildNumber = tc.buildNumber
			}
			if tc.majorVersion != 0 {
				osVersion.MajorVersion = tc.majorVersion
			}
			if tc.minorVersion != 0 {
				osVersion.MinorVersion = tc.minorVersion
			}
			val := isColorable("auto", tc.tty)
			if tc.expect != val {
				t.Errorf("%s => got=%v, want=%v", tn, val, tc.expect)
			}
		})
	}
}
