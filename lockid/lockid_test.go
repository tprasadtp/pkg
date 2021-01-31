package lockid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLockID(t *testing.T) {
	tests := []struct {
		name       string
		identifier string
		expect     string
	}{
		{name: "email", identifier: "security@dev.null", expect: "14e3c51b-ae25-aca0-5738-b3bd742d5ff2"},
		{name: "me", identifier: "me", expect: "2744ccd1-0c75-33bd-736a-d890f9dd5cab"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := GenerateLockID(tc.identifier)
			assert.Nil(t, err)
			assert.Equal(t, tc.expect, actual)
		})
	}
}

func TestLockIDErrors(t *testing.T) {
	tests := []struct {
		name       string
		identifier string
	}{
		{name: "empty", identifier: ""},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := GenerateLockID(tc.identifier)
			assert.NotNil(t, err)
			assert.Empty(t, actual)
		})
	}
}
