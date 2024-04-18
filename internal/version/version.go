// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: GPLv3-only

// Package version is a helper for processing
// VCS, build, and runtime information of the binary.
//
// You can inject them at build time via ld flags.
// If not already injected, this uses debug.ReadBuildInfo,
// to get version information if any.
package version

import (
	"runtime"
	"runtime/debug"
	"sync"
)

// Can override these at compile time.
var (
	// version is usually the git tag. MUST be semver compatible.
	//
	// You can override at build time using
	// 		-X github.com/pkg/version.version = "your-desired-version"
	version = ""

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
	buildDate = ""

	// gitTreeState is git tree state.
	// You can override at build time using
	// 		-X github.com/pkg/version.gitTreeState = "clean"
	gitTreeState = ""

	// once is sync.Once for getting build info from ReadBuildInfo.
	once sync.Once
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

	// Commit indicates which git sha1 commit hash.
	Commit string `json:"commit" yaml:"commit"`

	// GitTreeState
	GitTreeState string `json:"gitTreeState" yaml:"gitTreeState"`

	// BuildDate date of the build.
	// You can set this to CommitDate to get truly reproducible and verifiable builds.
	BuildDate string `json:"buildDate" yaml:"buildDate"`

	// GoVersion version of Go runtime.
	GoVersion string `json:"goVersion" yaml:"goVersion"`

	// OperatingSystem this is operating system in GOOS
	Os string `json:"os" yaml:"os"`

	// Arch this is system Arch
	Arch string `json:"platform" yaml:"arch"`

	// Compiler is Go compiler.
	// This is useful in determining if binary was built using gccgo.
	Compiler string `json:"compiler" yaml:"compiler"`
}

// GetInfo returns version information. This usually relies on
// build tools injecting version info via ld flags.
func GetInfo() Info {
	// Read from debug.ReadBuildInfo() if required.
	once.Do(func() {
		// only if commit or build date are not defined.
		if commit == "" || buildDate == "" {
			v, ok := debug.ReadBuildInfo()
			if ok {
				for _, item := range v.Settings {
					switch item.Key {
					case "vcs.revision":
						if commit == "" {
							commit = item.Value
						}
					case "vcs.time":
						if buildDate == "" {
							buildDate = item.Value
						}
					case "vcs.modified":
						if gitTreeState == "" {
							gitTreeState = item.Value
						}
					}
				}
			}
		}
	})

	return Info{
		Version:      version,
		Commit:       commit,
		BuildDate:    buildDate,
		GitTreeState: gitTreeState,
		GoVersion:    runtime.Version(),
		Os:           runtime.GOOS,
		Arch:         runtime.GOARCH,
		Compiler:     runtime.Compiler,
	}
}
