package testcli

import (
	"github.com/spf13/cobra"
)

// Command1 specific flags.
var (
	Command1PersistentFlag       bool
	Command1HiddenPersistentFlag string
)

// non runnable command with subcommands.
// also has a flag with deprecated short form.
func Command1() *cobra.Command {
	command1 := &cobra.Command{
		Use:   "command1",
		Short: "This is command1 short description",
		Long: `This is command1 long description.

command1 has subcommands of its own.
One subcommand is hidden. There are few persistent flags.`}
	command1.PersistentFlags().BoolVarP(
		&Command1PersistentFlag,
		"command1-persistent-flag",
		"p",
		false,
		"persistent-flag (from command1)",
	)
	_ = command1.Flags().MarkShorthandDeprecated("persistent-flag", "persistent-flag shorthand is deprecated")
	command1.PersistentFlags().StringVar(
		&Command1HiddenPersistentFlag,
		"command1-persistent-flag-hidden",
		"default-value",
		"persistent-flag-hidden (from command1) 54dac5afe1fcac2f65c059fc97b44a58",
	)
	//nolint:errcheck // testing code
	command1.PersistentFlags().MarkHidden("command1-persistent-flag-hidden")
	command1.AddCommand(Command1Subcommand1(), Command1Subcommand2())
	return command1
}
