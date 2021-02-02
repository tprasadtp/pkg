package cfgpath

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const pathSep = string(os.PathSeparator)

func TestGetTokenFilePath(t *testing.T) {
	configDir, _ := os.UserConfigDir()
	tests := []struct {
		name      string
		partition string
		expect    string
	}{
		{
			name:      "simple",
			partition: "gotee",
			expect:    strings.Join([]string{configDir, "gotee", "auth.yml"}, pathSep),
		},
		{
			name:      "with-hyphen",
			partition: "gotee-config",
			expect:    strings.Join([]string{configDir, "gotee-config", "auth.yml"}, pathSep),
		},
		{
			name:      "with-nums",
			partition: "gotee255",
			expect:    strings.Join([]string{configDir, "gotee255", "auth.yml"}, pathSep),
		},
		{
			name:      "only-nums",
			partition: "255",
			expect:    strings.Join([]string{configDir, "255", "auth.yml"}, pathSep),
		},
		{
			name:      "with-underscore",
			partition: "gotee_config",
			expect:    strings.Join([]string{configDir, "gotee_config", "auth.yml"}, pathSep),
		},
		{
			name:      "with-underscore-hyphen",
			partition: "gotee_cfg-base",
			expect:    strings.Join([]string{configDir, "gotee_cfg-base", "auth.yml"}, pathSep),
		},
		{
			name:      "with-uppercase",
			partition: "GOTEE",
			expect:    strings.Join([]string{configDir, "GOTEE", "auth.yml"}, pathSep),
		},
		{
			name:      "with-uppercase-hyphen",
			partition: "GOTEE-CONFIG",
			expect:    strings.Join([]string{configDir, "GOTEE-CONFIG", "auth.yml"}, pathSep),
		},
		{
			name:      "with-mixed",
			partition: "GoTee-Config2-Latest",
			expect:    strings.Join([]string{configDir, "GoTee-Config2-Latest", "auth.yml"}, pathSep),
		},
		{
			name:      "with-underscore-hyphen-nums",
			partition: "gotee_cfg2-base",
			expect:    strings.Join([]string{configDir, "gotee_cfg2-base", "auth.yml"}, pathSep),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := GetTokenFilePath(tc.partition)
			assert.Nil(t, err)
			assert.Equal(t, tc.expect, actual)
		})
	}
}

func TestInvalidGetTokenFilePath(t *testing.T) {
	tests := []struct {
		name      string
		partition string
	}{
		{
			name:      "special-chars",
			partition: "gotee#",
		},
		{
			name:      "empty",
			partition: "",
		},
		{
			name:      "spaces",
			partition: "     ",
		},
		{
			name:      "underscore",
			partition: "_",
		},
		{
			name:      "starts-with-hyphen",
			partition: "-gotee",
		},
		{
			name:      "starts-with-underscore",
			partition: "_gotee",
		},
		{
			name:      "starts-with-dot",
			partition: ".gotee",
		},
		{
			name:      "starts-with-special#",
			partition: "#gotee",
		},
		{
			name:      "starts-with-special$",
			partition: "$gotee",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := GetTokenFilePath(tc.partition)
			assert.ErrorIs(t, err, ErrInvalidPartition)
			assert.Equal(t, "", actual)
		})
	}
}

func TestGetUpcheckFilePath(t *testing.T) {
	cacheDir, _ := os.UserCacheDir()
	tests := []struct {
		name      string
		partition string
		expect    string
	}{
		{
			name:      "simple",
			partition: "gotee",
			expect:    strings.Join([]string{cacheDir, "gotee", "upcheck.yml"}, pathSep),
		},
		{
			name:      "with-hyphen",
			partition: "gotee-config",
			expect:    strings.Join([]string{cacheDir, "gotee-config", "upcheck.yml"}, pathSep),
		},
		{
			name:      "with-nums",
			partition: "gotee255",
			expect:    strings.Join([]string{cacheDir, "gotee255", "upcheck.yml"}, pathSep),
		},
		{
			name:      "only-nums",
			partition: "255",
			expect:    strings.Join([]string{cacheDir, "255", "upcheck.yml"}, pathSep),
		},
		{
			name:      "with-underscore",
			partition: "gotee_config",
			expect:    strings.Join([]string{cacheDir, "gotee_config", "upcheck.yml"}, pathSep),
		},
		{
			name:      "with-underscore-hyphen",
			partition: "gotee_cfg-base",
			expect:    strings.Join([]string{cacheDir, "gotee_cfg-base", "upcheck.yml"}, pathSep),
		},
		{
			name:      "with-uppercase",
			partition: "GOTEE",
			expect:    strings.Join([]string{cacheDir, "GOTEE", "upcheck.yml"}, pathSep),
		},
		{
			name:      "with-uppercase-hyphen",
			partition: "GOTEE-CONFIG",
			expect:    strings.Join([]string{cacheDir, "GOTEE-CONFIG", "upcheck.yml"}, pathSep),
		},
		{
			name:      "with-mixed",
			partition: "GoTee-Config2-Latest",
			expect:    strings.Join([]string{cacheDir, "GoTee-Config2-Latest", "upcheck.yml"}, pathSep),
		},
		{
			name:      "with-underscore-hyphen-nums",
			partition: "gotee_cfg2-base",
			expect:    strings.Join([]string{cacheDir, "gotee_cfg2-base", "upcheck.yml"}, pathSep),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := GetUpcheckFilePath(tc.partition)
			assert.Nil(t, err)
			assert.Equal(t, tc.expect, actual)
		})
	}
}

func TestInvalidGetUpcheckPath(t *testing.T) {
	tests := []struct {
		name      string
		partition string
	}{
		{
			name:      "special-chars",
			partition: "gotee#",
		},
		{
			name:      "empty",
			partition: "",
		},
		{
			name:      "spaces",
			partition: "     ",
		},
		{
			name:      "underscore",
			partition: "_",
		},
		{
			name:      "starts-with-hyphen",
			partition: "-gotee",
		},
		{
			name:      "starts-with-underscore",
			partition: "_gotee",
		},
		{
			name:      "starts-with-dot",
			partition: ".gotee",
		},
		{
			name:      "starts-with-special#",
			partition: "#gotee",
		},
		{
			name:      "starts-with-special$",
			partition: "$gotee",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := GetUpcheckFilePath(tc.partition)
			assert.ErrorIs(t, err, ErrInvalidPartition)
			assert.Equal(t, "", actual)
		})
	}
}
