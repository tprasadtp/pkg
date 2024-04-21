// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: MIT

package command

import (
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/tprasadtp/knit/internal/testutils"
)

func TestFixCobraBehavior(t *testing.T) {
	root := cobra.Command{
		Use:   "root",
		Short: "Root command",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	// Command has subcommands but does nothing.
	command := cobra.Command{
		Use:   "command",
		Short: "command",
	}

	// Subcommand of Command
	subcommand := cobra.Command{
		Use:   "sub-command",
		Short: "sub-command",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, _ []string) {
			fmt.Fprintln(cmd.OutOrStdout(), "Running root -> command -> sub-command")
		},
	}

	command.AddCommand(&subcommand)
	root.AddCommand(&command)

	tt := []struct {
		Name string
		ok   bool
		Args []string
	}{
		{
			Name: "root-command-no-args",
			ok:   true,
		},
		// command has subcommands, so it must error
		// when a subcommand is not provided.
		{
			Name: "command-only-when-it-has-subcommands",
			Args: []string{"command"},
		},
		// command has subcommands, so it must error
		// when a invalid sub-command is provided.
		{
			Name: "command-invalid-subcommand",
			Args: []string{"command", "invalid-subcommand"},
		},
		{
			Name: "command-when-it-has-subcommands-help-flag",
			Args: []string{"command", "--help"},
			ok:   true,
		},
		{
			Name: "command-when-it-has-valid-subcommand",
			Args: []string{"command", "sub-command"},
			ok:   true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			stdout := testutils.NewOutputLogger(t, "stdout")
			stderr := testutils.NewOutputLogger(t, "stderr")
			defer stdout.Close()
			defer stderr.Close()

			root.SetOut(stdout)
			root.SetErr(stderr)
			fixCobraBehavior(&root)

			root.SetArgs(tc.Args)

			err := root.Execute()
			if tc.ok {
				if err != nil {
					t.Errorf("expected no error, got %s", err)
				}
			} else {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
			}
		})
	}
}
