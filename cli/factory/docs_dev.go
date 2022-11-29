//go:build dev

package factory

import "github.com/tprasadtp/pkg/cli"

// getDocsCmd returns nil on builds not tagged docs
// otherwise it returns *cli.Command with subcommands
// generating man pages and markdown documentation.
func getDocsCmd(programName string) *cli.Command {
	return NewDocsCmd(programName, true)
}
