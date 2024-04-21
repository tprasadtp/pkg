// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: MIT

package auth

import (
	"context"
	"io"
)

type options struct {
	stdout io.Writer
}

func (o *options) Run(ctx context.Context, args []string) error {
	return nil
}
