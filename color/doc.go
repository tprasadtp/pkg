// This package only supports terminals which support true colors (24 bit).
// In case terminal does not support true color, input bytes/strings are
// left unmodified and returns a no-op equivalents.
//
// This is primarily intended for use with
// [tprasadtp/pkg/log] and [tprasadtp/pkg/cli] libraries
// and cannot provide API compatibility guarantees beyond those.
package color
