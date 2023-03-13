package factory

import (
	"strings"
	"testing"
)

func Test_replace(t *testing.T) {
	if strings.ReplaceAll("foo bar baz", " ", "-") != replace(" ", "-", "foo bar baz") {
		t.Error("replace() != strings.ReplaceAll()")
	}
}
