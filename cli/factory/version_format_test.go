package factory

import (
	"bytes"
	"encoding/json"
	"runtime"
	"testing"

	"github.com/tprasadtp/pkg/version"
)

func TestTemplateValid(t *testing.T) {
	buf := &bytes.Buffer{}
	tests := []struct {
		name     string
		template string
		expect   string
	}{
		{
			name:     "Version",
			template: "{{.Version}}",
			expect:   "v0.0.0+undefined",
		},
		{
			name:     "GitCommit",
			template: "{{.GitCommit}}",
			expect:   "",
		},
		{
			name:     "BuildDate",
			template: "{{.BuildDate}}",
			expect:   "1970-01-01T00:00+00:00",
		},
		{
			name:     "Compiler",
			template: "{{.Compiler}}",
			expect:   runtime.Compiler,
		},
		{
			name:     "GoVersion",
			template: "{{.GoVersion}}",
			expect:   runtime.Version(),
		},
		{
			name:     "Os",
			template: "{{.Os}}",
			expect:   runtime.GOOS,
		},
		{
			name:     "Arch",
			template: "{{.Arch}}",
			expect:   runtime.GOARCH,
		},
		{
			name:     "NO_VARIABLES",
			template: "NO_VARIABLES",
			expect:   "NO_VARIABLES",
		},
	}
	for _, tc := range tests {
		buf.Reset()
		t.Run(tc.name, func(t *testing.T) {
			err := renderVersionTemplate(tc.template, buf)
			if err != nil {
				t.Errorf("expected no error: %q", err)
			}
			if buf.String() != tc.expect {
				t.Errorf("output mismatch\nexpected(%q)\n(got)%q", tc.expect, buf.String())
			}
		})
	}
}

func TestTemplateInvalid(t *testing.T) {
	buf := &bytes.Buffer{}
	tests := []struct {
		name     string
		template string
		expect   string
	}{
		{
			name:     "Unclosed Bracket",
			template: "{{.Version}",
		},
		{
			name:     "No doT",
			template: "{{GitCommit}}",
		},
		{
			name:     "Empty",
			template: "",
		},
		{
			name:     "No such field in Info struct",
			template: "{{.NoSuchStructField}}",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buf.Reset()
			err := renderVersionTemplate(tc.template, buf)
			if err == nil {
				t.Errorf("invalid template must error")
			}
			if buf.Len() != 0 {
				t.Errorf("invalid template must not write anything to io.Writer")
			}
		})
	}
}

func Test_asText(t *testing.T) {
	buf := &bytes.Buffer{}
	err := asText(buf)
	if err != nil {
		t.Errorf("asText must not error: %q", err)
	}
	if buf.String() == "" {
		t.Errorf("asText must return non empty output")
	}
}

func Test_ShortText(t *testing.T) {
	buf := &bytes.Buffer{}
	err := asShortText(buf)
	if err != nil {
		t.Errorf("asText must not error: %q", err)
	}
	if buf.String() != version.GetInfo().Version+"\n" {
		t.Errorf("asText must return just the version ending with newline")
	}
}

func Test_asJSON(t *testing.T) {
	buf := &bytes.Buffer{}
	err := asJSON(buf)
	if err != nil {
		t.Errorf("asJSON must not error: %q", err)
	}
	if buf.Len() == 0 {
		t.Errorf("asJSON must return non empty output")
	}
	v := &version.Info{}
	if unmarshalErr := json.Unmarshal(buf.Bytes(), v); unmarshalErr != nil {
		t.Errorf("unmarshalling json ouput must not error, bit got %s", unmarshalErr)
	}
}
