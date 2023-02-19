// Package version is a helper for processing
// VCS, build, and runtime information of the binary.
// You can inject them at build time via ld flags.
package version

import (
	"runtime"
)

// Can override these at compile time.
var (
	// version is usually the git tag. MUST be semver compatible.
	//
	// You can override at build time using
	// 		-X github.com/pkg/version.version = "your-desired-version"
	version = "v0.0.0+undefined"
	// commit is git commit sha1 hash
	//
	// You can override at build time using
	// 		-X github.com/pkg/version.commit = "a-commit-hash"
	commit = ""
	// buildDate is build date.
	// For reproducible builds, set this to source epoch or commit date.
	//
	// You can override at build time using
	// 		-X github.com/pkg/version.buildDate = "build-date-in-format"
	buildDate = "1970-01-01T00:00+00:00"
)

// Info describes the build, revision and runtime information.
type Info struct {
	// Version  indicates which version of the binary is running.
	// In most cases this should be semver compatible string.
	//
	// Because we use go modules, it MUST include a prefix "v".
	// See [golang/go/issues/30146] as to why.
	//
	// [golang/go/issues/30146]: https://github.com/golang/go/issues/30146
	Version string `json:"version" yaml:"version"`
	// GitCommit indicates which git sha1 commit hash.
	GitCommit string `json:"gitCommit" yaml:"gitCommit"`
	// BuildDate date of the build.
	// You can set this to CommitDate to get truly reproducible and verifiable builds.
	BuildDate string `json:"buildDate" yaml:"buildDate"`
	// GoVersion version of Go runtime.
	GoVersion string `json:"goVersion" yaml:"goVersion"`
	// OperatingSystem this is operating system in GOOS
	Os string `json:"os" yaml:"os"`
	// Platform this is system platform
	Platform string `json:"platform" yaml:"platform"`
	// Compiler Go compiler.
	// This is useful in determining if binary was built using CGO.
	Compiler string `json:"compiler" yaml:"compiler"`
}

// Get returns version information. This usually relies on
// build tools injecting version info via ld flags.
func Get() Info {
	return Info{
		Version:   version,
		GitCommit: commit,
		BuildDate: buildDate,
		GoVersion: runtime.Version(),
		Os:        runtime.GOOS,
		Platform:  runtime.GOARCH,
		Compiler:  runtime.Compiler,
	}
}

// GetShort returns just the version information.
func GetShort() string {
	return version
}
