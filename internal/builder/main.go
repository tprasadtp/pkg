// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: GPLv3-only

package main

import (
	"context"
	"os"
	"os/signal"

	_ "github.com/tprasadtp/go-autotune"
	"github.com/tprasadtp/knit/internal/command"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	err := command.RootCommand().ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}
