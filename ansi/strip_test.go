package ansi

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/tprasadtp/pkg/apollo"
)

func TestStrip(t *testing.T) {
	a := apollo.New(t, apollo.WithDiffEngine(apollo.ColoredDiff))
	tests := []struct {
		name string
	}{
		{name: "styles"},
		{name: "neofetch"},
		{name: "plain"},
		{name: "cli-help"},
		{name: "termenv"},
		{name: "ps1"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			inputFile := path.Join("testdata", fmt.Sprintf("%s.input.txt", tc.name))
			if f, err := os.ReadFile(inputFile); err == nil {
				v := Strip(string(f))
				a.Assert(t, tc.name, []byte(v))
			} else {
				t.Errorf("failed-%s because:%+v", tc.name, err)
			}
		})
	}
}
