package factory

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Options for shell completion.
type completionCmdOpts struct {
	Output string
}

// Runner for complete command.
func (opts *completionCmdOpts) RunE(cmd *cobra.Command, args []string) error {
	root := cmd.Root()
	switch args[0] {
	case "bash":
		switch opts.Output {
		case "", "-":
			return root.GenBashCompletionV2(root.OutOrStdout(), true)
		default:
			return root.GenBashCompletionFileV2(opts.Output, true)
		}
	case "zsh":
		switch opts.Output {
		case "", "-":
			return root.GenZshCompletion(root.OutOrStdout())
		default:
			return root.GenZshCompletionFile(opts.Output)
		}
	case "fish":
		switch opts.Output {
		case "", "-":
			return root.GenFishCompletion(root.OutOrStdout(), true)
		default:
			return root.GenFishCompletionFile(opts.Output, true)
		}
	case "powershell", "pwsh":
		switch opts.Output {
		case "", "-":
			return root.GenPowerShellCompletionWithDesc(root.OutOrStdout())
		default:
			return root.GenPowerShellCompletionFileWithDesc(opts.Output)
		}
	default:
		return fmt.Errorf("cannot generate completion for unknown shell: %s", args[0])
	}
}

// NewCompletionCmd returns a cobra command named completion
// with all supported shells as subcommands.
func NewCompletionCmd(hidden ...bool) *cobra.Command {
	var o = &completionCmdOpts{}
	if len(hidden) == 0 {
		hidden = make([]bool, 1)
	}
	cmd := &cobra.Command{
		Use:     "completion [--output=FILE] SHELL",
		Short:   "Generate shell autocompletion for specified shell",
		Aliases: []string{"complete", "compgen"},
		ValidArgs: []string{
			"bash\tGenerate shell completions for bash",
			"zsh\tGenerate  shell completions for zsh",
			"fish\tGenerate shell completions for fish",
			"powershell\tGenerate shell completions for powershell",
		},
		Args:   cobra.ExactArgs(1),
		RunE:   o.RunE,
		Hidden: hidden[0],
		Long: `Generate shell autocompletion for specified shell.
It may require additional tasks within your shell to setup
loading shell completions by default.`,
	}
	cmd.Flags().StringVarP(&o.Output,
		"output", "o",
		"",
		"Output file. If not specified or if equals '-', writes to stdout",
	)
	//nolint: errcheck // ignore
	cmd.MarkFlagFilename("output")
	return cmd
}
