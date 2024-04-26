package version

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tprasadtp/knit/internal/command/common"
	"github.com/tprasadtp/knit/internal/version"
)

// Version command options.
type options struct {
	short        bool
	json         bool
	fmt          string
	template     string
	templateFile string
	output       string
	append       bool
	appendLF     bool
}

func (o *options) Run(_ context.Context, _ []string) error {
	return nil
}

// RunE for version command.
func (o *options) RunE(cmd *cobra.Command, _ []string) error {
	if o.template != "" {
		o.fmt = "template"
	}
	switch o.fmt {
	case "text", "pretty", "simple", "":
		return asText(cmd.OutOrStdout())
	case "short":
		return asShortText(cmd.OutOrStdout())
	case "json":
		return asJSON(cmd.OutOrStdout())
	case "template":
		return renderVersionTemplate(o.template, cmd.OutOrStdout())
	default:
		return fmt.Errorf("not a valid format - %s", o.fmt)
	}
}

func Version() string {
	return version.GetInfo().Version
}

// NewVersionCmd returns a version command with options to
// output in json and templated string format.
func NewVersionCmd() *cobra.Command {
	o := &options{}
	cmd := &cobra.Command{
		Use:   "version",
		Args:  cobra.NoArgs,
		Short: "Show version and build information",
		RunE:  o.RunE,
	}
	cmd.Long = `Show version and build information.

When using the --template flag the following properties are
available to use in the template:

- .Version contains the semantic version.
- .GitCommit is the git commit SHA1 hash.
- .GitTreeState is the git tree state.
- .BuildDate is build date.
- .GoVersion contains the version of Go that binary was compiled with.
- .Os is operating system (GOOS).
- .Arch is system architecture (GOARCH).
- .Compiler is the Go compiler used to build the binary.
`

	cmd.Flags().StringVar(&o.fmt, "format", "text", "output format")
	cmd.Flags().BoolVar(&o.json, "json", false, "output as json")
	cmd.Flags().BoolVar(&o.short, "short", false, "only show version and skip build info")
	common.AddTemplateFlags(cmd, &o.template, &o.templateFile)
	common.AddOutputFlagsWithAppend(cmd, &o.output, &o.append, &o.appendLF)
	//nolint:errcheck // ignore
	cmd.RegisterFlagCompletionFunc(
		"format",
		func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
			return []string{
					"text\tTextual format",
					"json\tOutput as JSON",
				},
				cobra.ShellCompDirectiveNoFileComp
		},
	)
	return cmd
}
