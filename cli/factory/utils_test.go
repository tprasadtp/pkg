package factory

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/tprasadtp/pkg/cli/internal/testcli"
)

// mustFindCmd attempts to find command from root. panics if it cannot.
func mustFindCmd(root *cobra.Command, args []string) *cobra.Command {
	c, _, err := root.Find(args)
	if err != nil || c == nil {
		panic(
			fmt.Sprintf("failed to find command %s", strings.Join(args, " ")),
		)
	}
	return c
}

func Test_getSeeAlso(t *testing.T) {
	type testCase struct {
		Name        string
		Cmd         *cobra.Command
		ExpectNames []string // this is ordered!
	}
	root := testcli.GetTestCLI()
	tt := []testCase{
		{
			Name:        "root-command",
			Cmd:         root,
			ExpectNames: []string{"command1", "command2", "command3"},
		},
		{
			Name:        "command1",
			Cmd:         mustFindCmd(root, []string{"command1"}),
			ExpectNames: []string{"command2", "command3", "subcommand1", "subcommand2"},
		},
		{
			Name:        "subcommand1",
			Cmd:         mustFindCmd(root, []string{"command1", "subcommand1"}),
			ExpectNames: []string{"command1", "subcommand2"},
		},
		{
			Name:        "subcommand2",
			Cmd:         mustFindCmd(root, []string{"command1", "subcommand2"}),
			ExpectNames: []string{"command1", "subcommand1"},
		},
		{
			Name:        "command2",
			Cmd:         mustFindCmd(root, []string{"command2"}),
			ExpectNames: []string{"command1", "command3"},
		},
		{
			Name: "hidden",
			Cmd:  mustFindCmd(root, []string{"hidden"}),
		},
		{
			Name: "deprecated",
			Cmd:  mustFindCmd(root, []string{"deprecated"}),
		},
		{
			Name: "nil",
		},
	}
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			got := getSeeAlso(tc.Cmd)
			if len(got) != len(tc.ExpectNames) {
				t.Fatalf(
					"expected %d results, got %d",
					len(tc.ExpectNames),
					len(got))
			}

			for i, item := range got {
				if tc.ExpectNames[i] != item.Name() {
					t.Errorf("at index %d expected %s, got %s", i, tc.ExpectNames[i], item.Name())
				}
			}
		})
	}
}

func Test_getFlags(t *testing.T) {
	type testCase struct {
		Name        string
		Cmd         *cobra.Command
		ExpectFlags []string // this is ordered!
	}
	root := testcli.GetTestCLI()
	// These test cases do not have -h/--help
	// This is because GetTestCLI does not call
	// InitDefaultHelpCmd and InitDefaultHelpFlag.
	tt := []testCase{
		{
			Name:        "root-command",
			Cmd:         root,
			ExpectFlags: []string{"global-flag", "global-string-flag"},
		},
		{
			Name: "command1",
			Cmd:  mustFindCmd(root, []string{"command1"}),
			ExpectFlags: []string{
				"command1-persistent-flag",
				"global-flag",
				"global-string-flag",
			},
		},
		{
			Name: "subcommand1",
			Cmd:  mustFindCmd(root, []string{"command1", "subcommand1"}),
			ExpectFlags: []string{
				"command1-persistent-flag",
				"global-flag",
				"global-string-flag",
				"subcommand1-flag",
			},
		},
		{
			Name: "subcommand2",
			Cmd:  mustFindCmd(root, []string{"command1", "subcommand2"}),
			ExpectFlags: []string{
				"command1-persistent-flag",
				"global-flag",
				"global-string-flag",
				"subcommand2-required-flag",
			},
		},
		{
			Name: "command2",
			Cmd:  mustFindCmd(root, []string{"command2"}),
			ExpectFlags: []string{
				"global-flag",
				"global-string-flag",
				"mutually-exclusive-flag1",
				"mutually-exclusive-flag2",
			},
		},
		{
			Name: "hidden",
			Cmd:  mustFindCmd(root, []string{"hidden"}),
		},
		{
			Name: "deprecated",
			Cmd:  mustFindCmd(root, []string{"deprecated"}),
		},
		{
			Name: "nil",
		},
	}
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			flags := getFlags(tc.Cmd)
			if len(flags) != len(tc.ExpectFlags) {
				t.Fatalf(
					"expected %d flags from %s flagset, got %d",
					len(tc.ExpectFlags),
					tc.Cmd.CommandPath(),
					len(flags),
				)
			}

			for i, item := range flags {
				if tc.ExpectFlags[i] != item.Name {
					t.Errorf("at index %d expected %s flag, got %s", i, tc.ExpectFlags[i], item.Name)
				}
			}
		})
	}
}

func Test_formatGeneratedAt_SOURCE_DATE_EPOCH_ValidUnixTS(t *testing.T) {
	t.Setenv("SOURCE_DATE_EPOCH", "1136239445")
	if formatGeneratedAt(defaultTimeFormat) != "2-Jan-2006" {
		t.Errorf("expected 2-Jan-2006 for SOURCE_DATE_EPOCH=1136239445")
	}
}

func Test_formatGeneratedAt_SOURCE_DATE_EPOCH_ValidUnixTS_NoFmtDefined(t *testing.T) {
	t.Setenv("SOURCE_DATE_EPOCH", "1136239445")
	if formatGeneratedAt("") != defaultTimeFormat {
		t.Errorf("expected 2-Jan-2006 for SOURCE_DATE_EPOCH=1136239445 and undefined format")
	}
}

func Test_formatGeneratedAt_SOURCE_DATE_EPOCH_Invalid(t *testing.T) {
	t.Setenv("SOURCE_DATE_EPOCH", "foo-bar")
	now := time.Now()
	output := formatGeneratedAt(defaultTimeFormat)
	// re-parse output back to time.
	outputTime, _ := time.Parse(defaultTimeFormat, output)
	if outputTime.Sub(now) > time.Second {
		t.Errorf("diff time wrt time.Now is > 1s, when SOURCE_DATE_EPOCH is invalid")
	}
}

func Test_isAutoGenDisabled(t *testing.T) {
	type testCase struct {
		Name     string
		Cmd      *cobra.Command
		Expected bool
	}
	tt := []testCase{
		{
			Name:     "not-disabled-at-root",
			Cmd:      testcli.GetTestCLI(),
			Expected: false,
		},
		{
			Name: "disabled-at-root",
			Cmd: func() *cobra.Command {
				root := testcli.GetTestCLI()
				root.DisableAutoGenTag = true
				return root
			}(),
			Expected: true,
		},
		{
			Name: "disabled-at-command1",
			Cmd: func() *cobra.Command {
				root := testcli.GetTestCLI()
				command1 := mustFindCmd(root, []string{"command1"})
				command1.DisableAutoGenTag = true
				subcommand1 := mustFindCmd(root, []string{"command1", "subcommand1"})
				return subcommand1
			}(),
			Expected: true,
		},
		{
			Name: "disabled-at-root",
			Cmd: func() *cobra.Command {
				root := testcli.GetTestCLI()
				root.DisableAutoGenTag = true
				subcommand1 := mustFindCmd(root, []string{"command1", "subcommand1"})
				return subcommand1
			}(),
			Expected: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			output := isAutoGenDisabled(tc.Cmd)
			if output != tc.Expected {
				t.Errorf("expected %v, got %v", tc.Expected, output)
			}
		})
	}
}
