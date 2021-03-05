package clicolors

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnableColors(t *testing.T) {
	origNoColor, origNoColorSet := os.LookupEnv("NO_COLOR")
	origClicolor, origClicolorSet := os.LookupEnv("CLICOLOR")
	origClicolorForce, origClicolorForceSet := os.LookupEnv("CLICOLOR_FORCE")
	origTerm := os.Getenv("TERM")

	t.Cleanup(func() {
		if origNoColorSet {
			os.Setenv("NO_COLOR", origNoColor)
		} else {
			os.Unsetenv("NO_COLOR")
		}

		if origClicolorSet {
			os.Setenv("NO_COLOR", origClicolor)
		} else {
			os.Unsetenv("NO_COLOR")
		}

		if origClicolorForceSet {
			os.Setenv("NO_COLOR", origClicolorForce)
		} else {
			os.Unsetenv("NO_COLOR")
		}

		os.Setenv("TERM", origTerm)
	})

	tests := []struct {
		name               string
		expected           bool
		isTerminal         bool
		disableFlag        bool
		clicolorIsSet      bool
		clicolorForceIsSet bool
		clicolorVal        string
		clicolorForceVal   string
		noColorIsSet       bool
		noColorVal         string
		termDumb           bool
	}{
		{
			name:     "Terminal=false",
			expected: false,
		},
		{
			name:       "Terminal=true",
			expected:   true,
			isTerminal: true,
		},
		{
			name:     "Terminal=false,Term=DUMB",
			termDumb: true,
			expected: false,
		},
		{
			name:       "Terminal=true,Term=DUMB",
			expected:   false,
			termDumb:   true,
			isTerminal: true,
		},
		{
			name:        "DisableFlag=true",
			expected:    false,
			isTerminal:  false,
			disableFlag: true,
		},
		{
			name:          "CLICOLOR=1+Terminal=false",
			expected:      false,
			isTerminal:    false,
			disableFlag:   false,
			clicolorIsSet: true,
			clicolorVal:   "1",
		},
		{
			name:          "CLICOLOR=1+Terminal=true",
			expected:      true,
			isTerminal:    true,
			disableFlag:   false,
			clicolorIsSet: true,
			clicolorVal:   "1",
		},
		{
			name:          "CLICOLOR=1+Terminal=true,DisableFlag=true",
			expected:      false,
			isTerminal:    true,
			disableFlag:   true,
			clicolorIsSet: true,
			clicolorVal:   "1",
		},
		{
			name:          "CLICOLOR=1+Terminal=false,DisableFlag=true",
			expected:      false,
			isTerminal:    true,
			disableFlag:   true,
			clicolorIsSet: true,
			clicolorVal:   "1",
		},
		{
			name:               "CLICOLOR=1+Terminal=true,DisableFlag=true,FORCE=1",
			expected:           false,
			isTerminal:         true,
			disableFlag:        true,
			clicolorIsSet:      true,
			clicolorForceIsSet: true,
			clicolorVal:        "1",
			clicolorForceVal:   "1",
		},
		{
			name:               "CLICOLOR=1+Terminal=true,FORCE=1",
			expected:           true,
			isTerminal:         true,
			disableFlag:        false,
			clicolorIsSet:      true,
			clicolorForceIsSet: true,
			clicolorVal:        "1",
			clicolorForceVal:   "1",
		},
		{
			name:               "CLICOLOR=1+Terminal=false,FORCE=1",
			expected:           true,
			isTerminal:         false,
			disableFlag:        false,
			clicolorIsSet:      true,
			clicolorForceIsSet: true,
			clicolorVal:        "1",
			clicolorForceVal:   "1",
		},
		{
			name:               "CLICOLOR=1+Terminal=false,FORCE=0",
			expected:           false,
			isTerminal:         false,
			disableFlag:        false,
			clicolorIsSet:      true,
			clicolorForceIsSet: true,
			clicolorVal:        "0",
			clicolorForceVal:   "0",
		},
		{
			name:               "CLICOLOR=1+Terminal=true,FORCE=0",
			expected:           true,
			isTerminal:         true,
			disableFlag:        false,
			clicolorIsSet:      true,
			clicolorForceIsSet: true,
			clicolorVal:        "1",
			clicolorForceVal:   "0",
		},
		{
			name:               "CLICOLOR=0+Terminal=false,FORCE=0",
			expected:           false,
			isTerminal:         false,
			disableFlag:        false,
			clicolorIsSet:      true,
			clicolorForceIsSet: true,
			clicolorVal:        "0",
			clicolorForceVal:   "0",
		},
		{
			name:               "CLICOLOR=0+Terminal=false,FORCE=1",
			expected:           true,
			isTerminal:         false,
			disableFlag:        false,
			clicolorIsSet:      true,
			clicolorForceIsSet: true,
			clicolorVal:        "0",
			clicolorForceVal:   "1",
		},
		{
			name:               "CLICOLOR=0+Terminal=true,FORCE=1",
			expected:           true,
			isTerminal:         true,
			disableFlag:        false,
			clicolorIsSet:      true,
			clicolorForceIsSet: true,
			clicolorVal:        "0",
			clicolorForceVal:   "1",
		},
		{
			name:               "CLICOLOR=0+Terminal=false,FORCE=1,Flag",
			expected:           false,
			isTerminal:         false,
			disableFlag:        true,
			clicolorIsSet:      true,
			clicolorForceIsSet: true,
			clicolorVal:        "0",
			clicolorForceVal:   "1",
		},
		{
			name:               "CLICOLOR=0+Terminal=true,FORCE=1,disableFalg=true",
			expected:           false,
			isTerminal:         true,
			disableFlag:        true,
			clicolorIsSet:      true,
			clicolorForceIsSet: true,
			clicolorVal:        "0",
			clicolorForceVal:   "1",
		},
		{
			name:         "NO_COLOR=0",
			expected:     false,
			noColorIsSet: true,
			noColorVal:   "0",
		},
		{
			name:         "NO_COLOR=0,Terminal=true",
			expected:     false,
			isTerminal:   true,
			noColorIsSet: true,
			noColorVal:   "0",
		},
		{
			name:               "NO_COLOR=0,FORCE=1(ForceOverride)",
			expected:           true,
			isTerminal:         false,
			noColorIsSet:       true,
			noColorVal:         "0",
			clicolorForceIsSet: true,
			clicolorForceVal:   "1",
		},
		{
			name:               "NO_COLOR=0,FORCE=0",
			expected:           false,
			isTerminal:         true,
			noColorIsSet:       true,
			noColorVal:         "0",
			clicolorForceIsSet: true,
			clicolorForceVal:   "0",
		},
		{
			name:          "CLICOLOR=1+Terminal=true,TERM=dumb",
			expected:      false,
			isTerminal:    true,
			termDumb:      true,
			disableFlag:   false,
			clicolorIsSet: true,
			clicolorVal:   "1",
		},
	}

	cleanenv := func() {
		os.Unsetenv("CLICOLOR")
		os.Unsetenv("CLICOLOR_FORCE")
		os.Unsetenv("NO_COLOR")
		// set TERM to original
		os.Setenv("TERM", origTerm)
	}

	defer cleanenv()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cleanenv()
			if tc.clicolorIsSet {
				os.Setenv("CLICOLOR", tc.clicolorVal)
			}

			if tc.noColorIsSet {
				os.Setenv("NO_COLOR", tc.noColorVal)
			}

			if tc.clicolorForceIsSet {
				os.Setenv("CLICOLOR_FORCE", tc.clicolorForceVal)
			}

			// sets TERM to dumb
			if tc.termDumb {
				os.Setenv("TERM", "dumb")
			}

			v := EnableColors(tc.disableFlag, tc.isTerminal)
			assert.Equal(t, tc.expected, v)
		})
	}
}

func TestEnvColorDisabled(t *testing.T) {
	origNoColor, origNoColorSet := os.LookupEnv("NO_COLOR")
	origClicolor, origClicolorSet := os.LookupEnv("CLICOLOR")
	t.Cleanup(func() {
		if origNoColorSet {
			os.Setenv("NO_COLOR", origNoColor)
		} else {
			os.Unsetenv("NO_COLOR")
		}

		if origClicolorSet {
			os.Setenv("NO_COLOR", origClicolor)
		} else {
			os.Unsetenv("NO_COLOR")
		}
	})

	tests := []struct {
		name     string
		expected bool
		cliColor string
		noColor  string
	}{
		{
			name:     "Default",
			expected: false,
			cliColor: "",
			noColor:  "",
		},
		{
			name:     "CLICOLOR=0",
			expected: true,
			cliColor: "0",
			noColor:  "",
		},
		{
			name:     "NO_COLOR=1",
			expected: true,
			cliColor: "",
			noColor:  "1",
		},
		{
			name:     "NO_COLOR=true",
			expected: true,
			cliColor: "",
			noColor:  "true",
		},
		{
			name:     "CLICOLOR=0,NO_COLOR=1",
			expected: true,
			cliColor: "0",
			noColor:  "1",
		},
		{
			name:     "CLICOLOR=1",
			expected: false,
			cliColor: "1",
			noColor:  "",
		},
	}
	cleanenv := func() {
		os.Unsetenv("NO_COLOR")
		os.Unsetenv("CLICOLOR")
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cleanenv()
			os.Setenv("CLICOLOR", tc.cliColor)
			os.Setenv("NO_COLOR", tc.noColor)

			v := EnvColorDisabled()
			assert.Equal(t, tc.expected, v)
		})
	}
}

func TestEnvColorForced(t *testing.T) {
	origClicolorForce, origClicolorForceSet := os.LookupEnv("CLICOLOR_FORCE")
	t.Cleanup(func() {
		if origClicolorForceSet {
			os.Setenv("NO_COLOR", origClicolorForce)
		} else {
			os.Unsetenv("NO_COLOR")
		}
	})

	tests := []struct {
		name               string
		expected           bool
		clicolorForceIsSet bool
		clicolorForceVal   string
	}{
		{
			name:     "Default",
			expected: false,
		},
		{
			name:               "FORCE=0",
			expected:           false,
			clicolorForceIsSet: true,
			clicolorForceVal:   "0",
		},
		{
			name:               "FORCE=1",
			expected:           true,
			clicolorForceIsSet: true,
			clicolorForceVal:   "1",
		},
		{
			name:               "FORCE=true",
			expected:           true,
			clicolorForceIsSet: true,
			clicolorForceVal:   "true",
		},
	}
	cleanenv := func() {
		os.Unsetenv("CLICOLOR_FORCE")
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cleanenv()
			if tc.clicolorForceIsSet {
				os.Setenv("CLICOLOR_FORCE", tc.clicolorForceVal)
			}
			v := EnvColorForced()
			assert.Equal(t, tc.expected, v)
		})
	}
}

func TestIsDumbTermSet(t *testing.T) {
	origTerm, origTermSet := os.LookupEnv("CLICOLOR_FORCE")
	t.Cleanup(func() {
		if origTermSet {
			os.Setenv("TERM", origTerm)
		} else {
			os.Unsetenv("TERM")
		}
	})

	tests := []struct {
		name     string
		expected bool
		termSet  bool
		termVal  string
	}{
		{
			name:     "clean",
			expected: false,
			termSet:  false,
		},
		{
			name:     "xterm-256color",
			expected: false,
			termSet:  true,
			termVal:  "xterm-256color",
		},
		{
			name:     "dumb",
			expected: true,
			termSet:  true,
			termVal:  "dumb",
		},
	}
	cleanenv := func() {
		os.Unsetenv("TERM")
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cleanenv()
			if tc.termSet {
				os.Setenv("TERM", tc.termVal)
			}
			v := IsDumbTerm()
			assert.Equal(t, tc.expected, v)
		})
	}
}
