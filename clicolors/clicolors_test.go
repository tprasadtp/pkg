package clicolors

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnableColors(t *testing.T) {
	tests := []struct {
		name               string
		expected           bool
		isTerminal         bool
		disableFlag        bool
		clicolorIsSet      bool
		clicolorForceIsSet bool
		clicolorVal        string
		clicolorForceVal   string
	}{
		{
			name:               "Terminal=false",
			expected:           false,
			isTerminal:         false,
			disableFlag:        false,
			clicolorIsSet:      false,
			clicolorForceIsSet: false,
			clicolorVal:        "",
			clicolorForceVal:   "",
		},
		{
			name:               "Terminal=true",
			expected:           true,
			isTerminal:         true,
			disableFlag:        false,
			clicolorIsSet:      false,
			clicolorForceIsSet: false,
			clicolorVal:        "",
			clicolorForceVal:   "",
		},
		{
			name:               "DisableFlag=true",
			expected:           false,
			isTerminal:         false,
			disableFlag:        true,
			clicolorIsSet:      false,
			clicolorForceIsSet: false,
			clicolorVal:        "",
			clicolorForceVal:   "",
		},
		{
			name:               "CLICOLOR=1",
			expected:           false,
			isTerminal:         false,
			disableFlag:        false,
			clicolorIsSet:      true,
			clicolorForceIsSet: false,
			clicolorVal:        "1",
			clicolorForceVal:   "",
		},
		{
			name:               "CLICOLOR=1+Terminal=true",
			expected:           true,
			isTerminal:         true,
			disableFlag:        false,
			clicolorIsSet:      true,
			clicolorForceIsSet: false,
			clicolorVal:        "1",
			clicolorForceVal:   "",
		},
		{
			name:               "CLICOLOR=1+Terminal=true,DisableFlag=true",
			expected:           false,
			isTerminal:         true,
			disableFlag:        true,
			clicolorIsSet:      true,
			clicolorForceIsSet: false,
			clicolorVal:        "1",
			clicolorForceVal:   "",
		},
		{
			name:               "CLICOLOR=1+Terminal=false,DisableFlag=true",
			expected:           false,
			isTerminal:         true,
			disableFlag:        true,
			clicolorIsSet:      true,
			clicolorForceIsSet: false,
			clicolorVal:        "1",
			clicolorForceVal:   "",
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
			expected:           false,
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
	}
	cleanenv := func() {
		os.Unsetenv("CLICOLOR")
		os.Unsetenv("CLICOLOR_FORCE")
	}

	defer cleanenv()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cleanenv()
			if tc.clicolorIsSet {
				os.Setenv("CLICOLOR", tc.clicolorVal)
			}

			if tc.clicolorForceIsSet {
				os.Setenv("CLICOLOR_FORCE", tc.clicolorForceVal)
			}
			v := EnableColors(tc.disableFlag, tc.isTerminal)
			assert.Equal(t, tc.expected, v)
		})
	}
}
