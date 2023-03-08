package factory

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"text/template"

	"github.com/tprasadtp/pkg/version"
)

// Renders given go template with version info to out.
func renderVersionTemplate(tpl string, out io.Writer) error {
	if tpl == "" {
		return errors.New("template is empty")
	}
	info := version.GetInfo()
	vTemplate, err := template.New("version").Parse(tpl)
	if err != nil {
		return fmt.Errorf("template is invalid: %w", err)
	}
	return vTemplate.Execute(out, info)
}

// asText writes text output to specified writer.
func asText(out io.Writer) error {
	const template = `• Version      : {{.Version}}
• GitCommit    : {{.GitCommit}}
• BuildDate    : {{.BuildDate}}
• GoVersion    : {{.GoVersion}}
• Os           : {{.Os}}
• Arch         : {{.Arch}}
• Compiler     : {{.Compiler}}
`
	return renderVersionTemplate(template, out)
}

// asJSON returns formatted JSON string to be printed.
func asJSON(out io.Writer) error {
	b, err := json.MarshalIndent(version.GetInfo(), "", "  ")
	if err == nil {
		_, err = out.Write(b)
	}
	return err
}

// as Short returns just the version info.
func asShortText(out io.Writer) error {
	_, err := out.Write([]byte(version.GetInfo().Version + "\n"))
	return err
}
