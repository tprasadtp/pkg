// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: MIT

package shared

import (
	"fmt"
	"io/fs"
	"math/rand/v2"
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/tprasadtp/knit/internal/testutils"
)

func TestDirPermissionFrom(t *testing.T) {
	tt := []struct {
		input  fs.FileMode
		expect fs.FileMode
	}{
		{0, 0},
		{0o400, 0o500},
		{0440, 0o550},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("%o", tc.input), func(t *testing.T) {
			v := DirPermissionFrom(tc.input)
			if v != tc.expect {
				t.Errorf("expected=%s(%o), got=%s(%o)", tc.expect, tc.expect, v, v)
			}
		})
	}
}

func TestReadSmallFile(t *testing.T) {
	tt := []struct {
		name     string
		contents []byte
		max      int64
		ok       bool
		pre      func(t *testing.T, path string)
	}{
		{
			name: "empty-file",
			ok:   true,
		},
		{
			name:     "small-file",
			contents: []byte("foo\nbar\n"),
			ok:       true,
		},
		{
			name:     "small-file-max-negative",
			contents: []byte("foo\nbar\n"),
			max:      -1,
			ok:       true,
		},
		{
			name:     "small-file-larger-than-specified",
			contents: []byte("foo\nbar\n"),
			max:      1,
		},
		{
			name:     "missing-file",
			contents: []byte("foo\nbar\n"),
			pre: func(t *testing.T, path string) {
				err := os.Remove(path)
				if err != nil {
					t.Fatalf("Failed to remove file: %s", err)
				}
			},
		},
	}

	dir := t.TempDir()
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			path := filepath.Join(dir, fmt.Sprintf("%d.test", rand.Int()))
			testutils.CreateFileWithContents(t, path, tc.contents)

			if tc.pre != nil {
				tc.pre(t, path)
			}

			contents, err := ReadSmallFile(path, tc.max)
			if tc.ok {
				if !slices.Equal(tc.contents, contents) {
					t.Errorf("expected=%q, got=%q", tc.contents, contents)
				}
				if err != nil {
					t.Errorf("expected no error, got %s", err)
				}
			} else {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
				if contents != nil {
					t.Errorf("expected nil, got %s", contents)
				}
			}
		})
	}
}
