package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tprasadtp/pkg/cli/factory"
)

func main() {
	root := &cobra.Command{
		Use:   "blackhole-entropy",
		Short: "Black Hole Entropy CLI",
		Long:  "CLI to seed system's PRNG with entropy from M87 Black Hole",
	}
	root.AddCommand(factory.NewCompletionCmd(root.Name()))
	root.AddCommand(factory.NewVersionCmd(root.Name()))
	factory.FixCobraBehavior(root)
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
