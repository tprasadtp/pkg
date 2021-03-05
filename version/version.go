// Package version is a helper for processing VCS, build, and runtime information of the binary.
package version

import (
	"flag"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

var version string = "0.0.0+dev"

var gitCommit string

var gitTreeState string = "unknown"

var gitBranch string

var buildDate string = "1970-01-01T00:00+00:00"

var buildHost string = "localhost"

var buildNumber string = "0"

var buildSystem string = "go"

// BuildInfo Build Information.
type BuildInfo struct {
	// CI/CD Build Number. For local systems,
	// which lack concept of build number this is set to zero.
	Number uint64 `json:"number" yaml:"number"`
	// Build-system which built the binary.
	// Useful to identify if binary was built locally or by a CI/CD pipeline
	// convention is to follow {CI_SYSTEM_NAME}-{PIPELINE_NAME}. Default("GO") is set to
	// represent that go run or go test being used.
	System string `json:"system" yaml:"system"`
	// Hostname of the system which built the binary.
	// To protect user's privacy this is populated only on CI/CD systems,
	// for locally built binaries this will be set to localhost.
	Host string `json:"host"  yaml:"host"`
	// Build timestamp in iso-8601. If not built using make or goreleaser,
	// defaults to "1970-01-01T00:00+00:00"
	Date string `json:"date" yaml:"date"`
}

// GoInfo Go version, platform and compiler
type GoInfo struct {
	// Runtime version of Go runtime.
	Runtime string `json:"version" yaml:"version"`
	// Platform this is of format GOOS/GORCH
	Platform string `json:"platform" yaml:"platform"`
	// Compiler Go compiler user. This is useful in determining if binary
	// was built using CGO. (Especially in containers)
	Compiler string `json:"compiler" yaml:"compiler"`
}

// GitInfo VCS info
type GitInfo struct {
	// Commit indicates which git hash the binary was built from
	// (Can be empty if not compiled with make/goreleaser)
	Commit string `json:"commit" yaml:"commit"`
	// Branch Git Branch. Can be empty if it cannot be determined,
	// not built using release of build scripts.
	Branch string `json:"branch" yaml:"branch"`
	// TreeState Git tree state. Can take clean/dirty or unknown.
	// By default this is set to unknown as go toolchain does not know about git tree state.
	TreeState string `json:"treeState" yaml:"treeState"`
}

// Info describes the build, revision and version information.
type Info struct {
	// VersionInfo API version. This is used by autopackager
	API uint8 `json:"api" yaml:"api"`
	// Version indicates which version of the binary is running.
	// semver compatible version string, of format
	// [MAJOR].[MINOR].[PATCH]-{OPTIONAL-SUFFIX}
	Version string    `json:"version" yaml:"version"`
	Go      GoInfo    `json:"go" yaml:"go"`
	Git     GitInfo   `json:"git" yaml:"git"`
	Build   BuildInfo `json:"build" yaml:"build"`
}

// GetShortVersion returns version prefixed with added build metadata.
// returned string is of format version+buildSystem-buildNumber.
// If version has prefix v, which is often the case with git tags,
// it is removed. If you want to get detailed version/build info,
// see version.Describe()
func GetShortVersion() string {
	return strings.TrimPrefix(version, "v")
}

// GetUserAgent Returns user agent with version string.
func GetUserAgent(program string) string {
	return fmt.Sprintf("%s/%s", program, strings.TrimPrefix(version, "v"))
}

// getBuildNumber this is necessary to convert string to uint64
func getBuildNumber() uint64 {
	bn, err := strconv.ParseUint(buildNumber, 10, 64)
	if err != nil {
		bn = 0
	}
	return bn
}

// Describe returns build, version and revision information in
// a structured format
func Describe() Info {
	info := Info{
		API:     2,
		Version: version,
		Go: GoInfo{
			Runtime:  runtime.Version(),
			Platform: runtime.GOOS + "/" + runtime.GOARCH,
			Compiler: runtime.Compiler,
		},
		Git: GitInfo{
			Commit:    gitCommit,
			Branch:    gitBranch,
			TreeState: gitTreeState,
		},
		Build: BuildInfo{
			Number: getBuildNumber(),
			System: buildSystem,
			Host:   buildHost,
			Date:   buildDate,
		},
	}

	// HACK to strip out Version during a test run for consistent test output
	// https://github.com/helm/helm/blob/f546ebb1aca7c45a09a71886b720b6e11d45e9d8/internal/version/version.go#L77
	if flag.Lookup("test.v") != nil {
		info.Go.Runtime = "go1.xx"
	}

	return info
}
