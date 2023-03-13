package testcli

import (
	"github.com/spf13/cobra"
)

// Global Persistent flags.
var (
	GlobalBoolFlag       bool
	GlobalStringFlag     string
	GlobalHiddenFlag     bool
	GlobalDeprecatedFlag bool
)

// Test CLI Root.
func GetTestCLI() *cobra.Command {
	root := &cobra.Command{
		Use:   "test-cli",
		Short: "This is root command short description",
		Long: `This is root command long description.

This can span multiple lines.

- Item 1
- Item 2
`,
	}
	root.PersistentFlags().BoolVar(
		&GlobalBoolFlag,
		"global-flag",
		false,
		"global-flag (from root)",
	)
	root.PersistentFlags().StringVarP(
		&GlobalStringFlag,
		"global-string-flag",
		"s",
		"string-value",
		"global-string-flag (from root)",
	)
	root.PersistentFlags().BoolVar(
		&GlobalHiddenFlag,
		"global-hidden-flag",
		false,
		"global-hidden-flag (from root) token=54dac5afe1fcac2f65c059fc97b44a58",
	)
	_ = root.PersistentFlags().MarkHidden("global-hidden-flag")

	root.PersistentFlags().BoolVar(
		&GlobalDeprecatedFlag,
		"global-deprecated-flag",
		false,
		"global-deprecated-flag (from root) token=54dac5afe1fcac2f65c059fc97b44a58",
	)
	// printf "deprecated-" | md5sum | cut -d ' ' -f 1
	_ = root.PersistentFlags().MarkDeprecated(
		"global-deprecated-flag",
		"global-deprecated-flag is deprecated",
	)
	root.AddCommand(
		Command1(),
		Command2(),
		Command3(),
		HiddenCmd(),
		DeprecatedCmd(),
	)
	return root
}
