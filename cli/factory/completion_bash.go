//go:build dev

package factory

import (
	"fmt"
	"os"

	"github.com/tprasadtp/pkg/cli"
)

// BashCompletionCommand returns a cobra command for generating bash completion.
func NewBashCompletionCmd(programName string) *cli.Command {
	cmd := &cli.Command{
		Use:          "bash [FILE]",
		Args:         cli.MaximumNArgs(1),
		SilenceUsage: true,
		Short:        "Generate autocompletion for bash",
		Long: fmt.Sprintf(`Generate autocompletion for bash.

To load completions in your current bash session:
    source <(%[1]s completion bash)

To load completions for every bash session, execute once:
- Linux:
    %[1]s completion bash /etc/bash_completion.d/%[1]s
- MacOS:
    %[1]s completion bash /usr/local/etc/bash_completion.d/%[1]s
`, programName),
		RunE: func(cmd *cli.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Root().GenBashCompletion(os.Stdout)
			} else {
				if args[0] == "" || args[0] == "-" {
					return cmd.Root().GenBashCompletion(os.Stdout)
				}
				return cmd.Root().GenBashCompletionFile(args[0])
			}
		},
	}
	return cmd
}
