package factory

import (
	"sort"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func Test_SortCmdByName(t *testing.T) {
	cmdApple := &cobra.Command{
		Use: "apple",
	}
	cmdOrange := &cobra.Command{
		Use: "orange",
	}
	cmdPineapple := &cobra.Command{
		Use: "pineapple",
	}
	cmdKiwi := &cobra.Command{
		Use: "kiwi",
	}

	fruits := byNameCmd{
		cmdPineapple,
		cmdOrange,
		cmdApple,
		cmdKiwi,
	}
	sort.Sort(fruits)
	for i, item := range []string{"apple", "kiwi", "orange", "pineapple"} {
		if fruits[i].Name() != item {
			t.Errorf("index[%d] expected %s but got %s", i, item, fruits[i].Name())
		}
	}
}

func Test_SortFlagsByName(t *testing.T) {
	apple := &pflag.Flag{
		Name:      "apple",
		Shorthand: "a",
	}
	orange := &pflag.Flag{
		Name:      "orange",
		Shorthand: "o",
	}
	pineapple := &pflag.Flag{
		Name:      "pineapple",
		Shorthand: "p",
	}
	kiwi := &pflag.Flag{
		Name:      "kiwi",
		Shorthand: "k",
	}

	fruits := byNameFlag{
		orange, kiwi, pineapple, apple,
	}
	sort.Sort(fruits)
	for i, item := range []string{"apple", "kiwi", "orange", "pineapple"} {
		if fruits[i].Name != item {
			t.Errorf("index[%d] expected %s but got %s", i, item, fruits[i].Name)
		}
	}
}
