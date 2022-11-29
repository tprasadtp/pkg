//go:build dev

package factory

import (
	"fmt"
	"os"

	"github.com/tprasadtp/pkg/cli"
)

// FishCompletionCmd returns a cobra command for generating fish completion.
func NewFishCompletionCmd(programName string) *cli.Command {
	cmd := &cli.Command{
		Use:          "fish [FILE]",
		Args:         cli.MaximumNArgs(1),
		SilenceUsage: true,
		Short:        "Generate autocompletion for fish",
		Long: fmt.Sprintf(`Generate autocompletion for fish.

To load completions in your current fish session:
    %[1]s completion fish | source
To load completions for every fish session, execute once:
    %[1]s completion fish ~/.config/fish/completions/%[1]s.fish

You will need to start a new shell for this setup to take effect.
`, programName),
		RunE: func(cmd *cli.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Root().GenFishCompletion(os.Stdout)
			} else {
				if args[0] == "" || args[0] == "-" {
					return cmd.Root().GenFishCompletion(os.Stdout)
				}
				return cmd.Root().GenFishCompletionFile(args[0])
			}
		},
	}
	return cmd
}
