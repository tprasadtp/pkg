package factory

import (
	"compress/gzip"
	_ "embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

//go:embed man.tpl
var defaultManpageTpl string

// GenManTree generates a man page for the command and all descendants.
// The pages are written to the output directory. This does not use markdown
// hints or format to equivalent roff format and is only here for ease of use
// and quick references to flags and commands. Long description in particular,
// might be an issue as it is usually designed for output to terminal and markdown.
// code nits, lists and tables cannot be rendered faithfully. Upstream does this
// via generating markdown and transforming into roff format.
// That has side effect of messing with flag output (it is not aligned correctly).
func GenManTree(cmd *cobra.Command, output string, compress bool) error {
	if cmd == nil {
		return fmt.Errorf("cannot generate manpages for nil command")
	}

	for _, c := range cmd.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}
		if err := GenManTree(c, output, compress); err != nil {
			return err
		}
	}

	var fileExtension = ".1"
	if compress {
		fileExtension = ".1.gz"
	}
	// replace spaces with dashes.
	basename := strings.ReplaceAll(cmd.CommandPath(), " ", "-")
	filename := filepath.Join(output, basename+fileExtension)
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create manpage - %s :%w", filename, err)
	}
	defer f.Close()
	return genMan(cmd, f, compress)
}

// genMan will generate a man page for the given command and write it to
// w. The header argument may be nil, however obviously w may not.
func genMan(cmd *cobra.Command, w io.Writer, compress bool) error {
	cmd.InitDefaultHelpCmd()
	cmd.InitDefaultHelpFlag()

	tpl := template.New("man.tpl").Funcs(funcMap)
	tpl, err := tpl.Parse(defaultManpageTpl)
	if err != nil {
		return fmt.Errorf("failed to parse embedded manpage template: %w", err)
	}

	if compress {
		gzWriter := gzip.NewWriter(w)
		err = tpl.Execute(gzWriter, cmd)
		if err != nil {
			return fmt.Errorf("failed to render embedded manpage template: %w", err)
		}
		defer gzWriter.Flush()
		defer gzWriter.Close()
	} else {
		err = tpl.Execute(w, cmd)
		if err != nil {
			return fmt.Errorf("failed to render embedded manpage template: %w", err)
		}
	}
	return nil
}
