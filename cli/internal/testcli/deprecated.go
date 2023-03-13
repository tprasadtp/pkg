package testcli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Tests search for this to ensure no hidden commands or flags
// are in docs.
const HiddenToken = "54dac5afe1fcac2f65c059fc97b44a58"

func DeprecatedCmd() *cobra.Command {
	// Root:test-cli -> command-hidden
	cmd := &cobra.Command{
		Use:   "deprecated",
		Short: "This is deprecated short description. 54dac5afe1fcac2f65c059fc97b44a58",
		Long: `This is deprecated long description

How dare thee useth a deprecated commad?
Thee dare to usetht a command which is not shown?

All bugs reports from its useth I shalt ignore!

token=54dac5afe1fcac2f65c059fc97b44a58
`,
		Example: `
test-cli deprecated
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(cmd.Root().OutOrStdout(), "running - %s", cmd.CommandPath())
			return nil
		},
		Deprecated: "deprecated-command",
	}
	return cmd
}
