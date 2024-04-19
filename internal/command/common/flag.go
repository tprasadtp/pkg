// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: MIT

package common

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// AddTemplateFlags adds
//   - String flag --template/-t
//   - String flag --template-file/-T
//
// and marks both flags as MutuallyExclusive.
func AddTemplateFlags(cmd *cobra.Command, template, templateFile *string) {
	cmd.Flags().StringVarP(template, "template", "t", "", "go template to render")
	cmd.Flags().StringVarP(templateFile, "template-file", "T", "", "go template file to render")
	cmd.MarkFlagsMutuallyExclusive("template", "template-file")
}

// AddOutputFlagsWithAppend adds
//   - String flag --output/-o for output file.
//   - Bool flag --append/-a bool flag to indicate that output file should be appended, not overwritten.
//   - Bool flag --append-at-newline bool flag to indicate that output should be written on a new line.
func AddOutputFlagsWithAppend(cmd *cobra.Command, output *string, append, lf *bool) {
	cmd.Flags().StringVarP(output, "output", "o", "", "output file path")
	cmd.Flags().BoolVarP(append, "append", "a", false, "append to output file")
	cmd.Flags().BoolVar(lf, "append-at-newline", true, "write output on a newline")
	cmd.Flags().SetNormalizeFunc(func(_ *pflag.FlagSet, name string) pflag.NormalizedName {
		if name == "append-at-lf" {
			name = "append-at-newline"
		}
		return pflag.NormalizedName(name)
	})
}
