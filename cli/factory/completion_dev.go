//go:build dev

package factory

import (
	"github.com/tprasadtp/pkg/cli/cobra"
)

var completionCmd *cobra.Command = &cobra.Command{
	Use:          "completion --shell [SHELL] --output [FILE]",
	Short:        "Generate shell autocompletion",
	Args:         cobra.NoArgs,
	Hidden:       true,
	SilenceUsage: true,
}
