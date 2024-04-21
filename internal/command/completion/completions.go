package completion

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// Options for shell completion.
type options struct {
	output string
	stdout io.Writer
}

// Runner for complete command.
func (opts *options) Run(ctx context.Context, args []string, cmd *cobra.Command) error {
	root := cmd.Root()
	switch args[0] {
	case "bash":
		switch opts.output {
		case "", "-":
			return root.GenBashCompletionV2(root.OutOrStdout(), true)
		default:
			return root.GenBashCompletionFileV2(opts.output, true)
		}
	case "zsh":
		switch opts.output {
		case "", "-":
			return root.GenZshCompletion(root.OutOrStdout())
		default:
			return root.GenZshCompletionFile(opts.output)
		}
	case "fish":
		switch opts.output {
		case "", "-":
			return root.GenFishCompletion(root.OutOrStdout(), true)
		default:
			return root.GenFishCompletionFile(opts.output, true)
		}
	case "powershell", "pwsh":
		switch opts.output {
		case "", "-":
			return root.GenPowerShellCompletionWithDesc(root.OutOrStdout())
		default:
			return root.GenPowerShellCompletionFileWithDesc(opts.output)
		}
	default:
		return fmt.Errorf("cannot generate completion for unknown shell: %s", args[0])
	}
}

func newCompletionCmd(hidden bool) *cobra.Command {
	var o = &options{}
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
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.Run(cmd.Context(), args, cmd.Root())
		},
		Hidden: hidden,
		Long: `Generate shell autocompletion for specified shell.
It may require additional tasks within your shell to setup
loading shell completions by default.`,
	}
	cmd.Flags().StringVarP(&o.output,
		"output", "o",
		"",
		"Output file. If not specified or if equals '-', writes to stdout",
	)
	return cmd
}

func NewCompletionCmd() *cobra.Command {
	return newCompletionCmd(false)
}

func NewHiddenCompletionCmd() *cobra.Command {
	return newCompletionCmd(true)
}
