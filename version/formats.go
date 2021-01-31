package version

import (
	"bytes"
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

// YAMLIndent Indentation for YAML output
const YAMLIndent = 2

// AsJSON returns formatted JSON string to be printed.
func (v *Info) AsJSON() (string, error) {
	p, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}

	return string(p), nil
}

// AsText returns formatted printable text.
//
// Short Flag
//
// If short flag is set, it just returns shortVersion.
func (v *Info) AsText(short bool) (string, error) {
	if short {
		return fmt.Sprintf("%s\n", GetShortVersion()), nil
	}

	return fmt.Sprintf("• Version    : %s\n", v.Version) +
		"• Git\n" +
		fmt.Sprintf(" • Commit    : %s\n", v.Git.Commit) +
		fmt.Sprintf(" • Branch    : %s\n", v.Git.Branch) +
		fmt.Sprintf(" • TreeState : %s\n", v.Git.TreeState) +
		"• Build\n" +
		fmt.Sprintf(" • Number    : %d\n", v.Build.Number) +
		fmt.Sprintf(" • System    : %s\n", v.Build.System) +
		fmt.Sprintf(" • Host      : %s\n", v.Build.Host) +
		fmt.Sprintf(" • Date      : %s\n", v.Build.Date) +
		"• Go\n" +
		fmt.Sprintf(" • Runtime   : %s\n", v.Go.Runtime) +
		fmt.Sprintf(" • Platform  : %s\n", v.Go.Platform), nil
}

// AsYAML retutns formattd JSON string to be printed.
func (v *Info) AsYAML() (string, error) {
	w := bytes.NewBuffer(nil)
	enc := yaml.NewEncoder(w)
	defer enc.Close()
	enc.SetIndent(YAMLIndent)

	if err := enc.Encode(v); err != nil {
		return "", err
	}

	return w.String(), nil
}
