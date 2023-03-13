package factory

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/tprasadtp/pkg/cli/internal/testcli"
)

func Test_GenMarkdownTree_NoLayout(t *testing.T) {
	output := t.TempDir()
	root := testcli.GetTestCLI()
	err := GenMarkdownTree(root, output)
	if err != nil {
		t.Fatalf("failed to generate markdown - %s", err)
	}
	token := []byte(testcli.HiddenToken)
	filesWhichShouldBeGenerated := []string{
		"test-cli.md",
		"test-cli-command1.md",
		"test-cli-command1-subcommand1.md",
		"test-cli-command1-subcommand2.md",
		"test-cli-command2.md",
	}

	// we know what output file is check if
	// any deprecated flags are present.
	// testcli.HiddenToken is in all usage messages.
	for _, f := range filesWhichShouldBeGenerated {
		filename := filepath.Join(output, f)
		file, e := os.Open(filename)
		if e != nil {
			t.Errorf("failed to open md file - %s: %e", filename, e)
			continue
		}
		content, e := io.ReadAll(file)
		if e != nil {
			t.Errorf("failed to read md file - %s : %s", filename, e)
			file.Close()
		}
		file.Close()
		if bytes.Contains(content, token) {
			t.Errorf("hidden flag or command in help output - %s", filename)
		}
	}

	dirList, err := os.ReadDir(output)
	if err != nil {
		t.Errorf("failed to list dir %s: %s", output, err)
	}
	filesWhichShouldNotBeGenerated := []string{
		"test-cli-hidden.md",
		"test-cli-deprecated.md",
	}
	for _, item := range dirList {
		for _, lookupItem := range filesWhichShouldNotBeGenerated {
			if item.Name() == lookupItem {
				t.Errorf("%s should not be present as it belongs to hidden/deprecated cmd", lookupItem)
			}
		}
	}
}

func Test_GenMarkdownTree_CustomLayout(t *testing.T) {
	output := t.TempDir()
	root := testcli.GetTestCLI()
	err := GenMarkdownTree(root, output, "customLayout")
	if err != nil {
		t.Fatalf("failed to generate markdown - %s", err)
	}
	filesWhichShouldBeGenerated := []string{
		"test-cli.md",
		"test-cli-command1.md",
		"test-cli-command1-subcommand1.md",
		"test-cli-command1-subcommand2.md",
		"test-cli-command2.md",
	}

	// we know what output file is check if
	// any deprecated flags are present.
	// testcli.HiddenToken is in all usage messages.
	for _, f := range filesWhichShouldBeGenerated {
		filename := filepath.Join(output, f)
		file, e := os.Open(filename)
		if e != nil {
			t.Errorf("failed to open md file - %s: %e", filename, e)
			continue
		}
		content, e := io.ReadAll(file)
		if e != nil {
			t.Errorf("failed to read md file - %s : %s", filename, e)
			file.Close()
		}
		file.Close()
		if bytes.Contains(content, []byte(testcli.HiddenToken)) {
			t.Errorf("hidden flag or command in help output - %s", filename)
		}

		if !bytes.Contains(content, []byte("customLayout")) {
			t.Errorf("custom layout not in rendered file")
		}
	}

	dirList, err := os.ReadDir(output)
	if err != nil {
		t.Errorf("failed to list dir %s: %s", output, err)
	}
	filesWhichShouldNotBeGenerated := []string{
		"test-cli-hidden.md",
		"test-cli-deprecated.md",
	}
	for _, item := range dirList {
		for _, lookupItem := range filesWhichShouldNotBeGenerated {
			if item.Name() == lookupItem {
				t.Errorf("%s should not be present as it belongs to hidden/deprecated cmd", lookupItem)
			}
		}
	}
}

func Test_GenMarkdownTree_NilCommand(t *testing.T) {
	if GenMarkdownTree(nil, "testdata") == nil {
		t.Errorf("expected to error when cmd is nil")
	}
}

func Test_GenMarkdownTree_OutputDirNotPresent(t *testing.T) {
	if GenMarkdownTree(nil, "testdata/no-such-dir-present") == nil {
		t.Errorf("expected to error when output dir is not present")
	}
}
