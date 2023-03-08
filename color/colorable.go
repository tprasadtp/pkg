package color

import "os"

// IsStderrColorable detects whether true colored(24 bit) output
// can be enabled for [os.Stdout].
//
// This supports both [CLICOLOR] and [NO_COLOR] standards, while ignoring
//
//   - Argument 'flag' ALWAYS takes priority and is NOT case sensitive.
//   - If flag is 'never', 'disable', no, 'none' or 'false', returns false
//   - If flag is 'always' or 'force', returns 'true'
//   - You should probably map this variable to your cli's --color flag.
//   - If you specify flag string to other than above specified, it is ignored.
//
// If environment variable CI is set to true, and none of the above conditions
// or environment variable in [CLICOLOR] and [NO_COLOR] standards are specified,
// this will return true to enable colored output in CI builds.
//
// On unsupported platforms, this always returns false.
//
// [NO_COLOR]: https://no-color.org/
// [CLICOLOR]: https://bixense.com/clicolors/
func IsStdoutColorable(flag string) bool {
	return isColorable(flag, IsTerminal(os.Stdout.Fd()))
}

// Detects whether true color(24 bit) output
// can be enabled for [os.Stderr].
//
// This supports both [CLICOLOR] and [NO_COLOR] standards.
//
//   - Argument 'flag' ALWAYS takes priority and is NOT case sensitive.
//   - If flag is 'never', 'disable' or 'none' or 'false', returns false
//   - If flag is 'always' or 'force', returns true
//   - You SHOULD map this variable to your cli's --color flag.
//   - If you specify flag string to other than above specified, it is ignored.
//   - This function will return false on unsupported platforms.
//
// If environment variable CI is set to true, and none of the above conditions
// or environment variable in [CLICOLOR] and [NO_COLOR] standards are specified,
// this will return true to enable colored output in CI builds.
// On unsupported platforms, this always returns false.
//
// [NO_COLOR]: https://no-color.org/
// [CLICOLOR]: https://bixense.com/clicolors/
func IsStderrColorable(flag string) bool {
	return isColorable(flag, IsTerminal(os.Stderr.Fd()))
}

// IsColorable detects whether true color(24 bit) output
// can be enabled for BOTH [os.Stderr] and [os.Stdout].
//
// This supports both [CLICOLOR] and [NO_COLOR] standards. ([Sigh!]...)
//
//   - Argument 'flag' ALWAYS takes priority and is NOT case sensitive.
//   - If flag is 'never', 'disable' or 'none' or 'false', returns false
//   - If flag is 'always' or 'force', returns 'true'
//   - You should probably map this variable to your cli's --color flag.
//   - If you specify flag string to other than above specified, it is ignored.
//
// If environment variable CI is set to true, and none of the above conditions
// or environment variable in [CLICOLOR] and [NO_COLOR] standards are specified,
// this will return true to enable colored output in CI builds.
//
// On unsupported platforms, this always returns false.
//
// [NO_COLOR]: https://no-color.org/
// [CLICOLOR]: https://bixense.com/clicolors/
// [Sigh!]: https://xkcd.com/927/
func IsColorable(flag string) bool {
	return isColorable(flag, IsTerminal(os.Stdout.Fd())) &&
		isColorable(flag, IsTerminal(os.Stderr.Fd()))
}
