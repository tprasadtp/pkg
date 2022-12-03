package cli

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func checkOmit(t *testing.T, found, unexpected string) {
	if strings.Contains(found, unexpected) {
		t.Errorf("Got: %q\nBut should not have!\n", unexpected)
	}
}

func check(t *testing.T, found, expected string) {
	if !strings.Contains(found, expected) {
		t.Errorf("Expecting to contain: \n %q\nGot:\n %q\n", expected, found)
	}
}

func TestBashCompletionV2WithActiveHelp(t *testing.T) {
	c := &Command{Use: "c", Run: emptyRun}

	buf := new(bytes.Buffer)
	AssertErrNil(t, c.GenBashCompletion(buf))
	output := buf.String()

	// check that active help is not being disabled
	activeHelpVar := activeHelpEnvVar(c.Name())
	checkOmit(t, output, fmt.Sprintf("%s=0", activeHelpVar))
}
