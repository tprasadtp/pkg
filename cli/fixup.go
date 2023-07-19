package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// This is a workaround for cobra bugs which are unlikely to be fixed
// or have not been fixed yet. You provide root command to this function
// and it will fixup bugs in cobra and annoyances.
// Though backward incompatible changes are avoided, it cannot be guaranteed.
//   - This will fix https://github.com/spf13/cobra/issues/706
//     which does not return an error on unknown sub-commands.
//     This function fixes it by adding a RunE to the command.
//   - This also sets DisableAutoGenTag to true on all commands
//     and sub-commands as is makes output of docs command
//     deterministic.
//   - Disables default completion command. Use [NewCompletionCmd]
//     instead. [NewCompletionCmd] can handle writing output to a file
//     without shell redirection. This is helpful as it can be
//     used with go generate on all platforms and helps building os packages.
func FixCobraBehavior(cmd *cobra.Command) {
	// Disable completion command
	cmd.Root().CompletionOptions.DisableDefaultCmd = true
	// Iterate over all child commands
	for _, cmd := range cmd.Commands() {
		// Only apply fix if child command does not define Run or RunE.
		if cmd.HasSubCommands() {
			if cmd.Run == nil && cmd.RunE == nil {
				cmd.RunE = func(c *cobra.Command, args []string) error {
					if len(args) == 0 {
						return fmt.Errorf("please provide a valid sub-command for %s", c.Name())
					}
					return fmt.Errorf("unknown sub-command %s for %s", args[0], c.Name())
				}
			}
		}
		// Recursively run this function.
		FixCobraBehavior(cmd)
	}
}
