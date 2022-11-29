package cli

import (
	"testing"
	"text/template"
)

func AssertErrNil(t *testing.T, e error) {
	if e != nil {
		t.Errorf("expected=nil, got=%e", e)
	}
}

func TestAddTemplateFunctions(t *testing.T) {
	AddTemplateFunc("t", func() bool { return true })
	AddTemplateFuncs(template.FuncMap{
		"f": func() bool { return false },
		"h": func() string { return "Hello," },
		"w": func() string { return "world." }})

	c := &Command{}
	c.SetUsageTemplate(`{{if t}}{{h}}{{end}}{{if f}}{{h}}{{end}} {{w}}`)

	const expected = "Hello, world."
	if got := c.UsageString(); got != expected {
		t.Errorf("Expected Usage String: %v\nGot: %v", expected, got)
	}
}
