package testcli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Root -> command1 -> subcommand1 Flags.
var (
	Subcommand1Flag bool
)

// Root -> command1 -> subcommand1.
func Command1Subcommand1() *cobra.Command {
	subcommand1 := &cobra.Command{
		Use:   "subcommand1",
		Short: "This is subcommand1 short description.",
		Long: `This is subcommand1 long description

This can span multiple lines.

- Item 1
- Item 2

> Markdown Hint

`,
		Example: `
test-cli command1 --persistent-flag subcommand1 --subcommand1-flag
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(cmd.Root().OutOrStdout(), "running - %s", cmd.CommandPath())
			return nil
		},
	}
	subcommand1.Flags().BoolVar(
		&Subcommand1Flag,
		"subcommand1-flag",
		false,
		"subcommand1-flag (from subcommand1)",
	)
	return subcommand1
}
