// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: MIT

package shared

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"text/template"
)

// RenderTemplateToFile renders the template with given data and functions to path specified.
// If append is true, then output will be appended to the file. Otherwise it will be overwritten.
// If directory does not exist, it will be created. In append mode, output is always written on a
// new line unless lf is set to false. Directory created will have permission based on file permissions
// specified. To avoid modifying files on template errors, This function renders the template into
// a buffer first and then modifies the file.
//
//nolint:predeclared // ignore
func RenderTemplateToFile(path, tpl string, funcs template.FuncMap, append, lf bool, mode fs.FileMode, data any) error {
	var buf bytes.Buffer
	var err error

	if path == "" {
		return fmt.Errorf("shared(template): path is empty")
	}

	err = RenderTemplate(&buf, tpl, funcs, data)
	if err != nil {
		return err
	}

	return WriteToFile(path, buf.Bytes(), append, lf, mode)
}

func RenderTemplate(w io.Writer, tpl string, funcs template.FuncMap, data any) error {
	t, err := template.New("semver").Funcs(funcs).Parse(tpl)
	if err != nil {
		return fmt.Errorf("shared(template): invalid template: %w", err)
	}
	err = t.Execute(w, data)
	if err != nil {
		return fmt.Errorf("shared(template): failed to render template: %w", err)
	}
	return nil
}
