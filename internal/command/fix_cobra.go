// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: MIT

package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

// This is a workaround for cobra bug(s) which are unlikely to be fixed
// or have not been fixed yet. You provide root command to this function
// and it will fixup bugs in cobra and annoyances.
// Though backward incompatible changes are avoided, it cannot be guaranteed.
//   - This will fix https://github.com/spf13/cobra/issues/706
//     which does not return an error on unknown sub-commands.
//     This function fixes it by adding a RunE to the command.
//
// Must be called after adding all the subcommands and flags.
func fixCobraBehavior(cmd *cobra.Command) {
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
		fixCobraBehavior(cmd)
	}
}
