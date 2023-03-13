package testcli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func HiddenCmd() *cobra.Command {
	// Root -> hidden
	cmd := &cobra.Command{
		Use:   "hidden",
		Short: "This is hidden short description 54dac5afe1fcac2f65c059fc97b44a58",
		Long: `This is hidden long description

This can span multiple lines.
Line 1
Line 2

54dac5afe1fcac2f65c059fc97b44a58
`,
		Example: `
test-cli hidden 54dac5afe1fcac2f65c059fc97b44a58
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(cmd.Root().OutOrStdout(), "running - %s", cmd.CommandPath())
			return nil
		},
		Hidden: true,
	}
	return cmd
}
