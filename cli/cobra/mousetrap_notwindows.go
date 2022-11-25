//go:build !windows
// +build !windows

package cobra

func preExecHook(cmd *Command) {}
