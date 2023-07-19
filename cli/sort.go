package cli

import (
	"sort"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Compile time check to ensure byName
// implements sort interface.
var _ sort.Interface = (*byNameCmd)(nil)
var _ sort.Interface = (*byNameFlag)(nil)

// sortableCommands is used to sort commands by name
// in SEE ALSO section of docs and implements
// [sort.Interface] for sorting alphabetically.
type byNameCmd []*cobra.Command

func (s byNameCmd) Len() int {
	return len(s)
}

func (s byNameCmd) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byNameCmd) Less(i, j int) bool {
	return s[i].Name() < s[j].Name()
}

// sortableCommands is used to sort commands by name
// in SEE ALSO section of docs and implements
// [sort.Interface] for sorting alphabetically.
type byNameFlag []*pflag.Flag

func (s byNameFlag) Len() int {
	return len(s)
}

func (s byNameFlag) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byNameFlag) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}
