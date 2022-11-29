//go:build dev

package factory

import (
	"github.com/tprasadtp/pkg/cli"
)

var completionCmd *cli.Command = &cli.Command{
	Use:          "completion --shell [SHELL] --output [FILE]",
	Short:        "Generate shell autocompletion",
	Args:         cli.NoArgs,
	Hidden:       true,
	SilenceUsage: true,
}
