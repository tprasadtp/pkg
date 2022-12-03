// Copyright (c) 2013 Steve Francia <spf@spf13.com>.
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
	"time"
	"unicode"
)

var templateFuncs = template.FuncMap{
	"trim":                    strings.TrimSpace,
	"trimRightSpace":          trimRightSpace,
	"trimTrailingWhitespaces": trimRightSpace,
	"rpad":                    rpad,
}

var initializers []func()

// EnablePrefixMatching allows to set automatic prefix matching.
// Automatic prefix matching can be a dangerous thing
// to automatically enable in CLI tools.
// Set this to true to enable it.
var EnablePrefixMatching = false

// EnableCommandSorting controls sorting of the slice of commands,
// which is turned on by default.
// To disable sorting, set it to false.
var EnableCommandSorting = true

// MousetrapHelpText enables an information splash screen on Windows
// if the CLI is started from explorer.exe.
// To disable the mousetrap, just set this variable to blank string ("").
// Works only on Microsoft Windows.
var MousetrapHelpText = `This is a command line tool.

You need to open Windows Terminal and run it from there.
`

// MousetrapDisplayDuration controls how long the
// MousetrapHelpText message is displayed on Windows
// if the CLI is started from explorer.exe.
// Set to 0 to wait for the return key to be pressed.
// To disable the mousetrap, just set MousetrapHelpText to blank string ("").
// Works only on Microsoft Windows.
var MousetrapDisplayDuration = 5 * time.Second

// AddTemplateFunc adds a template function that's available to Usage and Help
// template generation.
func AddTemplateFunc(name string, f any) {
	templateFuncs[name] = f
}

// AddTemplateFuncs adds multiple template functions that are available to Usage and
// Help template generation.
func AddTemplateFuncs(funcs template.FuncMap) {
	for k, v := range funcs {
		templateFuncs[k] = v
	}
}

// OnInitialize sets the passed functions to be run when each command's
// Execute method is called.
func OnInitialize(y ...func()) {
	initializers = append(initializers, y...)
}

func trimRightSpace(s string) string {
	return strings.TrimRightFunc(s, unicode.IsSpace)
}

// rpad adds padding to the right of a string.
func rpad(s string, padding int) string {
	template := fmt.Sprintf("%%-%ds", padding)
	return fmt.Sprintf(template, s)
}

// tmpl executes the given template text on data, writing the result to w.
func tmpl(w io.Writer, text string, data any) error {
	t := template.New("top")
	t.Funcs(templateFuncs)
	template.Must(t.Parse(text))
	return t.Execute(w, data)
}

// ld compares two strings and returns the levenshtein distance between them.
func ld(s, t string, ignoreCase bool) int {
	if ignoreCase {
		s = strings.ToLower(s)
		t = strings.ToLower(t)
	}
	d := make([][]int, len(s)+1)
	for i := range d {
		d[i] = make([]int, len(t)+1)
	}
	for i := range d {
		d[i][0] = i
	}
	for j := range d[0] {
		d[0][j] = j
	}
	for j := 1; j <= len(t); j++ {
		for i := 1; i <= len(s); i++ {
			if s[i-1] == t[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				min := d[i-1][j]
				if d[i][j-1] < min {
					min = d[i][j-1]
				}
				if d[i-1][j-1] < min {
					min = d[i-1][j-1]
				}
				d[i][j] = min + 1
			}
		}

	}
	return d[len(s)][len(t)]
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// CheckErr prints the msg with the prefix 'Error:' and exits
// with error code 1. If the msg is nil, it does nothing.
func CheckErr(msg any) {
	if msg != nil {
		fmt.Fprintln(os.Stderr, "Error:", msg)
		os.Exit(1)
	}
}

// WriteStringAndCheck writes a string into a buffer, and checks
// if the error is not nil.
func WriteStringAndCheck(b io.StringWriter, s string) {
	_, err := b.WriteString(s)
	CheckErr(err)
}
