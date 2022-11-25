//go:build !dev

package cli

import (
	"github.com/spf13/cobra"
)

var (
	generateCmd *cobra.Command = nil
)
