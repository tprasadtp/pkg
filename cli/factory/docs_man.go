//go:build dev

package factory

import (
	"github.com/tprasadtp/pkg/cli"
)

// NewManpagesCmd returns a cobra command for
// generating manpages.
func NewManpagesCmd(programName string) *cli.Command {
	// header := cobradoc.GenManHeader{
	// 	Title:   strings.ToUpper(programName),
	// 	Section: "1",
	// }
	cmd := &cli.Command{
		Use:          "manpages DIRECTORY",
		Aliases:      []string{"man"},
		Args:         cli.ExactArgs(1),
		SilenceUsage: true,
		Short:        "Generate manpages",
		// RunE: func(cmd *cli.Command, args []string) error {
		// 	if len(args) == 1 {
		// 		output := args[0]
		// 		if len(strings.TrimSpace(output)) < 1 {
		// 			return fmt.Errorf("output directory is empty")
		// 		}
		// 		return cobradoc.GenManTree(cmd.Root(), &header, output)
		// 	}
		// 	return fmt.Errorf("no output directory specified")
		// },
	}
	return cmd
}
