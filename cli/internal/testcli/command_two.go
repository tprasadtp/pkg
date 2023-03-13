package testcli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Command2 specific flags.
var (
	Command2MutuallyExclusiveFlag1 bool
	Command2MutuallyExclusiveFlag2 bool
)

// Root -> command2 with mutually exclusive flags.
func Command2() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "command2",
		Short: "This is command2 short description",
		Long: `This is command2 long description.

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
test-cli command2 --required-together-flag1 --required-together-flag2
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(cmd.Root().OutOrStdout(), "running - %s", cmd.CommandPath())
			return nil
		},
	}
	cmd.Flags().BoolVar(
		&Command2MutuallyExclusiveFlag1,
		"mutually-exclusive-flag1",
		false,
		"this flag is mutually-exclusive with --mutually-exclusive-flag2",
	)
	cmd.Flags().BoolVar(
		&Command2MutuallyExclusiveFlag1,
		"mutually-exclusive-flag2",
		false,
		"this flag is mutually-exclusive with --mutually-exclusive-flag1",
	)
	cmd.MarkFlagsMutuallyExclusive("mutually-exclusive-flag1", "mutually-exclusive-flag2")
	return cmd
}
