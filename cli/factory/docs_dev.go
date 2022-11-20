//go:build dev

package factory

import (
	"github.com/tprasadtp/pkg/cli/cobra"
)

// getDocsCmd returns nil on builds not tagged docs
// otherwise it returns *cobra.Command with subcommands
// generating man pages and markdown documentation.
func getDocsCmd(programName string) *cobra.Command {
	return NewDocsCmd(programName, true)
}
