package cfgpath

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
)

// cfgFileNameRegEx Config filename regex
const cfgFileNameRegEx = "^[A-Za-z0-9]([A-Za-z0-9_-]*[A-Za-z0-9])?$"

// ErrInvalidPartition is returned when partition name is invalid.
//
// A Partiton name MUST start and end with an alphanumeric characters
// and can ONLY contain alphanumeric characters along with hyphens and underscore.
var ErrInvalidPartition = errors.New("partition name must start and end with alnum and can only contain alnum, hyphens and underscores")

// GetTokenFilePath Get path of token file.
//
// Token file is a file to store authentication data like oauth tokens.
// This file MUST be protected with correct filesystem permissions.
func GetTokenFilePath(partitionName string) (string, error) {
	if match, _ := regexp.MatchString(cfgFileNameRegEx, partitionName); !match {
		return "", ErrInvalidPartition
	}

	xdgConfigDir, err := os.UserConfigDir()

	if err != nil {
		return "", err
	}
	return filepath.Join(xdgConfigDir, partitionName, "auth.yml"), nil
}

// GetUpcheckFilePath returns path of Updater config file.
// This returns ErrInvalidPartition if partition name is invalid.
// This file is used to cache update info which is usually fetched from
// its UpcheckEndpoint and Rollkeeperendpoint.
func GetUpcheckFilePath(partitionName string) (string, error) {
	if match, _ := regexp.MatchString(cfgFileNameRegEx, partitionName); !match {
		return "", ErrInvalidPartition
	}

	xdgCacheDir, err := os.UserCacheDir()

	if err != nil {
		return "", err
	}
	return filepath.Join(xdgCacheDir, partitionName, "upcheck.yml"), nil
}
