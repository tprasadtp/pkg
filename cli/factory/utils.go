package factory

import (
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const defaultTimeFormat = "2-Jan-2006"

// getSeeAlso returns elements that should be in see also section.
func getSeeAlso(cmd *cobra.Command) []*cobra.Command {
	// If command is  nil, hidden or deprecated return none.
	if cmd == nil {
		return nil
	}
	if cmd.Hidden || len(cmd.Deprecated) > 0 {
		return nil
	}

	var rv []*cobra.Command
	if cmd.HasParent() {
		if cmd.Parent().IsAvailableCommand() &&
			cmd.Parent().Name() != cmd.Root().Name() &&
			!cmd.Parent().IsAdditionalHelpTopicCommand() {
			rv = append(rv, cmd.Parent())
		}
		// Iterate over parent subcommands, aka side looker.
		// This finds commands which share same cmd.Parent().
		for _, c := range cmd.Parent().Commands() {
			if c.IsAvailableCommand() &&
				!c.IsAdditionalHelpTopicCommand() &&
				c.CommandPath() != cmd.CommandPath() {
				rv = append(rv, c)
			}
		}
	}

	// Iterate over subcommands, aka down looker.
	if cmd.HasAvailableSubCommands() {
		for _, c := range cmd.Commands() {
			if c.IsAvailableCommand() && !c.IsAdditionalHelpTopicCommand() {
				rv = append(rv, c)
			}
		}
	}
	sort.Sort(byNameCmd(rv))
	return rv
}

// getFlags returns all non hidden and non deprecated flags for a command.
// This may return flags whose shorthand form is deprecated, but flag itself is not.
// returned flags are sorted.
func getFlags(cmd *cobra.Command) []*pflag.Flag {
	if cmd == nil {
		return nil
	}

	if cmd.Hidden || len(cmd.Deprecated) > 0 {
		return nil
	}

	var rv []*pflag.Flag
	cmd.NonInheritedFlags().VisitAll(func(flag *pflag.Flag) {
		if len(flag.Deprecated) == 0 && !flag.Hidden {
			rv = append(rv, flag)
		}
	})
	cmd.InheritedFlags().VisitAll(func(flag *pflag.Flag) {
		if len(flag.Deprecated) == 0 && !flag.Hidden {
			rv = append(rv, flag)
		}
	})
	cmd.HasParent()
	sort.Sort(byNameFlag(rv))
	return rv
}

// Test if DisableAutogenTag is set on parent command or the root command or the current command.
func isAutoGenDisabled(cmd *cobra.Command) bool {
	if cmd.HasParent() {
		return cmd.Root().DisableAutoGenTag || cmd.Parent().DisableAutoGenTag || cmd.DisableAutoGenTag
	}
	return cmd.Root().DisableAutoGenTag || cmd.DisableAutoGenTag
}

// if SOURCE_DATE_EPOCH is defined, uses it for generated at timestamp.
// or current time is used. uses [time.Format] for formatting.
func formatGeneratedAt(format string) string {
	if format == "" {
		format = defaultTimeFormat
	}
	if epoch := os.Getenv("SOURCE_DATE_EPOCH"); epoch != "" {
		unixEpoch, err := strconv.ParseInt(epoch, 10, 64)
		if err == nil {
			return time.Unix(unixEpoch, 0).Format(format)
		}
	}
	return time.Now().Format(format)
}
