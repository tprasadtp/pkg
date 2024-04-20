// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: MIT

package version

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/tprasadtp/knit/internal/version"
)

func TestVersionCmd_Template(t *testing.T) {
	var stdout = new(bytes.Buffer)
	var stderr = new(bytes.Buffer)
	type testCase struct {
		Name     string
		Args     []string
		Verifier func(t *testing.T, stdout, stderr *bytes.Buffer, err error)
	}
	tt := []testCase{
		{
			Name: "version-short",
			Args: []string{"version", "--format=short"},
			Verifier: func(t *testing.T, stdout, stderr *bytes.Buffer, err error) {
				if stdout.String() != version.GetInfo().Version+"\n" {
					t.Errorf("stdout: must return version ending with a newline, got=%s", stdout.String())
				}
				if stderr.String() != "" {
					t.Errorf("stdout: expected empty, got=%s", stderr.String())
				}

				if err != nil {
					t.Errorf("must not return an error, got=%s", err)
				}
			},
		},
		{
			Name: "version-template-1",
			Args: []string{"version", "--template={{.Version}}"},
			Verifier: func(t *testing.T, stdout, stderr *bytes.Buffer, err error) {
				if stdout.String() != version.GetInfo().Version {
					t.Errorf("stdout: must return version, got=%s", stdout.String())
				}
				if stderr.String() != "" {
					t.Errorf("stdout: expected empty, got=%s", stderr.String())
				}

				if err != nil {
					t.Errorf("must not return an error, got=%s", err)
				}
			},
		},
		{
			Name: "version-format-text",
			Args: []string{"version", "--format=text"},
			Verifier: func(t *testing.T, stdout, stderr *bytes.Buffer, err error) {
				output := stdout.String()
				contains := "• GitCommit"
				if !strings.Contains(output, contains) {
					t.Errorf("stdout: must contain %s, got=%s", contains, output)
				}
				if stderr.String() != "" {
					t.Errorf("stdout: expected empty, got=%s", stderr.String())
				}

				if err != nil {
					t.Errorf("must not return an error, got=%s", err)
				}
			},
		},
		{
			Name: "version-format-json",
			Args: []string{"version", "--format=json"},
			Verifier: func(t *testing.T, stdout, stderr *bytes.Buffer, err error) {
				if jerr := json.Unmarshal(stdout.Bytes(), &version.Info{}); jerr != nil {
					t.Errorf("stdout: must return json output, got=%s", jerr)
				}
				if stderr.String() != "" {
					t.Errorf("stdout: expected empty, got=%s", stderr.String())
				}

				if err != nil {
					t.Errorf("must not return an error, got=%s", err)
				}
			},
		},
		{
			Name: "version-format-invalid",
			Args: []string{"version", "--format=latin"},
			Verifier: func(t *testing.T, stdout, stderr *bytes.Buffer, err error) {
				if err == nil {
					t.Errorf("must return an error on invalid format")
				}
			},
		},
		{
			Name: "version-template-invalid",
			Args: []string{"version", "--template={{.NoSuchField}}"},
			Verifier: func(t *testing.T, stdout, stderr *bytes.Buffer, err error) {
				if err == nil {
					t.Errorf("must return an error on invalid template")
				}
			},
		},
		{
			Name: "version-conflicting-flags",
			Args: []string{"version", "--template={{.Version}}", "--format=json"},
			Verifier: func(t *testing.T, stdout, stderr *bytes.Buffer, err error) {
				if err == nil {
					t.Errorf("must return an error on conflicting flags")
				}
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			stdout.Reset()
			stderr.Reset()
			root := &cobra.Command{
				Use:   "blackhole-entropy",
				Short: "Black Hole Entropy CLI",
				Long:  "CLI to seed system's PRNG with entropy from M87 Black Hole",
			}
			root.SetOut(stdout)
			root.SetErr(stderr)
			root.AddCommand(NewVersionCmd())
			root.SetArgs(tc.Args)
			err := root.Execute()

			if tc.Verifier == nil {
				t.Fatalf("no verifier specified!")
			} else {
				tc.Verifier(t, stdout, stderr, err)
			}
		})
	}
}
