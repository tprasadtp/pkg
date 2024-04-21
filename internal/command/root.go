// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: MIT

package command

import (
	"context"
	"log/slog"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/tprasadtp/knit/internal/command/semver"
	"github.com/tprasadtp/knit/internal/command/version"
)

func logFormatTextReplaceAttr() func(groups []string, a slog.Attr) slog.Attr {
	return func(groups []string, a slog.Attr) slog.Attr {
		if len(groups) == 0 {
			if a.Key == slog.TimeKey && a.Value.Kind() == slog.KindTime {
				a.Value = slog.StringValue(a.Value.Time().Format(time.Kitchen))
			}
		}
		return a
	}
}

// rootCommand is meant to be used as root command only and is not re-usable.
// This does some stuff with default loggers and global states thus not suitable
// for re-use outside knit main.
func rootCommand() *cobra.Command {
	var verbose bool
	var logger *slog.Logger
	var logFormat string

	cmd := &cobra.Command{
		Use:               "knit",
		Short:             "A Toolkit for building docker images",
		Version:           version.Version(),
		Args:              cobra.NoArgs,
		DisableAutoGenTag: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		// Set a logger based on slog for all go-containerregistry operations.
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			var level slog.Leveler = slog.LevelInfo
			if verbose {
				level = slog.LevelDebug
			}

			switch strings.ToLower(strings.TrimSpace(logFormat)) {
			case "json":
				logger = slog.New(
					slog.NewJSONHandler(
						cmd.ErrOrStderr(),
						&slog.HandlerOptions{
							Level: level,
						},
					),
				)
			default:
				logger = slog.New(
					slog.NewTextHandler(
						cmd.ErrOrStderr(),
						&slog.HandlerOptions{
							Level:       level,
							ReplaceAttr: logFormatTextReplaceAttr(),
						},
					),
				)
			}
			logger.DebugContext(cmd.Context(), "enabled debug level logs")
			return nil
		},
	}

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose logs")
	cmd.AddCommand(semver.NewCommand())
	cmd.AddCommand(version.NewVersionCmd())

	fixCobraBehavior(cmd)
	return cmd
}

// Entrypoint is entrypoint to the knit binary.
func Entrypoint(ctx context.Context) error {
	//nolint:wrapcheck // ignore cli entrypoint.
	return rootCommand().ExecuteContext(ctx)
}
