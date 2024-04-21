// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: MIT

package oci

import (
	"log/slog"

	"github.com/google/go-containerregistry/pkg/logs"
	"github.com/spf13/cobra"
)

func NewCommand(logger *slog.Logger) *cobra.Command {
	if logger == nil {
		logger = slog.Default()
	}
	cmd := &cobra.Command{
		Use:     "oci",
		Aliases: []string{"registry", "images"},
		// Set loggers for go-containerregistry based on the given slog handler.
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			logs.Debug = slog.NewLogLogger(logger.Handler(), slog.LevelDebug)
			logs.Progress = slog.NewLogLogger(logger.Handler(), slog.LevelDebug)
			logs.Warn = slog.NewLogLogger(logger.Handler(), slog.LevelWarn)
			return nil
		},
	}
	return cmd
}
