// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: MIT

package semver

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"text/template"

	"github.com/Masterminds/semver/v3"
	"github.com/spf13/cobra"
	"github.com/tprasadtp/knit/internal/command/common"
	"github.com/tprasadtp/knit/internal/shared"
	"github.com/tprasadtp/knit/internal/sprout"
)

type options struct {
	template     string    // template to render
	version      string    // version to parse
	templateFile string    // read template to render from file
	versionFile  string    // read version from file
	stdout       io.Writer // output io.writer
	append       bool      // append to file
	appendLF     bool      // append newline before writing to file
	output       string    // output file
	prefix       string    // prefix to remove
	noTrimSpaces bool      // do not trim spaces in version file or version
}

func (o *options) Run(_ context.Context, _ []string) error {
	var w io.Writer
	var buf bytes.Buffer
	var tpl *template.Template
	var input string
	var templateStr string

	if o.versionFile != "" && o.version != "" {
		return fmt.Errorf(
			"cmd(semver): both version(%q) and version-file(%q) are defined",
			o.version, o.versionFile,
		)
	}

	if o.templateFile != "" && o.template != "" {
		return fmt.Errorf(
			"cmd(semver): both template and template-file(%q) are defined",
			o.templateFile,
		)
	}

	// Try to read version from version file.
	if o.versionFile != "" {
		contents, err := shared.ReadSmallFile(o.versionFile, 1e3)
		if err != nil {
			return fmt.Errorf("cmd(semver): failed to read version file: %w", err)
		}
		input = string(contents)
	} else {
		input = o.version
	}

	// Try to read template from template file.
	if o.templateFile != "" {
		contents, err := shared.ReadSmallFile(o.templateFile, 1e3)
		if err != nil {
			return fmt.Errorf("cmd(semver): failed to read template file: %w", err)
		}
		templateStr = string(contents)
	} else {
		templateStr = o.template
	}

	// Trim whitespace unless disabled explicitly.
	if !o.noTrimSpaces {
		input = strings.TrimSpace(input)
	}

	// Validate version and strip prefix if any.
	v, err := semver.NewVersion(strings.TrimPrefix(input, o.prefix))
	if err != nil {
		return fmt.Errorf("cmd(semver): invalid version(%q): %w", input, err)
	}

	// Render the template into a buffer first.
	// This avoids modifying the output file if template is invalid or errors.
	if templateStr != "" {
		tpl, err = template.New("semver").Funcs(sprout.FuncMap()).Parse(templateStr)
		if err != nil {
			return fmt.Errorf("cmd(semver): invalid template: %w", err)
		}
		err = tpl.Execute(&buf, v)
		if err != nil {
			return fmt.Errorf("cmd(semver): failed to render template: %w", err)
		}
	} else {
		buf.WriteString(v.String())
	}

	if o.output != "" && o.output != "-" {
		var flag int
		if o.append {
			flag = os.O_CREATE | os.O_WRONLY | os.O_APPEND
		} else {
			flag = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
		}

		// Create a file if required. if append is not specified, existing file
		// will be truncated, if any.
		file, err := os.OpenFile(o.output, flag, 0o644)
		if err != nil {
			return fmt.Errorf("cmd(semver): failed to open file(%s): %w", o.output, err)
		}
		defer file.Close()

		// Checks if file already contains a newline.
		// If appending is required simply use the template.
		if o.append && o.appendLF {
			stat, err := file.Stat()
			if err != nil {
				return fmt.Errorf("cmd(semver): failed to stat file(%s): %w", o.output, err)
			}
			if size := stat.Size(); size > 0 {
				x := make([]byte, 1)
				if size == 1 {
					_, err = file.ReadAt(x, 0)
				} else {
					_, err = file.ReadAt(x, stat.Size()-2)
				}
				if err != nil {
					return fmt.Errorf("cmd(semver): failed to read existing file(%s): %w", o.output, err)
				}
				if x[0] != '\n' {
					bufCopy := slices.Clone(buf.Bytes())
					buf.Reset()
					buf.WriteByte('\n')
					buf.Write(bufCopy)
				}
			}
		}
		w = file
	} else {
		w = o.stdout
	}

	// Write to the writer.
	_, err = w.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("cmd(semver): failed to write output: %w", err)
	}
	return nil
}

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "semver",
		Short: "Semantic version parser",
		Long:  "Parses given semantic version and output in different formats.",
		Args:  cobra.NoArgs,
	}
	opts := &options{
		stdout: cmd.OutOrStdout(),
	}

	cmd.Flags().BoolVar(&opts.append, "prefix", false, "trim the prefix if present")

	// Version flags
	cmd.Flags().StringVarP(&opts.version, "version", "v", "", "version to parse")
	cmd.Flags().StringVar(&opts.versionFile, "version-file", "", "file to read semver from")
	cmd.MarkFlagsOneRequired("version", "version-file")
	_ = cmd.MarkFlagFilename("version-file")

	// Template flags
	common.AddTemplateFlags(cmd, &opts.template, &opts.templateFile)

	// Append flags. This also aliases
	common.AddOutputFlagsWithAppend(cmd, &opts.output, &opts.append, &opts.appendLF)

	// If --append-no-newline is enabled ensure that --append and --output are is marked as required.
	// If --append is enabled, ensure --output is marked as required.
	cmd.PreRunE = func(cmd *cobra.Command, _ []string) error {
		if v, _ := cmd.Flags().GetBool("append-no-newline"); v {
			_ = cmd.MarkFlagRequired("append")
			_ = cmd.MarkFlagRequired("output")
		}

		if v, _ := cmd.Flags().GetBool("append"); v {
			_ = cmd.MarkFlagRequired("output")
		}
		return nil
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return opts.Run(cmd.Context(), args)
	}

	return cmd
}
