package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewDocsCmd returns hidden commands to generate manpages and markdown docs.
func NewDocsCmd() *cobra.Command {
	opts := docsCmdOpts{}
	cmd := &cobra.Command{
		Use:     "docs markdown|man [--layout=NAME] --output DIR",
		Aliases: []string{"generate-docs", "gen-docs"},
		ValidArgs: []string{
			"manpages\tGenerate manpages",
			"markdown\tGenerate markdown docs",
		},
		Args:   cobra.ExactArgs(1),
		Hidden: true,
		Short:  "Generate documentation files of specified type",
		Long: `Generate documentation files of specified type

With markdown, if --layout is specified, generated markdown files
include hugo yaml headers with specified layout. With manpages,
this flag is ignored.

With manpages, if --gzip is used, generated manpages are gzip compressed.
With markdown, this flag is ignored.
`,
		RunE: opts.RunE,
	}
	cmd.Flags().StringVarP(
		&opts.OutputDir,
		"output", "o", "",
		"Output directory. If it does not exist, it will be created.")
	//nolint: errcheck // ignore
	cmd.MarkFlagRequired("output")
	cmd.Flags().StringVar(
		&opts.Layout,
		"layout",
		"",
		"Markdown layout name",
	)
	cmd.Flags().BoolVar(
		&opts.CompressManpages,
		"gzip",
		false,
		"Compress generated manpages",
	)
	return cmd
}

type docsCmdOpts struct {
	Layout           string
	CompressManpages bool
	OutputDir        string
}

func (opts *docsCmdOpts) RunE(cmd *cobra.Command, args []string) error {
	root := cmd.Root()
	switch args[0] {
	case "markdown", "md":
		return GenMarkdownTree(root, opts.OutputDir, opts.Layout)
	case "manpages", "man":
		return GenManTree(root, opts.OutputDir, opts.CompressManpages)
	default:
		return fmt.Errorf("unknown type of documentation - %s", args[0])
	}
}
