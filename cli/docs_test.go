package cli

import (
	"bytes"
	"testing"

	"github.com/tprasadtp/pkg/cli/internal/testcli"
)

func Test_Docs_Markdown(t *testing.T) {
	type testCase struct {
		Name string
		Args []string
		Err  bool
	}
	tt := []testCase{
		{
			Name: "markdown",
			Args: []string{"docs", "markdown", "--output"},
		},
		{
			Name: "markdown-alias",
			Args: []string{"gen-docs", "md", "--output"},
		},
		{
			Name: "markdown-with-layout",
			Args: []string{"docs", "markdown", "--layout", "default", "--output"},
		},
		{
			Name: "unknown-type",
			Args: []string{"docs", "shiny-javascript", "--output"},
			Err:  true,
		},
	}
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			root := testcli.GetTestCLI()
			root.AddCommand(NewDocsCmd())
			root.SetErr(stderr)
			root.SetOut(stdout)
			output := t.TempDir()
			root.SetArgs(append(tc.Args, output))
			err := root.Execute()
			if tc.Err {
				if err == nil {
					t.Errorf("expected an error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %s", err)
				}
			}
		})
	}
}

func Test_Docs_Manpages(t *testing.T) {
	type testCase struct {
		Name string
		Args []string
		Err  bool
	}
	tt := []testCase{
		{
			Name: "manpages",
			Args: []string{"docs", "manpages", "--output"},
		},
		{
			Name: "manpages-alias",
			Args: []string{"gen-docs", "man", "--output"},
		},
		{
			Name: "manpages-with-compress",
			Args: []string{"docs", "manpages", "--gzip", "--output"},
		},
	}
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			root := testcli.GetTestCLI()
			root.AddCommand(NewDocsCmd())
			root.SetErr(stderr)
			root.SetOut(stdout)
			output := t.TempDir()
			root.SetArgs(append(tc.Args, output))
			err := root.Execute()
			if tc.Err {
				if err == nil {
					t.Errorf("expected an error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %s", err)
				}
			}
		})
	}
}
