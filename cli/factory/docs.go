//go:build dev

package factory

import (
	"github.com/tprasadtp/pkg/cli"
)

// NewDocsCmd returns commands to generate manpages and markdown docs.
func NewDocsCmd(programName string, hidden bool) *cli.Command {
	cmd := &cli.Command{
		Use:     "docs",
		Aliases: []string{"gen"},
		Args:    cli.NoArgs,
		Hidden:  hidden,
		Short:   "Generate docs, manpages etc.",
		Long: `Generates documentation, manpages and API docs (if available)
This command is typically not available unless built with docs tag,
and should only be used for development purposes and provides
no stability guarantees.
`,
	}
	cmd.AddCommand(NewManpagesCmd(programName))
	// cmd.AddCommand(NewMarkdownCmd(programName))
	return cmd
}
