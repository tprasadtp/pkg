package testcli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Root -> command1 -> subcommand2 Flags.
var (
	Subcommand2RequiredFlag string
)

// Root -> command1 -> subcommand2.
func Command1Subcommand2() *cobra.Command {
	subcommand2 := &cobra.Command{
		Use:   "subcommand2",
		Short: "This is subcommand2 (from subcommand2) short description.",
		Long: `This is subcommand2 (from subcommand2) long description

This can span multiple lines.
Line 1
Line 2
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(cmd.Root().OutOrStdout(), "running - %s", cmd.CommandPath())
			return nil
		},
	}
	subcommand2.Flags().StringVarP(
		&Subcommand2RequiredFlag,
		"subcommand2-required-flag",
		"r",
		"",
		"subcommand2-required-flag (from subcommand2)",
	)
	_ = subcommand2.MarkFlagRequired("subcommand2-required-flag")
	return subcommand2
}
