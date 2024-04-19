// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: MIT

package testutils

import (
	"bytes"
	"crypto/rand"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

// CreateFileWithJunk is a helper which creates file specified and fills it with junk bytes
// of size specified. If size > 100MB test errors.
func CreateFileWithJunk(t *testing.T, path string, size uint64) {
	t.Helper()

	if size > 100e6 {
		t.Fatalf("refusing to write > 100MB of juk bytes.")
	}

	junk := make([]byte, size)
	_, err := rand.Reader.Read(junk)
	if err != nil {
		t.Fatalf("failed to read random %d bytes: %s", size, err)
	}
	CreateFileWithContents(t, path, junk)
}

// CreateFileWithContents is a helper which creates file specified and fills it with content
// specified.
func CreateFileWithContents[T string | []byte](t *testing.T, path string, contents T) {
	t.Helper()

	if !filepath.IsAbs(path) {
		t.Fatalf("path must be absolute: %s", path)
	}

	temp, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY|os.O_TRUNC, os.FileMode(0o600))
	if err != nil {
		t.Fatalf("failed to create file(%q): %s", path, err)
	}

	_, err = temp.Write([]byte(contents))
	if err != nil {
		t.Fatalf("failed write contents to file(%q): %s", path, err)
	}
	err = temp.Close()
	if err != nil {
		t.Fatalf("failed to close file(%q): %s", path, err)
	}
}

func AssertFileNotEmpty(t *testing.T, path string) {
	file, err := os.Open(path)
	if err != nil {
		t.Errorf("failed to open file(%q): %s", path, err)
		return
	}
	t.Cleanup(func() {
		file.Close()
	})
	stat, err := file.Stat()
	if err != nil {
		t.Errorf("failed to stat file(%q): %s", path, err)
		return
	}

	if stat.Size() <= 0 {
		t.Errorf("expected non empty file: %q", path)
	}
}

func RequireFileNotEmpty(t *testing.T, path string) {
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("failed to open file(%q): %s", path, err)
	}
	t.Cleanup(func() {
		file.Close()
	})
	stat, err := file.Stat()
	if err != nil {
		t.Fatalf("failed to stat file(%q): %s", path, err)
	}

	if stat.Size() <= 0 {
		t.Fatalf("expected non empty file: %q", path)
	}
}

func AssertFileMode(t *testing.T, path string, mode fs.FileMode) {
	stat, err := os.Stat(path)
	if err != nil {
		t.Errorf("failed to stat file(%q): %s", path, err)
		return
	}

	if stat.Mode() != mode {
		t.Errorf("expected file mode=%s(%d) got=%s(%d)", stat.Mode(), stat.Mode(), mode, mode)
	}
}

func RequireFileMode(t *testing.T, path string, mode fs.FileMode) {
	stat, err := os.Stat(path)
	if err != nil {
		t.Fatalf("failed to stat file(%q): %s", path, err)
	}

	if stat.Mode() != mode {
		t.Fatalf("expected file mode=%s(%d) got=%s(%d)", stat.Mode(), stat.Mode(), mode, mode)
	}
}

func AssertFileContents(t *testing.T, path string, contents []byte) {
	file, err := os.Open(path)
	if err != nil {
		t.Errorf("failed to open file(%q): %s", path, err)
		return
	}
	t.Cleanup(func() {
		file.Close()
	})
	stat, err := file.Stat()
	if err != nil {
		t.Errorf("failed to stat file(%q): %s", path, err)
		return
	}

	if stat.Size() > 50e3 {
		t.Errorf("file is too large to load in memory: %s", path)
		return
	}

	buf, err := io.ReadAll(file)
	if err != nil {
		t.Errorf("failed to read file(%q): %s", path, err)
		return
	}

	if !bytes.Equal(buf, contents) {
		t.Errorf("expected=%v, got=%v", contents, buf)
		return
	}
}

func RequireFileContents(t *testing.T, path string, contents []byte) {
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("failed to open file(%q): %s", path, err)
	}
	t.Cleanup(func() {
		file.Close()
	})
	stat, err := file.Stat()
	if err != nil {
		t.Fatalf("failed to stat file(%q): %s", path, err)
	}

	if stat.Size() > 50e3 {
		t.Fatalf("file is too large to load in memory: %s", path)
	}

	buf, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("failed to read file(%q): %s", path, err)
	}

	if !bytes.Equal(buf, contents) {
		t.Fatalf("expected=%v, got=%v", contents, buf)
	}
}
