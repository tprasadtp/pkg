// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: MIT

package parse

import (
	"fmt"
	"io/fs"
	"strconv"
	"strings"
)

func FileMode[T ~string](input T) (fs.FileMode, error) {
	v := string(input)
	var rv uint64
	var err error
	switch {
	case strings.HasPrefix(v, "0o"), strings.HasPrefix(v, "0"):
		rv, err = strconv.ParseUint(v, 8, 32)
	default:
		rv, err = strconv.ParseUint(v, 10, 32)
	}

	if err != nil {
		return 0, fmt.Errorf("parse(filemode): %w", err)
	}
	return fs.FileMode(uint32(rv)), nil
}
