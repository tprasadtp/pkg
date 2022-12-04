//go:build dev

package factory

import "github.com/spf13/cobra"

var completionCmd *cobra.Command = &cobra.Command{
	Use:          "completion --shell [SHELL] --output [FILE]",
	Short:        "Generate shell autocompletion",
	Args:         cobra.NoArgs,
	Hidden:       true,
	SilenceUsage: true,
}
