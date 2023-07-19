package cli

import (
	_ "embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

//go:embed markdown.tpl
var defaultMarkdownTpl string

// GenMarkdownCustom creates markdown output. Use [GenHugo] for hugo
// output as it handles filenames and links better.
func genMarkdown(cmd *cobra.Command, w io.Writer, layout string) error {
	cmd.InitDefaultHelpCmd()
	cmd.InitDefaultHelpFlag()

	getLayout := func() string {
		return layout
	}
	funcMap["getLayout"] = getLayout

	tpl := template.New("markdown.tpl").Funcs(funcMap)
	tpl, err := tpl.Parse(defaultMarkdownTpl)
	if err != nil {
		return fmt.Errorf("failed to parse embedded markdown template: %w", err)
	}

	err = tpl.Execute(w, cmd)
	if err != nil {
		return fmt.Errorf("failed to render embedded markdown template: %w", err)
	}
	return nil
}

// GenMarkdownTree will generate a markdown page for this command and all
// descendants in the directory given.
func GenMarkdownTree(cmd *cobra.Command, dir string, layout ...string) error {
	if cmd == nil {
		return fmt.Errorf("cannot generate markdown docs for nil command")
	}

	for _, c := range cmd.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}
		if err := GenMarkdownTree(c, dir, layout...); err != nil {
			return err
		}
	}

	basename := strings.ReplaceAll(cmd.CommandPath(), " ", "-") + ".md"
	filename := filepath.Join(dir, basename)
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if len(layout) == 0 {
		layout = make([]string, 1)
		layout[0] = ""
	}

	if err = genMarkdown(cmd, f, layout[0]); err != nil {
		return err
	}
	return nil
}
