// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/tprasadtp/knit/internal/command"

	_ "github.com/tprasadtp/go-autotune"       // GOMAXPROCS & GOMEMLIMIT
	_ "golang.org/x/crypto/x509roots/fallback" // Fallback Roots
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	err := command.Entrypoint(ctx)
	if err != nil {
		os.Exit(1)
	}
}
