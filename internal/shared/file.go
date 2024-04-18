// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: GPLv3-only

package shared

import (
	"fmt"
	"io"
	"os"
)

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
