package factory

import (
	"bytes"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/tprasadtp/pkg/cli/internal/testcli"
)

func Test_GenManTree_NotCompressed(t *testing.T) {
	output := t.TempDir()
	root := testcli.GetTestCLI()
	err := GenManTree(root, output, false)
	if err != nil {
		t.Fatalf("failed to generate markdown - %s", err)
	}
	token := []byte(testcli.HiddenToken)
	filesWhichShouldBeGenerated := []string{
		"test-cli.1",
		"test-cli-command1.1",
		"test-cli-command1-subcommand1.1",
		"test-cli-command1-subcommand2.1",
		"test-cli-command2.1",
	}

	// we know what output file is check if
	// any deprecated flags are present.
	// testcli.HiddenToken is in all usage messages.
	for _, f := range filesWhichShouldBeGenerated {
		filename := filepath.Join(output, f)
		file, e := os.Open(filename)
		if e != nil {
			t.Errorf("failed to open man file - %s: %e", filename, e)
			continue
		}
		content, e := io.ReadAll(file)
		if e != nil {
			t.Errorf("failed to read man file - %s : %s", filename, e)
			file.Close()
		}
		file.Close()
		if bytes.Contains(content, token) {
			t.Errorf("hidden flag or command in man output - %s", filename)
		}
	}

	dirList, err := os.ReadDir(output)
	if err != nil {
		t.Errorf("failed to list dir %s: %s", output, err)
	}
	filesWhichShouldNotBeGenerated := []string{
		"test-cli-hidden.1",
		"test-cli-deprecated.1",
	}
	for _, item := range dirList {
		for _, lookupItem := range filesWhichShouldNotBeGenerated {
			if item.Name() == lookupItem {
				t.Errorf("%s should not be present as it belongs to hidden/deprecated cmd", lookupItem)
			}
		}
	}
}

func Test_GenManTree_Compress(t *testing.T) {
	output := t.TempDir()
	root := testcli.GetTestCLI()
	err := GenManTree(root, output, true)
	if err != nil {
		t.Fatalf("failed to generate markdown - %s", err)
	}
	filesWhichShouldBeGenerated := []string{
		"test-cli.1.gz",
		"test-cli-command1.1.gz",
		"test-cli-command1-subcommand1.1.gz",
		"test-cli-command1-subcommand2.1.gz",
		"test-cli-command2.1.gz",
	}

	// we know what output file is check if
	// any deprecated flags are present.
	// testcli.HiddenToken is in all usage messages.
	for _, f := range filesWhichShouldBeGenerated {
		filename := filepath.Join(output, f)
		file, e := os.Open(filename)
		if e != nil {
			t.Errorf("failed to open compressed man file - %s: %e", filename, e)
			continue
		}
		gzReader, e := gzip.NewReader(file)
		if e != nil {
			file.Close()
			t.Errorf("failed to create gzip reader from %s : %s", filename, e)
		}
		content, e := io.ReadAll(gzReader)
		if e != nil {
			file.Close()
			t.Errorf("failed to read compressed man file - %s : %s", filename, e)
		}
		file.Close()
		if bytes.Contains(content, []byte(testcli.HiddenToken)) {
			t.Errorf("hidden flag or command in help output - %s", filename)
		}
	}

	dirList, err := os.ReadDir(output)
	if err != nil {
		t.Errorf("failed to list dir %s: %s", output, err)
	}
	filesWhichShouldNotBeGenerated := []string{
		"test-cli-hidden.1.gz",
		"test-cli-deprecated.1.gz",
	}
	for _, item := range dirList {
		for _, lookupItem := range filesWhichShouldNotBeGenerated {
			if item.Name() == lookupItem {
				t.Errorf("%s should not be present as it belongs to hidden/deprecated cmd", lookupItem)
			}
		}
	}
}

func Test_GenManTree_NilCommand(t *testing.T) {
	if GenManTree(nil, "testdata", false) == nil {
		t.Errorf("expected to error when cmd is nil")
	}
}

func Test_GenManTree_OutputDirNotPresent(t *testing.T) {
	if GenManTree(nil, "testdata/no-such-dir-present", false) == nil {
		t.Errorf("expected to error when output dir is not present")
	}
}
