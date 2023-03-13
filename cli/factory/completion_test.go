package factory_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/tprasadtp/pkg/cli/factory"
)

func Test_CompletionCmd_ToStdOut_Success(t *testing.T) {
	var stdout = new(bytes.Buffer)
	var stderr = new(bytes.Buffer)
	type testCase struct {
		Name         string
		StdOutString string
		ArgsVariants [][]string
	}
	tt := []testCase{
		{
			Name:         "command-bash",
			StdOutString: "# bash completion V2",
			ArgsVariants: [][]string{
				{"completion", "bash"},
				{"completion", "bash", "--output=-"},
			},
		},
		{
			Name: "command-zsh",
			ArgsVariants: [][]string{
				{"completion", "zsh"},
				{"completion", "zsh", "--output=-"},
			},
			StdOutString: "# zsh completion",
		},
		{
			Name: "command-fish",
			ArgsVariants: [][]string{
				{"completion", "fish"},
				{"completion", "fish", "--output=-"},
			},
			StdOutString: "# fish completion",
		},
		{
			Name: "command-pwsh-no-args",
			ArgsVariants: [][]string{
				{"completion", "powershell"},
				{"completion", "pwsh"},
				{"completion", "powershell", "--output=-"},
				{"completion", "pwsh", "--output=-"},
			},
			StdOutString: "# powershell completion",
		},
	}
	for _, tc := range tt {
		for _, cmdArgs := range tc.ArgsVariants {
			testName := fmt.Sprintf("%s-with-args-%s", tc.Name, strings.Join(cmdArgs, "-"))
			t.Run(testName, func(t *testing.T) {
				stdout.Reset()
				stderr.Reset()
				root := &cobra.Command{
					Use:   "blackhole-entropy",
					Short: "Black Hole Entropy CLI",
					Long:  "CLI to seed system's PRNG with entropy from M87 Black Hole",
				}
				root.SetOut(stdout)
				root.SetErr(stderr)
				root.AddCommand(factory.NewCompletionCmd())
				root.SetArgs(cmdArgs)
				err := root.Execute()
				if stdout.String() == "" {
					t.Errorf("when no args specified stdout must contain bash completion")
				}

				if !strings.Contains(stdout.String(), tc.StdOutString) {
					t.Errorf("with %s does not contain expected completion header %s, got \n %s",
						cmdArgs[1], tc.StdOutString, stdout.String())
				}

				if stderr.String() != "" {
					t.Errorf("when no args specified stderr must be empty")
				}

				if err != nil {
					t.Errorf("%s: must not return any error, but got %s", t.Name(), err)
				}
			})
		}
	}
}

func Test_CompletionCmd_InvalidSubCommands(t *testing.T) {
	var stdout = new(bytes.Buffer)
	var stderr = new(bytes.Buffer)
	type testCase struct {
		Name string
		Args []string
	}
	tt := []testCase{
		{
			Name: "command-invalid-shell",
			Args: []string{"completion", "ion-shell"},
		},
		{
			Name: "command-no-shell-specified",
			Args: []string{"completion"},
		},
	}
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			root := &cobra.Command{
				Use:   "blackhole-entropy",
				Short: "Black Hole Entropy CLI",
				Long:  "CLI to seed system's PRNG with entropy from M87 Black Hole",
			}
			root.SetOut(stdout)
			root.SetErr(stderr)
			root.AddCommand(factory.NewCompletionCmd())
			root.SetArgs(tc.Args)
			err := root.Execute()
			if err == nil {
				t.Errorf("%s: must return an error, but got nil", t.Name())
			}
		})
	}
}

func Test_CompletionCmd_ToFile(t *testing.T) {
	var stdout = new(bytes.Buffer)
	var stderr = new(bytes.Buffer)
	type testCase struct {
		Name       string
		Args       []string
		FindString string
	}
	tt := []testCase{
		{
			Name: "bash",
			Args: []string{
				"completion",
				"bash",
				"--output", "testdata/completions-test.bash",
			},
			FindString: "# bash completion V2",
		},
		{
			Name: "zsh",
			Args: []string{
				"completion",
				"zsh",
				"--output", "testdata/completions-test.zsh",
			},
			FindString: "# zsh completion",
		},
		{
			Name: "fish",
			Args: []string{
				"completion",
				"fish",
				"--output", "testdata/completions-test.fish",
			},
			FindString: "# fish completion",
		},
		{
			Name: "powershell",
			Args: []string{
				"completion",
				"powershell",
				"--output", "testdata/completions-test.powershell",
			},
			FindString: "# powershell completion",
		},
		{
			Name: "pwsh",
			Args: []string{
				"completion",
				"pwsh",
				"--output", "testdata/completions-test.pwsh",
			},
			FindString: "# powershell completion",
		},
	}
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			root := &cobra.Command{
				Use:   "blackhole-entropy",
				Short: "Black Hole Entropy CLI",
				Long:  "CLI to seed system's PRNG with entropy from M87 Black Hole",
			}
			root.SetOut(stdout)
			root.SetErr(stderr)
			root.AddCommand(factory.NewCompletionCmd())
			root.SetArgs(tc.Args)
			err := root.Execute()

			// Validate stdout is empty
			if stdout.String() != "" {
				t.Errorf("should not output anything when output is file")
			}

			if stderr.String() != "" {
				t.Errorf("when no args specified stderr must be empty")
			}

			if err != nil {
				t.Fatalf("%s: must not return any error, but got %s", t.Name(), err)
			}

			// Open file and check if it is correct completion
			outputFile, outputReadErr := os.Open(tc.Args[3])
			if outputReadErr != nil {
				t.Fatalf("failed to open file for verification: %s", outputReadErr)
			}
			// close and delete generated file.
			t.Cleanup(func() {
				outputFile.Close()
				os.Remove(outputFile.Name())
			})
			outputFileContents, outputReadErr := io.ReadAll(outputFile)
			if outputReadErr != nil {
				t.Fatalf("failed to read file for verification: %s", outputReadErr)
			}

			if !strings.Contains(string(outputFileContents), tc.FindString) {
				t.Fatalf("generated file does not contain string: %s", tc.FindString)
			}
		})
	}
}

func Test_NewCompletionCmd_HiddenAttrs(t *testing.T) {
	type testCase struct {
		Name               string
		Args               []bool
		ExpectHiddenStatus bool
	}
	tt := []testCase{
		{
			Name:               "no-flag-specified",
			ExpectHiddenStatus: false,
		},
		{
			Name:               "single-bool",
			Args:               []bool{true},
			ExpectHiddenStatus: true,
		},
		{
			Name:               "multiple-bool-only-use-first-1",
			Args:               []bool{true, false},
			ExpectHiddenStatus: true,
		},
		{
			Name:               "multiple-bool-only-use-first-1",
			Args:               []bool{false, true},
			ExpectHiddenStatus: false,
		},
	}
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			cmd := factory.NewCompletionCmd(tc.Args...)
			if cmd.Hidden != tc.ExpectHiddenStatus {
				t.Errorf("Expected Hidden status=%t, got=%t", tc.ExpectHiddenStatus, cmd.Hidden)
			}
		})
	}
}
