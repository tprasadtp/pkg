// Package version is a helper for processing
// VCS, build, and runtime information of the binary.
// You can inject them at build time via ld flags.
// If not already injected, this uses debug.ReadBuildInfo,
// to get version information if any.
package version

import (
	"runtime"
	"runtime/debug"
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

//nolint:gochecknoinits // Reads embedded build info.
func init() {
	v, ok := debug.ReadBuildInfo()
	if ok {
		for _, item := range v.Settings {
			switch item.Key {
			case "vcs.revision":
				if commit == "" {
					commit = item.Value
				}
			case "vcs.time":
				if buildDate == "1970-01-01T00:00+00:00" {
					buildDate = item.Value
				}
			}
		}
	}
}

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
	// Arch this is system Arch
	Arch string `json:"platform" yaml:"arch"`
	// Compiler Go compiler.
	// This is useful in determining if binary was built using CGO.
	Compiler string `json:"compiler" yaml:"compiler"`
}

// GetInfo returns version information. This usually relies on
// build tools injecting version info via ld flags.
func GetInfo() Info {
	return Info{
		Version:   version,
		GitCommit: commit,
		BuildDate: buildDate,
		GoVersion: runtime.Version(),
		Os:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		Compiler:  runtime.Compiler,
	}
}
