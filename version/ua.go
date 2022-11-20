package version

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetUserAgentSuffix returns HTTP User-Agent suffix based on version information.
// Intended to build user agent which includes version information.
func GetUserAgentSuffix() string {
	v := Get()
	return fmt.Sprintf("%s/%s/%s", v.Version, v.Os, v.Platform)
}

// GetUserAgent returns HTTP User-Agent based on version information
// and name of the binary. This is similar to GetUserAgentSuffix,
// except this includes name of the binary in user agent and intended to be used
// directly. It is recommended to set this to a package level variable
// than setting it manually for each http request.
func GetUserAgent() string {
	v := Get()
	bin, err := os.Executable()
	if err != nil {
		bin = "unknown"
	} else {
		bin = filepath.Base(bin)
	}
	return fmt.Sprintf("%s/%s/%s/%s", bin, v.Version, v.Os, v.Platform)
}
