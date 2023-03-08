package factory

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewCompletionCmd returns a cobra command named completion
// with all supported shells as subcommands.
func NewCompletionCmd(programName string, hidden ...bool) *cobra.Command {
	var h bool
	if len(hidden) > 0 {
		h = hidden[0]
	}
	cmd := &cobra.Command{
		Use:     "completion SHELL [FILE]",
		Short:   "Generate shell autocompletion",
		Aliases: []string{"complete", "compgen"},
		Args:    cobra.NoArgs,
		Hidden:  h,
	}
	cmd.AddCommand(NewBashCompletionCmd(programName))
	cmd.AddCommand(NewFishCompletionCmd(programName))
	cmd.AddCommand(NewZshCompletionCmd(programName))
	cmd.AddCommand(NewPwshCompletionCmd(programName))
	return cmd
}

// BashCompletionCommand returns a cobra command for generating bash completion.
func NewBashCompletionCmd(programName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "bash [FILE]",
		Args:         cobra.MaximumNArgs(1),
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
		RunE: func(cmd *cobra.Command, args []string) error {
			root := cmd.Root()
			if len(args) == 0 {
				return root.GenBashCompletionV2(root.OutOrStdout(), true)
			}
			if args[0] == "" || args[0] == "-" {
				return root.GenBashCompletionV2(root.OutOrStdout(), true)
			}
			return root.GenBashCompletionFileV2(args[0], true)
		},
	}
	return cmd
}

// ZshCompletionCmd returns a cobra command for generating zsh completion.
func NewZshCompletionCmd(programName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "zsh",
		Args:         cobra.MaximumNArgs(1),
		SilenceUsage: true,
		Short:        "Generate autocompletion for zsh",
		Long: fmt.Sprintf(`Generate autocompletion for zsh.

To load completions in your current zsh session:
    source <(%[1]s completion zsh)

To load completions for every zsh session, execute once:
	%[1]s completion zsh "${fpath[1]}/_%[1]s

You will need to start a new shell for this setup to take effect.
`, programName),
		RunE: func(cmd *cobra.Command, args []string) error {
			root := cmd.Root()
			if len(args) == 0 {
				return root.GenZshCompletion(root.OutOrStdout())
			}
			if args[0] == "" || args[0] == "-" {
				return root.GenZshCompletion(root.OutOrStdout())
			}
			return root.GenZshCompletionFile(args[0])
		},
	}
	return cmd
}

// FishCompletionCmd returns a cobra command for generating fish completion.
func NewFishCompletionCmd(programName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "fish [FILE]",
		Args:         cobra.MaximumNArgs(1),
		SilenceUsage: true,
		Short:        "Generate shell completions for fish",
		Long: fmt.Sprintf(`Generate shell completions for fish.

To load completions in your current fish session:
    %[1]s completion fish | source
To load completions for every fish session, execute once:
    %[1]s completion fish ~/.config/fish/completions/%[1]s.fish

You will need to start a new shell for this setup to take effect.
`, programName),
		RunE: func(cmd *cobra.Command, args []string) error {
			root := cmd.Root()
			if len(args) == 0 {
				return root.GenFishCompletion(root.OutOrStdout(), true)
			}
			if args[0] == "" || args[0] == "-" {
				return root.GenFishCompletion(root.OutOrStdout(), true)
			}
			return root.GenFishCompletionFile(args[0], true)
		},
	}
	return cmd
}

// PowershellCompletionCmd returns a cobra command for
// generating PowershellCompletionCmd completion.
func NewPwshCompletionCmd(programName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "powershell",
		Aliases:      []string{"pwsh", "ps"},
		Args:         cobra.MaximumNArgs(1),
		SilenceUsage: true,
		Short:        "Generate autocompletion for powershell",
		Long: fmt.Sprintf(`Generate autocompletion for powershell.

To load completions in your current shell session:
PS C:\> %[1]s completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of
the above command to your powershell profile.
`, programName),
		RunE: func(cmd *cobra.Command, args []string) error {
			root := cmd.Root()
			if len(args) == 0 {
				return root.GenPowerShellCompletionWithDesc(root.OutOrStdout())
			}
			if args[0] == "" || args[0] == "-" {
				return root.GenPowerShellCompletionWithDesc(root.OutOrStdout())
			}
			return root.GenPowerShellCompletionFileWithDesc(args[0])
		},
	}
	return cmd
}
