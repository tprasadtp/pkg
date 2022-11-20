//go:build !dev

package factory

import (
	"github.com/tprasadtp/pkg/cli/cobra"
)

var (
	generateCmd *cobra.Command = nil
)
