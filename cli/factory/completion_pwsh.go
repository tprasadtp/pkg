//go:build dev

package factory

import (
	"fmt"
	"os"

	"github.com/tprasadtp/pkg/cli"
)

// PowershellCompletionCmd returns a cobra command for
// generating PowershellCompletionCmd completion.
func NewPwshCompletionCmd(programName string) *cli.Command {
	cmd := &cli.Command{
		Use:          "powershell",
		Aliases:      []string{"pwsh", "ps"},
		Args:         cli.MaximumNArgs(1),
		SilenceUsage: true,
		Short:        "Generate autocompletion for powershell",
		Long: fmt.Sprintf(`Generate autocompletion for powershell.

To load completions in your current shell session:
PS C:\> %[1]s completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of
the above command to your powershell profile.
`, programName),
		RunE: func(cmd *cli.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
			} else {
				if args[0] == "" || args[0] == "-" {
					return cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
				}
				return cmd.Root().GenPowerShellCompletionFileWithDesc(args[0])
			}
		},
	}
	return cmd
}
