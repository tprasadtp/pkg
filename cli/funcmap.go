package cli

import (
	"strings"
	"text/template"
)

// Functions used in template.
var funcMap = template.FuncMap{
	// string functions
	//nolint:staticcheck // not relevant.
	"title": strings.Title,
	"upper": strings.ToUpper,
	// Switch order so that pipelining works.
	"replace": replace,
	// cobra specific optimization.
	"isAutoGenDisabled": isAutoGenDisabled,
	"formatGeneratedAt": formatGeneratedAt,
	"getSeeAlso":        getSeeAlso,
	"getFlags":          getFlags,
}

// Template suitable version of [strings.ReplaceAll].
func replace(oldstr, newstr, src string) string {
	return strings.ReplaceAll(src, oldstr, newstr)
}
