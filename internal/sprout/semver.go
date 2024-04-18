// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: GPLv3-only

package sprout

import sv "github.com/Masterminds/semver/v3"

func semver(version string) (*sv.Version, error) {
	return sv.NewVersion(version)
}
