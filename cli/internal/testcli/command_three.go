package testcli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Command3 specific flags.
var (
	Command3RequiredTogetherFlag1 string
	Command3RequiredTogetherFlag2 string
	Command3FlagWithDefault       string
)

func Command3() *cobra.Command {
	// Root -> command3 (with required together flags)
	cmd := &cobra.Command{
		Use:   "command3",
		Short: "This is command3 short description",
		Long: `This is command3 long description.

Lorem ipsum dolor sit amet, consectetur adipiscing elit.
Pellentesque ut nunc fermentum, porta arcu in, molestie ante.
Pellentesque ullamcorper, magna et feugiat semper, turpis nibh tempor diam,
ac sodales dui ligula eget ligula.

Quisque ullamcorper ornare nulla, id vestibulum velit eleifend in.
Praesent eu dignissim nulla. Suspendisse congue aliquet dolor,
vel ullamcorper massa placerat in. Ut sit amet magna lectus.

Mauris bibendum euismod enim quis pellentesque Cras sit amet dolor vitae
ligula blandit varius. Quisque porta ullamcorper pellentesque.
Nullam maximus tellus ac lectus vulputate, vel aliquam nisl gravida.
Aenean dictum in libero a molestie. Vestibulum in ante sit amet tortor
lobortis porttitor at eu neque. In sit amet vestibulum nisl. Nunc blandit
arcu lacus, at faucibus urna commodo vitae. Sed sit amet orci at purus
pretium lacinia vel a purus. Quisque posuere sapien massa, sed volutpat ipsum
venenatis vitae.
`,
		Example: `
test-cli command3 --required-together-flag1 --required-together-flag2
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(cmd.Root().OutOrStdout(), "running - %s", cmd.CommandPath())
			return nil
		},
	}
	cmd.Flags().StringVar(
		&Command3RequiredTogetherFlag1,
		"required-together-flag1",
		"",
		"this flag is required-together with --required-together-flag2",
	)
	cmd.Flags().StringVar(
		&Command3RequiredTogetherFlag2,
		"required-together-flag2",
		"",
		"this flag is required-together with --required-together-flag1",
	)
	cmd.Flags().StringVar(
		&Command3FlagWithDefault,
		"flag-with-default",
		"value",
		"this flag has a default value",
	)
	cmd.MarkFlagsRequiredTogether("required-together-flag1", "required-together-flag2")
	return cmd
}
