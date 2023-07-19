package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version command options.
type versionCmdOpts struct {
	format   string
	template string
}

// RunE for version command.
func (o *versionCmdOpts) RunE(cmd *cobra.Command, args []string) error {
	if o.template != "" {
		o.format = "template"
	}
	switch o.format {
	case "text", "pretty", "simple", "":
		return asText(cmd.OutOrStdout())
	case "short":
		return asShortText(cmd.OutOrStdout())
	case "json":
		return asJSON(cmd.OutOrStdout())
	case "template":
		return renderVersionTemplate(o.template, cmd.OutOrStdout())
	default:
		return fmt.Errorf("not a valid format - %s", o.format)
	}
}

// NewVersionCmd returns a version command with options to
// output in json and templated string format.
func NewVersionCmd() *cobra.Command {
	o := &versionCmdOpts{}
	cmd := &cobra.Command{
		Use:   "version",
		Args:  cobra.NoArgs,
		Short: "Show version and build information",
		RunE:  o.RunE,
	}
	cmd.Long = `Show version and build info.

When using the --template flag the following properties are
available to use in the template:

- .Version contains the semantic version
- .GitCommit is the git commit SHA1 hash.
- .BuildDate is build date
- .GoVersion contains the version of Go that binary was compiled with
- .Os is operating system (GOOS)
- .Arch is system architecture (GOARCH)
- .Compiler is the Go compiler used to build the binary.
`

	cmd.Flags().StringVarP(&o.template, "template", "t", "", "output as template")
	cmd.Flags().StringVar(&o.format, "format", "text", "output as format")
	cmd.MarkFlagsMutuallyExclusive("template", "format")
	//nolint: errcheck // ignore
	cmd.RegisterFlagCompletionFunc(
		"format",
		func(cmd *cobra.Command, args []string, _ string) ([]string, cobra.ShellCompDirective) {
			return []string{
					"text\tTextual format",
					"json\tOutput as JSON",
				},
				cobra.ShellCompDirectiveNoFileComp
		},
	)
	return cmd
}
