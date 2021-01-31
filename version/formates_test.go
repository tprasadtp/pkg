package version

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tprasadtp/pkg/apollo"
)

func TestJSON(t *testing.T) {
	a := apollo.New(t, apollo.WithDiffEngine(apollo.ColoredDiff))

	v := Describe()
	expected, err := v.AsJSON()
	assert.Nil(t, err)
	a.Assert(t, "json", []byte(expected))
}

func TestYAML(t *testing.T) {
	a := apollo.New(t, apollo.WithDiffEngine(apollo.ColoredDiff))
	v := Describe()
	expected, err := v.AsYAML()
	assert.Nil(t, err)
	a.Assert(t, "yaml", []byte(expected))
}

func TestVersionFormats(t *testing.T) {
	a := apollo.New(t, apollo.WithDiffEngine(apollo.ColoredDiff))
	tests := []struct {
		name  string
		short bool
	}{
		{name: "text", short: false},
		{name: "text-short", short: true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := Describe()
			expected, err := v.AsText(tc.short)
			assert.Nil(t, err)
			a.Assert(t, tc.name, []byte(expected))
		})
	}
}

func ExampleInfo_json() {
	v := Describe()
	j, err := v.AsJSON()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", j)
}

func ExampleInfo_yaml() {
	v := Describe()
	j, err := v.AsYAML()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", j)
}

func ExampleInfo_pretty() {
	v := Describe()
	j, err := v.AsText(false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", j)
}
