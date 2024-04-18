// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: GPLv3-only

package shared

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"text/template"

	"github.com/tprasadtp/knit/internal/sprout"
)

func RenderTemplateToFile(path, tpl string, append, lf bool, mode fs.FileMode, data any) error {
	var flag int

	// Truncate the file if append is not specified.
	if append {
		flag = os.O_CREATE | os.O_RDWR | os.O_APPEND
	} else {
		flag = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	}

	// If mode is not specified, default to read/write owner.
	if mode == 0 {
		mode = fs.FileMode(0o600)
	}

	file, err := os.OpenFile(path, flag, mode)
	if err != nil {
		return fmt.Errorf("cmd(semver): failed to open file(%q): %w", path, err)
	}
	defer file.Close()
	return RenderTemplate(file, tpl, data)
}

func RenderTemplate(w io.Writer, tpl string, data any) error {
	t, err := template.New("semver").Funcs(sprout.FuncMap()).Parse(tpl)
	if err != nil {
		return fmt.Errorf("shared(template): invalid template: %w", err)
	}
	err = t.Execute(w, data)
	if err != nil {
		return fmt.Errorf("shared(template): failed to render template: %w", err)
	}
	return nil
}
