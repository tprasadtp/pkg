// package cli provides a wrapper around [github.com/spf13/cobra/cli]
// which handles common tasks like version command, help text
// along with helpers to generate documentation, and shell completions.
package cli

import (
	"github.com/spf13/cobra"
	"github.com/tprasadtp/pkg/cli/internal/factory"
	"github.com/tprasadtp/pkg/version"
)

// New returns a new CLI with all bells and whistles attached.
// If built with dev tag, this root command includes hidden commands
// 'docs' and 'completion' to generate man pages, documentation and
// shell completion scripts. Its best to let go generate
// generate these. You can use snippet below in you 'main.go' file.
// Please ensure to have directories already created, sub commands
// will not do it fo you.
//
//	// go:generate go run -tags dev main.go completion bash completion/<name>.bash
//	// go:generate go run -tags dev main.go completion fish completion/<name>.fish
//	// go:generate go run -tags dev main.go completion zsh completion/<name>.zsh
//	// go:generate go run -tags dev main.go completion powershell completion/<name>.ps1
//	// go:generate go run -tags dev main.go docs man pages
//	// go:generate go run -tags dev main.go docs markdown docs/content/manual
func New(name, shortDesc, longDesc string, commands ...*cobra.Command) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     name,
		Version: version.GetShort(),
		Short:   shortDesc,
		Long:    longDesc,
		Args:    cobra.NoArgs,
	}
	if len(commands) > 0 {
		rootCmd.AddCommand(commands...)
	}
	rootCmd.AddCommand(factory.NewVersionCmd(name))
	return rootCmd
}
