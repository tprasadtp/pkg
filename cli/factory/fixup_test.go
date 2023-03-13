package factory_test

import (
	"io"
	"testing"

	"github.com/tprasadtp/pkg/cli/factory"
	"github.com/tprasadtp/pkg/cli/internal/testcli"
)

func Test_FixupCommands(t *testing.T) {
	type testCase struct {
		Name   string
		RError bool
		Args   []string
	}
	tt := []testCase{
		{
			Name: "root command no args",
		},
		// command1 has subcommands, so it must error
		// when a subcommand is not provided.
		{
			Name:   "command1-only-when-it-has-subcommands",
			RError: true,
			Args:   []string{"command1"},
		},
		{
			Name:   "command1-when-it-has-subcommands-help-flag",
			RError: false,
			Args:   []string{"command1", "--help"},
		},
		{
			Name:   "command-when-it-has-valid-subcommand",
			RError: false,
			Args:   []string{"command1", "subcommand1"},
		},
		// command1 has subcommands, so it must error
		// when a invalid sub-command is provided.
		{
			Name:   "command1-invalid-subcommand",
			RError: true,
			Args:   []string{"command1", "invalid-subcommand"},
		},
	}
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			root := testcli.GetTestCLI()
			root.SetOut(io.Discard)
			root.SetErr(io.Discard)
			factory.FixCobraBehavior(root)
			root.SetArgs(tc.Args)
			err := root.Execute()
			if tc.RError {
				if err == nil {
					t.Fatalf("expected an error, bit got nil")
				}
			} else {
				if err != nil {
					t.Fatalf("expected no error, bit got - %s", err)
				}
			}
		})
	}
}
