package ansi

import (
	_ "embed"
	"testing"
)

//go:embed testdata/styles.input.txt
var small string

//go:embed testdata/plain.input.txt
var longPlain string

//go:embed testdata/neofetch.input.txt
var neofetch string

//go:embed testdata/cli-help.input.txt
var cliHelpText string

//go:embed testdata/termenv.input.txt
var termenvTable string

func testStripBenchUnit(str string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Strip(str)
	}
}

func BenchmarkSmall(b *testing.B) {
	testStripBenchUnit(small, b)
}

func BenchmarkLongPlain(b *testing.B) {
	testStripBenchUnit(longPlain, b)
}

func BenchmarkNeofetch(b *testing.B) {
	testStripBenchUnit(neofetch, b)
}

func BenchmarkHelpText(b *testing.B) {
	testStripBenchUnit(cliHelpText, b)
}

func BenchmarkTermEnvTable(b *testing.B) {
	testStripBenchUnit(termenvTable, b)
}
