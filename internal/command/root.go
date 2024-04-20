// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: MIT

package command

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/spf13/cobra"
	"github.com/tprasadtp/knit/internal/command/semver"
	"github.com/tprasadtp/knit/internal/command/version"
	"github.com/tprasadtp/knit/internal/log"
)

func RootCommand(logger *slog.Logger) *cobra.Command {
	var verbose bool

	cmd := &cobra.Command{
		Use:               "knit",
		Short:             "A Toolkit for building docker images",
		Version:           version.Version(),
		DisableAutoGenTag: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		// Set logger based on slog for all go-containerregistry operations.
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			logger := slog.Default()
			ctx := cmd.Context()

			if testing.Testing() {

			}

			cmd.SetContext(log.WithContext(cmd.Context(), slog.Default()))
			return nil
		},
	}
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose logs")
	cmd.AddCommand(semver.NewCommand())
	cmd.AddCommand(version.NewVersionCmd())
	FixCobraBehavior(cmd)
	return cmd
}

// This is a workaround for cobra bug(s) which are unlikely to be fixed
// or have not been fixed yet. You provide root command to this function
// and it will fixup bugs in cobra and annoyances.
// Though backward incompatible changes are avoided, it cannot be guaranteed.
//   - This will fix https://github.com/spf13/cobra/issues/706
//     which does not return an error on unknown sub-commands.
//     This function fixes it by adding a RunE to the command.
func FixCobraBehavior(cmd *cobra.Command) {
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
