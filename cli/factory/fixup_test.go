package factory_test

import (
	"io"
	"testing"

	"github.com/spf13/cobra"
	"github.com/tprasadtp/pkg/cli/factory"
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
		// completion command has subcommands, so it must error
		// when a subcommand is not provided.
		{
			Name:   "command-only-when-it-has-subcommands",
			RError: true,
			Args:   []string{"completion"},
		},
		{
			Name:   "command-when-it-has-subcommands-help-flag",
			RError: false,
			Args:   []string{"completion", "--help"},
		},
		{
			Name:   "command-when-it-has-valid-subcommand",
			RError: false,
			Args:   []string{"completion", "bash", "-"},
		},
		// completion command has subcommands, so it must error
		// when a invalid sub-command is provided.
		{
			Name:   "command-invalid-subcommand",
			RError: true,
			Args:   []string{"completion", "invalid-shell"},
		},
	}
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			root := &cobra.Command{
				Use:   "blackhole-entropy",
				Short: "Black Hole Entropy CLI",
				Long:  "CLI to seed system's PRNG with entropy from M87 Black Hole",
			}
			root.SetOut(io.Discard)
			root.SetErr(io.Discard)
			root.AddCommand(factory.NewCompletionCmd(root.Name()))
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
