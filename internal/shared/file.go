// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: MIT

package shared

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
)

// DirPermissionFrom derives directory permission from the given
// file permissions. Directory is set to be traversable by a user
// only if file permissions allow read operations.
func DirPermissionFrom(mode fs.FileMode) fs.FileMode {
	if mode&0o400 == 0o400 {
		mode |= 0o100
	}
	if mode&0o040 == 0o040 {
		mode |= 0o010
	}
	if mode&0o004 == 0o004 {
		mode |= 0o001
	}
	return mode
}

func ReadSmallFile(path string, max int64) ([]byte, error) {
	vf, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("shared(file): failed to open file(%q): %w", path, err)
	}
	defer vf.Close()
	stat, err := vf.Stat()
	if err != nil {
		return nil, fmt.Errorf("shared(file): failed to stat file(%q): %w", path, err)
	}

	if max > 0 {
		if stat.Size() > max {
			return nil, fmt.Errorf("shared(file): file(%q) is too large(%dB)", path, stat.Size())
		}
	}

	contents, err := io.ReadAll(vf)
	if err != nil {
		return nil, fmt.Errorf("shared(file): error reading file(%q): %w", path, err)
	}
	return contents, nil
}

// WriteToFile writes given contents to the file specified.
//
//   - If append is true, then output will be appended to the file.
//     Otherwise it will be overwritten.
//   - If directory does not exist, it will be created. Directory created
//     will have permission based on file permissions specified.
//   - In append mode, output is always written on a
//     new line unless lf is set to false.
//
//nolint:gocognit // ignore
func WriteToFile[T ~[]byte](path string, contents T, appendFile, appendOnNewLine bool, mode fs.FileMode) error {
	if path == "" {
		return fmt.Errorf("shared(template): path is empty")
	}

	var flag int
	if appendFile {
		if appendOnNewLine {
			flag = os.O_CREATE | os.O_RDWR | os.O_APPEND
		} else {
			flag = os.O_CREATE | os.O_WRONLY | os.O_APPEND
		}
	} else {
		flag = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	}

	// If mode is not specified, default to read/write owner.
	if mode == 0 {
		mode = fs.FileMode(0o600)
	}

	// Create base directory if required. Permission on the directory
	// is derived from the permission on the file.
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, DirPermissionFrom(mode))
	if err != nil {
		return fmt.Errorf("shared(file): failed to create dir file(%q): %w", dir, err)
	}

	// Create/Open file.
	file, err := os.OpenFile(path, flag, mode)
	if err != nil {
		return fmt.Errorf("shared(file): failed to open file(%q): %w", path, err)
	}
	defer file.Close()

	// Ensure to write on a new line unless disabled.
	//nolint:nestif // ignore
	if appendFile && appendOnNewLine {
		stat, err := file.Stat()
		if err != nil {
			return fmt.Errorf("shared(file): failed to stat file(%q): %w", path, err)
		}
		if stat.Size() > 0 {
			prefix := make([]byte, 0, 2)
			if runtime.GOOS == "windows" {
				if stat.Size() > 2 {
					b := make([]byte, 2)
					_, err = file.ReadAt(b, stat.Size()-2)
					if err != nil {
						return fmt.Errorf("shared(file): failed to check for new-line in file(%q): %w", path, err)
					}
					if b[0] != '\r' || b[1] != '\n' {
						prefix = append(prefix, '\r')
						prefix = append(prefix, '\n')
					}
				} else {
					b := make([]byte, 1)
					_, err = file.ReadAt(b, stat.Size()-1)
					if err != nil {
						return fmt.Errorf("shared(file): failed to check for new-line in file(%q): %w", path, err)
					}
					if b[0] != '\n' {
						prefix = append(prefix, '\r')
						prefix = append(prefix, '\n')
					}
				}
			} else {
				b := make([]byte, 1)
				_, err = file.ReadAt(b, stat.Size()-1)
				if err != nil {
					return fmt.Errorf("shared(file): failed to check for new-line in file(%q): %w", path, err)
				}
				if b[0] != '\n' {
					prefix = append(prefix, '\n')
				}
			}

			// Prefix is set to newline char to write to ensure output is written on a newline.
			if len(prefix) > 0 {
				_, err = file.Write(prefix)
				if err != nil {
					return fmt.Errorf("shared(file): failed to append new-line to file(%q): %w", path, err)
				}
			}
		}
	}
	_, err = file.Write(contents)
	if err != nil {
		return fmt.Errorf("shared(file): failed to write contents to file(%q): %w", path, err)
	}
	return nil
}
